package AppStoreDiscoverRouter

import (
	"automation-suite/AppStoreRouter"
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestDiscoverPreviouslyInstalledHelmAppsViaRepoId() {
	log.Println("=== Here we are getting airflow chart repo ===")
	queryParams := map[string]string{"appStoreName": "airflow"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	DiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)

	suite.Run("A=1=DiscoverWithoutDeployment", func() {
		deploymentOfInstalledApp := HitGetDeploymentOfInstalledAppApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		assert.Nil(suite.T(), deploymentOfInstalledApp.Result)
	})

	suite.Run("A=2=DiscoverAfterDeployment", func() {
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
		log.Println("Hitting the InstallAppApi with valid payload")
		resp := AppStoreRouter.HitInstallAppApi(string(expectedPayload), suite.authToken)
		time.Sleep(5 * time.Second)
		log.Println("Hitting the GetDeploymentOfInstalledApp API with valid payload")
		deploymentOfInstalledApp := HitGetDeploymentOfInstalledAppApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		assert.NotNil(suite.T(), deploymentOfInstalledApp.Result[0].InstalledAppVersionId)
		assert.Equal(suite.T(), deploymentOfInstalledApp.Result[0].AppName, "deepak-airflow-test")
		log.Println("Removing the data created via API")
		respOfDeleteInstallAppApi := AppStoreRouter.HitDeleteInstalledAppApi(strconv.Itoa(resp.Result.InstalledAppId), suite.authToken)
		assert.Equal(suite.T(), resp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
	})
}

//todo need to check this app once issue get fixed for search API for a chart-repo added from global configurations
