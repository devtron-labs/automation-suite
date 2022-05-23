package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *PipelinesConfigRouterTestSuite) TestClass1CreateApp() {
	appName := "app" + strings.ToLower(Base.GetRandomStringOfGivenLength(10))

	suite.Run("A=1=CreateAppWithValidPayload", func() {
		createAppRequestDto := GetAppRequestDto(appName, 1, 0)
		byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
		createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, appName, 1, 0, suite.authToken)
		//Base.CreateFileAndEnterData("createApp", "appId", strconv.Itoa(createAppResponseDto.Result.Id))
		log.Println("=== Validating the Response of the CreateAppWithValidPayload API ===")
		assert.Equal(suite.T(), createAppRequestDto.AppName, createAppResponseDto.Result.AppName)
		log.Println("getting payload for Delete Team API")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
		log.Println("Hitting the Delete team API for Removing the data created via automation")
		HitDeleteAppApi(byteValueOfDeleteApp, createAppResponseDto.Result.Id, suite.authToken)
	})

	suite.Run("A=2=CreateAppWithInvalidTeamId", func() {
		inValidTeamId := Base.GetRandomNumberOf9Digit()
		createAppRequestDto := GetAppRequestDto(appName, inValidTeamId, 0)
		byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
		createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, appName, inValidTeamId, 0, suite.authToken)
		log.Println("Validating the Response of the Create Gitops Config API...")
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", createAppResponseDto.Errors[0].InternalMessage)
		log.Println("getting payload for Delete Team API")
		byteValueOfCreateApp = GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
		log.Println("Hitting the Delete team API for Removing the data created via automation")
		HitDeleteAppApi(byteValueOfCreateApp, createAppResponseDto.Result.Id, suite.authToken)
	})

	suite.Run("A=3=CreateAppWithInvalidTemplateId", func() {
		invalidTemplateId := Base.GetRandomNumberOf9Digit()
		createAppRequestDto := GetAppRequestDto(appName, 1, invalidTemplateId)
		byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
		createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, "", 0, invalidTemplateId, suite.authToken)
		log.Println("Validating the Response of the CreateAppWithInvalidTemplateId API...")
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", createAppResponseDto.Errors[0].InternalMessage)
		log.Println("getting payload for Delete Team API")
		byteValueOfCreateApp = GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
		log.Println("Hitting the Delete team API for Removing the data created via automation")
		HitDeleteAppApi(byteValueOfCreateApp, createAppResponseDto.Result.Id, suite.authToken)
	})
	// <tear-down code>
}
