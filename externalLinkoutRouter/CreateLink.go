package externalLinkoutRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"log"

	"github.com/stretchr/testify/assert"
)

func (suite *ExternalLinkOutRouterTestSuite) TestClassA1CreateExternalLink() {
	suite.Run("A=1=CreateLinkoutWithValidPayload", func() {
		log.Println("Fetching links before creating new")
		getAllExternalLinksResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinks := len(getAllExternalLinksResponseDto.Result)

		createLinkRequestDto := GetSaveLinkRequestDto(1, nil)
		byteValueOfCreateLink, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)

		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 200, createLinkResponseDto.Code)
		getAllExternalLinksAgainResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksAgainResponseDto.Result)
		log.Println("Cheking length of result")
		assert.Equal(suite.T(), noOfLinks+1, noOfLinksAfterCreation)
		log.Println("Checking external-link name ")
		assert.Equal(suite.T(), createLinkRequestDto[0].Name, getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Name)
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		HitDeleteLinkApi(getAllExternalLinksAgainResponseDto.Result[noOfLinksAfterCreation-1].Id, suite.authToken)

	})
	suite.Run("A=2=CreateLinkoutWithInvalidToolId", func() {
		log.Println("Getting random monitoring tool id")
		monitoringToolId := testUtils.GetRandomNumberOf9Digit()
		createLinkRequestDto := GetSaveLinkRequestDto(monitoringToolId, nil)
		byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
		assert.NotNil(suite.T(), "external link failed to create in db", createLinkResponseDto.Errors[0].UserMessage)

	})
	suite.Run("A=3=CreateLinkoutWithInvalidClusterId", func() {
		clusterIds := []int{testUtils.GetRandomNumberOf9Digit()}
		createLinkRequestDto := GetSaveLinkRequestDto(1, clusterIds)
		byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
		assert.NotNil(suite.T(), "cluster id failed to create in db", createLinkResponseDto.Errors[0].UserMessage)
	})
	suite.Run("A=4=CreateLinkoutWithOneValidOneInvalidClusterId", func() {
		clusterIds := []int{1, testUtils.GetRandomNumberOf9Digit()}
		createLinkRequestDto := GetSaveLinkRequestDto(1, clusterIds)
		byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
		assert.NotNil(suite.T(), "cluster id failed to create in db", createLinkResponseDto.Errors[0].UserMessage)
	})
}
