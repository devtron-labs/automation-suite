package PipelineConfigRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassA8GetAppTemplate() {
	createAppApiResponse := suite.createAppResponseDto.Result
	appId := strconv.Itoa(createAppApiResponse.Id)
	latestChartRef := testUtils.ReadDataByFilenameAndKey("OutputDataGetChartReferenceViaAppId", "latestChartRef")
	suite.Run("A=1=GetTemplateViaValidArgs", func() {
		getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(appId, latestChartRef[0], suite.authToken)
		assert.NotNil(suite.T(), getTemplateResponse.Result.GlobalConfig.DefaultAppOverride)
	})

	suite.Run("A=2=GetTemplateViaInvalidChartRefId", func() {
		invalidChartRefId := strconv.Itoa(testUtils.GetRandomNumberOf9Digit())
		getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(appId, invalidChartRefId, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", getTemplateResponse.Errors[0].UserMessage)
	})
	testUtils.DeleteFile("OutputDataGetChartReferenceViaAppId")
}

//todo need to add one more case for invalid AppId as well once dev will fix the issue for invalid app-id
