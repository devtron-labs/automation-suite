package TeamRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *TeamTestSuite) TestSaveTeamWithValidPayload() {
	saveTeamRequestDto := GetSaveTeamRequestDto()
	byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)

	log.Println("Hitting The Save Team API")
	saveTeamResponseDto := HitSaveTeamApi(byteValueOfStruct, suite.authToken)

	log.Println("Validating the Response of the Save API...")
	assert.Equal(suite.T(), saveTeamRequestDto.Name, saveTeamResponseDto.Result.Name)
	assert.NotNil(suite.T(), saveTeamResponseDto.Result.Id)

	log.Println("getting payload for Delete Team API")
	byteValueOfStruct = GetPayLoadForDeleteAPI(saveTeamResponseDto.Result.Id, saveTeamResponseDto.Result.Name, true)
	log.Println("Hitting the Delete Team API for Removing the data created via automation")
	HitDeleteTeamApi(byteValueOfStruct, suite.authToken)
}

func (suite *TeamTestSuite) TestSaveTeamWithExistingId() {
	saveTeamRequestDto := GetSaveTeamRequestDto()
	byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)
	saveTeamResponseDto := HitSaveTeamApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting The Save Team API Again with already generated Id")
	saveTeamRequestDto.Id = saveTeamResponseDto.Result.Id
	updatedByteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)
	saveTeamResponseDtoAfterHittingWIthExistingId := HitSaveTeamApi(updatedByteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), 500, saveTeamResponseDtoAfterHittingWIthExistingId.Code)
	assert.Equal(suite.T(), "team failed to create in db", saveTeamResponseDtoAfterHittingWIthExistingId.Errors[0].UserMessage)

	byteValueOfStruct = GetPayLoadForDeleteAPI(saveTeamResponseDto.Result.Id, saveTeamResponseDto.Result.Name, true)
	HitDeleteTeamApi(byteValueOfStruct, suite.authToken)
}
