package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *AppStoreDiscoverTestSuite) TestSaveTemplateValuesApi() {
	log.Println("=== Here we are getting airflow chart repo ===")
	queryParams := map[string]string{"appStoreName": "airflow"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	DiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)

	suite.Run("A=1=SaveTemplateValuesWithValidPayload", func() {
		appName := "automation-preset-" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		valueForPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/SaveTemplateValuesRequestPayload.txt")
		requestPayload := getPayloadForSaveTemplateValues(appName, string(valueForPayload), DiscoveredApps.Result[0].AppStoreApplicationVersionId)
		payloadByteArray, _ := json.Marshal(requestPayload)
		responseOfSaveTemplateApi := HitSaveTemplateValuesApi(string(payloadByteArray), suite.authToken)
		assert.Equal(suite.T(), string(valueForPayload), responseOfSaveTemplateApi.Result.Values)
		assert.Equal(suite.T(), appName, responseOfSaveTemplateApi.Result.Name)
		log.Println("===   Here We are Deleting template after verification   ===")
		deleteTemplateValueApiResponse := HitDeleteTemplateValuesApi(strconv.Itoa(responseOfSaveTemplateApi.Result.Id), suite.authToken)
		assert.True(suite.T(), deleteTemplateValueApiResponse.Result)
	})

	suite.Run("A=2=SaveTemplateWithRandomStringInValue", func() {
		appName := "automation-preset-" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		values := "This is a random string" + Base.GetRandomStringOfGivenLength(20)
		requestPayload := getPayloadForSaveTemplateValues(appName, values, DiscoveredApps.Result[0].AppStoreApplicationVersionId)
		payloadByteArray, _ := json.Marshal(requestPayload)
		responseOfSaveTemplateApi := HitSaveTemplateValuesApi(string(payloadByteArray), suite.authToken)
		assert.Equal(suite.T(), values, responseOfSaveTemplateApi.Result.Values)
		assert.Equal(suite.T(), appName, responseOfSaveTemplateApi.Result.Name)
	})

	suite.Run("A=3=SaveTemplateWithRandomAppStoreApplicationVersionId", func() {
		appName := "automation-preset-" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		valueForPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/SaveTemplateValuesRequestPayload.txt")
		requestPayload := getPayloadForSaveTemplateValues(appName, string(valueForPayload), Base.GetRandomNumberOf9Digit())
		payloadByteArray, _ := json.Marshal(requestPayload)
		responseOfSaveTemplateApi := HitSaveTemplateValuesApi(string(payloadByteArray), suite.authToken)
		assert.Equal(suite.T(), 500, responseOfSaveTemplateApi.Code)
	})

}

//todo need to add test cases for appStoreVersionId and empty string for name in payload after bug fix from dev-side
//todo need to fix A=3 test case after fixing the issue from dev side
