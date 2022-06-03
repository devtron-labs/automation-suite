package AppStoreRouter

import (
	"automation-suite/AppStoreRouter/RequestDTOs"
	"automation-suite/AppStoreRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructAppStoreRouter struct {
	getApplicationValuesListResponseDto ResponseDTOs.GetApplicationValuesListResponseDTO
	installAppResponseDto               ResponseDTOs.InstallAppResponseDTO
	installAppRequestDto                RequestDTOs.InstallAppRequestDTO
}

func HitGetApplicationValuesList(appStoreId string, authToken string) ResponseDTOs.GetApplicationValuesListResponseDTO {
	resp, err := Base.MakeApiCall(GetApplicationValuesListApiUrl+appStoreId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetApplicationValuesListApi)
	structAppStoreRouter := StructAppStoreRouter{}
	appStoreRouter := structAppStoreRouter.UnmarshalGivenResponseBody(resp.Body(), GetApplicationValuesListApi)
	return appStoreRouter.getApplicationValuesListResponseDto
}

func HitInstallAppApi(requestPayload string, authToken string) ResponseDTOs.InstallAppResponseDTO {
	resp, err := Base.MakeApiCall(InstallAppApiUrl, http.MethodPost, requestPayload, nil, authToken)
	Base.HandleError(err, InstallAppApi)
	structAppStoreRouter := StructAppStoreRouter{}
	appStoreRouter := structAppStoreRouter.UnmarshalGivenResponseBody(resp.Body(), InstallAppApi)
	return appStoreRouter.installAppResponseDto
}

func HitDeleteInstalledAppApi(id string, authToken string) ResponseDTOs.InstallAppResponseDTO {
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
