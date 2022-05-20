package externalLinkout

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *LinkTestSuite) TestUpdateTeamWithValidPayload() {
	saveLinkRequestDto := GetSaveLinkRequestDto()
	byteValueOfStruct, _ := json.Marshal(saveLinkRequestDto)
	saveLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
	id := saveLinkResponseDto.Result.Id
	monitoringToolId := saveLinkResponseDto.Result.MonitoringToolId
	updateLinkRequestPayload := GetUpdateLinkRequestPayload(id, "UpdatedNameViaAutomation", monitoringToolId)
	log.Println("Hitting The Update Link API")
	updateLinkResponseDto := HitUpdateLinkApi(updateLinkRequestPayload, suite.authToken)
	assert.Equal(suite.T(), "UpdatedNameViaAutomation", updateLinkResponseDto.Result.Name)
	byteValueOfStruct = GetPayLoadForDeleteLinkAPI(id, "UpdatedNameViaAutomation", updateLinkResponseDto.Result.MonitoringToolId, updateLinkResponseDto.Result.Url, updateLinkResponseDto.Result.Active)
	HitDeleteLinkApi(updateLinkRequestPayload, suite.authToken)
}
func (suite *LinkTestSuite) TestUpdateTeamWithInvalidMonitoringToolId() {
	saveLinkRequestDto := GetSaveLinkRequestDto()
	byteValueOfStruct, _ := json.Marshal(saveLinkRequestDto)
	saveLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
	id := saveLinkResponseDto.Result.Id

	updateLinkRequestPayload := GetUpdateLinkRequestPayloadInvalidMonitorigId(id, "UpdatedNameViaAutomation")
	log.Println("Hitting The Update Link API")
	updateLinkResponseDto := HitUpdateLinkApi(updateLinkRequestPayload, suite.authToken)
	assert.Equal(suite.T(), 400, updateLinkResponseDto.Code)

	log.Println("Cross verifying the response of update API via getLinkById API")
	responseBodyGetLinkById := HitGetLinkByIdApi(strconv.Itoa(id), suite.authToken)
	assert.Equal(suite.T(), "UpdatedNameViaAutomation", responseBodyGetLinkById.Result.Name)

	byteValueOfStruct = GetPayLoadForDeleteLinkAPI(id, "UpdatedNameViaAutomation", responseBodyGetLinkById.Result.MonitoringToolId, responseBodyGetLinkById.Result.Url, responseBodyGetLinkById.Result.Active)
	HitDeleteLinkApi(updateLinkRequestPayload, suite.authToken)
}

func (suite *LinkTestSuite) TestUpdateTeamWithInvalidClusterId() {
	saveLinkRequestDto := GetSaveLinkRequestDto()
	byteValueOfStruct, _ := json.Marshal(saveLinkRequestDto)
	saveLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
	id := saveLinkResponseDto.Result.Id

	updateLinkRequestPayload := GetUpdateLinkRequestPayloadInvalidClusterId(id, "UpdatedNameViaAutomation")
	log.Println("Hitting The Update Link API")
	updateLinkResponseDto := HitUpdateLinkApi(updateLinkRequestPayload, suite.authToken)
	assert.Equal(suite.T(), 400, updateLinkResponseDto.Code)

	log.Println("Cross verifying the response of update API via getLinkById API")
	responseBodyGetLinkById := HitGetLinkByIdApi(strconv.Itoa(id), suite.authToken)
	assert.Equal(suite.T(), "UpdatedNameViaAutomation", responseBodyGetLinkById.Result.Name)

	byteValueOfStruct = GetPayLoadForDeleteLinkAPI(id, "UpdatedNameViaAutomation", responseBodyGetLinkById.Result.MonitoringToolId, responseBodyGetLinkById.Result.Url, responseBodyGetLinkById.Result.Active)
	HitDeleteLinkApi(updateLinkRequestPayload, suite.authToken)
}
