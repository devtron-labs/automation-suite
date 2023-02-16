package ResponseDTOs

import "automation-suite/EnvironmentRouter/RequestDTOs"

type CreateEnvironmentResponseDTO struct {
	Code   int                                     `json:"code"`
	Status string                                  `json:"status"`
	Result RequestDTOs.CreateEnvironmentRequestDTO `json:"result"`
}
