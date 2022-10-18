package AttributesRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func (suite *AttributeRouterTestSuite) TestClassA1AddAttributes() {
	envConfig := Base.ReadBaseEnvConfig()
	baseCredentials := Base.ReadAnyJsonFile(envConfig.BaseCredentialsFile)
	attributesDTO := GetPayloadForAddAttributes(baseCredentials.BaseServerUrl)
	attributesBytePayload, _ := json.Marshal(attributesDTO)

	suite.Run("A=1=AddAttributesWithValidPayload", func() {
		ApiResp := HitAddAttributesApi(attributesBytePayload, suite.authToken)
		assert.Equal(suite.T(), baseCredentials.BaseServerUrl, ApiResp.Result.Value)
		queryParams := map[string]string{"key": "url"}
		attributesApiResp := HitGetAttributesApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), baseCredentials.BaseServerUrl, attributesApiResp.Result.Value)
		query := "delete from \"attributes\" where id =" + strconv.Itoa(ApiResp.Result.Id)
		Base.ConnectToDB("UpdateOrDeleteData", query)
	})
}

//todo need to check if we can add some more test cases
