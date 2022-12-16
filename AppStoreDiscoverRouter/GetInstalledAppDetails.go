package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestGetInstalledAppDetails() {

	responseAfterInstallingApp, _, _, _ := CreateHelmApp(suite.authToken)
	installedAppId := responseAfterInstallingApp.Result.InstalledAppId
	environmentId := strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)

	suite.Run("A=1=GetDetailsWithCorrectAppIdAndEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(installedAppId)
		queryParamsOfApi["env-id"] = environmentId
		PollForAppStatusInAppDetails(queryParamsOfApi, suite.authToken)
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree["status"])
		assert.Equal(suite.T(), installedAppId, installedAppDetails.Result.InstalledAppId)
		assert.Equal(suite.T(), "apache", installedAppDetails.Result.AppStoreAppName)
		assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree["podMetadata"])
		assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree["nodes"])
		//assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree["hosts"].([]map[string]interface{})[0]["resourcesInfo"])
	})

	suite.Run("A=2=GetDetailsWithCorrectAppIdAndIncorrectEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(installedAppId)
		queryParamsOfApi["env-id"] = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), 404, installedAppDetails.Code)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", installedAppDetails.Error[0].InternalMessage)
	})

	suite.Run("A=3=GetDetailsWithIncorrectAppIdAndCorrectEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		queryParamsOfApi["env-id"] = environmentId
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), 404, installedAppDetails.Code)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", installedAppDetails.Error[0].InternalMessage)
	})

	suite.Run("A=4=GetDetailsWithIncorrectAppIdAndEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		queryParamsOfApi["env-id"] = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), 404, installedAppDetails.Code)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", installedAppDetails.Error[0].InternalMessage)
	})

	//log.Println("Removing the data created via API")
	//respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	//assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}

func PollForAppStatusInAppDetails(queryParams map[string]string, authToken string) bool {
	count := 0
	for {
		respOfGetApplicationDetailApi := HitGetInstalledAppDetailsApi(queryParams, authToken)
		appStatus := respOfGetApplicationDetailApi.Result.ResourceTree["status"]
		time.Sleep(1 * time.Second)
		count = count + 1
		if appStatus == "Healthy" || count >= 500 {
			break
		}
	}
	return true
}
