package ResponseDTOs

import "automation-suite/testUtils"

type RollbackApplicationApiResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Success bool `json:"success"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
