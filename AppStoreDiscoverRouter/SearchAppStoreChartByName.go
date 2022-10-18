package AppStoreDiscoverRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func (suite *AppStoreDiscoverTestSuite) TestSearchAppStoreChartByName() {
	queryParams := map[string]string{"chartName": "airflow"}

	suite.Run("A=1=SearchAppStoreChartByValidName", func() {
		appStoreChart := HitSearchAppStoreChartByNameApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), "bitnami", appStoreChart.Result[0].ChartRepoName)
		assert.NotNil(suite.T(), appStoreChart.Result[0].AppStoreApplicationVersionId)
		assert.NotNil(suite.T(), appStoreChart.Result[0].ChartId)
		assert.NotNil(suite.T(), appStoreChart.Result[0].ChartRepoId)
	})

	suite.Run("A=2=SearchAppStoreChartByInvalidName", func() {
		randomName := strconv.Itoa(testUtils.GetRandomNumberOf9Digit())
		queryParams = map[string]string{"chartName": randomName}
		appStoreChart := HitSearchAppStoreChartByNameApi(queryParams, suite.authToken)
		assert.Nil(suite.T(), appStoreChart.Result)
	})

	suite.Run("A=3=SearchAppStoreChartByMatchingString", func() {
		queryParams = map[string]string{"chartName": "air"}
		appStoreChart := HitSearchAppStoreChartByNameApi(queryParams, suite.authToken)
		assert.True(suite.T(), len(appStoreChart.Result) > 1)
	})
}
