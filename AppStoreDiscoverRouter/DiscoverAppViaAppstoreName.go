package AppStoreDiscoverRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestDiscoverHelmAppsViaAppstoreName() {
	chartRepoName := "airflow"

	suite.Run("A=1=DiscoverWithCorrectRepoName", func() {
		queryParams := map[string]string{"appStoreName": chartRepoName}
		PollForGettingHelmAppData(queryParams, suite.authToken)
		ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), chartRepoName, ActiveDiscoveredApps.Result[0].Name)
		assert.Equal(suite.T(), "bitnami", ActiveDiscoveredApps.Result[0].ChartName)
		assert.False(suite.T(), ActiveDiscoveredApps.Result[0].Deprecated)
	})

	suite.Run("A=2=DiscoverWithInCorrectRepoName", func() {
		randomAppstoreName := Base.GetRandomStringOfGivenLength(8)
		queryParams := map[string]string{"appStoreName": randomAppstoreName}
		time.Sleep(10 * time.Second)
		ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
		assert.Nil(suite.T(), ActiveDiscoveredApps.Result)
	})
}
