package AppStoreDiscoverRouter

import (
	"automation-suite/AppStoreDiscoverRouter/RequestDTOs"
	"automation-suite/AppStoreDiscoverRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"strings"
	"time"
)

type StructAppStoreDiscoverRouter struct {
	discoverAppApiResponse                  ResponseDTOs.DiscoverAppApiResponse
	deploymentOfInstalledAppResponseDTO     ResponseDTOs.DeploymentOfInstalledAppResponseDTO
	helmAppVersionsDTO                      ResponseDTOs.HelmAppVersionsDTO
	helmAppViaAppVersionIDResponseDTO       ResponseDTOs.HelmAppViaVersionIdResponseDTO
	helmEnvAutocompleteResponseDTO          ResponseDTOs.HelmEnvAutocompleteResponseDTO
	templateValuesViaReferenceIdResponseDTO ResponseDTOs.TemplateValuesResponseDTO
	getApplicationValuesListResponseDTO     ResponseDTOs.GetApplicationValuesListResponseDTO
	saveTemplateValuesResponseDTO           ResponseDTOs.SaveTemplateValuesResponseDTO
	saveTemplateValuesRequestDTO            RequestDTOs.SaveTemplateValuesRequestDTO
	deleteTemplateValuesResponseDTO         ResponseDTOs.DeleteTemplateValuesResponseDTO
	appStoreChartByNameResponseDTO          ResponseDTOs.AppStoreChartByNameResponseDTO
	installAppResponseDto                   ResponseDTOs.InstallAppResponseDTO
	installedAppVersionResponseDTO          ResponseDTOs.InstalledAppVersionResponseDTO
	installedAppDetailsResponseDTO          ResponseDTOs.InstalledAppDetailsResponseDTO
	checkAppExistsResponseDTO               ResponseDTOs.CheckAppExistsResponseDTO
	checkAppExistsRequestDTO                RequestDTOs.CheckAppExistsRequestDTO
	getAllInstalledAppResponseDTO           ResponseDTOs.GetAllInstalledAppResponseDTO
}

func HitDiscoverAppApi(queryParams map[string]string, authToken string) ResponseDTOs.DiscoverAppApiResponse {
	resp, err := Base.MakeApiCall(DiscoverAppApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, DiscoverAppApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), DiscoverAppApi)
	return appStoreDiscoverRouter.discoverAppApiResponse
}

func GetInstalledAppsByAppStoreId(appStoreId string, authToken string) ResponseDTOs.DeploymentOfInstalledAppResponseDTO {
	resp, err := Base.MakeApiCall(GetInstalledAppsByAppStoreIdApiUrl+appStoreId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetInstalledAppsByAppStoreIdApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetInstalledAppsByAppStoreIdApi)
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

func HitGetApplicationValuesListApi(appId string, authToken string) ResponseDTOs.GetApplicationValuesListResponseDTO {
	resp, err := Base.MakeApiCall(GetApplicationValuesListApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetApplicationValuesListApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetApplicationValuesListApi)
	return appStoreDiscoverRouter.getApplicationValuesListResponseDTO
}

func HitSaveTemplateValuesApi(requestPayload string, authToken string) ResponseDTOs.SaveTemplateValuesResponseDTO {
	resp, err := Base.MakeApiCall(SaveTemplateValuesApiUrl, http.MethodPost, requestPayload, nil, authToken)
	Base.HandleError(err, SaveTemplateValuesApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), SaveTemplateValuesApi)
	return appStoreDiscoverRouter.saveTemplateValuesResponseDTO
}

func HitDeleteTemplateValuesApi(Id string, authToken string) ResponseDTOs.DeleteTemplateValuesResponseDTO {
	resp, err := Base.MakeApiCall(SaveTemplateValuesApiUrl+"/"+Id, http.MethodDelete, "", nil, authToken)
	Base.HandleError(err, DeleteTemplateValuesApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteTemplateValuesApi)
	return appStoreDiscoverRouter.deleteTemplateValuesResponseDTO
}

func getPayloadForSaveTemplateValues(name string, values string, appStoreVersionId int) RequestDTOs.SaveTemplateValuesRequestDTO {
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	structAppStoreDiscoverRouter.saveTemplateValuesRequestDTO.Name = name
	structAppStoreDiscoverRouter.saveTemplateValuesRequestDTO.AppStoreVersionId = appStoreVersionId
	structAppStoreDiscoverRouter.saveTemplateValuesRequestDTO.Values = values
	return structAppStoreDiscoverRouter.saveTemplateValuesRequestDTO
}

func HitSearchAppStoreChartByNameApi(queryParams map[string]string, authToken string) ResponseDTOs.AppStoreChartByNameResponseDTO {
	resp, err := Base.MakeApiCall(SearchAppStoreChartByNameApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, SearchAppStoreChartByNameApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), SearchAppStoreChartByNameApi)
	return appStoreDiscoverRouter.appStoreChartByNameResponseDTO
}

func HitInstallAppApi(requestPayload string, authToken string) ResponseDTOs.InstallAppResponseDTO {
	resp, err := Base.MakeApiCall(InstallAppApiUrl, http.MethodPost, requestPayload, nil, authToken)
	Base.HandleError(err, InstallAppApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), InstallAppApi)
	return appStoreDiscoverRouter.installAppResponseDto
}

