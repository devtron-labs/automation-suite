package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

// TestClassB6GetContainerRegistry  todo app ID is not required in URL for this API,I will remove this once dev will fix it
func (suite *PipelinesConfigRouterTestSuite) TestClassB6GetContainerRegistry() {
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appId := strconv.Itoa(createAppApiResponse.Id)

	suite.Run("A=1=GetContainerRegistryWithValidAppId", func() {
		getContainerRegistryResponse := HitGetContainerRegistry(appId, suite.authToken)
		assert.NotNil(suite.T(), getContainerRegistryResponse.Result[0].RegistryUrl)
	})

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId)

	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfDeleteApp, createAppApiResponse.Id, suite.authToken)
}
