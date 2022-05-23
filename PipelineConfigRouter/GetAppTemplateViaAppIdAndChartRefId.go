package PipelineConfigRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func (suite *PipelinesConfigRouterTestSuite) TestClass8GetAppTemplate() {
	createAppApiResponse := suite.createAppResponseDto.Result
	appId := strconv.Itoa(createAppApiResponse.Id)
	latestChartRef := testUtils.ReadDataByFilenameAndKey("OutputDataGetChartReferenceViaAppId", "latestChartRef")
	suite.Run("A=1=GetTemplateViaValidArgs", func() {
		getChartReferenceResponse := HitGetTemplateViaAppIdAndChartRefId(appId, latestChartRef[0], suite.authToken)
		indexOfLastResult := len(getChartReferenceResponse.Result.ChartRefs) - 1
		assert.NotNil(suite.T(), getChartReferenceResponse.Result.ChartRefs[indexOfLastResult].Id)
	})
}
