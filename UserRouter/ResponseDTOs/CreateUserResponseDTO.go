package ResponseDTOs

import "automation-suite/testUtils"

type CreateUserResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Errors []testUtils.Errors `json:"errors"`
	Result []struct {
		Id          int          `json:"id"`
		EmailId     string       `json:"email_id"`
		RoleFilters []RoleFilter `json:"roleFilters"`
		Groups      []string     `json:"groups"`
		SuperAdmin  bool         `json:"superAdmin"`
	} `json:"result"`
}

type RoleFilter struct {
	Entity      string `json:"entity"`
	Team        string `json:"team"`
	EntityName  string `json:"entityName"`
	Environment string `json:"environment"`
	Action      string `json:"action"`
	AccessType  string `json:"accessType"`
}
