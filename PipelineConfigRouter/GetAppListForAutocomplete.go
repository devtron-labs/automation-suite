package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassGetAppListForAutocomplete() {

	suite.Run("A=1=GetAppListForAutocomplete", func() {
		log.Println("=== Here we are getting app-list before creating new app ===")
		getAppListForAutocompleteResponse := HitGetAppListForAutocomplete(suite.authToken)
		noOfAppsBeforeCreationNewApp := len(getAppListForAutocompleteResponse.Result)
		log.Println("=== Here we are creating new App ===")
		createAppApiResponse := Base.CreateApp(suite.authToken).Result
		log.Println("=== created App name is ===>", createAppApiResponse.AppName)
		log.Println("=== Here we are getting app-list after creating new app for asserting the API ===")
		getAppListForAutocompleteResponse = HitGetAppListForAutocomplete(suite.authToken)
		noOfAppsAfterCreationNewApp := len(getAppListForAutocompleteResponse.Result)
		assert.Equal(suite.T(), noOfAppsBeforeCreationNewApp+1, noOfAppsAfterCreationNewApp)
		assert.Equal(suite.T(), createAppApiResponse.AppName, getAppListForAutocompleteResponse.Result[noOfAppsAfterCreationNewApp-1].Name)
		log.Println("=== Here we are deleting the newly created app after verification ===")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId)
		HitDeleteAppApi(byteValueOfDeleteApp, createAppApiResponse.Id, suite.authToken)
	})
}
