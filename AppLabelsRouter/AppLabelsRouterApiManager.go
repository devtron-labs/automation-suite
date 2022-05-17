package AppLabelsRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
	"time"
)

type AppMetaInfoResponseDto struct {
	Code   int            `json:"code"`
	Status string         `json:"status"`
	Result AppMetaInfoDto `json:"result"`
	Errors []Error        `json:"errors"`
}

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
type Error struct {
	Code            string `json:"code"`
	InternalMessage string `json:"internalMessage"`
	UserMessage     string `json:"userMessage"`
}

type StructAppLabelsRouter struct {
	appMetaInfoResponseDto AppMetaInfoResponseDto
}

func HitGetAppMetaInfoByIdApi(appId string, authToken string) AppMetaInfoResponseDto {
	resp, err := Base.MakeApiCall(GetAppMetaInfoByIdApiUrl+"/"+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAppMetaInfoByIdApi)
	structAppLabelsRouter := StructAppLabelsRouter{}
	appLabelRepoRouter := structAppLabelsRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppMetaInfoByIdApi)
	return appLabelRepoRouter.appMetaInfoResponseDto
}

func (structAppLabelsRouter StructAppLabelsRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppLabelsRouter {
	switch apiName {
	case GetAppMetaInfoByIdApi:
		json.Unmarshal(response, &structAppLabelsRouter.appMetaInfoResponseDto)
	}
	return structAppLabelsRouter
}

type EnvironmentConfigAppLabelsRouter struct {
	AppIdForAppLabelRouter string `env:"APP_ID_APP_LABELS_ROUTER" envDefault:"193"`
}

func GetEnvironmentConfigForAppLabelsRouter() (*EnvironmentConfigAppLabelsRouter, error) {
	cfg := &EnvironmentConfigAppLabelsRouter{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

type AppLabelsSuite struct {
	suite.Suite
	authToken string
}

func (suite *AppLabelsSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
