package dockerRegRouter

import (
	RequestDTOs "automation-suite/dockerRegRouter/DockerRequestDTOs"
	"automation-suite/dockerRegRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
)

type StructDockerRegRouter struct {
	saveDockerRegistryResponseDto ResponseDTOs.SaveDockerRegistryResponseDto
	deleteDockerRegistryResponse  ResponseDTOs.DeleteDockerRegistryResponse
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

type DockerRegistry struct {
	Id           string `env:"ID" envDefault:""`
	PluginId     string `env:"PLUGINID" envDefault:""`
	RegistryType string `env:"REGISTRYTYPE" envDefault:""`
	RegistryUrl  string `env:"REGISTRYURL" envDefault:""`
	Username     string `env:"USERNAME" envDefault:""`
	Password     string `env:"PASSWORD" envDefault:""`
}

func GetDockerRegistry() (*DockerRegistry, error) {
	cfg := &DockerRegistry{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from ChartRepoRouterConfig")
	}
	return cfg, err
}

func GetDockerRegistryRequestDto(isRepeat bool, id string, pluginId string, regType string, regUrl string, isDefault bool, username string, password string) RequestDTOs.SaveDockerRegistryRequestDto {
	if isRepeat == false {
		dockerRegistry, _ := GetDockerRegistry()
		var saveDockerRegistryRequestDto RequestDTOs.SaveDockerRegistryRequestDto
		saveDockerRegistryRequestDto.Id = Base.GetRandomStringOfGivenLength(10)
		saveDockerRegistryRequestDto.PluginId = dockerRegistry.PluginId
		saveDockerRegistryRequestDto.RegistryType = dockerRegistry.RegistryType
		saveDockerRegistryRequestDto.RegistryUrl = dockerRegistry.RegistryUrl
		saveDockerRegistryRequestDto.IsDefault = isDefault
		saveDockerRegistryRequestDto.Username = dockerRegistry.Username
		saveDockerRegistryRequestDto.Password = dockerRegistry.Password
		return saveDockerRegistryRequestDto
	}

	var saveDockerRegistryRequestDto RequestDTOs.SaveDockerRegistryRequestDto
	saveDockerRegistryRequestDto.Id = id
	saveDockerRegistryRequestDto.PluginId = pluginId
	saveDockerRegistryRequestDto.RegistryType = regType
	saveDockerRegistryRequestDto.RegistryUrl = regUrl
	saveDockerRegistryRequestDto.IsDefault = isDefault
	saveDockerRegistryRequestDto.Username = username
	saveDockerRegistryRequestDto.Password = password
	return saveDockerRegistryRequestDto
}

func HitSaveDockerRegistryApi(isRepeat bool, payload []byte, id string, pluginId string, regUrl string, regType string, username string, password string, isDefault bool, authToken string) ResponseDTOs.SaveDockerRegistryResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		if isRepeat == false {
			dockerRegistry, _ := GetDockerRegistry()
			var saveDockerRegistryRequestDto RequestDTOs.SaveDockerRegistryRequestDto
			saveDockerRegistryRequestDto.Id = dockerRegistry.Id
			saveDockerRegistryRequestDto.PluginId = dockerRegistry.PluginId
			saveDockerRegistryRequestDto.RegistryType = dockerRegistry.RegistryType
			saveDockerRegistryRequestDto.RegistryUrl = dockerRegistry.RegistryUrl
			saveDockerRegistryRequestDto.IsDefault = isDefault
			saveDockerRegistryRequestDto.Username = dockerRegistry.Username
			saveDockerRegistryRequestDto.Password = dockerRegistry.Password
			byteValueOfStruct, _ := json.Marshal(saveDockerRegistryRequestDto)
			payloadOfApi = string(byteValueOfStruct)
		} else {
			var saveDockerRegistryRequestDto RequestDTOs.SaveDockerRegistryRequestDto
			saveDockerRegistryRequestDto.Id = id
			saveDockerRegistryRequestDto.PluginId = pluginId
			saveDockerRegistryRequestDto.RegistryType = regType
			saveDockerRegistryRequestDto.RegistryUrl = regUrl
			saveDockerRegistryRequestDto.IsDefault = isDefault
			saveDockerRegistryRequestDto.Username = username
			saveDockerRegistryRequestDto.Password = password
			byteValueOfStruct, _ := json.Marshal(saveDockerRegistryRequestDto)
			payloadOfApi = string(byteValueOfStruct)
		}
	}

	resp, err := Base.MakeApiCall(SaveDockerRegistryApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, SaveDockerRegistryApi)

	structDockerRegRouter := StructDockerRegRouter{}
	dockerRegRouter := structDockerRegRouter.UnmarshalGivenResponseBody(resp.Body(), SaveDockerRegistryApi)
	return dockerRegRouter.saveDockerRegistryResponseDto
}
func GetPayLoadForDeleteDockerRegistryAPI(id string, pluginId string, regUrl string, regType string, username string, password string, isDefault bool) []byte {
	var saveDockerRegistryRequestDto RequestDTOs.SaveDockerRegistryRequestDto
	saveDockerRegistryRequestDto.Id = id
	saveDockerRegistryRequestDto.PluginId = pluginId
	saveDockerRegistryRequestDto.RegistryUrl = regUrl
	saveDockerRegistryRequestDto.RegistryType = regType
	saveDockerRegistryRequestDto.Username = username
	saveDockerRegistryRequestDto.Password = password
	saveDockerRegistryRequestDto.IsDefault = isDefault
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
