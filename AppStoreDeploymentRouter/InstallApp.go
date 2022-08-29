package AppStoreDeploymentRouter

import (
	"automation-suite/AppStoreDeploymentRouter/RequestDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *AppStoreDeploymentTestSuite) TestInstallApp() {

	suite.Run("A=1=InstallAppWithValidPayload", func() {
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
		log.Println("Hitting the InstallAppApi with valid payload")
		installAppRequestDTO := RequestDTOs.InstallAppRequestDTO{}
		json.Unmarshal(expectedPayload, &installAppRequestDTO)

		AppName := "automation" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		log.Println("=== Helm AppName for this Test Case is :===", AppName)
		installAppRequestDTO.AppName = AppName
		requestPayload, _ := json.Marshal(installAppRequestDTO)
		resp := HitInstallAppApi(string(requestPayload), suite.authToken)
		expectedValuesOverrideYaml, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/expectedValuesOverrideYaml.txt")

		log.Println("Validating the InstallAppApi response with with valid payload")
		assert.Equal(suite.T(), string(expectedValuesOverrideYaml), resp.Result.ValuesOverrideYaml)
		assert.NotNil(suite.T(), resp.Result.InstalledAppId)

		log.Println("Removing the data created via API")
		respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(resp.Result.InstalledAppId), suite.authToken)
		assert.Equal(suite.T(), resp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
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
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayloadWithAlreadyExistsName.json")

		log.Println("Hitting the install App API with already installed app name")
		latestResponse := HitInstallAppApi(string(expectedPayload), suite.authToken)

		log.Println("Validating the InstallAppApi response with already existed name in payload")
		assert.Equal(suite.T(), "applications.argoproj.io \"deepak-airflow-test-already-installed-devtron-demo\" already exists", latestResponse.Errors[0].UserMessage)
	})
}
