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

	suite.Run("A=1=SaveDockerRegistryWithValidPayload", func() {
		saveDockerRegistryRequestDto := GetDockerRegistryRequestDto(false)
		byteValueOfSaveDockerRegistry, _ = json.Marshal(saveDockerRegistryRequestDto)

		log.Println("Hitting The post Docker registry API")
		saveDockerRegistryResponseDto = HitSaveContainerRegistryApi(byteValueOfSaveDockerRegistry, suite.authToken)

		log.Println("Validating the Response of the save docker registry API...")
		assert.Equal(suite.T(), saveDockerRegistryRequestDto.Id, saveDockerRegistryResponseDto.Result.Id)
		assert.Equal(suite.T(), saveDockerRegistryRequestDto.IsDefault, saveDockerRegistryResponseDto.Result.IsDefault)

	})

	suite.Run("A=2=SaveDockerRegistryWithPreviousId", func() {
		log.Println("Hitting DockerRegistryApi with same payload again")
		saveDockerRegistryOnceAgainResponseDto := HitSaveContainerRegistryApi(byteValueOfSaveDockerRegistry, suite.authToken)

		log.Println("Validating the Response of the save docker registry API...")
		assert.Equal(suite.T(), "docker registry failed to create in db", saveDockerRegistryOnceAgainResponseDto.Errors[0].InternalMessage)
	})

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteDockerRegistry := GetPayLoadForDeleteDockerRegistryAPI(saveDockerRegistryResponseDto.Result.Id, saveDockerRegistryResponseDto.Result.PluginId, saveDockerRegistryResponseDto.Result.RegistryUrl, saveDockerRegistryResponseDto.Result.RegistryType, saveDockerRegistryResponseDto.Result.Username, saveDockerRegistryResponseDto.Result.Password, saveDockerRegistryResponseDto.Result.IsDefault)
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteDockerRegistryApi(byteValueOfDeleteDockerRegistry, suite.authToken)
}
