package AppListingRouter

import (
	"automation-suite/testUtils"
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *AppsListingRouterTestSuite) TestClassA2FetchOtherEnvironment() {
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appId := createAppApiResponse.Id

	suite.Run("A=1=FetchOtherEnvWithValidAppId", func() {
		fetchOtherEnvResponseDto := FetchOtherEnv(appId, suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), 200, fetchOtherEnvResponseDto.Code)
	})
	suite.Run("A=2=FetchOtherEnvWithInvalidAppId", func() {
		fetchOtherEnvResponseDto := FetchOtherEnv(testUtils.GetRandomNumberOf9Digit(), suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), 404, fetchOtherEnvResponseDto.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", fetchOtherEnvResponseDto.Errors[0].UserMessage)

	})
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
