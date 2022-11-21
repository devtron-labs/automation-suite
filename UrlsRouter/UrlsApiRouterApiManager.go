package UrlsRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"os"
)

type UrlsDto struct {
	Code   int            `json:"code"`
	Status string         `json:"status"`
	Result []UrlsResponse `json:"result"`
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

type UrlsResponse struct {
	Kind     string   `json:"kind"`
	Name     string   `json:"name"`
	PointsTo string   `json:"pointsTo"`
	Urls     []string `json:"urls"`
}

func (suite *UrlsTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}

func HitGetUrls(queryParams map[string]string, authToken string, url string) UrlsDto {
	resp, err := Base.MakeApiCall(url, http.MethodGet, "", queryParams, authToken)
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

func GetTestExpectedUrlsData() []UrlsResponse {
	testFile, err := os.Open("../testdata/UrlsRouter/urlsTestData.json")
	if err != nil {
		fmt.Println(err)
		return []UrlsResponse{}
	}
	defer testFile.Close()
	data, _ := ioutil.ReadAll(testFile)
	res := make([]UrlsResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
		return []UrlsResponse{}
	}
	return res
}
