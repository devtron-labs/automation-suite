package externalLinkoutRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/stretchr/testify/assert"
)

func (suite *LinkOutRouterTestSuite) TestClassA1CreateExternalLink() {
	log.Println("=== Here we are creating an App ===")
	createAppApiResponse := testUtils.CreateApp(suite.authToken).Result
	log.Println("=== Here we are printing AppName ===>", createAppApiResponse.AppName)

	suite.Run("A=1=CreateLinkForExternalHelmApp", func() {
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

	//todo suite.Run("A=2=CreateLinkWithInvalidIdentifierTypeInPayload", func(){}, need to handle this in code

	suite.Run("A=3=VerifyCreateLinkWithInvalidLevel", func() {
		identifier := "ea-app-" + strings.ToLower(testUtils.GetRandomStringOfGivenLength(5))
		listOfCreateLinkRequestDto := GetSaveLinkRequestDtoList(1, identifier, 8, true, "appLevelInvalid", "https://www.google.co.in", "external-helm-app", 0)
		byteValueOfCreateLink, _ := json.Marshal(listOfCreateLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfCreateLink, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
		assert.NotNil(suite.T(), "external link failed to create", createLinkResponseDto.Errors[0].UserMessage)
	})

	//todo suite.Run("A=4=CreateLinkWithInvalidUrlInPayload", func(){}, need to handle this in code

	suite.Run("A=5=CreateLinkOutWithInvalidToolId", func() {
		log.Println("Getting random monitoring tool id")
		monitoringToolId := testUtils.GetRandomNumberOf9Digit()
		identifier := "ea-app-" + strings.ToLower(testUtils.GetRandomStringOfGivenLength(5))
		createLinkRequestDto := GetSaveLinkRequestDtoList(1, identifier, monitoringToolId, true, "appLevel", "https://www.google.co.in", "external-helm-app", 0)
		byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
		assert.NotNil(suite.T(), "external link failed to create ", createLinkResponseDto.Errors[0].UserMessage)
	})

	suite.Run("A=6=CreateLinkOutWithValidAppIdAndEditRight", func() {
		log.Println("Getting random monitoring tool id")
		identifier := strconv.Itoa(createAppApiResponse.Id)
		createLinkRequestDto := GetSaveLinkRequestDtoList(1, identifier, 19, true, "appLevel", "https://www.google.co.in", "devtron-app", 0)
		byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 200, createLinkResponseDto.Code)
		log.Println("Hitting the Delete link API for Removing the data created via automation")
		getAllExternalLinksResponseDto := HitFetchAllLinkApi(suite.authToken)
		noOfLinksAfterCreation := len(getAllExternalLinksResponseDto.Result)
		HitDeleteLinkApi(getAllExternalLinksResponseDto.Result[noOfLinksAfterCreation-1].Id, suite.authToken)

	})

	//todo need to handle this case once dev will provide fix
	/*suite.Run("A=3=CreateLinkOutWithInvalidClusterId", func() {
		clusterId := testUtils.GetRandomNumberOf9Digit()
		createLinkRequestDto := GetSaveLinkRequestDtoList(8, true, "appLevel", "https://www.google.co.in", "external-helm-app", clusterId)
		byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
		assert.NotNil(suite.T(), "cluster id failed to create", createLinkResponseDto.Errors[0].UserMessage)
	})*/

	/*suite.Run("A=4=CreateLinkOutWithOneValidOneInvalidClusterId", func() {
		clusterIds := []int{1, testUtils.GetRandomNumberOf9Digit()}
		createLinkRequestDto := GetSaveLinkRequestDtoList(8, true, "appLevel", "https://www.google.co.in", "external-helm-app", clusterId)
		byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)
		log.Println("Hitting The Save Link API")
		createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)
		log.Println("Validating the Response of the Create API...")
		assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
		assert.NotNil(suite.T(), "cluster id failed to create in db", createLinkResponseDto.Errors[0].UserMessage)
	})*/
	log.Println("=== Here we are Deleting the app after all verifications ===")
	testUtils.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
