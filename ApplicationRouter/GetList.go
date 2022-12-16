package ApplicationRouter

import (
	"automation-suite/AppStoreDiscoverRouter"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *ApplicationsRouterTestSuite) TestGetList() {

	responseAfterInstallingApp, _, _, _ := AppStoreDiscoverRouter.CreateHelmApp(suite.authToken)
	installedAppId := responseAfterInstallingApp.Result.InstalledAppId
	environmentId := strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)

	log.Println("=== Here we are installing helm App ===")
	queryParamsForGettingHealthStatus := make(map[string]string)
	queryParamsForGettingHealthStatus["installed-app-id"] = strconv.Itoa(installedAppId)
	queryParamsForGettingHealthStatus["env-id"] = environmentId
	AppStoreDiscoverRouter.PollForAppStatusInAppDetails(queryParamsForGettingHealthStatus, suite.authToken)
	installedAppDetails := AppStoreDiscoverRouter.HitGetInstalledAppDetailsApi(queryParamsForGettingHealthStatus, suite.authToken)
	assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree["status"])

	suite.Run("A=1=GetListWithCorrectArguments", func() {
		queryParams := make(map[string]string)
		queryParams["name"] = installedAppDetails.Result.AppName + "-devtron-demo"
		queryParams["refresh"] = "5"
		queryParams["project"] = "devtron-demo"
		ResponseOfGetListApi := HitGetListApi(queryParams, suite.authToken)
		assert.NotNil(suite.T(), ResponseOfGetListApi.Result.Metadata.ResourceVersion)
		assert.NotEqual(suite.T(), ResponseOfGetListApi.Result.Metadata.ResourceVersion, "")
	})

	suite.Run("A=2=GetListWithIncorrectName", func() {
		queryParams := make(map[string]string)
		queryParams["name"] = installedAppDetails.Result.AppName + "wrong"
		queryParams["refresh"] = "5"
		queryParams["project"] = "devtron-demo"
		ResponseOfGetListApi := HitGetListApi(queryParams, suite.authToken)
		assert.True(suite.T(), strings.Contains(ResponseOfGetListApi.Errors[0].InternalMessage, "error: code = NotFound desc = application"))
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := AppStoreDiscoverRouter.HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
