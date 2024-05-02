package ContainersEngine

type NewScanRequest struct {
	ScanId        string `json:"ScanId"`
	ProjectId     string `json:"ProjectId"`
	TenantId      string `json:"TenantId"`
	Url           string `json:"Url"`
	CorrelationId string `json:"CorrelationId"`
	ScanType      string `json:"ScanType"`
}
