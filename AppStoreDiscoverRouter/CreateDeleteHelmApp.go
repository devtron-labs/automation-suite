package AppStoreDiscoverRouter

import (
	"automation-suite/AppStoreDiscoverRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/tidwall/sjson"
	"log"
	"sigs.k8s.io/yaml"
	"strconv"
)

var responseAfterInstallingAppPtr *ResponseDTOs.InstallAppResponseDTO
var updatedInstallAppRequestPayload string
var activeDiscoveredAppsPtr *ResponseDTOs.DiscoverAppApiResponse
var appStoreIdPtr int

func CreateHelmApp(authToken string) (ResponseDTOs.InstallAppResponseDTO, string, ResponseDTOs.DiscoverAppApiResponse, int) {
	if responseAfterInstallingAppPtr != nil && activeDiscoveredAppsPtr != nil && len(updatedInstallAppRequestPayload) != 0 {
		return *responseAfterInstallingAppPtr, updatedInstallAppRequestPayload, *activeDiscoveredAppsPtr, appStoreIdPtr
	}
	var valuesOverrideInterface interface{}
	log.Println("=== Getting apache chart repo via DiscoverApp API ===")
	queryParamsForGettingHelmAppData := map[string]string{"appStoreName": "apache"}
	PollForGettingHelmAppData(queryParamsForGettingHelmAppData, authToken)
	ActiveDiscoveredApps := HitDiscoverAppApi(queryParamsForGettingHelmAppData, authToken)
	var appStoreId int
	var requiredReferenceId int
	for _, DiscoveredApp := range ActiveDiscoveredApps.Result {
		if DiscoveredApp.ChartName == "bitnami" {
			requiredReferenceId = DiscoveredApp.AppStoreApplicationVersionId
			appStoreId = DiscoveredApp.Id
			break
		}
	}
	log.Println("=== Getting Template values for apache chart===")
	queryParamsOfApi := map[string]string{"referenceId": strconv.Itoa(requiredReferenceId), "kind": "DEFAULT"}
	referenceTemplate := HitGetTemplateValuesViaReferenceIdApi(queryParamsOfApi, authToken)
	valuesOverrideYaml := referenceTemplate.Result.Values
	if err := yaml.Unmarshal([]byte(valuesOverrideYaml), &valuesOverrideInterface); err != nil {
		panic(err)
	}
	Base.ConvertYamlIntoJson(valuesOverrideInterface)
	valuesOverrideJson, _ := json.Marshal(valuesOverrideInterface)
	jsonOfSaveDeploymentTemp := string(valuesOverrideJson)
	jsonWithTypeAsClusterIP, _ := sjson.Set(jsonOfSaveDeploymentTemp, "service.type", "ClusterIP")
	updatedValuesOverrideJson := []byte(jsonWithTypeAsClusterIP)
	log.Println("=== converting Json into YAML for Values Override in Install API===")
	updatedValuesOverrideYaml, _ := yaml.JSONToYAML(updatedValuesOverrideJson)
	installAppRequestDTO := GetRequestDtoForInstallApp(requiredReferenceId, requiredReferenceId, valuesOverrideInterface, string(updatedValuesOverrideYaml))
	byteValueOfInstallAppRequestPayload, _ := json.Marshal(installAppRequestDTO)
	jsonOfSaveDeploymentTemp1 := string(byteValueOfInstallAppRequestPayload)
	jsonWithTypeAsClusterIP1, _ := sjson.Set(jsonOfSaveDeploymentTemp1, "valuesOverride.service.type", "ClusterIP")
	updatedByteValueOfInstallAppRequestPayload := []byte(jsonWithTypeAsClusterIP1)
	responseAfterInstallingApp := HitInstallAppApi(string(updatedByteValueOfInstallAppRequestPayload), authToken)
	updatedInstallAppRequestPayload = string(updatedByteValueOfInstallAppRequestPayload)
	responseAfterInstallingAppPtr = &responseAfterInstallingApp
	activeDiscoveredAppsPtr = &ActiveDiscoveredApps
	appStoreIdPtr = appStoreId
	return responseAfterInstallingApp, updatedInstallAppRequestPayload, ActiveDiscoveredApps, appStoreId
}

func DeleteHelmApp(authToken string) (int, int) {
	log.Println("Removing the data created for helm app via API")
	if responseAfterInstallingAppPtr == nil {
		return 0, 0
	}
	respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingAppPtr.Result.InstalledAppId), authToken)
	return responseAfterInstallingAppPtr.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId
}
