package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *AppStoreDiscoverTestSuite) TestDeleteTemplateValuesApi() {
	log.Println("=== Here we are getting airflow chart repo ===")
	queryParams := map[string]string{"appStoreName": "airflow"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	DiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
	log.Println("=== Here we are saving template values ===")
	appName := "automation-preset-" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
	valueForPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/SaveTemplateValuesRequestPayload.txt")
	requestPayload := getPayloadForSaveTemplateValues(appName, string(valueForPayload), DiscoveredApps.Result[0].AppStoreApplicationVersionId)
	payloadByteArray, _ := json.Marshal(requestPayload)
	responseOfSaveTemplateApi := HitSaveTemplateValuesApi(string(payloadByteArray), suite.authToken)

	suite.Run("A=1=DeleteTemplateValuesWithValidId", func() {
		log.Println("=== Here We are getting noOfTemplate before deleting it ===")
		ApplicationValuesList := HitGetApplicationValuesListApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		noOfDeployedCharts := len(ApplicationValuesList.Result.Values[1].Values)
		deleteTemplateValueApiResponse := HitDeleteTemplateValuesApi(strconv.Itoa(responseOfSaveTemplateApi.Result.Id), suite.authToken)
		assert.True(suite.T(), deleteTemplateValueApiResponse.Result)
		log.Println("=== Here We are getting noOfTemplate after deleting it ===")
		ApplicationValuesList = HitGetApplicationValuesListApi(strconv.Itoa(DiscoveredApps.Result[0].Id), suite.authToken)
		assert.Equal(suite.T(), noOfDeployedCharts-1, len(ApplicationValuesList.Result.Values[1].Values))
	})

	suite.Run("A=2=DeleteTemplateValuesWithInvalidId", func() {
		randomId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		deleteTemplateValueApiResponse := HitDeleteTemplateValuesApi(randomId, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", deleteTemplateValueApiResponse.Errors[0].UserMessage)
	})
}
