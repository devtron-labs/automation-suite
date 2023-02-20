package ResponseDTOs

import (
	"automation-suite/testUtils"
)

type CreateRoleGroupResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id          int          `json:"id"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		RoleFilters []RoleFilter `json:"roleFilters"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
