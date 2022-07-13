package ResponseDTOs

import "automation-suite/AppStoreRouter/RequestDTOs"

type InstallAppResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	Result *RequestDTOs.InstallAppRequestDTO `json:"result"`
}
