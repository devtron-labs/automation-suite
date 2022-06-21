package dockerRegRouter

import (
	"automation-suite/dockerRegRouter/RequestDTOs"
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
	dockerRequestDTOs             RequestDTOs.SaveDockerRegistryRequestDto
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
	PluginId     string `env:"PLUGINID" envDefault:"cd.go.artifact.docker.registry"`
	RegistryType string `env:"REGISTRYTYPE" envDefault:"docker-hub"`
	RegistryUrl  string `env:"REGISTRYURL" envDefault:"docker.io"`
	Username     string `env:"DOCKER_USERNAME" envDefault:""`
	Password     string `env:"DOCKER_PASSWORD" envDefault:""`
}

func GetDockerRegistry() (*DockerRegistry, error) {
	cfg := &DockerRegistry{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from ChartRepoRouterConfig")
	}
	return cfg, err
}

func GetDockerRegistryRequestDto(isDefault bool) RequestDTOs.SaveDockerRegistryRequestDto {
	var saveDockerRegistryRequestDto RequestDTOs.SaveDockerRegistryRequestDto
	dockerRegistry, _ := GetDockerRegistry()
	saveDockerRegistryRequestDto.Id = dockerRegistry.Username
	saveDockerRegistryRequestDto.PluginId = dockerRegistry.PluginId
	saveDockerRegistryRequestDto.RegistryType = dockerRegistry.RegistryType
	saveDockerRegistryRequestDto.RegistryUrl = dockerRegistry.RegistryUrl
	saveDockerRegistryRequestDto.IsDefault = isDefault
	saveDockerRegistryRequestDto.Username = dockerRegistry.Username
	saveDockerRegistryRequestDto.Password = dockerRegistry.Password
	return saveDockerRegistryRequestDto
}

func HitSaveContainerRegistryApi(payloadOfApi []byte, authToken string) ResponseDTOs.SaveDockerRegistryResponseDto {
	resp, err := Base.MakeApiCall(SaveDockerRegistryApiUrl, http.MethodPost, string(payloadOfApi), nil, authToken)
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
