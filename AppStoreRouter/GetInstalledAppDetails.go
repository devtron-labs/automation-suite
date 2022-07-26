package AppStoreRouter

import (
	"automation-suite/AppStoreDeploymentRouter"
	"automation-suite/AppStoreRouter/RequestDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *AppStoreTestSuite) TestGetInstalledAppDetails() {
	log.Println("=== Here We are installing Helm chart from chart-store ===")
	expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
	log.Println("Hitting the InstallAppApi with valid payload")
	installAppRequestDTO := RequestDTOs.InstallAppRequestDTO{}
	json.Unmarshal(expectedPayload, &installAppRequestDTO)
	installAppRequestDTO.AppName = "automation-helm-airflow" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
	requestPayload, _ := json.Marshal(installAppRequestDTO)
	responseAfterInstallingApp := AppStoreDeploymentRouter.HitInstallAppApi(string(requestPayload), suite.authToken)
	time.Sleep(2 * time.Second)
	installedAppId := responseAfterInstallingApp.Result.InstalledAppId
	environmentId := strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)
	suite.Run("A=1=GetDetailsWithCorrectAppIdAndEnvId", func() {
		queryParamsOfApi := make(map[string]string)
		queryParamsOfApi["installed-app-id"] = strconv.Itoa(installedAppId)
		queryParamsOfApi["env-id"] = environmentId
		PollForAppStatusInAppDetails(queryParamsOfApi, suite.authToken)
		installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree.Status)
		assert.Equal(suite.T(), installedAppId, installedAppDetails.Result.InstalledAppId)
		assert.Equal(suite.T(), "airflow", installedAppDetails.Result.AppStoreAppName)
		assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree.PodMetadata)
		assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree.Nodes)
		assert.NotNil(suite.T(), installedAppDetails.Result.ResourceTree.Hosts[0].ResourcesInfo)
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

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := AppStoreDeploymentRouter.HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}

func PollForAppStatusInAppDetails(queryParams map[string]string, authToken string) bool {
	count := 0
	for {
		respOfGetApplicationDetailApi := HitGetInstalledAppDetailsApi(queryParams, authToken)
		appStatus := respOfGetApplicationDetailApi.Result.ResourceTree.Status
		time.Sleep(1 * time.Second)
		count = count + 1
		if appStatus == "Healthy" || count >= 500 {
			break
		}
	}
	return true
}
