package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func (suite *PipelineConfigSuite) TestCreateAppWithValidPayload() {
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	createAppRequestDto := GetAppRequestDto("app"+appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, "app"+appName, 1, 0, suite.authToken)
	Base.CreateFileAndEnterData("createApp", "app_id", strconv.Itoa(createAppResponseDto.Result.Id))
	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), createAppRequestDto.AppName, createAppResponseDto.Result.AppName)
	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfDeleteApp, createAppResponseDto.Result.Id, suite.authToken)
}

func (suite *PipelineConfigSuite) TestCreateAppWithInvalidTeamId() {
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	teamId := rand.Intn(89-10) + 10
	createAppRequestDto := GetAppRequestDto("app"+appName, teamId, 0)
	byteValueOfStruct, _ := json.Marshal(createAppRequestDto)

	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfStruct, "app"+appName, teamId, 0, suite.authToken)
	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), "[{pg: no rows in result set}]", createAppResponseDto.Errors[0].InternalMessage)
	log.Println("getting payload for Delete Team API")
	byteValueOfStruct = GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfStruct, createAppResponseDto.Result.Id, suite.authToken)
}
func (suite *PipelineConfigSuite) TestCreateAppWithInvalidTemplateId() {
	appName := Base.GetRandomStringOfGivenLength(10)
	templateId := rand.Intn(89-10) + 10
	createAppRequestDto := GetAppRequestDto("app"+appName, 1, templateId)
	byteValueOfStruct, _ := json.Marshal(createAppRequestDto)
	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfStruct, "app"+appName, 1, templateId, suite.authToken)
	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), "Key: 'CreateAppDTO.AppName' Error:Field validation for 'AppName' failed on the 'name-component' tag", createAppResponseDto.Errors[0].InternalMessage)
	log.Println("getting payload for Delete Team API")
	byteValueOfStruct = GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfStruct, createAppResponseDto.Result.Id, suite.authToken)
}
