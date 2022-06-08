package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassC5GetPipelineSuggestedCiCd() {
	appId := suite.createAppResponseDto.Result.Id

	suite.Run("A=1=GetPipelineSuggestedCDWithValidAppId", func() {
		pipelineSuggestedCDResponse := HitGetPipelineSuggestedCiCd("cd", appId, suite.authToken)
		log.Println("Validating the response of pipelineSuggestedCD API")
		assert.NotNil(suite.T(), pipelineSuggestedCDResponse.Result)
	})

	suite.Run("A=2=GetPipelineSuggestedCDWithInvalidAppId", func() {
		pipelineSuggestedCDResponse := HitGetPipelineSuggestedCiCd("cd", Base.GetRandomNumberOf9Digit(), suite.authToken)
		log.Println("Validating the response of GetPipelineSuggestedCD API")
		assert.Equal(suite.T(), 404, pipelineSuggestedCDResponse.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", pipelineSuggestedCDResponse.Errors[0].UserMessage)

	})
	suite.Run("A=3=GetPipelineSuggestedCIWithValidAppId", func() {
		pipelineSuggestedCDResponse := HitGetPipelineSuggestedCiCd("ci", appId, suite.authToken)
		log.Println("Validating the response of pipelineSuggestedCD API")
		assert.NotNil(suite.T(), pipelineSuggestedCDResponse.Result)
	})

	suite.Run("A=4=GetPipelineSuggestedCiWithInvalidAppId", func() {
		pipelineSuggestedCDResponse := HitGetPipelineSuggestedCiCd("ci", Base.GetRandomNumberOf9Digit(), suite.authToken)
		log.Println("Validating the response of GetPipelineSuggestedCD API")
		assert.Equal(suite.T(), 404, pipelineSuggestedCDResponse.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", pipelineSuggestedCDResponse.Errors[0].UserMessage)

	})
}
