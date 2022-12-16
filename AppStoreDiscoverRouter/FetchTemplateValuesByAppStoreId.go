package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestGetTemplateValuesList() {

	responseAfterInstallingApp, _, ActiveDiscoveredApps, appStoreId := CreateHelmApp(suite.authToken)
	//appStoreId := responseAfterInstallingApp.Result.AppStoreId
	suite.Run("A=1=GetValuesListForKindDeploy", func() {
		log.Println("=== Here We are getting noOfDeployedCharts after new deployment ===")
		ApplicationValuesList := HitGetApplicationValuesListApi(strconv.Itoa(appStoreId), suite.authToken)
		noOfDeployedCharts := len(ApplicationValuesList.Result.Values[2].Values)
		time.Sleep(5 * time.Second)
		appName := responseAfterInstallingApp.Result.AppName
		ApplicationValuesList = HitGetApplicationValuesListApi(strconv.Itoa(appStoreId), suite.authToken)
		indexOfLastDeployed := 0
		latestIdOfValue := 0
		for i, val := range ApplicationValuesList.Result.Values[2].Values {
			maxIndex := int(math.Max(float64(val.Id), float64(latestIdOfValue)))
			if maxIndex > latestIdOfValue {
				latestIdOfValue = maxIndex
				indexOfLastDeployed = i
			}
		}
		assert.Equal(suite.T(), appName, ApplicationValuesList.Result.Values[2].Values[indexOfLastDeployed].Name)
		assert.Equal(suite.T(), noOfDeployedCharts, len(ApplicationValuesList.Result.Values[2].Values))

		//log.Println("Removing the data created via API")
		//respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
		//assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
	})

	suite.Run("A=2=GetValuesListForKindDefault", func() {
		deploymentOfInstalledApp := HitGetApplicationValuesListApi(strconv.Itoa(appStoreId), suite.authToken)
		assert.True(suite.T(), len(deploymentOfInstalledApp.Result.Values[0].Values) >= 1)
	})

	suite.Run("A=3=GetValuesListForKindTemplate", func() {
		log.Println("=== Here we are saving template values   ===")
		appName := "automation-preset-" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		valueForPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/SaveTemplateValuesRequestPayload.txt")
		requestPayload := getPayloadForSaveTemplateValues(appName, string(valueForPayload), ActiveDiscoveredApps.Result[0].AppStoreApplicationVersionId)
		payloadByteArray, _ := json.Marshal(requestPayload)
		responseOfSaveTemplateApi := HitSaveTemplateValuesApi(string(payloadByteArray), suite.authToken)
		log.Println("=== Here We are getting noOfTemplate before deleting it ===")
		ApplicationValuesList := HitGetApplicationValuesListApi(strconv.Itoa(appStoreId), suite.authToken)
		noOfDeployedCharts := len(ApplicationValuesList.Result.Values[1].Values)
		deleteTemplateValueApiResponse := HitDeleteTemplateValuesApi(strconv.Itoa(responseOfSaveTemplateApi.Result.Id), suite.authToken)
		assert.True(suite.T(), deleteTemplateValueApiResponse.Result)
		log.Println("=== Here We are getting noOfTemplate after deleting it ===")
		ApplicationValuesList = HitGetApplicationValuesListApi(strconv.Itoa(appStoreId), suite.authToken)
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
