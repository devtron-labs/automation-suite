package TeamRouter

import (
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *TeamTestSuite) TestClassA7FetchAllTeam() {

	suite.Run("A=1=FetchAllTeamWithValidPayload", func() {
		log.Println("Hitting the 'FetchAllTeam' Api before creating any new entry")

		fetchAllTeamResponseDto := HitFetchAllTeamApi(suite.authToken)
		noOfTeams := len(fetchAllTeamResponseDto.Result)

		log.Println("Hitting the 'Save Team' Api for creating a new entry")
		saveTeamResponseDto := HitSaveTeamApi(nil, suite.authToken)

		log.Println("Hitting the FetchAllTeam API again for verifying the functionality of it")
		fetchAllTeamResponseDto = HitFetchAllTeamApi(suite.authToken)

		log.Println("Validating the response of FetchAllTeam API")
		assert.Equal(suite.T(), noOfTeams+1, len(fetchAllTeamResponseDto.Result))
		assert.Equal(suite.T(), saveTeamResponseDto.Result.Name, fetchAllTeamResponseDto.Result[len(fetchAllTeamResponseDto.Result)-1].Name)

		log.Println("getting payload for Delete Team API")
		byteValueOfStruct := GetPayLoadForDeleteAPI(saveTeamResponseDto.Result.Id, saveTeamResponseDto.Result.Name, true)
		log.Println("Hitting the Delete Team API for Removing the data created via automation")
		HitDeleteTeamApi(byteValueOfStruct, suite.authToken)
	})
}
