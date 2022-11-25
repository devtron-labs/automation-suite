package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/sjson"
	"log"
	"sigs.k8s.io/yaml"
	"strconv"
	"strings"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestGetTemplateValuesList() {
	var valuesOverrideInterface interface{}
	var AppStoreId int
	log.Println("=== Getting apache chart repo via DiscoverApp API ===")
	queryParams := map[string]string{"appStoreName": "apache"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
	var requiredReferenceId int
	for _, DiscoveredApp := range ActiveDiscoveredApps.Result {
		if DiscoveredApp.ChartName == "bitnami" {
			requiredReferenceId = DiscoveredApp.AppStoreApplicationVersionId
			AppStoreId = DiscoveredApp.Id
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

	suite.Run("A=1=GetValuesListForKindDeploy", func() {
		log.Println("=== Here We are getting noOfDeployedCharts before new deployment ===")
		ApplicationValuesList := HitGetApplicationValuesListApi(strconv.Itoa(AppStoreId), suite.authToken)
		noOfDeployedCharts := len(ApplicationValuesList.Result.Values[2].Values)

		installAppRequestDTO := GetRequestDtoForInstallApp(requiredReferenceId, requiredReferenceId, valuesOverrideInterface, string(updatedValuesOverrideYaml))
		byteValueOfInstallAppRequestPayload, _ := json.Marshal(installAppRequestDTO)
		jsonOfSaveDeploymentTemp1 := string(byteValueOfInstallAppRequestPayload)
		jsonWithTypeAsClusterIP1, _ := sjson.Set(jsonOfSaveDeploymentTemp1, "valuesOverride.service.type", "ClusterIP")
		updatedByteValueOfInstallAppRequestPayload := []byte(jsonWithTypeAsClusterIP1)
		responseAfterInstallingApp := HitInstallAppApi(string(updatedByteValueOfInstallAppRequestPayload), suite.authToken)
		time.Sleep(5 * time.Second)
		appName := responseAfterInstallingApp.Result.AppName
		log.Println("=== Here We are getting noOfDeployedCharts after new deployment ===")
		ApplicationValuesList = HitGetApplicationValuesListApi(strconv.Itoa(AppStoreId), suite.authToken)
		indexOfLastDeployed := len(ApplicationValuesList.Result.Values[2].Values)
		assert.Equal(suite.T(), appName, ApplicationValuesList.Result.Values[2].Values[indexOfLastDeployed-1].Name)
		assert.Equal(suite.T(), noOfDeployedCharts+1, len(ApplicationValuesList.Result.Values[2].Values))

		log.Println("Removing the data created via API")
		respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
	})

	suite.Run("A=2=GetValuesListForKindDefault", func() {
		deploymentOfInstalledApp := HitGetApplicationValuesListApi(strconv.Itoa(AppStoreId), suite.authToken)
		assert.True(suite.T(), len(deploymentOfInstalledApp.Result.Values[0].Values) >= 1)
	})

	suite.Run("A=3=GetValuesListForKindTemplate", func() {
		log.Println("=== Here we are saving template values   ===")
		appName := "automation-preset-" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		valueForPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/SaveTemplateValuesRequestPayload.txt")
		requestPayload := getPayloadForSaveTemplateValues(appName, string(valueForPayload), ActiveDiscoveredApps.Result[0].AppStoreApplicationVersionId)
		payloadByteArray, _ := json.Marshal(requestPayload)
		responseOfSaveTemplateApi := HitSaveTemplateValuesApi(string(payloadByteArray), suite.authToken)
		log.Println("=== Here We are getting noOfTemplate before deleting it ===")
		ApplicationValuesList := HitGetApplicationValuesListApi(strconv.Itoa(AppStoreId), suite.authToken)
		noOfDeployedCharts := len(ApplicationValuesList.Result.Values[1].Values)
		deleteTemplateValueApiResponse := HitDeleteTemplateValuesApi(strconv.Itoa(responseOfSaveTemplateApi.Result.Id), suite.authToken)
		assert.True(suite.T(), deleteTemplateValueApiResponse.Result)
		log.Println("=== Here We are getting noOfTemplate after deleting it ===")
		ApplicationValuesList = HitGetApplicationValuesListApi(strconv.Itoa(AppStoreId), suite.authToken)
		assert.Equal(suite.T(), noOfDeployedCharts-1, len(ApplicationValuesList.Result.Values[1].Values))
	})

	suite.Run("A=4=GetValuesListForKindExisting", func() {

	})

	suite.Run("A=5=GetListWithInvalidAppId", func() {
		randomAppId := Base.GetRandomNumberOf9Digit()
		appValuesList := HitGetApplicationValuesListApi(strconv.Itoa(randomAppId), suite.authToken)
		assert.Nil(suite.T(), appValuesList.Result.Values[0].Values)
		assert.Nil(suite.T(), appValuesList.Result.Values[1].Values)
		assert.Equal(suite.T(), 0, len(appValuesList.Result.Values[2].Values))
		assert.Equal(suite.T(), 0, len(appValuesList.Result.Values[3].Values))
	})
}

//todo need to add a test case around GetValuesListForKindExisting after understanding the flow of this
