package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strings"
)

func (suite *regressionTestSuite) TestCreateAppWithValidPayload() {
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	//appName := "test-app-2"
	createAppRequestDto := GetAppRequestDto(appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)

	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, appName, 1, 0, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 200, createAppResponseDto.Code)
	assert.Equal(suite.T(), createAppRequestDto.AppName, createAppResponseDto.Result.AppName)

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfDeleteApp, suite.authToken)
}

func (suite *regressionTestSuite) TestCreateAppWithInvalidTeamId() {
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	teamId := rand.Intn(89-10) + 10
	createAppRequestDto := GetAppRequestDto(appName, teamId, 0)
	byteValueOfStruct, _ := json.Marshal(createAppRequestDto)

	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfStruct, appName, teamId, 0, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 404, createAppResponseDto.Code)
	assert.Equal(suite.T(), "[{pg: no rows in result set}]", createAppResponseDto.Errors[0].InternalMessage)
	// 404
	// "Key: 'CreateAppDTO.AppName' Error:Field validation for 'AppName' failed on the 'name-component' tag"

	log.Println("getting payload for Delete Team API")
	byteValueOfStruct = GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfStruct, suite.authToken)
}
func (suite *regressionTestSuite) TestCreateAppWithInvalidTemplateId() {
	appName := Base.GetRandomStringOfGivenLength(10)
	templateId := rand.Intn(89-10) + 10
	createAppRequestDto := GetAppRequestDto(appName, 1, templateId)
	byteValueOfStruct, _ := json.Marshal(createAppRequestDto)

	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfStruct, appName, 1, templateId, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 404, createAppResponseDto.Code)
	assert.Equal(suite.T(), "[{pg: no rows in result set}]", createAppResponseDto.Errors[0].InternalMessage)

	log.Println("getting payload for Delete Team API")
	byteValueOfStruct = GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfStruct, suite.authToken)
}
