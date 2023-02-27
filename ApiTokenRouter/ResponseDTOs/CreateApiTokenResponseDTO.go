package ResponseDTOs

import "automation-suite/testUtils"

type CreateApiTokenResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id             int    `json:"id"`
		Success        bool   `json:"success"`
		Token          string `json:"token"`
		UserId         int    `json:"userId"`
		UserIdentifier string `json:"userIdentifier"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
