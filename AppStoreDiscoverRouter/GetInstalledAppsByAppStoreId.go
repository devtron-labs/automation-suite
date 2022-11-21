package AppStoreDiscoverRouter

import (
	"automation-suite/AppStoreDeploymentRouter"
	"automation-suite/AppStoreDeploymentRouter/RequestDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestDiscoverPreviouslyInstalledHelmAppsViaRepoId() {
	log.Println("=== Here we are getting apache chart repo ===")
	queryParams := map[string]string{"appStoreName": "apache"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	DiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
	var installedAppId int

	suite.Run("A=1=GetInstalledAppsByAppStoreId", func() {
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
		log.Println("Hitting the InstallAppApi with valid payload")
		installAppRequestDTO := RequestDTOs.InstallAppRequestDTO{}
		json.Unmarshal(expectedPayload, &installAppRequestDTO)
		AppName := "automation" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))

		log.Println("=====Helm AppName used in this test Case is====", AppName)
		installAppRequestDTO.AppName = AppName
		requestPayload, _ := json.Marshal(installAppRequestDTO)
		resp := AppStoreDeploymentRouter.HitInstallAppApi(string(requestPayload), suite.authToken)
		installedAppId = resp.Result.InstalledAppId
		time.Sleep(5 * time.Second)
		log.Println("Hitting the GetDeploymentOfInstalledApp API with valid payload")
		deploymentOfInstalledApp := GetInstalledAppsByAppStoreId(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		assert.NotNil(suite.T(), deploymentOfInstalledApp.Result[len(deploymentOfInstalledApp.Result)-1].InstalledAppVersionId)
		assert.Equal(suite.T(), deploymentOfInstalledApp.Result[len(deploymentOfInstalledApp.Result)-1].AppName, AppName)
	})
	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := AppStoreDeploymentRouter.HitDeleteInstalledAppApi(strconv.Itoa(installedAppId), suite.authToken)
	assert.Equal(suite.T(), installedAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)

}

//todo need to check this app once issue get fixed for search API for a chart-repo added from global configurations
