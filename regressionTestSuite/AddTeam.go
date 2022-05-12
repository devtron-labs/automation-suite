package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *regressionTestSuite) TestCreateTeamWithValidPayload() {
	name := Base.GetRandomStringOfGivenLength(10)
	createTeamRequestDto := GetTeamRequestDto(name, true)
	byteValueOfCreateTeam, _ := json.Marshal(createTeamRequestDto)

	log.Println("Hitting The post team API")
	createTeamResponseDto := HitCreateTeamApi(byteValueOfCreateTeam, name, true, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), createTeamRequestDto.Name, createTeamResponseDto.Result.Name)
}
