package ResponseDTOs

import "automation-suite/testUtils"

type DeleteChartRepoResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Result string             `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
