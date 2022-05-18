package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

//(suite *regressionTestSuite)
func TestCreateAppWithValidPayload(t *testing.T) {
	authToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjUwOTIxNWQ0OGY4NDcwMDIzNjY3MmJlMmVkOWQwNWZkYWQ4MmI2NzMifQ.eyJpc3MiOiJodHRwczovL3N0YWdpbmcuZGV2dHJvbi5pbmZvL29yY2hlc3RyYXRvci9hcGkvZGV4Iiwic3ViIjoiQ2hVeE1ESTJPVEk1TmpBd056RXhNekU0TVRFMU5qY1NCbWR2YjJkc1pRIiwiYXVkIjoiYXJnby1jZCIsImV4cCI6MTY1Mjg1NDQyMSwiaWF0IjoxNjUyNzY4MDIxLCJhdF9oYXNoIjoiZWlMVHFoZFdVT083SFpYMXpnNWxNZyIsImVtYWlsIjoibmlrZXNoQGRldnRyb24uYWkiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6Ik5pa2VzaCBSYXRob2QifQ.DDIcLIG2jaQYJfgYjemSQ3fKxpFu-I6e-UvKdPQ1c3mjlTBjwCA8JKgOYPHaLbwrOrXsHQBg-48NrKE5dZNYJg0lAKEnGo7feSeZ4ZFIQtp3lC-XxXGMfodkziRZ_D36s0ePbjo5Q2zB0q3i6BEfokXxESE-31aH_jZDOCs2px_1OxMORCX29j8MyiGUPPFR3h3NYG3UhvompEl7eRTGygtft1_9rhHv9BPxL_yXqHU2IGVtaRz--zH6V6IknA_Ha9kkl7ZEh8LJLlr6a787MBxQfSXlOlqmfAeOACVbnlbCvWPBHqCwn6zgjIn8nVxtmNnShn_sbDeeyZy7aJjS4Q"
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	createAppRequestDto := GetAppRequestDto(appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)

	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, appName, 1, 0, authToken)

	Base.CreateFileAndEnterData("createApp", "app_id", strconv.Itoa(createAppResponseDto.Result.Id))

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(t, createAppRequestDto.AppName, createAppResponseDto.Result.AppName)

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")

	HitDeleteAppApi(byteValueOfDeleteApp, createAppResponseDto.Result.Id, authToken)
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
	log.Println("getting payload for Delete Team API")
	byteValueOfStruct = GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfStruct, createAppResponseDto.Result.Id, suite.authToken)
}
func (suite *regressionTestSuite) TestCreateAppWithInvalidTemplateId() {
	appName := Base.GetRandomStringOfGivenLength(10)
	templateId := rand.Intn(89-10) + 10
	createAppRequestDto := GetAppRequestDto(appName, 1, templateId)
	byteValueOfStruct, _ := json.Marshal(createAppRequestDto)

	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfStruct, appName, 1, templateId, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 400, createAppResponseDto.Code)
	assert.Equal(suite.T(), "Key: 'CreateAppDTO.AppName' Error:Field validation for 'AppName' failed on the 'name-component' tag", createAppResponseDto.Errors[0].InternalMessage)

	log.Println("getting payload for Delete Team API")
	byteValueOfStruct = GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfStruct, createAppResponseDto.Result.Id, suite.authToken)
}
