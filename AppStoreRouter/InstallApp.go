package AppStoreRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

//todo disabling this test case until latest build not deployed on stage cluster
/*func (suite *AppStoreTestSuite) TestInstallAppApiWithPayloadHavingAlreadyInstalledAppName() {
	expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayloadWithAlreadyExistsName.json")
	resp := HitInstallAppApi(string(expectedPayload))

	log.Println("Hitting the install App API with already installed app name")
	latestResponse := HitInstallAppApi(string(expectedPayload))

	log.Println("Validating the InstallAppApi response with already existed name in payload")
	assert.Equal(suite.T(), "applications.argoproj.io \"deepak-airflow-test-already-installed-devtron-demo\" already exists", latestResponse.Errors[0].UserMessage)

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(resp.InstallAppRequestDto.InstalledAppId))
	assert.Equal(suite.T(), resp.InstallAppRequestDto.InstalledAppId, respOfDeleteInstallAppApi.InstallAppRequestDto.InstalledAppId)
}*/

func (suite *AppStoreTestSuite) TestInstallApp() {

	suite.Run("A=1=InstallAppWithValidPayload", func() {
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
		log.Println("Hitting the InstallAppApi with valid payload")
		resp := HitInstallAppApi(string(expectedPayload), suite.authToken)
		expectedValuesOverrideYaml, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/expectedValuesOverrideYaml.txt")

		log.Println("Validating the InstallAppApi response with with valid payload")
		assert.Equal(suite.T(), string(expectedValuesOverrideYaml), resp.InstallAppRequestDto.ValuesOverrideYaml)
		assert.NotNil(suite.T(), resp.InstallAppRequestDto.InstalledAppId)

		log.Println("Removing the data created via API")
		respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(resp.InstallAppRequestDto.InstalledAppId), suite.authToken)
		assert.Equal(suite.T(), resp.InstallAppRequestDto.InstalledAppId, respOfDeleteInstallAppApi.InstallAppRequestDto.InstalledAppId)
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
}
