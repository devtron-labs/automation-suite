package externalLinkout

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
)

func (suite *LinkTestSuite) FetchAllLinkouts() {
	log.Println("Hitting the 'FetchAllLink' Api before creating any new entry")
	fetchAllLinkResponseDto := HitFetchAllLinkApi()
	noOfTeams := len(fetchAllLinkResponseDto.Result)

	log.Println("Hitting the 'Save Link' Api for creating a new entry")
	saveLinkResponseDto := HitCreateLinkApi(nil, suite.authToken)

	log.Println("Hitting the FetchAllTeam API again for verifying the functionality of it")
	clusterId := map[string]string{
		"id": saveLinkResponseDto.Result.ClusterIds[0],
	}
	fetchAllLinkResponseDto = HitFetchAllLinkByClusterIdApi(clusterId)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), noOfTeams+1, len(fetchAllLinkResponseDto.Result))
	assert.Equal(suite.T(), saveLinkResponseDto.Result.Name, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Name)

	log.Println("getting payload for Delete Link API")
	byteValueOfStruct := GetPayLoadForDeleteLinkAPI(saveLinkResponseDto.Result.Id, saveLinkResponseDto.Result.Name, saveLinkResponseDto.Result.MonitoringToolId, saveLinkResponseDto.Result.Url, true)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfStruct, suite.authToken)
}

func (suite *LinkTestSuite) FetchAllLinkoutsWithValidClusterId() {
	log.Println("Hitting the 'FetchAllLink' Api before creating any new entry")
	fetchAllLinkResponseDto := HitFetchAllLinkApi()
	noOfTeams := len(fetchAllLinkResponseDto.Result)

	log.Println("Hitting the 'Save Link' Api for creating a new entry")
	saveLinkResponseDto := HitCreateLinkApi(nil, suite.authToken)

	log.Println("Hitting the FetchAllTeam API again for verifying the functionality of it")
	clusterId := map[string]string{
		"id": saveLinkResponseDto.Result.ClusterIds[0],
	}
	fetchAllLinkResponseDto = HitFetchAllLinkByClusterIdApi(clusterId)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), noOfTeams+1, len(fetchAllLinkResponseDto.Result))
	assert.Equal(suite.T(), saveLinkResponseDto.Result.Name, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Name)

	log.Println("getting payload for Delete Link API")
	byteValueOfStruct := GetPayLoadForDeleteLinkAPI(saveLinkResponseDto.Result.Id, saveLinkResponseDto.Result.Name, saveLinkResponseDto.Result.MonitoringToolId, saveLinkResponseDto.Result.Url, true)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfStruct, suite.authToken)
}

func (suite *LinkTestSuite) FetchAllLinkoutsWithInvalidClusterId() {
	log.Println("Hitting the 'FetchAllLink' Api before creating any new entry")
	fetchAllLinkResponseDto := HitFetchAllLinkApi()
	noOfTeams := len(fetchAllLinkResponseDto.Result)

	log.Println("Hitting the 'Save Link' Api for creating a new entry")
	saveLinkResponseDto := HitCreateLinkApi(nil, suite.authToken)

	log.Println("Hitting the FetchAllTeam API again for verifying the functionality of it")
	clusterId := map[string]string{
		"id": strconv.Itoa(rand.Intn(89-10) + 10),
	}
	fetchAllLinkResponseDto = HitFetchAllLinkByClusterIdApi(clusterId)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), noOfTeams+1, len(fetchAllLinkResponseDto.Result))
	assert.Equal(suite.T(), saveLinkResponseDto.Result.Name, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Name)

	log.Println("getting payload for Delete Link API")
	byteValueOfStruct := GetPayLoadForDeleteLinkAPI(saveLinkResponseDto.Result.Id, saveLinkResponseDto.Result.Name, saveLinkResponseDto.Result.MonitoringToolId, saveLinkResponseDto.Result.Url, true)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfStruct, suite.authToken)
}
