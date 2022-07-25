package AppStoreDiscoverRouter

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

func (suite *AppStoreDiscoverTestSuite) TestGetApplicationValuesList() {
	log.Println("=== Here we are getting airflow chart repo ===")
	queryParams := map[string]string{"appStoreName": "airflow"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	DiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)

	suite.Run("A=1=GetValuesListForKindDeploy", func() {
		log.Println("=== Here We are getting noOfDeployedCharts before new deployment ===")
		ApplicationValuesList := HitGetApplicationValuesListApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		noOfDeployedCharts := len(ApplicationValuesList.Result.Values[2].Values)
		log.Println("=== Here We are installing Helm chart from chart-store ===")
		expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
		log.Println("=== Here we are installing the app ===")
		installAppRequestDTO := RequestDTOs.InstallAppRequestDTO{}
		json.Unmarshal(expectedPayload, &installAppRequestDTO)
		installAppRequestDTO.AppName = "deepak-helm-airflow" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		requestPayload, _ := json.Marshal(installAppRequestDTO)
		responseAfterInstallingApp := AppStoreDeploymentRouter.HitInstallAppApi(string(requestPayload), suite.authToken)
		time.Sleep(2 * time.Second)
		appName := responseAfterInstallingApp.Result.AppName
		log.Println("=== Here We are getting noOfDeployedCharts after new deployment ===")
		ApplicationValuesList = HitGetApplicationValuesListApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		assert.Equal(suite.T(), "DEFAULT", ApplicationValuesList.Result.Values[0].Kind)
		assert.Equal(suite.T(), "TEMPLATE", ApplicationValuesList.Result.Values[1].Kind)
		indexOfLastDeployed := len(ApplicationValuesList.Result.Values[2].Values)
		assert.Equal(suite.T(), appName, ApplicationValuesList.Result.Values[2].Values[indexOfLastDeployed-1].Name)
		assert.Equal(suite.T(), noOfDeployedCharts+1, len(ApplicationValuesList.Result.Values[2].Values))
		assert.Equal(suite.T(), "EXISTING", ApplicationValuesList.Result.Values[3].Kind)

		log.Println("Removing the data created via API")
		respOfDeleteInstallAppApi := AppStoreDeploymentRouter.HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
	})

	suite.Run("A=2=GetValuesListForKindDefault", func() {
		deploymentOfInstalledApp := HitGetApplicationValuesListApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		assert.True(suite.T(), len(deploymentOfInstalledApp.Result.Values[0].Values) >= 1)
	})

	suite.Run("A=3=GetValuesListForKindTemplate", func() {
		log.Println("=== Here we are saving template values   ===")
		appName := "automation-preset-" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		valueForPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/SaveTemplateValuesRequestPayload.txt")
		requestPayload := getPayloadForSaveTemplateValues(appName, string(valueForPayload), DiscoveredApps.Result[0].AppStoreApplicationVersionId)
		payloadByteArray, _ := json.Marshal(requestPayload)
		responseOfSaveTemplateApi := HitSaveTemplateValuesApi(string(payloadByteArray), suite.authToken)
		log.Println("=== Here We are getting noOfTemplate before deleting it ===")
		ApplicationValuesList := HitGetApplicationValuesListApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		noOfDeployedCharts := len(ApplicationValuesList.Result.Values[1].Values)
		deleteTemplateValueApiResponse := HitDeleteTemplateValuesApi(strconv.Itoa(responseOfSaveTemplateApi.Result.Id), suite.authToken)
		assert.True(suite.T(), deleteTemplateValueApiResponse.Result)
		log.Println("=== Here We are getting noOfTemplate after deleting it ===")
		ApplicationValuesList = HitGetApplicationValuesListApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		assert.Equal(suite.T(), noOfDeployedCharts-1, len(ApplicationValuesList.Result.Values[1].Values))
	})

	suite.Run("A=4=GetValuesListForKindExisting", func() {

	})

	suite.Run("A=5=GetListWithInvalidAppId", func() {
		randomAppId := Base.GetRandomNumberOf9Digit()
		appValuesList := HitGetApplicationValuesListApi(strconv.Itoa(randomAppId), suite.authToken)
		assert.Nil(suite.T(), appValuesList.Result.Values[0].Values)
		assert.Nil(suite.T(), appValuesList.Result.Values[1].Values)
		assert.Equal(suite.T(), 0, len(appValuesList.Result.Values[2].Values))
		assert.Equal(suite.T(), 0, len(appValuesList.Result.Values[3].Values))
	})
}

//todo need to add a test case around GetValuesListForKindExisting after understanding the flow of this
