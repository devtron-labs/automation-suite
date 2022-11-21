package ResponseDTOs

import (
	"automation-suite/dockerRegRouter/RequestDTOs"
	Base "automation-suite/testUtils"
)

type SaveDockerRegistryResponseDto struct {
	Code   int                                      `json:"code"`
	Status string                                   `json:"status"`
	Result RequestDTOs.SaveDockerRegistryRequestDto `json:"result"`
	Errors []Base.Errors                            `json:"errors"`
}

type DeleteDockerRegistryResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}
