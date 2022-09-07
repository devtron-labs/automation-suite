package ResponseDTOs

import "automation-suite/testUtils"

type RefreshMaterialsResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Message       string `json:"message"`
		ErrorMsg      string `json:"errorMsg"`
		LastFetchTime string `json:"lastFetchTime"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
