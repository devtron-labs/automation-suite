package ResponseDTOs

import (
	"automation-suite/testUtils"
	"time"
)

type Label struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type AppMetaInfoDto struct {
	AppId       int       `json:"appId"`
	AppName     string    `json:"appName"`
	ProjectId   int       `json:"projectId"`
	ProjectName string    `json:"projectName"`
	CreatedBy   string    `json:"createdBy"`
	CreatedOn   time.Time `json:"createdOn"`
	Active      bool      `json:"active,notnull"`
	Labels      []*Label  `json:"labels"`
	UserId      int32     `json:"-"`
}

type AppMetaInfoResponseDto struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Result AppMetaInfoDto     `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
