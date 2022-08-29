package AppStoreDiscoverRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *AppStoreDiscoverTestSuite) TestGetTemplateValuesViaReferenceId() {
	log.Println("=== Here we are getting airflow chart repo ===")
	queryParams := map[string]string{"appStoreName": "apache"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)

	suite.Run("A=1=TemplateValuesViaCorrectReferenceId", func() {
		queryParamsOfApi := map[string]string{"referenceId": strconv.Itoa(ActiveDiscoveredApps.Result[0].AppStoreApplicationVersionId), "kind": "DEFAULT"}
		templateValues := HitGetTemplateValuesViaReferenceIdApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), ActiveDiscoveredApps.Result[0].AppStoreApplicationVersionId, templateValues.Result.Id)
		assert.NotNil(suite.T(), templateValues.Result.Values)
		assert.NotNil(suite.T(), templateValues.Result.ChartVersion)
	})

	suite.Run("A=2=TemplateValuesViaIncorrectReferenceId", func() {
		randomId := strconv.Itoa(testUtils.GetRandomNumberOf9Digit())
		queryParamsOfApi := map[string]string{"referenceId": randomId, "kind": "DEFAULT"}
		templateValues := HitGetTemplateValuesViaReferenceIdApi(queryParamsOfApi, suite.authToken)
		assert.Equal(suite.T(), 404, templateValues.Code)
	})
}
