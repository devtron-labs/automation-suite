package AppStoreRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type GetApplicationValuesListResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Values []struct {
			Values []struct {
				Id                int    `json:"id"`
				Name              string `json:"name"`
				ChartVersion      string `json:"chartVersion"`
				AppStoreVersionId int    `json:"appStoreVersionId,omitempty"`
				EnvironmentName   string `json:"environmentName,omitempty"`
			} `json:"values"`
			Kind string `json:"kind"`
		} `json:"values"`
	} `json:"result"`
}

type InstallAppResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	InstallAppRequestDto *InstallAppRequestDto `json:"result"`
}

type InstallAppRequestDto struct {
	Id                    int    `json:"id"`
	AppId                 int    `json:"appId"`
	AppName               string `json:"appName"`
	TeamId                int    `json:"teamId"`
	EnvironmentId         int    `json:"environmentId"`
	InstalledAppId        int    `json:"installedAppId"`
	InstalledAppVersionId int    `json:"installedAppVersionId"`
	AppStoreVersion       int    `json:"appStoreVersion"`
	ValuesOverrideYaml    string `json:"valuesOverrideYaml"`
	ReferenceValueId      int    `json:"referenceValueId"`
	ReferenceValueKind    string `json:"referenceValueKind"`
	AppStoreId            int    `json:"appStoreId"`
	AppStoreName          string `json:"appStoreName"`
	Deprecated            bool   `json:"deprecated"`
	ClusterId             int    `json:"clusterId"`
	Namespace             string `json:"namespace"`
	AppOfferingMode       string `json:"appOfferingMode"`
	GitOpsRepoName        string `json:"gitOpsRepoName"`
	GitOpsPath            string `json:"gitOpsPath"`
	GitHash               string `json:"gitHash"`
}

type StructAppStoreRouter struct {
	getApplicationValuesListResponseDto GetApplicationValuesListResponseDto
	installAppResponseDto               InstallAppResponseDto
	installAppRequestDto                InstallAppRequestDto
}

func HitGetApplicationValuesList(appStoreId string, authToken string) GetApplicationValuesListResponseDto {
	resp, err := Base.MakeApiCall(GetApplicationValuesListApiUrl+appStoreId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetApplicationValuesListApi)
	structAppStoreRouter := StructAppStoreRouter{}
	appStoreRouter := structAppStoreRouter.UnmarshalGivenResponseBody(resp.Body(), GetApplicationValuesListApi)
	return appStoreRouter.getApplicationValuesListResponseDto
}

func HitInstallAppApi(requestPayload string, authToken string) InstallAppResponseDto {
	resp, err := Base.MakeApiCall(InstallAppApiUrl, http.MethodPost, requestPayload, nil, authToken)
	Base.HandleError(err, InstallAppApi)
	structAppStoreRouter := StructAppStoreRouter{}
	appStoreRouter := structAppStoreRouter.UnmarshalGivenResponseBody(resp.Body(), InstallAppApi)
	return appStoreRouter.installAppResponseDto
}

func HitDeleteInstalledAppApi(id string, authToken string) InstallAppResponseDto {
	resp, err := Base.MakeApiCall(DeleteInstalledAppApiUrl+id, http.MethodDelete, " ", nil, authToken)
	Base.HandleError(err, DeleteInstalledAppApi)
	structAppStoreRouter := StructAppStoreRouter{}
	appStoreRouter := structAppStoreRouter.UnmarshalGivenResponseBody(resp.Body(), InstallAppApi)
	return appStoreRouter.installAppResponseDto
}

func (structAppStoreRouter StructAppStoreRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppStoreRouter {
	switch apiName {
	case GetApplicationValuesListApi:
		json.Unmarshal(response, &structAppStoreRouter.getApplicationValuesListResponseDto)
	case InstallAppApi:
		json.Unmarshal(response, &structAppStoreRouter.installAppResponseDto)
	}
	return structAppStoreRouter
}

type EnvironmentConfigAppStoreRouter struct {
	AppStoreId string `env:"APP_STORE_ID" envDefault:"2514"`
}

func GetEnvironmentConfigForAppStoreRouter() (*EnvironmentConfigAppStoreRouter, error) {
	cfg := &EnvironmentConfigAppStoreRouter{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

type AppStoreTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *AppStoreTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
