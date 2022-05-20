package externalLinkout

import (
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *LinkTestSuite) FetchAllToolsWithValidPayload() {
	log.Println("Hitting the 'FetchAllTools' Api before creating any new entry")
	fetchAllToolsResponseDto := HitFetchAllToolsApi()
	noOfTools := len(fetchAllToolsResponseDto.Result)

	log.Println("Hitting the 'Save Tools' Api for creating a new entry")
	createToolResponseDto := HitCreateToolApi(nil)

	log.Println("Hitting the FetchAllTool API again for verifying the functionality of it")
	fetchAllToolsResponseDto = HitFetchAllToolsApi()

	log.Println("Validating the response of FetchAllTool API")
	assert.Equal(suite.T(), noOfTools+1, len(fetchAllToolsResponseDto.Result))
	assert.Equal(suite.T(), createToolResponseDto.Result.Name, fetchAllToolsResponseDto.Result[len(fetchAllToolsResponseDto.Result)-1].Name)

	log.Println("getting payload for Delete Tool API")
	byteValueOfStruct := GetPayLoadForDeleteToolAPI(createToolResponseDto.Result.Id, createToolResponseDto.Result.Name, createToolResponseDto.Result.Icon)
	log.Println("Hitting the Delete Tool API for Removing the data created via automation")
	HitDeleteToolApi(byteValueOfStruct)
}
