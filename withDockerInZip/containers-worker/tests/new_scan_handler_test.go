package tests

import (
	"context"
	"github.com/Checkmarx-Containers/containers-worker/internal/ContainersEngine"
	handlers2 "github.com/Checkmarx-Containers/containers-worker/internal/handlers"
	"github.com/Checkmarx-Containers/containers-worker/internal/scanStatus"
	"github.com/Checkmarx-Containers/containers-worker/pkg/handlers"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	"github.com/go-redis/redismock/v8"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScanHandler_Handle(t *testing.T) {
	tests := []struct {
		name              string
		rabbitMockFunc    func(ctx context.Context, newScanRequest *ContainersEngine.NewScanRequest) error
		statusServiceFunc func(ctx context.Context, scanStatusMoqParam scan_status_service.ScanStatus) error
		expectedJobStatus handlers.WorkState
	}{
		{
			name: "Happy Flow",
			rabbitMockFunc: func(ctx context.Context, newScanRequest *ContainersEngine.NewScanRequest) error {
				return nil
			},
			statusServiceFunc: func(ctx context.Context, scanStatusMoqParam scan_status_service.ScanStatus) error {
				return nil
			},
			expectedJobStatus: handlers.JobStatusCompleted,
		},
		{
			name: "Failed Flow",
			// Define mock behaviors for the second test case
			rabbitMockFunc: func(ctx context.Context, newScanRequest *ContainersEngine.NewScanRequest) error {
				return errors.New("Some error")
			},
			statusServiceFunc: func(ctx context.Context, scanStatusMoqParam scan_status_service.ScanStatus) error {
				return errors.New("Some error")
			},
			expectedJobStatus: handlers.JobStatusFailed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := extendedlogger.GetExtendedLogger()

			mockClient, _ := redismock.NewClientMock()

			rabbitMock := &RabbitMqIntMock{
				SendNewScanRequestFunc: tt.rabbitMockFunc,
			}

			statusServiceMock := &ScanStatusServiceIntMock{
				SaveScanStatusFunc: tt.statusServiceFunc,
			}

			sh, err := handlers2.NewScanHandler(logger, rabbitMock, mockClient, statusServiceMock)
			assert.NoError(t, err)

			statusInfoCh := sh.Handle(GetMessage())
			msg := <-statusInfoCh
			assert.Equal(t, tt.expectedJobStatus, msg.State)
		})
	}
}
