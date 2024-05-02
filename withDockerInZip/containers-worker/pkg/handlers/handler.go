package handlers

import (
	"fmt"
	"github.com/checkmarxDev/scans-flow/v8/pkg/api/workflow"
)

type WorkState int32

const (
	JobStatusQueued         WorkState = 0
	JobStatusRunning        WorkState = 1
	JobStatusCompleted      WorkState = 2
	JobStatusFailed         WorkState = 3
	JobStatusCanceled       WorkState = 4
	JobStatusRetry          WorkState = 5
	JobStatusUpdateProgress WorkState = 5
)

func (w WorkState) String() string {
	switch w {
	case JobStatusQueued:
		return "Queued"
	case JobStatusRunning:
		return "Running"
	case JobStatusCompleted:
		return "Completed"
	case JobStatusFailed:
		return "Failed"
	case JobStatusCanceled:
		return "Canceled"
	case JobStatusUpdateProgress:
		return "UpdateProgress"
	default:
		return fmt.Sprintf("%d", int(w))
	}
}

type StatusInfo struct {
	ScanID, Info string
	Progress     []StepUpdate
	State        WorkState
	Err          error
	ErrMsg       string
	ErrCode      int
}

type StepUpdate struct {
	Name   string
	Status string
}

type WorkHandler interface {
	Handle(msg *workflow.Message) <-chan StatusInfo
}

func ReturnJobStatus(statusInfo chan StatusInfo, scanId string, workState WorkState, stepUpdates []StepUpdate, errorMsg string, err error) {
	statusInfo <- StatusInfo{
		ScanID:   scanId,
		State:    workState,
		Progress: stepUpdates,
		Err:      err,
		ErrMsg:   errorMsg,
	}
}
