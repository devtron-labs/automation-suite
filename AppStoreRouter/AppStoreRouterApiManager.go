package AppStoreRouter

import (
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
	installedAppDetailsResponseDTO      ResponseDTOs.InstalledAppDetailsResponseDTO
}

func HitGetApplicationValuesList(appStoreId string, authToken string) ResponseDTOs.GetApplicationValuesListResponseDTO {
	resp, err := Base.MakeApiCall(GetApplicationValuesListApiUrl+appStoreId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetApplicationValuesListApi)
	structAppStoreRouter := StructAppStoreRouter{}
	appStoreRouter := structAppStoreRouter.UnmarshalGivenResponseBody(resp.Body(), GetApplicationValuesListApi)
	return appStoreRouter.getApplicationValuesListResponseDto
}

func HitGetInstalledAppDetailsApi(queryParams map[string]string, authToken string) ResponseDTOs.InstalledAppDetailsResponseDTO {
	resp, err := Base.MakeApiCall(GetInstalledAppDetailsApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetInstalledAppDetailsApi)
	structAppStoreDiscoverRouter := StructAppStoreRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetInstalledAppDetailsApi)
	return appStoreDiscoverRouter.installedAppDetailsResponseDTO
}

func (structAppStoreRouter StructAppStoreRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppStoreRouter {
	switch apiName {
	case GetApplicationValuesListApi:
		json.Unmarshal(response, &structAppStoreRouter.getApplicationValuesListResponseDto)
	case GetInstalledAppDetailsApi:
		json.Unmarshal(response, &structAppStoreRouter.installedAppDetailsResponseDTO)
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
