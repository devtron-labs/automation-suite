package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/sjson"
	"log"
	"sigs.k8s.io/yaml"
	"strconv"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestGetInstalledAppDetails() {
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
	queryParamsOfApiGetTemplateValuesViaReferenceId := map[string]string{"referenceId": strconv.Itoa(requiredReferenceId), "kind": "DEFAULT"}
	referenceTemplate := HitGetTemplateValuesViaReferenceIdApi(queryParamsOfApiGetTemplateValuesViaReferenceId, suite.authToken)
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

	suite.Run("A=1=GetDetailsWithCorrectAppIdAndEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(installedAppId)
		queryParamsOfApi["env-id"] = environmentId
		PollForAppStatusInAppDetails(queryParamsOfApi, suite.authToken)
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree.Status)
		assert.Equal(suite.T(), installedAppId, installedAppDetails.Result.InstalledAppId)
		assert.Equal(suite.T(), "apache", installedAppDetails.Result.AppStoreAppName)
		assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree.PodMetadata)
		assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree.Nodes)
		assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree.Hosts[0].ResourcesInfo)
	})

	suite.Run("A=2=GetDetailsWithCorrectAppIdAndIncorrectEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(installedAppId)
		queryParamsOfApi["env-id"] = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), 404, installedAppDetails.Code)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", installedAppDetails.Error[0].InternalMessage)
	})

	suite.Run("A=3=GetDetailsWithIncorrectAppIdAndCorrectEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		queryParamsOfApi["env-id"] = environmentId
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), 404, installedAppDetails.Code)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", installedAppDetails.Error[0].InternalMessage)
	})

	suite.Run("A=4=GetDetailsWithIncorrectAppIdAndEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		queryParamsOfApi["env-id"] = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), 404, installedAppDetails.Code)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", installedAppDetails.Error[0].InternalMessage)
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}

func PollForAppStatusInAppDetails(queryParams map[string]string, authToken string) bool {
	count := 0
	for {
		respOfGetApplicationDetailApi := HitGetInstalledAppDetailsApi(queryParams, authToken)
		appStatus := respOfGetApplicationDetailApi.Result.ResourceTree.Status
		time.Sleep(1 * time.Second)
		count = count + 1
		if appStatus == "Healthy" || count >= 500 {
			break
		}
	}
	return true
}
