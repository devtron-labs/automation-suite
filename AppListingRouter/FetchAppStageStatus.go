package AppListingRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *AppsListingRouterTestSuite) TestClassA1FetchAppStageStatus() {
	appId := suite.createAppResponseDto.Result.Id
	createAppApiResponse := suite.createAppResponseDto.Result

	suite.Run("A=1=FetchAllStageStatusWithValidAppId", func() {
		fetchAllLinkResponseDto := FetchAllStageStatus(appId, suite.authToken)

		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), 200, fetchAllLinkResponseDto.Code)
		assert.Equal(suite.T(), true, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Required)

	})
	suite.Run("A=2=FetchAllStageStatusWithInvalidAppId", func() {
		fetchAllLinkResponseDto := FetchAllStageStatus(Base.GetRandomNumberOf9Digit(), suite.authToken)

		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), 404, fetchAllLinkResponseDto.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", fetchAllLinkResponseDto.Errors[0].UserMessage)

	})
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
