package AppStoreDiscoverRouter

import (
	"github.com/stretchr/testify/assert"
)

func (suite *AppStoreDiscoverTestSuite) TestDiscoverHelmApps() {
	noOfNonDeprecatedApp := 0

	suite.Run("A=1=DiscoverNonDeprecatedAppOnly", func() {
		queryParams := map[string]string{"includeDeprecated": "0"}
		ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
		noOfNonDeprecatedApp = len(ActiveDiscoveredApps.Result)
		assert.NotNil(suite.T(), ActiveDiscoveredApps.Result[0].Id)
		assert.False(suite.T(), ActiveDiscoveredApps.Result[0].Deprecated)
	})

	suite.Run("A=2=DiscoverAppIncludingDeprecatedApps", func() {
		queryParams := map[string]string{"includeDeprecated": "1"}
		AllDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
		assert.True(suite.T(), noOfNonDeprecatedApp <= len(AllDiscoveredApps.Result))
	})
}
