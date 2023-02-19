package ResponseDTOs

import "automation-suite/SSOLoginRouter/RequestDTOs"

type GetSSODetailsResponseDTO struct {
	Code                       int                                     `json:"code"`
	Status                     string                                  `json:"status"`
	CreateSSODetailsRequestDto *RequestDTOs.CreateSSODetailsRequestDTO `json:"result"`
}
