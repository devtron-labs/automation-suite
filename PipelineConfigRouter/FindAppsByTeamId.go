package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassFindAppsByTeamId() {

	suite.Run("A=1=FindAppsByValidTeamId", func() {
		log.Println("=== Here we are getting app-list before creating new app ===")
		appsByTeamId := HitFindAppsByTeamId("1", suite.authToken)
		noOfAppsBeforeCreationNewApp := len(appsByTeamId.Result)
		log.Println("=== Here we are creating new App ===")
		createAppApiResponse := Base.CreateApp(suite.authToken).Result
		log.Println("=== created App name is ===>", createAppApiResponse.AppName)
		log.Println("=== Here we are getting app-list after creating new app for asserting the API ===")
		appsByTeamId = HitFindAppsByTeamId("1", suite.authToken)
		noOfAppsAfterCreationNewApp := len(appsByTeamId.Result)
		assert.Equal(suite.T(), noOfAppsBeforeCreationNewApp+1, noOfAppsAfterCreationNewApp)
		assert.Equal(suite.T(), createAppApiResponse.AppName, appsByTeamId.Result[noOfAppsAfterCreationNewApp-1].Name)
		log.Println("=== Here we are deleting the newly created app after verification ===")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId)
		HitDeleteAppApi(byteValueOfDeleteApp, createAppApiResponse.Id, suite.authToken)
	})

	suite.Run("A=2=FindAppsByValidTeamId", func() {
		invalidTeamId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		log.Println("=== Here we are getting app-list via invalid TeamId ===")
		appsByTeamId := HitFindAppsByTeamId(invalidTeamId, suite.authToken)
		assert.Nil(suite.T(), appsByTeamId.Result)
	})
}
