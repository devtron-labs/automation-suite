package PipelineConfigRouter

import (
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassC6GetEnvironmentAutocompleteDetails() {
	queryParams := map[string]string{"auth": "true"}
	suite.Run("A=1=GetEnvironmentDetailsWithAuthAsTrue", func() {
		allEnvironmentDetailsResponse := HitGetAllEnvironmentDetails(queryParams, suite.authToken)
		log.Println("Validating the response of GetAllEnvironmentDetails API")
		assert.NotNil(suite.T(), allEnvironmentDetailsResponse.Result)
		assert.Equal(suite.T(), 200, allEnvironmentDetailsResponse.Code)
		/*assert.Equal(suite.T(), 1, allEnvironmentDetailsResponse.Result[0].Id)
		assert.Equal(suite.T(), "devtron-demo", allEnvironmentDetailsResponse.Result[0].EnvironmentName)
		assert.Equal(suite.T(), "devtron-demo", allEnvironmentDetailsResponse.Result[0].Namespace)
		assert.Equal(suite.T(), "devtron-demo", allEnvironmentDetailsResponse.Result[0].EnvironmentIdentifier)
		*/
	})

	// todo will enable this test case once bug will fix from dev's side
	/*suite.Run("A=2=GetEnvironmentDetailsWithAuthAsFalse", func() {
		pipelineSuggestedCDResponse := HitGetMaterial(Base.GetRandomNumberOf9Digit(), suite.authToken)
		log.Println("Validating the response of GetPipelineSuggestedCD API")
		assert.Equal(suite.T(), 404, pipelineSuggestedCDResponse.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", pipelineSuggestedCDResponse.Errors[0].UserMessage)

	})*/
}