func HitDeleteInstalledAppApi(id string, authToken string) ResponseDTOs.InstallAppResponseDTO {
	resp, err := Base.MakeApiCall(DeleteInstalledAppApiUrl+id, http.MethodDelete, " ", nil, authToken)
	Base.HandleError(err, DeleteInstalledAppApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), InstallAppApi)
	return appStoreDiscoverRouter.installAppResponseDto
}

func HitGetInstalledAppVersionApi(appId string, authToken string) ResponseDTOs.InstalledAppVersionResponseDTO {
	resp, err := Base.MakeApiCall(GetInstalledAppVersionApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetInstalledAppVersionApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetInstalledAppVersionApi)
	return appStoreDiscoverRouter.installedAppVersionResponseDTO
}

func GetRequestDtoForInstallApp(ReferenceValueId int, AppStoreVersion int, ValuesOverride interface{}, ValuesOverrideYaml string) RequestDTOs.InstallAppRequestDTO {
	installAppRequestDTO1 := RequestDTOs.InstallAppRequestDTO{}
	installAppRequestDTO1.TeamId = 1
	installAppRequestDTO1.ReferenceValueId = ReferenceValueId
	installAppRequestDTO1.ReferenceValueKind = "DEFAULT"
	installAppRequestDTO1.EnvironmentId = 1
	installAppRequestDTO1.Namespace = "devtron-demo"
	installAppRequestDTO1.AppStoreVersion = AppStoreVersion
	installAppRequestDTO1.ValuesOverride = ValuesOverride
	installAppRequestDTO1.ValuesOverrideYaml = ValuesOverrideYaml
	appName := "automation" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
	installAppRequestDTO1.AppName = appName
	log.Println("=== name for the helm app is ===", appName)
	return installAppRequestDTO1
}

func HitGetApplicationValuesList(appStoreId string, authToken string) ResponseDTOs.GetApplicationValuesListResponseDTO {
	resp, err := Base.MakeApiCall(GetApplicationValuesListApiUrl+appStoreId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetApplicationValuesListApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetApplicationValuesListApi)
	return appStoreDiscoverRouter.getApplicationValuesListResponseDTO
}

func HitGetInstalledAppDetailsApi(queryParams map[string]string, authToken string) ResponseDTOs.InstalledAppDetailsResponseDTO {
	resp, err := Base.MakeApiCall(GetInstalledAppDetailsApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetInstalledAppDetailsApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetInstalledAppDetailsApi)
	return appStoreDiscoverRouter.installedAppDetailsResponseDTO
}

func HitCheckAppExistsOrNot(payload string, authToken string) ResponseDTOs.CheckAppExistsResponseDTO {
	resp, err := Base.MakeApiCall(CheckAppExistsApiUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, CheckAppExistsApi)
	structAppStoreRouter := StructAppStoreDiscoverRouter{}
	appStoreRouter := structAppStoreRouter.UnmarshalGivenResponseBody(resp.Body(), CheckAppExistsApi)
	return appStoreRouter.checkAppExistsResponseDTO
}

func getCheckAppExistsApi(AppNames []string) []RequestDTOs.CheckAppExistsRequestDTO {
	var listOfCheckAppExistsRequestDTOs []RequestDTOs.CheckAppExistsRequestDTO
	for _, name := range AppNames {
		CheckAppExistsRequestDTO := RequestDTOs.CheckAppExistsRequestDTO{}
		CheckAppExistsRequestDTO.Name = name
		listOfCheckAppExistsRequestDTOs = append(listOfCheckAppExistsRequestDTOs, CheckAppExistsRequestDTO)
	}
	return listOfCheckAppExistsRequestDTOs
}

func HitApiGetAllInstalledApps(authToken string) ResponseDTOs.GetAllInstalledAppResponseDTO {
	resp, err := Base.MakeApiCall(GetAllInstalledAppApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAllInstalledAppApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), GetAllInstalledAppApi)
	return appStoreDiscoverRouter.getAllInstalledAppResponseDTO
}

func (structAppStoreDiscoverRouter StructAppStoreDiscoverRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppStoreDiscoverRouter {
	switch apiName {
	case DiscoverAppApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.discoverAppApiResponse)
	case GetInstalledAppsByAppStoreIdApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.deploymentOfInstalledAppResponseDTO)
	case GetVersionsAutocompleteApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.helmAppVersionsDTO)
	case DiscoverAppViaAppstoreApplicationVersionIdApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.helmAppViaAppVersionIDResponseDTO)
	case GetHelmEnvAutocompleteApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.helmEnvAutocompleteResponseDTO)
	case GetTemplateValuesViaReferenceIdApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.templateValuesViaReferenceIdResponseDTO)
	case GetApplicationValuesListApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.getApplicationValuesListResponseDTO)
	case SaveTemplateValuesApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.saveTemplateValuesResponseDTO)
	case DeleteTemplateValuesApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.deleteTemplateValuesResponseDTO)
	case SearchAppStoreChartByNameApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.appStoreChartByNameResponseDTO)
	case InstallAppApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.installAppResponseDto)
	case GetInstalledAppVersionApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.installedAppVersionResponseDTO)
	case GetInstalledAppDetailsApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.installedAppDetailsResponseDTO)
	case CheckAppExistsApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.checkAppExistsResponseDTO)
	case GetAllInstalledAppApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.getAllInstalledAppResponseDTO)
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
func (suite *AppStoreDiscoverTestSuite) AfterSuite() {
	installedAppId, respOfDeleteInstalledAppId := DeleteHelmApp(suite.authToken)
	assert.Equal(suite.T(), installedAppId, respOfDeleteInstalledAppId)
}
