package externalLinkoutRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *LinkOutRouterTestSuite) TestClassA3GetAllExternalLink() {
	suite.Run("A=1=FetchAllLinkOuts", func() {
		var listOfLinkName []string
		var listOfActualLinkNames []string
		log.Println("Fetching links before creating new")
		var noOfLinkRequired = 1
		getAllExternalLinksResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinks := len(getAllExternalLinksResponseDto.Result)
		identifier := "ea-app-" + strings.ToLower(testUtils.GetRandomStringOfGivenLength(5))
		listOfCreateLinkRequestDto := GetSaveLinkRequestDtoList(noOfLinkRequired, identifier, 8, true, "appLevel", "https://www.google.co.in", "external-helm-app", 0)
		for _, createLinkRequestDto := range listOfCreateLinkRequestDto {
			listOfLinkName = append(listOfLinkName, createLinkRequestDto.Name)
		}
		byteValueOfCreateLink, _ := json.Marshal(listOfCreateLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 200, createLinkResponseDto.Code)
		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		for _, ExternalLinks := range getAllExternalLinksAgainResponseDto.Result {
			listOfActualLinkNames = append(listOfActualLinkNames, ExternalLinks.Name)
		}
		assert.Subset(suite.T(), listOfActualLinkNames, listOfLinkName)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)
		log.Println("Checking length of result")
		assert.Equal(suite.T(), noOfLinks+noOfLinkRequired, noOfLinksAfterCreation)
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		for x := 0; x < noOfLinkRequired; x++ {
			HitDeleteLinkApi(getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id, suite.authToken)
		}
	})
}
