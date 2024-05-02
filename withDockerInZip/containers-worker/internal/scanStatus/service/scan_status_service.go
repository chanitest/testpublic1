package scan_status_service

import (
	"context"
	"fmt"
	"github.com/Checkmarx-Containers/containers-worker/internal/scanStatus"
	pb "github.com/Checkmarx-Containers/containers-worker/pkg/api/scan_status"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScanStepStatus struct {
	Status     string
	StepName   string
	UpdateTime *timestamppb.Timestamp
}

//go:generate moq -out ../../../tests/scan_status_service_test.go . ScanStatusServiceInt
type ScanStatusServiceInt interface {
	UpdateScanStatus(ctx context.Context, scanId string, status string, updateTime *timestamppb.Timestamp) error
	UpdateScanStatusDetails(ctx context.Context, scanId string, stepName string, status string, updateTime *timestamppb.Timestamp) error
	GetScanStatusDetails(ctx context.Context, scanId string) (*scan_status_service.ScanStatus, error)
	SaveScanStatus(ctx context.Context, scanStatus scan_status_service.ScanStatus) error
}

type ScanStatusService struct {
	logger      extendedlogger.ExtendedLogger
	redisClient redis.UniversalClient
}

func NewScanStatusService(logger *extendedlogger.ExtendedLogger, redisClient redis.UniversalClient) *ScanStatusService {
	return &ScanStatusService{
		logger:      *logger,
		redisClient: redisClient,
	}
}

func (s *ScanStatusService) UpdateScanStatus(ctx context.Context, scanId string, status string, updateTime *timestamppb.Timestamp) error {
	statusObj, err := GetScanStatusFromRedis(ctx, s.redisClient, scanId)
	if err != nil {
		return err
	}
	if statusObj == nil {
		s.logger.Warn(fmt.Sprintf("Scan status could not be found in cache. scanId= %s", scanId))
		return fmt.Errorf("scan status does not exist in cache DB")
	}
	//check if status update is the same status
	if statusObj.Status == status {
		s.logger.Info(fmt.Sprintf("Trying to update the same status. no need for update. scanId= %s", scanId))
		return nil
	}

	//update new status and save
	statusObj.Status = status
	statusObj.Seen = false

	err = s.SaveScanStatus(ctx, *statusObj)
	if err != nil {
		return err
	}
	return nil
}

func (s *ScanStatusService) UpdateScanStatusDetails(ctx context.Context, scanId string, stepName string, status string, updateTime *timestamppb.Timestamp) error {
	statusObj, err := GetScanStatusFromRedis(ctx, s.redisClient, scanId)
	if err != nil {
		return err
	}

	if statusObj == nil {
		s.logger.Warn(fmt.Sprintf("Scan status details could not be found in cache. scanId= %s", scanId))
		return fmt.Errorf("scan status details does not exist in cache DB")
	}

	//check if step status changed
	for _, stepStatus := range statusObj.StepStatus {
		if stepStatus.StepName == stepName {
			if stepStatus.Status == status {
				return nil
			} else {
				stepStatus.Status = status
				err = s.SaveScanStatus(ctx, *statusObj)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}

	//add new scan step status
	stepStatusesUpdated := append(statusObj.StepStatus, scan_status_service.ScanStepStatus{
		Status:     status,
		StepName:   stepName,
		UpdateTime: timestamppb.Now(),
		Seen:       false,
	})

	//update new status and save
	statusObj.StepStatus = stepStatusesUpdated
	err = s.SaveScanStatus(ctx, *statusObj)
	if err != nil {
		return err
	}
	return nil
}

func (s *ScanStatusService) GetScanStatusDetails(ctx context.Context, scanId string) (*scan_status_service.ScanStatus, error) {
	status, err := GetScanStatusFromRedis(ctx, s.redisClient, scanId)
	if err != nil {
		return nil, err
	}
	if status == nil {
		s.logger.Warn("Scan Status was not found in cache DB. ScanId = %s", scanId)
		return nil, nil
	}
	return status, nil
}

func (s *ScanStatusService) SaveScanStatus(ctx context.Context, scanStatus scan_status_service.ScanStatus) error {
	s.logger.Info(fmt.Sprintf("saving scan status for scanId = %s", scanStatus.Id))
	err := UpdateScanStatusToRedis(ctx, s.redisClient, scanStatus)
	if err != nil {
		s.logger.Error("Could not update scan status", err)
		return err
	}
	return nil
}

func ConvertScanStatusToProto(scanStatus *scan_status_service.ScanStatus) *pb.GetScanStatusDetailsResponse {
	stepStatus := make([]*pb.StepStatus, len(scanStatus.StepStatus))
	for i, step := range scanStatus.StepStatus {
		stepStatus[i] = &pb.StepStatus{
			Status:     step.Status,
			Name:       step.StepName,
			UpdateTime: step.UpdateTime,
		}
	}

	return &pb.GetScanStatusDetailsResponse{
		Id:            scanStatus.Id,
		Status:        scanStatus.Status,
		StepStatus:    stepStatus,
		CorrelationId: scanStatus.CorrelationId,
		UpdateTime:    scanStatus.UpdateTime,
	}
}
