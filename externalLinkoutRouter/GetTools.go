package externalLinkoutRouter

import (
	"log"

	"github.com/stretchr/testify/assert"
)

func (suite *LinkOutRouterTestSuite) TestClassA4GetTools() {
	suite.Run("A=1=GetAllTools", func() {
		log.Println("Hitting the FetchAllTool API again for verifying the functionality of it")
		fetchAllToolsResponseDto := HitFetchAllToolsApi(suite.authToken)
		log.Println("Validating the response of FetchAllTool API")
		assert.Equal(suite.T(), 200, fetchAllToolsResponseDto.Code)
		assert.NotNil(suite.T(), fetchAllToolsResponseDto.Result[0].Id)
		assert.NotNil(suite.T(), fetchAllToolsResponseDto.Result)
	})
}
