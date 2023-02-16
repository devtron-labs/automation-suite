package ResponseDTOs

import "automation-suite/testUtils"

type CreateRoleGroupResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		RoleFilters []struct {
			Entity      string `json:"entity"`
			Team        string `json:"team"`
			EntityName  string `json:"entityName"`
			Environment string `json:"environment"`
			Action      string `json:"action"`
			AccessType  string `json:"accessType"`
		} `json:"roleFilters"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
