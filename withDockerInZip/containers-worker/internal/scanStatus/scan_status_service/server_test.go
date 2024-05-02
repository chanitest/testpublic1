package scan_status_service

import (
	"context"
	scanStatusService "github.com/Checkmarx-Containers/containers-worker/internal/scanStatus"
	pb1 "github.com/Checkmarx-Containers/containers-worker/pkg/api/scan_status"
	pb "google.golang.org/protobuf/types/known/timestamppb"

	"testing"
	"time"

	"github.com/checkmarxDev/lumologger/extendedlogger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockScanStatusService is a mock implementation for scanStatusService.ScanStatusService
type MockScanStatusService struct {
	mock.Mock
}

func (m *MockScanStatusService) UpdateScanStatus(ctx context.Context, id, status string, updateTime *pb.Timestamp) error {
	args := m.Called(ctx, id, status, updateTime)
	return args.Error(0)
}

func (m *MockScanStatusService) UpdateScanStatusDetails(ctx context.Context, id, stepName, status string, updateTime *pb.Timestamp) error {
	args := m.Called(ctx, id, stepName, status, updateTime)
	return args.Error(0)
}

func (m *MockScanStatusService) GetScanStatusDetails(ctx context.Context, id string) (*scanStatusService.ScanStatus, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*scanStatusService.ScanStatus), args.Error(1)
}
func (m *MockScanStatusService) SaveScanStatus(ctx context.Context, scanStatus scanStatusService.ScanStatus) error {
	return nil
}

func TestGRPCServer_UpdateScanStatus(t *testing.T) {
	// Create a logger
	logger := extendedlogger.GetExtendedLogger()

	// Create a mock for ScanStatusService
	statusServiceMock := &MockScanStatusService{}

	server := NewScanStatusGRPCServer(1243, statusServiceMock, logger)

	// Set up expectations for the mock
	expectedID := "123"
	expectedStatus := "COMPLETED"
	expectedUpdateTime := &pb.Timestamp{Seconds: time.Now().Unix(), Nanos: 0}
	statusServiceMock.On("UpdateScanStatus", mock.Anything, expectedID, expectedStatus, expectedUpdateTime).Return(nil)

	// Call the method being tested
	response, err := server.UpdateScanStatus(context.Background(), &pb1.UpdateScanStatusRequest{
		Id:         expectedID,
		Status:     expectedStatus,
		UpdateTime: expectedUpdateTime,
	})

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "SUCCESS", response.UpdateStatus)

	// Assert that expectations for the mock were met
	statusServiceMock.AssertExpectations(t)
}
