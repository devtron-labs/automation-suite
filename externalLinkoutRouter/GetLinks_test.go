package externalLinkoutRouter

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/stretchr/testify/assert"
)

func (suite *ExternalLinkOutRouterTestSuite) TestClassA3GetExternalLink() {
	suite.Run("A=1=FetchAllLinkouts", func() {
		log.Println("Hitting the 'FetchAllLink' Api before creating any new entry")
		fetchAllLinkResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinks := len(fetchAllLinkResponseDto.Result)

		createLinkRequestDto := GetSaveLinkRequestDto(1, nil)
		byteValueOfCreateLink, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)

		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)
		log.Println("Cheking length of result")
		assert.Equal(suite.T(), noOfLinks+1, noOfLinksAfterCreation)
		log.Println("Checking external-link name ")
		assert.Equal(suite.T(), createLinkRequestDto[0].Name, getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Name)
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		HitDeleteLinkApi(getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id, suite.authToken)

	})

	suite.Run("A=2=FetchAllLinkoutsWithValidClusterId", func() {
		log.Println("Hitting the 'FetchAllLink' Api before creating any new entry")

		createLinkRequestDto := GetSaveLinkRequestDto(1, nil)
		byteValueOfCreateLink, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)

		log.Println("Hitting the FetchAllTeam API again for verifying the functionality of it")
		clusterId := map[string]string{
			"id": strconv.Itoa(1),
		}
		fetchAllLinkResponseDto := HitFetchAllLinkByClusterIdApi(clusterId, suite.authToken)

		log.Println("Validating the response of FetchAllLink API")
		assert.NotNil(suite.T(), fetchAllLinkResponseDto.Result)

		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		HitDeleteLinkApi(getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id, suite.authToken)

	})
}
