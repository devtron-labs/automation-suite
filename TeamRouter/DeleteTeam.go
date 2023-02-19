package TeamRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *TeamTestSuite) TestClass6DeleteTeam() {
	suite.Run("A=1=DeleteTeamWithValidPayload", func() {
		saveTeamRequestDto := GetSaveTeamRequestDto()
		byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)

		log.Println("Hitting The Save Team API")
		saveTeamResponseDto := HitSaveTeamApi(byteValueOfStruct, suite.authToken)

		log.Println("getting payload for Delete Team API")
		byteValueOfStruct = GetPayLoadForDeleteAPI(saveTeamResponseDto.Result.Id, saveTeamResponseDto.Result.Name, true)
		log.Println("Hitting the Delete Team API for Removing the data created via automation")
		deleteTeamResponse := HitDeleteTeamApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Project deleted successfully.", deleteTeamResponse.Result)
		assert.Equal(suite.T(), "OK", deleteTeamResponse.Status)
	})
}
