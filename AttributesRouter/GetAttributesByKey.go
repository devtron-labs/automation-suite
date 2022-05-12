package AttributesRouter

import (
	"github.com/stretchr/testify/assert"
)

func (suite AttributeRouterTestSuite) TestGetAttributesByKeyWithValidValue() {
	queryParams := map[string]string{"key": "url"}
	envConf, _ := GetEnvironmentConfigForHelmApp()
	attributesApiResp := HitGetAttributesApi(queryParams, suite.authToken)
	assert.Equal(suite.T(), envConf.ValueAttribute, attributesApiResp.Result.Value)
}

func (suite AttributeRouterTestSuite) TestGetAttributesByKeyWithInvalidValue() {
	queryParams := map[string]string{"key": "InvalidUrl"}
	attributesApiResp := HitGetAttributesApi(queryParams, suite.authToken)
	assert.Nil(suite.T(), attributesApiResp.Result)
}
