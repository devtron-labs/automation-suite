package ResponseDTOs

import Base "automation-suite/testUtils"

type FetchAllStageStatusResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Stage     int    `json:"stage"`
		StageName string `json:"stageName"`
		Status    bool   `json:"status"`
		Required  bool   `json:"required"`
	} `json:"result"`
	Errors []Base.Errors `json:"errors"`
}
