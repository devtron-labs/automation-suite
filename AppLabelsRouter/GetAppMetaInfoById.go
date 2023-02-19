package AppLabelsRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *AppLabelRouterTestSuite) TestGetAppMetaInfo() {
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result

	suite.Run("A=1=MetaInfoWithCorrectAppId", func() {
		appMetaInfo := HitGetAppMetaInfoByIdApi(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
		assert.NotNil(suite.T(), appMetaInfo.Result.CreatedOn)
		assert.True(suite.T(), appMetaInfo.Result.Active)
	})

	suite.Run("A=2=MetaInfoWithIncorrectAppId", func() {
		randomNumber := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		appMetaInfo := HitGetAppMetaInfoByIdApi(randomNumber, suite.authToken)
		assert.Equal(suite.T(), appMetaInfo.Errors[0].UserMessage, "pg: no rows in result set")
	})

	log.Println("=== Here we are Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
