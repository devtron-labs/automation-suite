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

func (suite *AppStoreDiscoverTestSuite) TestGetInstalledAppVersion() {
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
	installedAppVersionId := responseAfterInstallingApp.Result.InstalledAppVersionId

	suite.Run("A=1=GetDetailsWithCorrectAppId", func() {
		installedAppVersion := HitGetInstalledAppVersionApi(strconv.Itoa(installedAppVersionId), suite.authToken)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.AppName, installedAppVersion.Result.AppName)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.AppStoreVersion, installedAppVersion.Result.AppStoreVersion)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.GitOpsRepoName, installedAppVersion.Result.GitOpsRepoName)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.ValuesOverrideYaml, installedAppVersion.Result.ValuesOverrideYaml)
	})

	suite.Run("A=2=GetDetailsWithIncorrectAppId", func() {
		randomAppId := Base.GetRandomNumberOf9Digit()
		installedAppVersion := HitGetInstalledAppVersionApi(strconv.Itoa(randomAppId), suite.authToken)
		assert.Equal(suite.T(), 404, installedAppVersion.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", installedAppVersion.Error[0].UserMessage)
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
