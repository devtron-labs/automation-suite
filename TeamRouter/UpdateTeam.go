package TeamRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *TeamTestSuite) TestUpdateTeamWithValidPayload() {
	saveTeamRequestDto := GetSaveTeamRequestDto()
	byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)
	saveTeamResponseDto := HitSaveTeamApi(byteValueOfStruct, suite.authToken)
	id := saveTeamResponseDto.Result.Id

	updateTeamRequestPayload := GetUpdateTeamRequestPayload(id, "UpdatedNameViaAutomation")
	log.Println("Hitting The Update Team API")
	updateTeamResponseDto := HitUpdateTeamApi(updateTeamRequestPayload, suite.authToken)
	assert.Equal(suite.T(), "UpdatedNameViaAutomation", updateTeamResponseDto.Result.Name)

	log.Println("Cross verifying the response of update API via getTeamById API")
	responseBodyGetTeamById := HitGetTeamByIdApi(strconv.Itoa(id), suite.authToken)
	assert.Equal(suite.T(), "UpdatedNameViaAutomation", responseBodyGetTeamById.Result.Name)

	byteValueOfStruct = GetPayLoadForDeleteAPI(id, "UpdatedNameViaAutomation", true)
	HitDeleteTeamApi(updateTeamRequestPayload, suite.authToken)
}

//todo there are some doubt as well ,will add other test cases after the clarification
