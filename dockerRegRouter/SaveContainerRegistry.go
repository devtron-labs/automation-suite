package dockerRegRouter

import (
	"automation-suite/dockerRegRouter/ResponseDTOs"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *DockersRegRouterTestSuite) TestClassA6GetApp() {

	var byteValueOfSaveDockerRegistry []byte
	var saveDockerRegistryResponseDto ResponseDTOs.SaveDockerRegistryResponseDto

	suite.Run("A=1=SaveContainerRegistryWithValidPayload", func() {
		saveDockerRegistryRequestDto := GetDockerRegistryRequestDto(false)
		byteValueOfSaveDockerRegistry, _ = json.Marshal(saveDockerRegistryRequestDto)

		log.Println("Hitting The post Docker registry API")
		saveDockerRegistryResponseDto = HitSaveContainerRegistryApi(byteValueOfSaveDockerRegistry, suite.authToken)

		log.Println("Validating the Response of the save docker registry API...")
		assert.Equal(suite.T(), saveDockerRegistryRequestDto.Id, saveDockerRegistryResponseDto.Result.Id)
		assert.Equal(suite.T(), saveDockerRegistryRequestDto.IsDefault, saveDockerRegistryResponseDto.Result.IsDefault)
	})

	suite.Run("A=2=SaveContainerRegistryWithExistingName", func() {
		saveDockerRegistryRequestPayload := GetDockerRegistryRequestDto(false)
		byteValueOfSaveDockerRegistryPayload, _ := json.Marshal(saveDockerRegistryRequestPayload)

		log.Println("Hitting The save container registry Api First time")
		saveDockerRegistryResponse := HitSaveContainerRegistryApi(byteValueOfSaveDockerRegistryPayload, suite.authToken)

		log.Println("Hitting The save container registry Api second time with existing registry name")
		finalApiResponse := HitSaveContainerRegistryApi(byteValueOfSaveDockerRegistryPayload, suite.authToken)

		log.Println("Validating the Response of the save docker registry API...")
		assert.Equal(suite.T(), finalApiResponse.Errors[0].InternalMessage, "docker registry failed to create in db")
		assert.Equal(suite.T(), finalApiResponse.Errors[0].UserMessage, "requested by 2")

		log.Println("getting payload for Delete registry API")
		byteValueOfDeleteDockerRegistry := GetPayLoadForDeleteDockerRegistryAPI(saveDockerRegistryResponse.Result.Id, saveDockerRegistryResponse.Result.IpsConfig.Id, saveDockerRegistryResponse.Result.PluginId, saveDockerRegistryResponse.Result.RegistryUrl, saveDockerRegistryResponse.Result.RegistryType, saveDockerRegistryResponse.Result.Username, saveDockerRegistryResponse.Result.Password, saveDockerRegistryResponse.Result.IsDefault)
		log.Println("Hitting the Delete team API for Removing the data created via automation")
		HitDeleteDockerRegistryApi(byteValueOfDeleteDockerRegistry, suite.authToken)
	})

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteDockerRegistry := GetPayLoadForDeleteDockerRegistryAPI(saveDockerRegistryResponseDto.Result.Id, saveDockerRegistryResponseDto.Result.IpsConfig.Id, saveDockerRegistryResponseDto.Result.PluginId, saveDockerRegistryResponseDto.Result.RegistryUrl, saveDockerRegistryResponseDto.Result.RegistryType, saveDockerRegistryResponseDto.Result.Username, saveDockerRegistryResponseDto.Result.Password, saveDockerRegistryResponseDto.Result.IsDefault)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteDockerRegistryApi(byteValueOfDeleteDockerRegistry, suite.authToken)
}

//todo need to add assertion via getting the list of container registry before and after saving new registry
