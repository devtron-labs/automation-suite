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
	ValueAttribute string `env:"VALUE_ATTRIBUTE" envDefault:"https://staging.devtron.info"`
}

func GetEnvironmentConfigForHelmApp() (*EnvironmentConfigAppLabelsRouter, error) {
	cfg := &EnvironmentConfigAppLabelsRouter{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

type AppLabelRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *AppLabelRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
