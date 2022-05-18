package DockerRegRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type DockerRegRouter struct {
	suite.Suite
	authToken string
}

func (suite *DockerRegRouter) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
	fmt.Println(suite.authToken)
}

type InstallationScriptStruct struct {
	saveDockerRegistryResponseDto SaveDockerRegistryResponseDto
	deleteDockerRegistryResponse  DeleteDockerRegistryResponse
}

func (installationScriptStruct InstallationScriptStruct) UnmarshalGivenResponseBody(response []byte, apiName string) InstallationScriptStruct {
	switch apiName {
	case DeleteDockerRegistry:
		json.Unmarshal(response, &installationScriptStruct.deleteDockerRegistryResponse)
	case SaveDockerRegistryApi:
		json.Unmarshal(response, &installationScriptStruct.saveDockerRegistryResponseDto)
	}
	return installationScriptStruct
}

type SaveDockerRegistryRequestDto struct {
	Id           string `json:"id"`
	PluginId     string `json:"pluginId"`
	RegistryType string `json:"registryType"`
	IsDefault    bool   `json:"isDefault"`
	RegistryUrl  string `json:"registryUrl"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type DeleteDockerRegistryResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}
type SaveDockerRegistryResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id           string `json:"id"`
		PluginId     string `json:"pluginId"`
		RegistryUrl  string `json:"registryUrl"`
		RegistryType string `json:"registryType"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		IsDefault    bool   `json:"isDefault"`
		Connection   string `json:"connection"`
		Cert         string `json:"cert"`
		Active       bool   `json:"active"`
	} `json:"result"`

	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
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

func GetDockerRegistryRequestDto(isRepeat bool, id string, pluginId string, regType string, regUrl string, isDefault bool, username string, password string) SaveDockerRegistryRequestDto {
	if isRepeat == false {
		dockerRegistry, _ := GetDockerRegistry()
		var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
		saveDockerRegistryRequestDto.Id = Base.GetRandomStringOfGivenLength(10)
		saveDockerRegistryRequestDto.PluginId = dockerRegistry.PluginId
		saveDockerRegistryRequestDto.RegistryType = dockerRegistry.RegistryType
		saveDockerRegistryRequestDto.RegistryUrl = dockerRegistry.RegistryUrl
		saveDockerRegistryRequestDto.IsDefault = false
		saveDockerRegistryRequestDto.Username = dockerRegistry.Username
		saveDockerRegistryRequestDto.Password = dockerRegistry.Password
		return saveDockerRegistryRequestDto
	}

	var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
	saveDockerRegistryRequestDto.Id = id
	saveDockerRegistryRequestDto.PluginId = pluginId
	saveDockerRegistryRequestDto.RegistryType = regType
	saveDockerRegistryRequestDto.RegistryUrl = regUrl
	saveDockerRegistryRequestDto.IsDefault = isDefault
	saveDockerRegistryRequestDto.Username = username
	saveDockerRegistryRequestDto.Password = password
	return saveDockerRegistryRequestDto
}

func HitSaveDockerRegistryApi(isRepeat bool, payload []byte, id string, pluginId string, regUrl string, regType string, username string, password string, isDefault bool, authToken string) SaveDockerRegistryResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		if isRepeat == false {
			dockerRegistry, _ := GetDockerRegistry()
			var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
			saveDockerRegistryRequestDto.Id = dockerRegistry.Id
			saveDockerRegistryRequestDto.PluginId = dockerRegistry.PluginId
			saveDockerRegistryRequestDto.RegistryType = dockerRegistry.RegistryType
			saveDockerRegistryRequestDto.RegistryUrl = dockerRegistry.RegistryUrl
			saveDockerRegistryRequestDto.IsDefault = false
			saveDockerRegistryRequestDto.Username = dockerRegistry.Username
			saveDockerRegistryRequestDto.Password = dockerRegistry.Password
			byteValueOfStruct, _ := json.Marshal(saveDockerRegistryRequestDto)
			payloadOfApi = string(byteValueOfStruct)
		} else {
			var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
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

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), SaveDockerRegistryApi)
	return installationScriptRouter.saveDockerRegistryResponseDto
}
func GetPayLoadForDeleteDockerRegistryAPI(id string, pluginId string, regUrl string, regType string, username string, password string, isDefault bool) []byte {
	var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
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

func HitDeleteDockerRegistryApi(byteValueOfStruct []byte, authToken string) DeleteDockerRegistryResponse {
	resp, err := Base.MakeApiCall(SaveDockerRegistryApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, DeleteDockerRegistry)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteDockerRegistry)
	return apiRouter.deleteDockerRegistryResponse
}
