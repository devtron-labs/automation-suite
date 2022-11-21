package ResponseDTOs

import "automation-suite/testUtils"

type SaveTeamResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
