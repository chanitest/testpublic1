package handlers

import (
	"context"
	"fmt"
	"github.com/Checkmarx-Containers/containers-worker/internal/ContainersEngine"
	"github.com/Checkmarx-Containers/containers-worker/internal/messages"
	"github.com/Checkmarx-Containers/containers-worker/internal/scanStatus"
	scanStatusService "github.com/Checkmarx-Containers/containers-worker/internal/scanStatus/service"
	"github.com/Checkmarx-Containers/containers-worker/pkg/handlers"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	"github.com/checkmarxDev/scans-flow/v8/pkg/api/workflow"
	"github.com/checkmarxDev/scans/pkg/api/scans"
	redisV8 "github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScanHandler struct {
	logger        extendedlogger.ExtendedLogger
	rabbit        messages.RabbitMqInt
	redisClient   redisV8.UniversalClient
	statusService scanStatusService.ScanStatusServiceInt
}

func NewScanHandler(logger *extendedlogger.ExtendedLogger, rabbit messages.RabbitMqInt, redisClient redisV8.UniversalClient, statusService scanStatusService.ScanStatusServiceInt) (*ScanHandler, error) {
	wh := &ScanHandler{
		logger:        *logger,
		rabbit:        rabbit,
		redisClient:   redisClient,
		statusService: statusService,
	}
	return wh, nil
}

func (sh *ScanHandler) Handle(msg *workflow.Message) <-chan handlers.StatusInfo {
	statusInfoCh := make(chan handlers.StatusInfo)
	go func(workStatChan chan handlers.StatusInfo) {
		defer close(workStatChan)
		sh.handleWork(msg, workStatChan)
	}(statusInfoCh)
	return statusInfoCh
}

func (sh *ScanHandler) handleWork(msg *workflow.Message, statusInfo chan handlers.StatusInfo) {
	ctx := context.Background()

	sh.logger.Info(fmt.Sprintf("got new scan message, msg: `%s`", msg.String()))

	scan := scans.Scan{}
	if err := proto.Unmarshal(msg.Input.Value, &scan); err != nil {
		sh.HandleError(scan.Id, "Could not unmarshall scan object", err, msg, statusInfo)
	}

	value := &scan_status_service.ScanStatus{
		Id:            scan.Id,
		Status:        scan_status_service.ContainersScanStatusQueued,
		StepStatus:    []scan_status_service.ScanStepStatus{},
		CorrelationId: "",
		UpdateTime:    timestamppb.Now(),
		Seen:          true,
	}

	err := sh.statusService.SaveScanStatus(ctx, *value)
	if err != nil {
		sh.HandleError(scan.Id, "Could not set cache key", err, msg, statusInfo)
	}

	newScanRequest := &ContainersEngine.NewScanRequest{
		ScanId:        scan.Id,
		ProjectId:     scan.Project.Id,
		TenantId:      msg.Tenant,
		Url:           msg.CorrelationId,
		CorrelationId: msg.CorrelationId,
		ScanType:      scan.Type,
	}

	err = sh.rabbit.SendNewScanRequest(ctx, newScanRequest)
	if err != nil {
		sh.HandleError(scan.Id, "Could not publish new scan message to messages", err, msg, statusInfo)
	}

	sh.logger.Info("Successfully published new scan request to containers engine.")
	handlers.ReturnJobStatus(statusInfo, scan.Id, handlers.JobStatusCompleted, nil, "", nil)
}

func (sh *ScanHandler) HandleError(scanId, errorMsg string, err error, msg *workflow.Message, statusInfo chan handlers.StatusInfo) {
	sh.logger.Error(fmt.Sprintf("%s. Message: %s", errorMsg, msg.String()), err)
	handlers.ReturnJobStatus(statusInfo, scanId, handlers.JobStatusFailed, nil, errorMsg, err)
}
