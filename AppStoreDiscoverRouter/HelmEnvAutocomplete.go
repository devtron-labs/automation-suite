package AppStoreDiscoverRouter

import (
	"github.com/stretchr/testify/assert"
)

func (suite *AppStoreDiscoverTestSuite) TestHelmEnvAutocomplete() {

	suite.Run("A=1=GetHelmEnv", func() {
		HelmEnvAndClusters := HitHelmEnvAutocompleteApi(suite.authToken)
		assert.NotNil(suite.T(), HelmEnvAndClusters.Result)
		assert.NotNil(suite.T(), HelmEnvAndClusters.Result[0].Environments)
	})
}

//todo will make this test case more generic after automating the AddClusterAndEnv API and bug fix of this API #2001
