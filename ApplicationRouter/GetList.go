package ApplicationRouter

import (
	"automation-suite/AppStoreDeploymentRouter"
	"automation-suite/AppStoreRouter"
	"automation-suite/AppStoreRouter/RequestDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *ApplicationsRouterTestSuite) TestGetInstalledAppDetails() {
	log.Println("=== Here We are installing Helm chart from chart-store ===")
	expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
	log.Println("Hitting the InstallAppApi with valid payload")
	installAppRequestDTO := RequestDTOs.InstallAppRequestDTO{}
	json.Unmarshal(expectedPayload, &installAppRequestDTO)

	AppName := "automation" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
	log.Println("=====Helm AppName used in this test Case is====", AppName)
	installAppRequestDTO.AppName = AppName
	requestPayload, _ := json.Marshal(installAppRequestDTO)
	responseAfterInstallingApp := AppStoreDeploymentRouter.HitInstallAppApi(string(requestPayload), suite.authToken)
	time.Sleep(2 * time.Second)
	installedAppId := responseAfterInstallingApp.Result.InstalledAppId
	environmentId := strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)

	log.Println("=== Here we are installing helm App ===")
	queryParamsOfApi := make(map[string]string)
	queryParamsOfApi["installed-app-id"] = strconv.Itoa(installedAppId)
	queryParamsOfApi["env-id"] = environmentId
	PollForAppStatusInAppDetails(queryParamsOfApi, suite.authToken)
	installedAppDetails := AppStoreRouter.HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
	assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree.Status)

	suite.Run("A=1=GetListWithCorrectArguments", func() {
		queryParams := make(map[string]string)
		queryParams["name"] = installedAppDetails.Result.AppName + "-devtron-demo"
		queryParams["refresh"] = "5"
		queryParams["project"] = "devtron-demo"
		ResponseOfGetListApi := HitGetListApi(queryParams, suite.authToken)
		assert.NotNil(suite.T(), ResponseOfGetListApi.Result.Metadata.ResourceVersion)
		assert.NotEqual(suite.T(), ResponseOfGetListApi.Result.Metadata.ResourceVersion, "")
	})

	suite.Run("A=1=GetListWithInCorrectProject", func() {
		queryParams := make(map[string]string)
		queryParams["name"] = installedAppDetails.Result.AppName + "-devtron-demo"
		queryParams["refresh"] = "5"
		queryParams["project"] = "devtron-demo"
		ResponseOfGetListApi := HitGetListApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), ResponseOfGetListApi.Errors[0].InternalMessage, "[{rpc error: code = NotFound desc = application 'automation3guoj-devtron-demoautomation3guoj' not found}]")
	})

	suite.Run("A=3=GetListWithIncorrectName", func() {
		queryParams := make(map[string]string)
		queryParams["name"] = installedAppDetails.Result.AppName
		queryParams["refresh"] = "5"
		queryParams["project"] = "devtron-demo"
		ResponseOfGetListApi := HitGetListApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), ResponseOfGetListApi.Result.Metadata.ResourceVersion, "[{rpc error: code = NotFound desc = application 'automation3guoj-devtron-demoautomation3guoj' not found}]")
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := AppStoreDeploymentRouter.HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}

func PollForAppStatusInAppDetails(queryParams map[string]string, authToken string) bool {
	count := 0
	for {
		respOfGetApplicationDetailApi := AppStoreRouter.HitGetInstalledAppDetailsApi(queryParams, authToken)
		appStatus := respOfGetApplicationDetailApi.Result.ResourceTree.Status
		time.Sleep(1 * time.Second)
		count = count + 1
		if appStatus == "Healthy" || count >= 500 {
			break
		}
	}
	return true
}
