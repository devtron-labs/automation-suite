package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassA6GetApp() {
	createAppApiResponse := Base.CreateApp(suite.authToken).Result

	suite.Run("A=1=FetchAppWithValidAppId", func() {
		fetchAppGetResponseDto := HitGetMaterial(createAppApiResponse.Id, suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), createAppApiResponse.Id, fetchAppGetResponseDto.Result.Id)

	})
	suite.Run("A=2=FetchAppWithInvalidAppId", func() {
		fetchAppGetResponseDto := HitGetMaterial(Base.GetRandomNumberOf9Digit(), suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), 404, fetchAppGetResponseDto.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", fetchAppGetResponseDto.Errors[0].UserMessage)

	})

	log.Println("=== Here we are Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
