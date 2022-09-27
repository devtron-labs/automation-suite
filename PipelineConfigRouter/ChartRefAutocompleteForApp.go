package PipelineConfigRouter

import (
	"github.com/stretchr/testify/assert"
	"strconv"
)

// TestClassA7GetChartReference todo need to add one more test case  for ValidAppIdAfterSavingTemplate
// TestClassA7GetChartReference todo need to add test case of invalid App Id as well once issue fixed from dev side
func (suite *PipelinesConfigRouterTestSuite) TestClassA7GetChartReference() {
	createAppApiResponse := suite.createAppResponseDto.Result
	appId := strconv.Itoa(createAppApiResponse.Id)
	suite.Run("A=1=GetChartReferenceWithValidAppIdBeforeSavingTemplate", func() {
		getChartReferenceResponse := HitGetChartReferenceViaAppId(appId, suite.authToken)
		indexOfLastResult := len(getChartReferenceResponse.Result.ChartRefs) - 1
		assert.NotNil(suite.T(), getChartReferenceResponse.Result.ChartRefs[indexOfLastResult].Id)
		assert.NotNil(suite.T(), getChartReferenceResponse.Result.ChartRefs[indexOfLastResult].Version)
		assert.Equal(suite.T(), 0, getChartReferenceResponse.Result.LatestAppChartRef)
	})
}
