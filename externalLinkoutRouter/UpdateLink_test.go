package externalLinkoutRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"log"

	"github.com/stretchr/testify/assert"
)

func (suite *ExternalLinkOutRouterTestSuite) TestClassA5UpdateExternalLink() {
	suite.Run("A=1=UpdateTeamWithValidPayload", func() {
		createLinkRequestDto := GetSaveLinkRequestDto(1, nil)
		byteValueOfCreateLink, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)
		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)

		id := getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id
		monitoringToolId := getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].MonitoringToolId
		url := getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Url
		updateLinkRequestPayload := GetUpdateLinkRequestPayload(id, "UpdatedNameViaAutomation", monitoringToolId, url)

		byteValueOfUpdateLink, _ := json.Marshal(updateLinkRequestPayload)
		log.Println("Hitting The Update Link API")
		updateLinkResponseDto := HitUpdateLinkApi(byteValueOfUpdateLink, suite.authToken)

		assert.Equal(suite.T(), 200, updateLinkResponseDto.Code)
		HitDeleteLinkApi(id, suite.authToken)
	})
	suite.Run("A=2=UpdateTeamWithInvalidMonitoringToolId", func() {
		createLinkRequestDto := GetSaveLinkRequestDto(1, nil)
		byteValueOfCreateLink, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)
		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)

		id := getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id
		url := getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Url

		monitoringToolId := testUtils.GetRandomNumberOf9Digit()

		updateLinkRequestPayload := GetUpdateLinkRequestPayload(id, "UpdatedNameViaAutomation", monitoringToolId, url)
		byteValueOfUpdateLink, _ := json.Marshal(updateLinkRequestPayload)
		log.Println("Hitting The Update Link API")
		updateLinkResponseDto := HitUpdateLinkApi(byteValueOfUpdateLink, suite.authToken)

		assert.Equal(suite.T(), 500, updateLinkResponseDto.Code)
		assert.Equal(suite.T(), "ERROR #23503 insert or update on table \"external_link\" violates foreign key constraint \"external_link_external_link_monitoring_tool_id_fkey\"", updateLinkResponseDto.Errors[0].UserMessage)
		HitDeleteLinkApi(id, suite.authToken)
	})

	suite.Run("A=3=UpdateTeamWithInvalidClusterId", func() {
		createLinkRequestDto := GetSaveLinkRequestDto(1, nil)
		byteValueOfCreateLink, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)
		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)

		id := getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id
		url := getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Url
		monitoringToolId := getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].MonitoringToolId

		updateLinkRequestPayload := GetUpdateLinkRequestPayload(id, "UpdatedNameViaAutomation", monitoringToolId, url)
		clusterIds := []int{testUtils.GetRandomNumberOf9Digit()}
		updateLinkRequestPayload.ClusterIds = append(updateLinkRequestPayload.ClusterIds, clusterIds...)

		byteValueOfUpdateLink, _ := json.Marshal(updateLinkRequestPayload)
		log.Println("Hitting The Update Link API")
		updateLinkResponseDto := HitUpdateLinkApi(byteValueOfUpdateLink, suite.authToken)

		assert.Equal(suite.T(), 500, updateLinkResponseDto.Code)
		assert.Equal(suite.T(), "cluster id failed to create in db", updateLinkResponseDto.Errors[0].UserMessage)
		HitDeleteLinkApi(id, suite.authToken)
	})

}
