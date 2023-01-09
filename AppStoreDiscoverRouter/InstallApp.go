package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *AppStoreDiscoverTestSuite) TestInstallApp() {
	responseAfterInstallingApp, updatedByteValueOfInstallAppRequestPayload, _, _ := CreateHelmApp(suite.authToken)
	//InstalledAppId := responseAfterInstallingApp.Result.InstalledAppId
	suite.Run("A=1=InstallAppWithValidPayload", func() {
		log.Println("=== Validating the response of install Helm-chart API ===")
		assert.NotNil(suite.T(), responseAfterInstallingApp.Result.InstalledAppId)
		queryParamsForAppStatus := make(map[string]string)
		queryParamsForAppStatus["installed-app-id"] = strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId)
		queryParamsForAppStatus["env-id"] = strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)
		PollForAppStatusInAppDetails(queryParamsForAppStatus, suite.authToken)
		respOfGetApplicationDetailApi := HitGetInstalledAppDetailsApi(queryParamsForAppStatus, suite.authToken)
		assert.Equal(suite.T(), "Healthy", respOfGetApplicationDetailApi.Result.ResourceTree["status"])
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

	//log.Println("Removing the data created via API")
	//respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(InstalledAppId), suite.authToken)
	//assert.Equal(suite.T(), InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)

}
