package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *AppStoreDiscoverTestSuite) TestCheckAppExists() {
	responseAfterInstallingApp, _, _, _ := CreateHelmApp(suite.authToken)
	installedAppId := responseAfterInstallingApp.Result.InstalledAppId
	environmentId := strconv.Itoa(responseAfterInstallingApp.Result.EnvironmentId)
	AppName := responseAfterInstallingApp.Result.AppName
	log.Println("=== Here we are installing helm App ===")
	queryParamsOfApi := make(map[string]string)
	queryParamsOfApi["installed-app-id"] = strconv.Itoa(installedAppId)
	queryParamsOfApi["env-id"] = environmentId
	PollForAppStatusInAppDetails(queryParamsOfApi, suite.authToken)
	installedAppDetails := HitGetInstalledAppDetailsApi(queryParamsOfApi, suite.authToken)
	assert.Equal(suite.T(), "Healthy", installedAppDetails.Result.ResourceTree["status"].(string))

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

	//log.Println("Removing the data created via API")
	//respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	//assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
