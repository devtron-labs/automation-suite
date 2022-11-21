package UrlsRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type UrlsDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Kind     string   `json:"kind"`
		Name     string   `json:"name"`
		PointsTo string   `json:"pointsTo"`
		Urls     []string `json:"urls"`
	} `json:"result"`
}

type DevtronAppConfig struct {
	AppId string `env:"APP_ID" envDefault:"1"`
	EnvId string `env:"ENV_ID" envDefault:"1"`
}

type InstalledAppConfig struct {
	InstalledAppId string `env:"INSTALLED_APP_ID" envDefault:"1"`
	EnvId          string `env:"ENV_ID" envDefault:"1"`
}

type StructUrlsRouter struct {
	urlsResponse UrlsDto
}

type UrlsTestSuite struct {
	suite.Suite
	authToken string
}

func HitGetUrls(queryParams map[string]string, authToken string) UrlsDto {
	resp, err := Base.MakeApiCall(GetUrlsUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetUrls)

	structUrlsRouter := StructUrlsRouter{}
	urlsRouter := structUrlsRouter.UnmarshalGivenResponseBody(resp.Body())
	return urlsRouter.urlsResponse
}
func (r StructUrlsRouter) UnmarshalGivenResponseBody(response []byte) StructUrlsRouter {
	json.Unmarshal(response, &r.urlsResponse)
	return r
}
func GetEnvironmentConfigForDevtronApp() (*DevtronAppConfig, error) {
	cfg := &DevtronAppConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}
func GetEnvironmentConfigForInstalledApp() (*InstalledAppConfig, error) {
	cfg := &InstalledAppConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}
