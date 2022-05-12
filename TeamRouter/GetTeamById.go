package TeamRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *TeamTestSuite) TestGetTeamByIdWithValidId() {
	log.Println("Hitting the 'Save Team' Api for creating a new entry")
	saveTeamResponseDto := HitSaveTeamApi(nil, suite.authToken)

	log.Println("Hitting the GetTeamById API with Valid Id")
	responseBodyGetTeamById := HitGetTeamByIdApi(strconv.Itoa(saveTeamResponseDto.Result.Id), suite.authToken)
	log.Println("Validating the response of GetTeamById API")
	assert.Equal(suite.T(), saveTeamResponseDto.Result.Id, responseBodyGetTeamById.Result.Id)
	assert.Equal(suite.T(), saveTeamResponseDto.Result.Name, responseBodyGetTeamById.Result.Name)
	assert.Equal(suite.T(), saveTeamResponseDto.Result.Active, responseBodyGetTeamById.Result.Active)

	log.Println("getting payload for Delete Team API")
	byteValueOfStruct := GetPayLoadForDeleteAPI(saveTeamResponseDto.Result.Id, saveTeamResponseDto.Result.Name, true)
	log.Println("Hitting the Delete Team API for Removing the data created via automation")
	HitDeleteTeamApi(byteValueOfStruct, suite.authToken)
}

func (suite *TeamTestSuite) TestGetTeamByIdWithInvalidId() {
	randomId := Base.GetRandomNumberOf9Digit()
	log.Println("Hitting the GetTeamById API with Invalid Random Id")
	resp := HitGetTeamByIdApi(strconv.Itoa(randomId), suite.authToken)
	assert.Equal(suite.T(), 404, resp.Code)
	assert.Equal(suite.T(), "Not Found", resp.Status)
}
