package externalLinkoutRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *LinkOutRouterTestSuite) TestClassA2DeleteExternalLink() {
	suite.Run("A=1=DeleteExternalLinkWithValidId", func() {
		log.Println("Fetching links before creating new")
		getAllExternalLinksResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinks := len(getAllExternalLinksResponseDto.Result)
		identifier := "ea-app-" + strings.ToLower(testUtils.GetRandomStringOfGivenLength(5))
		listOfCreateLinkRequestDto := GetSaveLinkRequestDtoList(1, identifier, 8, true, "appLevel", "https://www.google.co.in", "external-helm-app", 0)
		byteValueOfCreateLink, _ := json.Marshal(listOfCreateLinkRequestDto)
		log.Println("Hitting The Save Link API")
		HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)
		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		DeletedLinkResponse := HitDeleteLinkApi(getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id, suite.authToken)
		assert.Equal(suite.T(), true, DeletedLinkResponse.Result.Success)
		getAllExternalLinksAgainToCheckResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksToCheck := len(getAllExternalLinksAgainToCheckResponseDto.Result)
		assert.Equal(suite.T(), noOfLinks, noOfLinksToCheck)
	})
	suite.Run("A=2=DeleteExternalLinkOutWithInValidId", func() {
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		DeletedLinkResponseDto := HitDeleteLinkApi(testUtils.GetRandomNumberOf9Digit(), suite.authToken)
		assert.Equal(suite.T(), 500, DeletedLinkResponseDto.Code)
		assert.Equal(suite.T(), "external_link failed to delete", DeletedLinkResponseDto.Errors[0].UserMessage)

	})
}
