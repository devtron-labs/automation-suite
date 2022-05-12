package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *regressionTestSuite) TestFetchAutocomplete() {
	name := Base.GetRandomStringOfGivenLength(10)
	createTeamRequestDto := GetTeamRequestDto(name, true)
	byteValueOfCreateTeam, _ := json.Marshal(createTeamRequestDto)

	log.Println("Hitting The post team API")
	createTeamResponseDto := HitCreateTeamApi(byteValueOfCreateTeam, name, true, suite.authToken)

	log.Println("Validating the Response of the get autocomplete API...")
	fetchAutocompleteResponseDto := HitFetchAllTeamApi(suite.authToken)

	log.Println("Validating the response of fetchAutocomplete API")
	assert.Equal(suite.T(), createTeamResponseDto.Result.Id, fetchAutocompleteResponseDto.Result[len(fetchAutocompleteResponseDto.Result)-1].Id)
	assert.Equal(suite.T(), createTeamResponseDto.Result.Name, fetchAutocompleteResponseDto.Result[len(fetchAutocompleteResponseDto.Result)-1].Name)
	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteTeam := GetPayLoadForDeleteTeamAPI(createTeamResponseDto.Result.Name, createTeamResponseDto.Result.Active)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteTeamApi(byteValueOfDeleteTeam, suite.authToken)
}
