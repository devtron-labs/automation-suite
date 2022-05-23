package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClass6GetApp() {
	appId := suite.createAppResponseDto.Result.Id

	suite.Run("A=1=FetchAppGetWithValidAppId", func() {
		fetchAppGetResponseDto := HitGetMaterial(appId, suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), appId, fetchAppGetResponseDto.Result.Id)

	})
	suite.Run("A=2=FetchAppGetWithInvalidAppId", func() {
		fetchAppGetResponseDto := HitGetMaterial(Base.GetRandomNumberOf9Digit(), suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), 404, fetchAppGetResponseDto.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", fetchAppGetResponseDto.Errors[0].UserMessage)

	})
}
