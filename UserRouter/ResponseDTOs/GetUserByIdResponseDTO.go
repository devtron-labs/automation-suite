package ResponseDTOs

import "automation-suite/testUtils"

type GetUserByIdResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id          int          `json:"id"`
		EmailId     string       `json:"email_id"`
		RoleFilters []RoleFilter `json:"roleFilters"`
		Groups      []string     `json:"groups"`
		SuperAdmin  bool         `json:"superAdmin"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
