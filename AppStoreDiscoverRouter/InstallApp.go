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
)

func (suite *AppStoreDiscoverTestSuite) TestInstallApp() {
	var valuesOverrideInterface interface{}
	var InstalledAppId int
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

	var updatedByteValueOfInstallAppRequestPayload []byte
	suite.Run("A=1=InstallAppWithValidPayload", func() {
		installAppRequestDTO := GetRequestDtoForInstallApp(requiredReferenceId, requiredReferenceId, valuesOverrideInterface, string(updatedValuesOverrideYaml))
		byteValueOfInstallAppRequestPayload, _ := json.Marshal(installAppRequestDTO)
		jsonOfSaveDeploymentTemp1 := string(byteValueOfInstallAppRequestPayload)
		jsonWithTypeAsClusterIP1, _ := sjson.Set(jsonOfSaveDeploymentTemp1, "valuesOverride.service.type", "ClusterIP")
		updatedByteValueOfInstallAppRequestPayload = []byte(jsonWithTypeAsClusterIP1)
		responseAfterInstallingApp := HitInstallAppApi(string(updatedByteValueOfInstallAppRequestPayload), suite.authToken)
		log.Println("=== Validating the response of install Helm-chart API ===")
		assert.NotNil(suite.T(), responseAfterInstallingApp.Result.InstalledAppId)
		queryParamsForAppStatus := make(map[string]string)
		queryParamsForAppStatus["installed-app-id"] = strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId)
		queryParamsForAppStatus["env-id"] = strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)
		PollForAppStatusInAppDetails(queryParamsForAppStatus, suite.authToken)
		respOfGetApplicationDetailApi := HitGetInstalledAppDetailsApi(queryParamsForAppStatus, suite.authToken)
		InstalledAppId = respOfGetApplicationDetailApi.Result.InstalledAppId
		assert.Equal(suite.T(), "Healthy", respOfGetApplicationDetailApi.Result.ResourceTree.Status)
		assert.Equal(suite.T(), "apache", respOfGetApplicationDetailApi.Result.AppStoreAppName)

	})
	suite.Run("A=2=InstallAppWithInvalidTeamId", func() {
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstalledAppRequestPayloadWithInvalidTeamId.json")
		log.Println("Hitting the InstallAppApi with InvalidTeamId in Payload")
		resp := HitInstallAppApi(string(expectedPayload), suite.authToken)
		assert.Equal(suite.T(), "[{ERROR #23503 insert or update on table \"app\" violates foreign key constraint \"app_team_id_fkey\"}]", resp.Errors[0].InternalMessage)

	})
	suite.Run("A=3=InstallAppWithInvalidAppStoreVersion", func() {
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstalledAppRequestPayloadWithInvalidAppStoreVersion.json")
		log.Println("Hitting the InstallAppApi with invalid AppStoreVersion in Payload")
		resp := HitInstallAppApi(string(expectedPayload), suite.authToken)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", resp.Errors[0].InternalMessage)
	})
	suite.Run("A=4=InstallAppWithInvalidEnvId", func() {
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstalledAppRequestPayloadWithInvalidEnvId.json")
		log.Println("Hitting the InstallAppApi with invalid EnvId in Payload")
		resp := HitInstallAppApi(string(expectedPayload), suite.authToken)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", resp.Errors[0].InternalMessage)
	})
	suite.Run("A=5=InstallAppWithInvalidReferenceValueKind", func() {
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstalledAppRequestPayloadWithInvalidReferenceValueKind.json")
		log.Println("Hitting the InstallAppApi with valid payload")
		resp := HitInstallAppApi(string(expectedPayload), suite.authToken)
		assert.Equal(suite.T(), "Key: 'InstallAppVersionDTO.ReferenceValueKind' Error:Field validation for 'ReferenceValueKind' failed on the 'oneof' tag", resp.Errors[0].UserMessage)
	})

	suite.Run("A=6=InstallAppWithAlreadyExistingName", func() {
		log.Println("Hitting the install App API with already installed app name")
		latestResponse := HitInstallAppApi(string(updatedByteValueOfInstallAppRequestPayload), suite.authToken)

		log.Println("Validating the InstallAppApi response with already existed name in payload")
		assert.True(suite.T(), strings.Contains(latestResponse.Errors[0].UserMessage, "app already exists"))
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)

}
