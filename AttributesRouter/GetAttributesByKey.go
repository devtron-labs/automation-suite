package AttributesRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
)

func (suite *AttributeRouterTestSuite) TestGetAttributesByKey() {
	envConf, _ := Base.GetEnvironmentConfig()
	suite.Run("A=1=AttributesWithValidValueOfKey", func() {
		queryParams := map[string]string{"key": "url"}
		attributesApiResp := HitGetAttributesApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), envConf.BaseServerUrl, attributesApiResp.Result.Value)
	})

	suite.Run("A=2=AttributesWithInvalidValueOfKey", func() {
		queryParams := map[string]string{"key": "InvalidUrl"}
		attributesApiResp := HitGetAttributesApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), "", attributesApiResp.Result.Key)
	})
}
