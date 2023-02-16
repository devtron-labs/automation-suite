package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassFindAppsByTeamName() {

	suite.Run("A=1=FindAppsByValidTeamName", func() {
		log.Println("=== Here we are getting app-list before creating new app ===")
		appsByTeamId := HitFindAppsByTeamName("devtron-demo", suite.authToken)
		noOfAppsBeforeCreationNewApp := len(appsByTeamId.Result)
		log.Println("=== Here we are creating new App ===")
		createAppApiResponse := Base.CreateApp(suite.authToken).Result
		log.Println("=== created App name is ===>", createAppApiResponse.AppName)
		log.Println("=== Here we are getting app-list after creating new app for asserting the API ===")
		appsByTeamId = HitFindAppsByTeamName("devtron-demo", suite.authToken)
		noOfAppsAfterCreationNewApp := len(appsByTeamId.Result)
		assert.Equal(suite.T(), noOfAppsBeforeCreationNewApp+1, noOfAppsAfterCreationNewApp)
		assert.Equal(suite.T(), createAppApiResponse.AppName, appsByTeamId.Result[noOfAppsAfterCreationNewApp-1].Name)
		log.Println("=== Here we are deleting the newly created app after verification ===")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId)
		HitDeleteAppApi(byteValueOfDeleteApp, createAppApiResponse.Id, suite.authToken)
	})

	suite.Run("A=2=FindAppsByValidTeamName", func() {
		invalidTeamId := "TeamName" + Base.GetRandomStringOfGivenLength(5)
		log.Println("=== Here we are getting app-list via invalid TeamName ===")
		appsByTeamId := HitFindAppsByTeamName(invalidTeamId, suite.authToken)
		assert.Nil(suite.T(), appsByTeamId.Result)
	})
}
