package AppStoreRouter

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

func (suite *AppStoreTestSuite) TestCheckAppExists() {
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
	installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
	assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree.Status)

	suite.Run("A=1=TestApiWithNonExistingNameOnly", func() {
		listOfNames := []string{Base.GetRandomStringOfGivenLength(9)}
		checkAppExistsRequestDTO := getCheckAppExistsApi(listOfNames)
		payloadForCheckAppExistsRequestDTO, _ := json.Marshal(checkAppExistsRequestDTO)
		CheckAppExistsApiResponse := HitCheckAppExistsOrNot(string(payloadForCheckAppExistsRequestDTO), suite.authToken)
		assert.False(suite.T(), CheckAppExistsApiResponse.Result[0].Exists)
	})

	suite.Run("A=2=TestApiWithExistingNameOnly", func() {
		listOfNames := []string{AppName}
		checkAppExistsRequestDTO := getCheckAppExistsApi(listOfNames)
		payloadForCheckAppExistsRequestDTO, _ := json.Marshal(checkAppExistsRequestDTO)
		CheckAppExistsApiResponse := HitCheckAppExistsOrNot(string(payloadForCheckAppExistsRequestDTO), suite.authToken)
		assert.True(suite.T(), CheckAppExistsApiResponse.Result[0].Exists)
	})

	suite.Run("A=3=TestApiWithExistingAndNonExistingNames", func() {
		listOfNames := []string{Base.GetRandomStringOfGivenLength(9), AppName}
		checkAppExistsRequestDTO := getCheckAppExistsApi(listOfNames)
		payloadForCheckAppExistsRequestDTO, _ := json.Marshal(checkAppExistsRequestDTO)
		CheckAppExistsApiResponse := HitCheckAppExistsOrNot(string(payloadForCheckAppExistsRequestDTO), suite.authToken)
		assert.False(suite.T(), CheckAppExistsApiResponse.Result[0].Exists)
		assert.True(suite.T(), CheckAppExistsApiResponse.Result[1].Exists)
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := AppStoreDeploymentRouter.HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
