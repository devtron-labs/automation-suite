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
	discoverAppApiResponse                  ResponseDTOs.DiscoverAppApiResponse
	deploymentOfInstalledAppResponseDTO     ResponseDTOs.DeploymentOfInstalledAppResponseDTO
	helmAppVersionsDTO                      ResponseDTOs.HelmAppVersionsDTO
	helmAppViaAppVersionIDResponseDTO       ResponseDTOs.HelmAppViaVersionIdResponseDTO
	helmEnvAutocompleteResponseDTO          ResponseDTOs.HelmEnvAutocompleteResponseDTO
	templateValuesViaReferenceIdResponseDTO ResponseDTOs.TemplateValuesResponseDTO
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

func GetAppVersionsAutocomplete(chartStorId string, authToken string) ResponseDTOs.HelmAppVersionsDTO {
	resp, err := Base.MakeApiCall(GetVersionsAutocompleteApiUrl+chartStorId+"/version/autocomplete", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetVersionsAutocompleteApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetVersionsAutocompleteApi)
	return appStoreDiscoverRouter.helmAppVersionsDTO
}

func DiscoverAppViaAppStoreApplicationVersionId(appStoreApplicationVersionId string, authToken string) ResponseDTOs.HelmAppViaVersionIdResponseDTO {
	resp, err := Base.MakeApiCall(GetVersionsAutocompleteApiUrl+appStoreApplicationVersionId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, DiscoverAppViaAppstoreApplicationVersionIdApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), DiscoverAppViaAppstoreApplicationVersionIdApi)
	return appStoreDiscoverRouter.helmAppViaAppVersionIDResponseDTO
}

func HitHelmEnvAutocompleteApi(authToken string) ResponseDTOs.HelmEnvAutocompleteResponseDTO {
	resp, err := Base.MakeApiCall(GetHelmEnvAutocompleteApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetHelmEnvAutocompleteApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetHelmEnvAutocompleteApi)
	return appStoreDiscoverRouter.helmEnvAutocompleteResponseDTO
}

func HitGetTemplateValuesViaReferenceIdApi(queryParams map[string]string, authToken string) ResponseDTOs.TemplateValuesResponseDTO {
	resp, err := Base.MakeApiCall(GetTemplateValuesViaReferenceIdApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetTemplateValuesViaReferenceIdApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetTemplateValuesViaReferenceIdApi)
	return appStoreDiscoverRouter.templateValuesViaReferenceIdResponseDTO
}

func (structAppStoreDiscoverRouter StructAppStoreDiscoverRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppStoreDiscoverRouter {
	switch apiName {
	case DiscoverAppApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.discoverAppApiResponse)
	case GetDeploymentOfInstalledAppApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.deploymentOfInstalledAppResponseDTO)
	case GetVersionsAutocompleteApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.helmAppVersionsDTO)
	case DiscoverAppViaAppstoreApplicationVersionIdApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.helmAppViaAppVersionIDResponseDTO)
	case GetHelmEnvAutocompleteApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.helmEnvAutocompleteResponseDTO)
	case GetTemplateValuesViaReferenceIdApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.templateValuesViaReferenceIdResponseDTO)

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
