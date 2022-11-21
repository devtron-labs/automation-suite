package ResponseDTOs

import Base "automation-suite/testUtils"

type FetchOtherEnvResponseDTO struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Result []OtherEnvResult `json:"result"`
	Errors []Base.Errors    `json:"errors"`
}

type OtherEnvResult struct {
	EnvironmentId   int    `json:"environmentId"`
	EnvironmentName string `json:"environmentName"`
	AppMetrics      bool   `json:"appMetrics"`
	InfraMetrics    bool   `json:"infraMetrics"`
	Prod            bool   `json:"prod"`
}
