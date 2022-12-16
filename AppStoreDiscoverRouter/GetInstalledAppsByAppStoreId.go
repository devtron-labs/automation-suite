package AppStoreDiscoverRouter

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestDiscoverPreviouslyInstalledHelmAppsViaRepoId() {
	responseAfterInstallingApp, _, _, appStoreId := CreateHelmApp(suite.authToken)
	installedAppId := responseAfterInstallingApp.Result.InstalledAppId
	time.Sleep(10 * time.Second)
	suite.Run("A=1=GetInstalledAppsByAppStoreId", func() {
		log.Println("Hitting the GetDeploymentOfInstalledApp API with valid payload")
		deploymentOfInstalledApp := GetInstalledAppsByAppStoreId(strconv.Itoa(appStoreId), suite.authToken)
		lastDeployedId := 0
		for i, res := range deploymentOfInstalledApp.Result {
			if res.InstalledAppId == installedAppId {
				lastDeployedId = i
			}
		}
		assert.NotNil(suite.T(), deploymentOfInstalledApp.Result[lastDeployedId].InstalledAppVersionId)
		assert.Equal(suite.T(), installedAppId, deploymentOfInstalledApp.Result[lastDeployedId].InstalledAppId)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.AppName, deploymentOfInstalledApp.Result[lastDeployedId].AppName)
	})
	//log.Println("Removing the data created via API")
	//respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(installedAppId), suite.authToken)
	//assert.Equal(suite.T(), installedAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
