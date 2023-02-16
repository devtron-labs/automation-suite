package AppStoreDiscoverRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *AppStoreDiscoverTestSuite) TestVersionsAutocomplete() {
	log.Println("=== Here we are getting airflow chart repo ===")
	queryParams := map[string]string{"appStoreName": "airflow"}
	PollForGettingHelmAppData(queryParams, suite.authToken)
	ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)

	suite.Run("A=1=GetVersionsWithValidRepoId", func() {
		appVersions := GetAppVersionsAutocomplete(strconv.Itoa(ActiveDiscoveredApps.Result[0].Id), suite.authToken)
		assert.NotNil(suite.T(), appVersions.Result)
	})

	suite.Run("A=2=GetVersionsWithInvalidRepoId", func() {
		randomId := strconv.Itoa(testUtils.GetRandomNumberOf9Digit())
		appVersions := GetAppVersionsAutocomplete(randomId, suite.authToken)
		assert.Nil(suite.T(), appVersions.Result)
	})
}
