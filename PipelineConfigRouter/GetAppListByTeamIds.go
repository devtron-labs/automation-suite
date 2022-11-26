package PipelineConfigRouter

import (
	"automation-suite/AppStoreDiscoverRouter"
	"automation-suite/AppStoreDiscoverRouter/RequestDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassGetAppListByTeamIds() {

	suite.Run("A=1=GetDevtronAppListByTeamIds", func() {
		log.Println("=== Here we are getting app-list before creating new app ===")
		queryParams := make(map[string]string)
		queryParams["teamIds"] = "1"
		getAppListForAutocompleteResponse := HitGetAppListByTeamIds(queryParams, suite.authToken)
		noOfAppsBeforeCreationNewApp := len(getAppListForAutocompleteResponse.Result[0].AppList)
		log.Println("=== Here we are creating new App ===")
		createAppApiResponse := Base.CreateApp(suite.authToken).Result
		log.Println("=== created App name is ===>", createAppApiResponse.AppName)
		log.Println("=== Here we are getting app-list after creating new app for asserting the API ===")
		getAppListForAutocompleteResponse = HitGetAppListByTeamIds(queryParams, suite.authToken)
		noOfAppsAfterCreationNewApp := len(getAppListForAutocompleteResponse.Result[0].AppList)
		assert.Equal(suite.T(), noOfAppsBeforeCreationNewApp+1, noOfAppsAfterCreationNewApp)
		assert.Equal(suite.T(), createAppApiResponse.AppName, getAppListForAutocompleteResponse.Result[0].AppList[noOfAppsAfterCreationNewApp-1].Name)
		log.Println("=== Here we are deleting the newly created app after verification ===")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId)
		HitDeleteAppApi(byteValueOfDeleteApp, createAppApiResponse.Id, suite.authToken)
	})

	suite.Run("A=2=GetHelmAppListByTeamIds", func() {
		log.Println("=== Here we are getting app-list before creating new app ===")
		queryParams := make(map[string]string)
		queryParams["teamIds"] = "1"
		queryParams["appType"] = "DevtronChart"
		getAppListForAutocompleteResponse := HitGetAppListByTeamIds(queryParams, suite.authToken)
		var noOfAppsBeforeCreationNewApp int
		if len(getAppListForAutocompleteResponse.Result) != 0 {
			noOfAppsBeforeCreationNewApp = len(getAppListForAutocompleteResponse.Result[0].AppList)
		}

		log.Println("=== Here We are installing Helm chart from chart-store ===")
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
		log.Println("Hitting the InstallAppApi with valid payload")
		installAppRequestDTO := RequestDTOs.InstallAppRequestDTO{}
		json.Unmarshal(expectedPayload, &installAppRequestDTO)

		AppName := "automation" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		log.Println("=====Helm AppName used in this test Case is====", AppName)
		installAppRequestDTO.AppName = AppName
		requestPayload, _ := json.Marshal(installAppRequestDTO)
		responseAfterInstallingApp := AppStoreDiscoverRouter.HitInstallAppApi(string(requestPayload), suite.authToken)
		time.Sleep(2 * time.Second)
		installedAppId := responseAfterInstallingApp.Result.InstalledAppId

		log.Println("=== Here we are getting app-list after creating new app for asserting the API ===")
		getAppListForAutocompleteResponse = HitGetAppListByTeamIds(queryParams, suite.authToken)

		if noOfAppsBeforeCreationNewApp == 0 {
			assert.Equal(suite.T(), len(getAppListForAutocompleteResponse.Result[0].AppList), 1)
			assert.Equal(suite.T(), getAppListForAutocompleteResponse.Result[0].AppList[0].Name, AppName)
		} else if noOfAppsBeforeCreationNewApp > 0 {
			index := len(getAppListForAutocompleteResponse.Result[0].AppList) - 1
			assert.Equal(suite.T(), noOfAppsBeforeCreationNewApp+1, len(getAppListForAutocompleteResponse.Result[0].AppList))
			assert.Equal(suite.T(), getAppListForAutocompleteResponse.Result[0].AppList[index].Name, AppName)
		}
		log.Println("=== Here we are deleting the newly created app after verification ===")
		log.Println("Removing the data created via API")
		AppStoreDiscoverRouter.HitDeleteInstalledAppApi(strconv.Itoa(installedAppId), suite.authToken)
	})

	suite.Run("A=3=GetDevtronAppListByInvalidTeamId", func() {
		log.Println("=== Here we are getting app-list before creating new app ===")
		invalidTeamId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		queryParams := make(map[string]string)
		queryParams["teamIds"] = invalidTeamId
		getAppListForAutocompleteResponse := HitGetAppListByTeamIds(queryParams, suite.authToken)
		assert.Equal(suite.T(), len(getAppListForAutocompleteResponse.Result), 0)
	})

	suite.Run("A=4=GetHelmAppListByInvalidTeamId", func() {
		log.Println("=== Here we are getting app-list before creating new app ===")
		invalidTeamId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		queryParams := make(map[string]string)
		queryParams["teamIds"] = invalidTeamId
		queryParams["appType"] = "DevtronChart"
		getAppListForAutocompleteResponse := HitGetAppListByTeamIds(queryParams, suite.authToken)
		assert.Equal(suite.T(), len(getAppListForAutocompleteResponse.Result), 0)
	})
}
