package AttributesRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type AttributesDto struct {
	Id     int    `json:"id"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
	Active bool   `json:"active"`
	UserId int32  `json:"-"`
}

type GetAttributesRespDto struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result AttributesDto `json:"result"`
}

type StructAttributesRouter struct {
	getAttributesRespDto GetAttributesRespDto
}

func HitGetAttributesApi(queryParams map[string]string, authToken string) GetAttributesRespDto {
	resp, err := Base.MakeApiCall(GetAttributesApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetAttributesApi)
	structAttributesRouter := StructAttributesRouter{}
	chartRepoRouter := structAttributesRouter.UnmarshalGivenResponseBody(resp.Body(), GetAttributesApi)
	return chartRepoRouter.getAttributesRespDto
}

func (structAttributesRouter StructAttributesRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAttributesRouter {
	switch apiName {
	case GetAttributesApi:
		json.Unmarshal(response, &structAttributesRouter.getAttributesRespDto)
	}
	return structAttributesRouter
}

type EnvironmentConfigAttributesRouter struct {
	ValueAttribute string `env:"VALUE_ATTRIBUTE" envDefault:"https://staging.devtron.info"`
}

func GetEnvironmentConfigForHelmApp() (*EnvironmentConfigAttributesRouter, error) {
	cfg := &EnvironmentConfigAttributesRouter{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

type AttributeRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *AttributeRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
