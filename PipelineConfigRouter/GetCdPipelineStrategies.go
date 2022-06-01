package PipelineConfigRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"strconv"
)

// TestClassC4GetCdPipelineStrategies todo will add assertion once workflow editor test cases will be completed
func (suite *PipelinesConfigRouterTestSuite) TestClassC4GetCdPipelineStrategies() {
	createAppApiResponse := suite.createAppResponseDto.Result
	appId := strconv.Itoa(createAppApiResponse.Id)
	suite.Run("A=1=GetCdPipelineStrategiesWithValidAppId", func() {
		cdPipelineStrategiesResponse := HitGetCdPipelineStrategies(appId, suite.authToken)
		assert.NotNil(suite.T(), cdPipelineStrategiesResponse)
	})

	suite.Run("A=2=GetCdPipelineStrategiesWithInvalidAppId", func() {
		invalidAppId := strconv.Itoa(testUtils.GetRandomNumberOf9Digit())
		cdPipelineStrategiesResponse := HitGetCdPipelineStrategies(invalidAppId, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", cdPipelineStrategiesResponse.Errors[0].UserMessage)
	})

	// todo need to proper error handling here instead of 500
	suite.Run("A=3=GetCdPipelineStrategiesWithoutSavingCI", func() {
		/*appNameForNewCreation := "app" + strings.ToLower(testUtils.GetRandomStringOfGivenLength(6))
		log.Println("=== Here We are creating a new App ===")
		createAppRequestDto := GetAppRequestDto(appNameForNewCreation, 1, 0)
		byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
		createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, appNameForNewCreation, 1, 0, suite.authToken)

		AppId := strconv.Itoa(createAppResponseDto.Result.Id)
		cdPipelineStrategiesResponse := HitGetCdPipelineStrategies(AppId, suite.authToken)
		assert.Equal(suite.T(), 500, cdPipelineStrategiesResponse.Code)

		log.Println("getting payload for Delete Team API")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
		log.Println("Hitting the Delete team API for Removing the data created via automation")
		HitDeleteAppApi(byteValueOfDeleteApp, createAppResponseDto.Result.Id, suite.authToken)
		*/
	})
}
