package AttributesRouter

import (
	"github.com/stretchr/testify/assert"
)

func (suite *AttributeRouterTestSuite) TestGetAttributesActiveList() {
	suite.Run("A=1=GetAttributesActiveList", func() {
		attributesApiResp := HitGetAttributesActiveListApi(suite.authToken)
		assert.NotNil(suite.T(), attributesApiResp.Result)
	})

}
