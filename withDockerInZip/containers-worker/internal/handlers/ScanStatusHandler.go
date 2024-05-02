package handlers

import (
	"context"
	"fmt"
	scanstatus "github.com/Checkmarx-Containers/containers-worker/internal/scanStatus"
	scanStatusService "github.com/Checkmarx-Containers/containers-worker/internal/scanStatus/service"
	"github.com/Checkmarx-Containers/containers-worker/pkg/handlers"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	"github.com/checkmarxDev/scans-flow/v8/pkg/api/workflow"
	"github.com/checkmarxDev/scans/pkg/api/scans"
	redisV8 "github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
)

type ScanStatusHandler struct {
	logger        extendedlogger.ExtendedLogger
	redisClient   redisV8.UniversalClient
	statusService scanStatusService.ScanStatusServiceInt
}

func NewScanStatusHandler(logger *extendedlogger.ExtendedLogger, redisClient redisV8.UniversalClient, statusService scanStatusService.ScanStatusServiceInt) (*ScanStatusHandler, error) {
	ssh := &ScanStatusHandler{
		logger:        *logger,
		redisClient:   redisClient,
		statusService: statusService,
	}
	return ssh, nil
}

func (sh *ScanStatusHandler) Handle(msg *workflow.Message) <-chan handlers.StatusInfo {
	statusInfoCh := make(chan handlers.StatusInfo)
	go func(workStatChan chan handlers.StatusInfo) {
		defer close(workStatChan)
		sh.handleWork(msg, workStatChan)
	}(statusInfoCh)
	return statusInfoCh
}

func (sh *ScanStatusHandler) handleWork(msg *workflow.Message, statusInfo chan handlers.StatusInfo) {

	sh.logger.Info(fmt.Sprintf("got new message in scans status handler, msg = `%s`", msg.String()))
	ctx := context.Background()
	scan := scans.Scan{}
	if err := proto.Unmarshal(msg.Input.Value, &scan); err != nil {
		sh.logger.Error("Unable to unmarshal scan input", err)
		handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusFailed, nil, "Unable to unmarshal scan input", err)
		return
	}
	status, err := sh.statusService.GetScanStatusDetails(ctx, scan.Id)
	if err != nil {
		sh.logger.Error("Unable to get scan status from cache", err)
		handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusFailed, nil, "Unable to get scan status from cache", err)
		return
	}
	if status == nil {
		sh.logger.Error("Could not find scan status in cache db", fmt.Errorf("cache key could not be found for scanId = %s", scan.Id))
		handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusFailed, nil, "Could not find scan status in cache db", fmt.Errorf("cache key could not be found for scanId = %s", scan.Id))
		return
	}

	//update steps that were not updated yet
	var stepUpdates []handlers.StepUpdate
	for i, stepStatus := range status.StepStatus {
		if !stepStatus.Seen {
			sh.logger.Info(fmt.Sprintf("Found step status update, StepName = %s, StepStatus = %s", stepStatus.StepName, stepStatus.Status))
			stepUpdates = append(stepUpdates, handlers.StepUpdate{Name: stepStatus.StepName, Status: stepStatus.Status})
			stepStatus.Seen = true
			status.StepStatus[i] = stepStatus
		}
	}
	if len(stepUpdates) != 0 {
		err = sh.statusService.SaveScanStatus(ctx, *status)
		if err != nil {
			sh.logger.Error("Unable to update scan step statuses", err)
		}
		handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusUpdateProgress, stepUpdates, "", nil)
		return
	}

	//check if status changed - if not retry
	if status.Seen {
		sh.logger.Info(fmt.Sprintf("No Status update for ScanId: %s", scan.Id))
		handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusRetry, nil, "", nil)
		return
	}

	//return scan status
	switch status.Status {
	case scanstatus.ContainersScanStatusQueued, scanstatus.ContainersScanStatusWorking:
		sh.UpdateStatus(ctx, status, scan.Id)
		handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusRetry, nil, "", nil)
		return
	case scanstatus.ContainersScanStatusCompleted:
		sh.UpdateStatus(ctx, status, scan.Id)
		handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusCompleted, nil, "", nil)
		return
	case scanstatus.ContainersScanStatusFailed, scanstatus.ContainersScanStatusCanceled, scanstatus.ContainersScanStatusTimeOut:
		sh.UpdateStatus(ctx, status, scan.Id)
		handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusFailed, nil, status.Error, nil)
		return
	}
}

func (sh *ScanStatusHandler) UpdateStatus(ctx context.Context, status *scanstatus.ScanStatus, id string) {
	sh.logger.Info(fmt.Sprintf("Found scan status update, Status = %s, scanId = %s", status.Status, id))
	status.Seen = true
	err := sh.statusService.SaveScanStatus(ctx, *status)
	if err != nil {
		sh.logger.Error("Unable to update scan step statuses", err)
	}
}
