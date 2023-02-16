package dockerRegRouter

import (
	"automation-suite/dockerRegRouter/RequestDTOs"
	"automation-suite/dockerRegRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"net/http"

	"github.com/stretchr/testify/suite"
)

type StructDockerRegRouter struct {
	saveDockerRegistryResponseDto ResponseDTOs.SaveDockerRegistryResponseDto
	deleteDockerRegistryResponse  ResponseDTOs.DeleteDockerRegistryResponse
	dockerRequestDTOs             RequestDTOs.SaveDockerRegistryRequestDTO
}

func (structDockerRegRouter StructDockerRegRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructDockerRegRouter {
	switch apiName {
	case DeleteDockerRegistry:
		json.Unmarshal(response, &structDockerRegRouter.deleteDockerRegistryResponse)
	case SaveDockerRegistryApi:
		json.Unmarshal(response, &structDockerRegRouter.saveDockerRegistryResponseDto)
	}
	return structDockerRegRouter
}

func GetDockerRegistryRequestDto(isDefault bool) RequestDTOs.SaveDockerRegistryRequestDTO {
	var saveDockerRegistryRequestDto RequestDTOs.SaveDockerRegistryRequestDTO
	ipsConfig := GetIpsConfig(0)
	//dockerRegistry, _ := GetDockerRegistry()
	envConf := Base.ReadBaseEnvConfig()
	file := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)
	saveDockerRegistryRequestDto.Id = "automation" + Base.GetRandomStringOfGivenLength(5)
	saveDockerRegistryRequestDto.IpsConfig = ipsConfig
	saveDockerRegistryRequestDto.IsDefault = isDefault
	saveDockerRegistryRequestDto.PluginId = file.PluginId
	saveDockerRegistryRequestDto.RegistryType = file.RegistryType
	saveDockerRegistryRequestDto.RegistryUrl = file.RegistryUrl
	saveDockerRegistryRequestDto.Username = file.DockerUsername
	saveDockerRegistryRequestDto.Password = file.Password
	return saveDockerRegistryRequestDto
}

func GetIpsConfig(id int) RequestDTOs.IpsConfig {
	var IpsConfig RequestDTOs.IpsConfig
	IpsConfig.Id = id
	IpsConfig.CredentialType = "SAME_AS_REGISTRY"
	IpsConfig.IgnoredClusterIdsCsv = "-1"
	return IpsConfig
}

func HitSaveContainerRegistryApi(payloadOfApi []byte, authToken string) ResponseDTOs.SaveDockerRegistryResponseDto {
	resp, err := Base.MakeApiCall(SaveDockerRegistryApiUrl, http.MethodPost, string(payloadOfApi), nil, authToken)
	Base.HandleError(err, SaveDockerRegistryApi)

	structDockerRegRouter := StructDockerRegRouter{}
	dockerRegRouter := structDockerRegRouter.UnmarshalGivenResponseBody(resp.Body(), SaveDockerRegistryApi)
	return dockerRegRouter.saveDockerRegistryResponseDto
}

func GetPayLoadForDeleteDockerRegistryAPI(registryName string, id int, pluginId string, regUrl string, regType string, username string, password string, isDefault bool) []byte {
	var saveDockerRegistryRequestDto RequestDTOs.SaveDockerRegistryRequestDTO
	ipsConfig := GetIpsConfig(id)
	saveDockerRegistryRequestDto.Id = registryName
	saveDockerRegistryRequestDto.IpsConfig = ipsConfig
	saveDockerRegistryRequestDto.IsDefault = isDefault
	saveDockerRegistryRequestDto.PluginId = pluginId
	saveDockerRegistryRequestDto.RegistryType = regType
	saveDockerRegistryRequestDto.RegistryUrl = regUrl
	saveDockerRegistryRequestDto.Username = username
	saveDockerRegistryRequestDto.Password = password
	byteValueOfStruct, _ := json.Marshal(saveDockerRegistryRequestDto)
	return byteValueOfStruct
}

func HitDeleteDockerRegistryApi(byteValueOfStruct []byte, authToken string) ResponseDTOs.DeleteDockerRegistryResponse {
	resp, err := Base.MakeApiCall(SaveDockerRegistryApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, DeleteDockerRegistry)

	structDockerRegRouter := StructDockerRegRouter{}
	dockerRegRouter := structDockerRegRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteDockerRegistry)
	return dockerRegRouter.deleteDockerRegistryResponse
}

type DockersRegRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *DockersRegRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
