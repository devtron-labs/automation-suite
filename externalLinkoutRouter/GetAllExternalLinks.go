package externalLinkoutRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *LinkOutRouterTestSuite) TestClassA3GetAllExternalLink() {
	suite.Run("A=1=FetchAllLinkOuts", func() {
		log.Println("Hitting the 'FetchAllLink' Api before creating any new entry")
		fetchAllLinkResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinks := len(fetchAllLinkResponseDto.Result)
		log.Println("=== Here we are adding new link ===")
		createLinkRequestDto := GetSaveLinkRequestDto(1, nil)
		byteValueOfCreateLink, _ := json.Marshal(createLinkRequestDto)
		HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)
		log.Println("Hitting the 'FetchAllLink' Api after creating a new entry")
		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)
		assert.Equal(suite.T(), noOfLinks+1, noOfLinksAfterCreation)
		log.Println("Checking external-link name ")
		assert.Equal(suite.T(), createLinkRequestDto[0].Name, getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Name)
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		HitDeleteLinkApi(getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id, suite.authToken)
	})
}
