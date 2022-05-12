package external_linkout

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *LinkTestSuite) TestDeleteTeamWithValidPayload() {
	createLinkRequestDto := GetSaveLinkRequestDto()
	byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)

	log.Println("Hitting The Save Link API")
	createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)

	log.Println("getting payload for Delete Team API")
	byteValueOfStruct = GetPayLoadForDeleteLinkAPI(createLinkResponseDto.Result.Id, createLinkResponseDto.Result.Name, createLinkResponseDto.Result.MonitoringToolId, createLinkResponseDto.Result.Url, createLinkResponseDto.Result.Active)
	log.Println("Hitting the Delete Linkout API for Removing the data created via automation")
	deleteLinkResponse := HitDeleteLinkApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), "Project deleted successfully.", deleteLinkResponse.Result)
	assert.Equal(suite.T(), "OK", deleteLinkResponse.Status)
}
