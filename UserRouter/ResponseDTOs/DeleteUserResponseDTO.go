package ResponseDTOs

import "automation-suite/testUtils"

type DeleteUserResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Result bool               `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
