package worker

import (
	"context"
	"fmt"
	"github.com/Checkmarx-Containers/containers-worker/pkg/handlers"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	"github.com/checkmarxDev/scans-flow/v8/pkg/zbAbstractionLayer/workflow/updater"
	wf "github.com/checkmarxDev/scans-flow/v8/pkg/zbAbstractionLayer/workflow/worker"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	ServiceName = "ContainerWorker"
)

type Worker struct {
	ZBWorker    wf.Worker
	workHandler handlers.WorkHandler
	logger      extendedlogger.ExtendedLogger
}

func NewWorker(ctx context.Context,
	zeebeClient zbc.Client,
	workerName string,
	workTimeoutMinutes, maxJobCount uint,
	jobType,
	rabbitConnectionString, exchange string,
	rabbitSkipTLS bool,
	rabbitRetries int,
	rabbitRetryWaitDuration time.Duration,
	workHandler handlers.WorkHandler,
) (*Worker, error) {

	logger := extendedlogger.GetExtendedLogger()

	log.Debug().Msgf("Connecting to rabbitmq. addr=%s, skipTLS=%t, exchange=%s", rabbitConnectionString, rabbitSkipTLS, "")

	rup := updater.Empty()
	if rabbitConnectionString != "" {
		rup = updater.NewRabbitUpdater(rabbitConnectionString, exchange, ServiceName, rabbitSkipTLS, rabbitRetries, rabbitRetryWaitDuration)
	}
	log.Debug().Msgf("Create new workload, jobType=%s", jobType)

	zbWorker, err := wf.NewZBWorker(ctx, zeebeClient, nil, jobType, workerName, workTimeoutMinutes, maxJobCount, workTimeoutMinutes, int(maxJobCount), rup, nil)
	if err != nil {
		return nil, fmt.Errorf("cant connect to scan exchange : %w", err)
	}

	return InitWorker(zbWorker, workHandler, logger), nil
}

func InitWorker(wo wf.Worker, workHandler handlers.WorkHandler, logger *extendedlogger.ExtendedLogger) *Worker {
	w := &Worker{
		ZBWorker:    wo,
		workHandler: workHandler,
		logger:      *logger,
	}
	w.ZBWorker.JobHandler(w.handle)
	return w
}

func (w *Worker) ListenAndServe() error {
	log.Info().Msgf("New scan worker is listening for new messages")
	w.ZBWorker.JobHandler(w.handle)
	return w.ZBWorker.Start()
}

func (w *Worker) Close() error {
	log.Info().Msg("Closing new scan worker")
	return w.ZBWorker.Close()
}

func (w *Worker) handle(job *wf.Job) {
	for message := range w.workHandler.Handle(job.Msg) {

		w.logger.Info(fmt.Sprintf("finished job. elementId = %s with state - %s for scan id - %s", job.ElementID, message.State, job.ScanID()))

		switch message.State {
		case handlers.JobStatusCompleted:
			w.logger.Info(fmt.Sprintf("Work on job completed."))
			if err := w.ZBWorker.CompleteJob(context.Background(), job); err != nil {
				w.logger.Error("Failed to send update that job completed.", err)
			}
		case handlers.JobStatusFailed:
			w.logger.Error(fmt.Sprintf("Work on job failed. Error message: %s", message.ErrMsg), message.Err)
			if err := w.ZBWorker.FailJob(context.Background(), job, message.ErrMsg, message.ErrCode); err != nil {
				w.logger.Error(fmt.Sprintf("Failed to send update that job failed."), message.Err)
			}
		case handlers.JobStatusUpdateProgress:
			w.handleProgressUpdate(job, message)
		default:
			w.handleProgressUpdate(job, message)
		}
	}
}

func (w *Worker) handleProgressUpdate(job *wf.Job, message handlers.StatusInfo) {
	w.logger.Info(fmt.Sprintf("Update scan progress."))
	for _, update := range message.Progress {
		if err := w.ZBWorker.UpdateStatus(context.Background(), job.ScanID(), update.Name, update.Status, job.Tenant()); err != nil {
			w.logger.Error(fmt.Sprintf("Failed to update scan progress."), err)
		}
	}

	if err := w.ZBWorker.ReturnJob(context.Background(), job, false, "", 0); err != nil {
		w.logger.Error(fmt.Sprintf("Failed to send update that job completed."), err)
	}
}
