package ResponseDTOs

import "automation-suite/testUtils"

type TriggerCiPipelineResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		ApiResponse string `json:"apiResponse"`
		AuthStatus  string `json:"authStatus"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
