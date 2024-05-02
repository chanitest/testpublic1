package tests

import (
	"context"
	handlers2 "github.com/Checkmarx-Containers/containers-worker/internal/handlers"
	"github.com/Checkmarx-Containers/containers-worker/internal/scanStatus"
	"github.com/Checkmarx-Containers/containers-worker/pkg/handlers"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	"github.com/checkmarxDev/scans-flow/v8/pkg/api/workflow"
	"github.com/checkmarxDev/scans-flow/v8/pkg/protoApi/scan_request"
	"github.com/go-redis/redismock/v8"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanStatusHandler_Handle(t *testing.T) {
	tests := []struct {
		name                     string
		getScanStatusDetailsFunc func(ctx context.Context, scanId string) (*scan_status_service.ScanStatus, error)
		expectedJobStatus        handlers.WorkState
	}{
		{
			name: "Happy Flow",

			getScanStatusDetailsFunc: func(ctx context.Context, scanId string) (*scan_status_service.ScanStatus, error) {
				return GenerateStatus(scan_status_service.ContainersScanStatusCompleted, false), nil
			},
			expectedJobStatus: handlers.JobStatusCompleted,
		},
		{
			name: "Happy Flow - no update",

			getScanStatusDetailsFunc: func(ctx context.Context, scanId string) (*scan_status_service.ScanStatus, error) {
				return GenerateStatus(scan_status_service.ContainersScanStatusQueued, true), nil
			},
			expectedJobStatus: handlers.JobStatusRetry,
		},
		{
			name: "Failed Flow - status was not found",
			// Define mock behaviors for the second test case

			getScanStatusDetailsFunc: func(ctx context.Context, scanId string) (*scan_status_service.ScanStatus, error) {
				return nil, errors.New("Some error")
			},
			expectedJobStatus: handlers.JobStatusFailed,
		},
		{
			name: "Failed Flow - status is null",
			// Define mock behaviors for the second test case

			getScanStatusDetailsFunc: func(ctx context.Context, scanId string) (*scan_status_service.ScanStatus, error) {
				return nil, nil
			},
			expectedJobStatus: handlers.JobStatusFailed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := extendedlogger.GetExtendedLogger()

			mockClient, _ := redismock.NewClientMock()

			statusServiceMock := &ScanStatusServiceIntMock{
				GetScanStatusDetailsFunc: tt.getScanStatusDetailsFunc,
				SaveScanStatusFunc: func(ctx context.Context, scanStatusMoqParam scan_status_service.ScanStatus) error {
					return nil
				},
			}

			sh, err := handlers2.NewScanStatusHandler(logger, mockClient, statusServiceMock)
			assert.NoError(t, err)

			statusInfoCh := sh.Handle(GetMessage())
			msg := <-statusInfoCh
			assert.Equal(t, tt.expectedJobStatus, msg.State)
		})
	}
}

func GenerateStatus(status string, seen bool) *scan_status_service.ScanStatus {
	return &scan_status_service.ScanStatus{
		Id:                     "",
		Status:                 status,
		StepStatus:             nil,
		WorkflowCorrelationKey: "",
		CorrelationId:          "",
		UpdateTime:             nil,
		Seen:                   seen,
		Error:                  "",
	}
}

func GetMessage() *workflow.Message {
	r := &scan_request.Scan{
		Id: "87654",
		Project: &scan_request.Project{
			Id:   "9753",
			Name: "test-project",
			Tags: nil,
		},
		Configs:   nil,
		Tags:      nil,
		CreatedAt: nil,
		Type:      "upload",
		Handler:   nil,
	}

	anyTypeURL := r.ProtoReflect().Descriptor().FullName()
	serializedScan, _ := proto.Marshal(r)

	message := &workflow.Message{
		WorkflowCorrelationKey: "4321",
		Tenant:                 "test_tenant",
		TenantName:             "test_tenant",
		CorrelationId:          "4321",
		Input: &any.Any{
			TypeUrl: string(anyTypeURL),
			Value:   serializedScan,
		},
	}
	return message
}
