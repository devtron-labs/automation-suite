package RequestDTOs

import "automation-suite/UserRouter/ResponseDTOs"

type UserRole struct {
	Id      int32  `json:"id" validate:"number"`
	EmailId string `json:"email_id" validate:"email"`
	Role    string `json:"role"`
}

type UserInfo struct {
	Id          int32                     `json:"id" validate:"number"`
	EmailId     string                    `json:"email_id" validate:"required"`
	Roles       []string                  `json:"roles,omitempty"`
	AccessToken string                    `json:"access_token,omitempty"`
	Exist       bool                      `json:"-"`
	UserId      int32                     `json:"-"` // created or modified user id
	RoleFilters []ResponseDTOs.RoleFilter `json:"roleFilters"`
	Status      string                    `json:"status,omitempty"`
	Groups      []string                  `json:"groups"`
	SuperAdmin  bool                      `json:"superAdmin,notnull"`
}

type RoleGroup struct {
	Id          int32                     `json:"id" validate:"number"`
	Name        string                    `json:"name,omitempty"`
	Description string                    `json:"description,omitempty"`
	RoleFilters []ResponseDTOs.RoleFilter `json:"roleFilters"`
	Status      string                    `json:"status,omitempty"`
	UserId      int32                     `json:"-"` // created or modified user id
}

type Role struct {
	Id   int    `json:"id" validate:"number"`
	Role string `json:"role" validate:"required"`
}

type RoleData struct {
	Id          int    `json:"id" validate:"number"`
	Role        string `json:"role" validate:"required"`
	Entity      string `json:"entity"`
	Team        string `json:"team"`
	EntityName  string `json:"entityName"`
	Environment string `json:"environment"`
	Action      string `json:"action"`
	AccessType  string `json:"accessType"`
}
