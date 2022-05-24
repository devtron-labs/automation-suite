package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassC5GetPipelineSuggestedCD() {
	appId := suite.createAppResponseDto.Result.Id

	suite.Run("A=1=GetPipelineSuggestedCDWithValidAppId", func() {
		pipelineSuggestedCDResponse := HitGetPipelineSuggestedCD(strconv.Itoa(appId), suite.authToken)
		log.Println("Validating the response of pipelineSuggestedCD API")
		assert.NotNil(suite.T(), pipelineSuggestedCDResponse.Result)
	})

	suite.Run("A=2=GetPipelineSuggestedCDWithInvalidAppId", func() {
		pipelineSuggestedCDResponse := HitGetMaterial(Base.GetRandomNumberOf9Digit(), suite.authToken)
		log.Println("Validating the response of GetPipelineSuggestedCD API")
		assert.Equal(suite.T(), 404, pipelineSuggestedCDResponse.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", pipelineSuggestedCDResponse.Errors[0].UserMessage)

	})
}
