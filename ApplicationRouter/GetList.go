package ApplicationRouter

import (
	"automation-suite/AppStoreDiscoverRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/sjson"
	"log"
	"sigs.k8s.io/yaml"
	"strconv"
	"strings"
)

func (suite *ApplicationsRouterTestSuite) TestGetList() {
	var valuesOverrideInterface interface{}
	log.Println("=== Getting apache chart repo via DiscoverApp API ===")
	queryParamsForGettingHelmAppData := map[string]string{"appStoreName": "apache"}
	AppStoreDiscoverRouter.PollForGettingHelmAppData(queryParamsForGettingHelmAppData, suite.authToken)
	ActiveDiscoveredApps := AppStoreDiscoverRouter.HitDiscoverAppApi(queryParamsForGettingHelmAppData, suite.authToken)
	var requiredReferenceId int
	for _, DiscoveredApp := range ActiveDiscoveredApps.Result {
		if DiscoveredApp.ChartName == "bitnami" {
			requiredReferenceId = DiscoveredApp.AppStoreApplicationVersionId
			break
		}
	}
	log.Println("=== Getting Template values for apache chart===")
	queryParamsOfApi := map[string]string{"referenceId": strconv.Itoa(requiredReferenceId), "kind": "DEFAULT"}
	referenceTemplate := AppStoreDiscoverRouter.HitGetTemplateValuesViaReferenceIdApi(queryParamsOfApi, suite.authToken)
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
	installAppRequestDTO := AppStoreDiscoverRouter.GetRequestDtoForInstallApp(requiredReferenceId, requiredReferenceId, valuesOverrideInterface, string(updatedValuesOverrideYaml))
	byteValueOfInstallAppRequestPayload, _ := json.Marshal(installAppRequestDTO)
	jsonOfSaveDeploymentTemp1 := string(byteValueOfInstallAppRequestPayload)
	jsonWithTypeAsClusterIP1, _ := sjson.Set(jsonOfSaveDeploymentTemp1, "valuesOverride.service.type", "ClusterIP")
	updatedByteValueOfInstallAppRequestPayload := []byte(jsonWithTypeAsClusterIP1)
	responseAfterInstallingApp := AppStoreDiscoverRouter.HitInstallAppApi(string(updatedByteValueOfInstallAppRequestPayload), suite.authToken)
	installedAppId := responseAfterInstallingApp.Result.InstalledAppId
	environmentId := strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)

	log.Println("=== Here we are installing helm App ===")
	queryParamsForGettingHealthStatus := make(map[string]string)
	queryParamsForGettingHealthStatus["installed-app-id"] = strconv.Itoa(installedAppId)
	queryParamsForGettingHealthStatus["env-id"] = environmentId
	AppStoreDiscoverRouter.PollForAppStatusInAppDetails(queryParamsForGettingHealthStatus, suite.authToken)
	installedAppDetails := AppStoreDiscoverRouter.HitGetInstalledAppDetailsApi(queryParamsForGettingHealthStatus, suite.authToken)
	assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree.Status)

	suite.Run("A=1=GetListWithCorrectArguments", func() {
		queryParams := make(map[string]string)
		queryParams["name"] = installedAppDetails.Result.AppName + "-devtron-demo"
		queryParams["refresh"] = "5"
		queryParams["project"] = "devtron-demo"
		ResponseOfGetListApi := HitGetListApi(queryParams, suite.authToken)
		assert.NotNil(suite.T(), ResponseOfGetListApi.Result.Metadata.ResourceVersion)
		assert.NotEqual(suite.T(), ResponseOfGetListApi.Result.Metadata.ResourceVersion, "")
	})

	suite.Run("A=2=GetListWithIncorrectName", func() {
		queryParams := make(map[string]string)
		queryParams["name"] = installedAppDetails.Result.AppName + "wrong"
		queryParams["refresh"] = "5"
		queryParams["project"] = "devtron-demo"
		ResponseOfGetListApi := HitGetListApi(queryParams, suite.authToken)
		assert.True(suite.T(), strings.Contains(ResponseOfGetListApi.Errors[0].InternalMessage, "error: code = NotFound desc = application"))
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := AppStoreDiscoverRouter.HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
