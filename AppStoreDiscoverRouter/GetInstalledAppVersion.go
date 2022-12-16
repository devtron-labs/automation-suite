package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func (suite *AppStoreDiscoverTestSuite) TestGetInstalledAppVersion() {

	responseAfterInstallingApp, _, _, _ := CreateHelmApp(suite.authToken)
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

	//log.Println("Removing the data created via API")
	//respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	//assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
