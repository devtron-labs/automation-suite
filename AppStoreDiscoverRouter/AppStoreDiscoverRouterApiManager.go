package AppStoreDiscoverRouter

import (
	"automation-suite/AppStoreDiscoverRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
	"time"
)

type StructAppStoreDiscoverRouter struct {
	discoverAppApiResponse              ResponseDTOs.DiscoverAppApiResponse
	deploymentOfInstalledAppResponseDTO ResponseDTOs.DeploymentOfInstalledAppResponseDTO
}

func HitDiscoverAppApi(queryParams map[string]string, authToken string) ResponseDTOs.DiscoverAppApiResponse {
	resp, err := Base.MakeApiCall(DiscoverAppApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, DiscoverAppApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), DiscoverAppApi)
	return appStoreDiscoverRouter.discoverAppApiResponse
}

func HitGetDeploymentOfInstalledAppApi(appId string, authToken string) ResponseDTOs.DeploymentOfInstalledAppResponseDTO {
	resp, err := Base.MakeApiCall(GetDeploymentOfInstalledAppApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetDeploymentOfInstalledAppApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetDeploymentOfInstalledAppApi)
	return appStoreDiscoverRouter.deploymentOfInstalledAppResponseDTO
}

func (structAppStoreDiscoverRouter StructAppStoreDiscoverRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppStoreDiscoverRouter {
	switch apiName {
	case DiscoverAppApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.discoverAppApiResponse)
	case GetDeploymentOfInstalledAppApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.deploymentOfInstalledAppResponseDTO)
	}
	return structAppStoreDiscoverRouter
}

func PollForGettingHelmAppData(queryParams map[string]string, authToken string) bool {
	count := 0
	for {
		appData := HitDiscoverAppApi(queryParams, authToken)
		helmAppData := appData.Result
		time.Sleep(1 * time.Second)
		count = count + 1
		if helmAppData != nil || count >= 25 {
			break
		}
	}
	return true
}

type AppStoreDiscoverTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *AppStoreDiscoverTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
