package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/sjson"
	"log"
	"sigs.k8s.io/yaml"
	"strconv"
)

func (suite *AppStoreDiscoverTestSuite) TestCheckAppExists() {
	var valuesOverrideInterface interface{}
	log.Println("=== Getting apache chart repo via DiscoverApp API ===")
	queryParams := map[string]string{"appStoreName": "apache"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
	var requiredReferenceId int
	for _, DiscoveredApp := range ActiveDiscoveredApps.Result {
		if DiscoveredApp.ChartName == "bitnami" {
			requiredReferenceId = DiscoveredApp.AppStoreApplicationVersionId
			break
		}
	}
	log.Println("=== Getting Template values for apache chart===")
	queryParamsOfApi := map[string]string{"referenceId": strconv.Itoa(requiredReferenceId), "kind": "DEFAULT"}
	referenceTemplate := HitGetTemplateValuesViaReferenceIdApi(queryParamsOfApi, suite.authToken)
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
	responseAfterInstallingApp := HitInstallAppApi(string(updatedByteValueOfInstallAppRequestPayload), suite.authToken)
	installedAppId := responseAfterInstallingApp.Result.InstalledAppId
	environmentId := strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)
	AppName := responseAfterInstallingApp.Result.AppName
	log.Println("=== Here we are installing helm App ===")
	queryParamsOfApi = make(map[string]string)
	queryParamsOfApi["installed-app-id"] = strconv.Itoa(installedAppId)
	queryParamsOfApi["env-id"] = environmentId
	PollForAppStatusInAppDetails(queryParamsOfApi, suite.authToken)
	installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
	assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree.Status)

	suite.Run("A=1=TestApiWithNonExistingNameOnly", func() {
		listOfNames := []string{Base.GetRandomStringOfGivenLength(9)}
		checkAppExistsRequestDTO := getCheckAppExistsApi(listOfNames)
		payloadForCheckAppExistsRequestDTO, _ := json.Marshal(checkAppExistsRequestDTO)
		CheckAppExistsApiResponse := HitCheckAppExistsOrNot(string(payloadForCheckAppExistsRequestDTO), suite.authToken)
		assert.False(suite.T(), CheckAppExistsApiResponse.Result[0].Exists)
	})

	suite.Run("A=2=TestApiWithExistingNameOnly", func() {
		listOfNames := []string{AppName}
		checkAppExistsRequestDTO := getCheckAppExistsApi(listOfNames)
		payloadForCheckAppExistsRequestDTO, _ := json.Marshal(checkAppExistsRequestDTO)
		CheckAppExistsApiResponse := HitCheckAppExistsOrNot(string(payloadForCheckAppExistsRequestDTO), suite.authToken)
		assert.True(suite.T(), CheckAppExistsApiResponse.Result[0].Exists)
	})

	suite.Run("A=3=TestApiWithExistingAndNonExistingNames", func() {
		listOfNames := []string{Base.GetRandomStringOfGivenLength(9), AppName}
		checkAppExistsRequestDTO := getCheckAppExistsApi(listOfNames)
		payloadForCheckAppExistsRequestDTO, _ := json.Marshal(checkAppExistsRequestDTO)
		CheckAppExistsApiResponse := HitCheckAppExistsOrNot(string(payloadForCheckAppExistsRequestDTO), suite.authToken)
		assert.False(suite.T(), CheckAppExistsApiResponse.Result[0].Exists)
		assert.True(suite.T(), CheckAppExistsApiResponse.Result[1].Exists)
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
