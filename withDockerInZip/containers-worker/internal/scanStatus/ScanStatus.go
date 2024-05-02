package scan_status_service

import "google.golang.org/protobuf/types/known/timestamppb"

type ScanStatus struct {
	Id                     string                 `json:"ScanId"`
	Status                 string                 `json:"Status"`
	StepStatus             []ScanStepStatus       `json:"StepStatus"`
	WorkflowCorrelationKey string                 `json:"WorkflowCorrelationKey"`
	CorrelationId          string                 `json:"CorrelationId"`
	UpdateTime             *timestamppb.Timestamp `json:"UpdateTime"`
	Seen                   bool                   `json:"Seen"`
	Error                  string                 `json:"Error"`
}

type ScanStepStatus struct {
	Status     string                 `json:"Status"`
	StepName   string                 `json:"StepName"`
	UpdateTime *timestamppb.Timestamp `json:"UpdateTime"`
	Seen       bool                   `json:"Seen"`
}

// containers engine status
const (
	ContainersScanStatusKeyPrefix = "ContainersScanStatus"

	ContainersScanStatusQueued    = "Queued"
	ContainersScanStatusWorking   = "Working"
	ContainersScanStatusCompleted = "Success"
	ContainersScanStatusFailed    = "Failed"

	ContainersScanStatusCanceled = "Canceled"
	ContainersScanStatusTimeOut  = "Timeout"
)
