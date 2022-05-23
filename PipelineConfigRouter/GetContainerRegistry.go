package PipelineConfigRouter

import (
	"github.com/stretchr/testify/assert"
	"strconv"
)

// TestClass4GetContainerRegistry todo app ID is not required in URL for this API,I will remove this once dev will fix it
func (suite *PipelinesConfigRouterTestSuite) TestClass4GetContainerRegistry() {
	createAppApiResponse := suite.createAppResponseDto.Result
	appId := strconv.Itoa(createAppApiResponse.Id)
	suite.Run("A=1=GetContainerRegistryWithValidAppId", func() {
		getContainerRegistryResponse := HitGetContainerRegistry(appId, suite.authToken)
		indexOfLastResult := len(getContainerRegistryResponse.Result) - 1
		assert.True(suite.T(), getContainerRegistryResponse.Result[indexOfLastResult].IsDefault)
	})
}
