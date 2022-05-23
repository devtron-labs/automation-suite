package externalLinkout

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *LinkOutRouterTestSuite) CreateLinkoutWithValidPayload() {

	createLinkRequestDto := GetSaveLinkRequestDto()
	byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)

	log.Println("Hitting The Save Link API")
	createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)

	log.Println("Validating the Response of the Create API...")
	assert.Equal(suite.T(), createLinkRequestDto.Name, createLinkResponseDto.Result.Name)
	assert.Equal(suite.T(), 200, createLinkResponseDto.Code)
	assert.NotNil(suite.T(), createLinkResponseDto.Result.Id)

	log.Println("getting payload for Delete Link API")
	byteValueOfStruct = GetPayLoadForDeleteLinkAPI(createLinkResponseDto.Result.Id, createLinkResponseDto.Result.Name, createLinkResponseDto.Result.MonitoringToolId, createLinkResponseDto.Result.Url, true)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfStruct, suite.authToken)
}

func (suite *LinkOutRouterTestSuite) CreateLinkoutWithInvalidToolId() {
	createLinkRequestDto := GetSaveLinkRequestInvalidMonitoringToolIdDto()
	byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)

	log.Println("Hitting The Save Link API")
	createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)

	log.Println("Validating the Response of the Create API...")
	assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
	assert.NotNil(suite.T(), "Internal Server Error", createLinkResponseDto.Status)
	log.Println("getting payload for Delete Link API")
	byteValueOfStruct = GetPayLoadForDeleteLinkAPI(createLinkResponseDto.Result.Id, createLinkResponseDto.Result.Name, createLinkResponseDto.Result.MonitoringToolId, createLinkResponseDto.Result.Url, true)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfStruct, suite.authToken)
}

func (suite *LinkOutRouterTestSuite) CreateLinkoutWithInvalidClusterId() {
	createLinkRequestDto := GetSaveLinkRequestInvalidClusterIdDto()
	byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)

	log.Println("Hitting The Save Link API")
	createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)

	log.Println("Validating the Response of the Create API...")
	assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
	assert.NotNil(suite.T(), "cluster id failed to create in db", createLinkResponseDto.Errors[0].InternalMessage)
	log.Println("getting payload for Delete Link API")
	byteValueOfStruct = GetPayLoadForDeleteLinkAPI(createLinkResponseDto.Result.Id, createLinkResponseDto.Result.Name, createLinkResponseDto.Result.MonitoringToolId, createLinkResponseDto.Result.Url, true)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfStruct, suite.authToken)
}

func (suite *LinkOutRouterTestSuite) CreateLinkoutWithOneValidOneInvalidClusterId() {
	createLinkRequestDto := GetSaveLinkRequestOneValidOneInvalidClusterId()
	byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)

	log.Println("Hitting The Save Link API")
	createLinkResponseDto := HitCreateLinkApi(byteValueOfStruct, suite.authToken)

	log.Println("Validating the Response of the Create API...")
	assert.Equal(suite.T(), 500, createLinkResponseDto.Code)
	assert.NotNil(suite.T(), "cluster id failed to create in db", createLinkResponseDto.Errors[0].InternalMessage)
	log.Println("getting payload for Delete Link API")
	byteValueOfStruct = GetPayLoadForDeleteLinkAPI(createLinkResponseDto.Result.Id, createLinkResponseDto.Result.Name, createLinkResponseDto.Result.MonitoringToolId, createLinkResponseDto.Result.Url, true)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfStruct, suite.authToken)
}
