package AppStoreDiscoverRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *AppStoreDiscoverTestSuite) TestDiscoverAppViaAppStoreApplicationVersionId() {
	log.Println("=== Here we are getting memcached chart repo ===")
	queryParams := map[string]string{"appStoreName": "airflow"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)

	suite.Run("A=1=DiscoverAppWithValidId", func() {
		helmApp := DiscoverAppViaAppStoreApplicationVersionId(strconv.Itoa(ActiveDiscoveredApps.Result[0].AppStoreApplicationVersionId), suite.authToken)
		assert.Equal(suite.T(), ActiveDiscoveredApps.Result[0].AppStoreApplicationVersionId, helmApp.Result.Id)
		assert.NotEqual(suite.T(), "Airflow", helmApp.Result.AppStoreApplicationName)
		assert.Equal(suite.T(), ActiveDiscoveredApps.Result[0].Id, helmApp.Result.AppStoreId)
	})

	suite.Run("A=2=DiscoverAppWithInvalidId", func() {
		randomId := strconv.Itoa(testUtils.GetRandomNumberOf9Digit())
		helmApp := DiscoverAppViaAppStoreApplicationVersionId(randomId, suite.authToken)
		assert.Equal(suite.T(), 0, helmApp.Code)
	})
}
