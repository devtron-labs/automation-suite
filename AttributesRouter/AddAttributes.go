package AttributesRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func (suite *AttributeRouterTestSuite) TestClassA1AddAttributes() {
	fileData := Base.ReadAnyJsonFile("../testUtils/credentials.json")
	attributesDTO := GetPayloadForAddAttributes(fileData.BaseServerUrl)
	attributesBytePayload, _ := json.Marshal(attributesDTO)

	suite.Run("A=1=AddAttributesWithValidPayload", func() {
		ApiResp := HitAddAttributesApi(attributesBytePayload, suite.authToken)
		assert.Equal(suite.T(), fileData.BaseServerUrl, ApiResp.Result.Value)
		queryParams := map[string]string{"key": "url"}
		attributesApiResp := HitGetAttributesApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), fileData.BaseServerUrl, attributesApiResp.Result.Value)
	})
}

//todo need to check if we can add some more test cases
