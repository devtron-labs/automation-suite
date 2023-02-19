package ResponseDTOs

import "automation-suite/testUtils"

type DeleteClusterResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Result string             `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
