package AppStoreDeploymentRouter

import (
	"automation-suite/AppStoreDeploymentRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructAppStoreDeploymentRouter struct {
	getApplicationValuesListResponseDto ResponseDTOs.GetApplicationValuesListResponseDTO
	installAppResponseDto               ResponseDTOs.InstallAppResponseDTO
	installedAppVersionResponseDTO      ResponseDTOs.InstalledAppVersionResponseDTO
}

func HitInstallAppApi(requestPayload string, authToken string) ResponseDTOs.InstallAppResponseDTO {
	resp, err := Base.MakeApiCall(InstallAppApiUrl, http.MethodPost, requestPayload, nil, authToken)
	Base.HandleError(err, InstallAppApi)
	structAppStoreDepRouter := StructAppStoreDeploymentRouter{}
	appStoreDepRouter := structAppStoreDepRouter.UnmarshalGivenResponseBody(resp.Body(), InstallAppApi)
	return appStoreDepRouter.installAppResponseDto
}

func HitDeleteInstalledAppApi(id string, authToken string) ResponseDTOs.InstallAppResponseDTO {
	resp, err := Base.MakeApiCall(DeleteInstalledAppApiUrl+id, http.MethodDelete, " ", nil, authToken)
	Base.HandleError(err, DeleteInstalledAppApi)
	structAppStoreRouter := StructAppStoreDeploymentRouter{}
	appStoreRouter := structAppStoreRouter.UnmarshalGivenResponseBody(resp.Body(), InstallAppApi)
	return appStoreRouter.installAppResponseDto
}

func HitGetInstalledAppVersionApi(appId string, authToken string) ResponseDTOs.InstalledAppVersionResponseDTO {
	resp, err := Base.MakeApiCall(GetInstalledAppVersionApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetInstalledAppVersionApi)
	structAppStoreDeploymentRouter := StructAppStoreDeploymentRouter{}
	appStoreDeploymentRouter := structAppStoreDeploymentRouter.UnmarshalGivenResponseBody(resp.Body(), GetInstalledAppVersionApi)
	return appStoreDeploymentRouter.installedAppVersionResponseDTO
}

func (structAppStoreDeploymentRouter StructAppStoreDeploymentRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppStoreDeploymentRouter {
	switch apiName {
	case InstallAppApi:
		json.Unmarshal(response, &structAppStoreDeploymentRouter.installAppResponseDto)
	case GetInstalledAppVersionApi:
		json.Unmarshal(response, &structAppStoreDeploymentRouter.installedAppVersionResponseDTO)
	}
	return structAppStoreDeploymentRouter
}

type AppStoreDeploymentTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *AppStoreDeploymentTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
