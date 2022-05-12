package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
)

func (suite *regressionTestSuite) TestFetchOtherEnvWithValidAppId() {

	appName := Base.GetRandomStringOfGivenLength(10)
	createAppRequestDto := GetAppRequestDto(appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)

	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, appName, 1, 0, suite.authToken)

	AppId := map[string]string{
		"id": strconv.Itoa(createAppResponseDto.Result.Id),
	}

	fetchOtherEnvResponseDto := FetchOtherEnv(AppId, suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 200, fetchOtherEnvResponseDto.Code)

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfDeleteApp, suite.authToken)
}
func (suite *regressionTestSuite) TestFetchOtherEnvWithInvalidAppId() {
	AppId := map[string]string{
		"id": strconv.Itoa(rand.Intn(899-100) + 100),
	}
	fetchOtherEnvResponseDto := FetchOtherEnv(AppId, suite.authToken)
	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 404, fetchOtherEnvResponseDto.Code)

}
