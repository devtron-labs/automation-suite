package externalLinkoutRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *LinkOutRouterTestSuite) TestClassGetExternalLinksViaClusterId() {
	clusterId := map[string]string{
		"clusterId": strconv.Itoa(1),
	}
	suite.Run("A=1=GetExternalLinksViaValidClusterId", func() {
		log.Println("=== Here we are getting no of links in clusterId 1 ===")
		fetchLinksResponseDto := HitFetchLinksByClusterIdApi(clusterId, suite.authToken)
		noOfLinksBeforeCreatingNew := len(fetchLinksResponseDto.Result)
		log.Println("=== Here we are adding link in clusterId 1 ===")
		createLinkRequestDto := GetSaveLinkRequestDto(1, nil)
		byteValueOfCreateLink, _ := json.Marshal(createLinkRequestDto)
		HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)
		log.Println("=== Here we are getting no of links in clusterId 1 after creating new one ===")
		fetchLinksResponseDto = HitFetchLinksByClusterIdApi(clusterId, suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		noOfLinksAfterCreatingNew := len(fetchLinksResponseDto.Result)
		assert.Equal(suite.T(), noOfLinksBeforeCreatingNew+1, noOfLinksAfterCreatingNew)
		assert.NotNil(suite.T(), fetchLinksResponseDto.Result)
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)
		HitDeleteLinkApi(getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id, suite.authToken)
	})

	suite.Run("A=2=GetExternalLinksViaValidClusterId", func() {
		invalidClusterId := map[string]string{
			"clusterId": strconv.Itoa(Base.GetRandomNumberOf9Digit()),
		}
		fetchAllLinkResponseDto := HitFetchLinksByClusterIdApi(invalidClusterId, suite.authToken)
		log.Println("Validating the response of FetchAllLinkByClusterIdApi API")
		assert.Equal(suite.T(), 0, len(fetchAllLinkResponseDto.Result))
	})
}
