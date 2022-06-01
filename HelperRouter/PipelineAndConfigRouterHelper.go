package HelperRouter

import (
	Base "automation-suite/testUtils"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
)

type ConfigMapAndSecretDataRequestDTO struct {
	Id            int          `json:"id"`
	AppId         int          `json:"appId"`
	EnvironmentId int          `json:"environmentId"`
	ConfigData    []ConfigData `json:"configData"`
}

type Data struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

type SecretData struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Property string `json:"property,omitempty"`
	IsBinary bool   `json:"isBinary"`
}
type ConfigData struct {
	Name               string       `json:"name"`
	Data               Data         `json:"data"`
	DefaultData        Data         `json:"defaultData"`
	Global             bool         `json:"global"`
	SecretData         []SecretData `json:"secretData"`
	Type               string       `json:"type"`
	External           bool         `json:"external"`
	ExternalSecretType string       `json:"externalType"`
	MountPath          string       `json:"mountPath"`
	DefaultMountPath   string       `json:"defaultMountPath,omitempty"`
	RoleARN            string       `json:"roleARN"`
	SubPath            bool         `json:"subPath"`
	FilePermission     string       `json:"filePermission"`
}

type SaveConfigMapResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id         int          `json:"id"`
		AppId      int          `json:"appId"`
		ConfigData []ConfigData `json:"configData"`
	} `json:"result"`
}

type StructConfigMapRouter struct {
	saveConfigMapResponseDTO SaveConfigMapResponseDTO
}

func GetRequestPayloadForSecretOrConfig(configId int, configName string, appId int, userOfSecretAs string, externalType string, isSubPathNeeded bool, isFilePermissionNeeded bool, isSecret bool) ConfigMapAndSecretDataRequestDTO {
	configMapDataRequestDTO := ConfigMapAndSecretDataRequestDTO{}
	var configDataList = make([]ConfigData, 0)
	conf := ConfigData{}
	configMapDataRequestDTO.AppId = appId
	configMapDataRequestDTO.Id = configId
	configMapDataRequestDTO.EnvironmentId = 1
	switch externalType {
	case awsSystemManager:
		{
			data := GetConfigData(awsSystemManager+configName, userOfSecretAs, true, awsSystemManager, isSubPathNeeded, isFilePermissionNeeded, isSecret)
			data.RoleARN = "RoleARNAdmin"
			conf = data
		}
	case hashiCorpVault:
		{
			data := GetConfigData(hashiCorpVault+configName, userOfSecretAs, true, hashiCorpVault, isSubPathNeeded, isFilePermissionNeeded, isSecret)
			data.RoleARN = ""
			conf = data
		}
	case awsSecretsManager:
		{
			data := GetConfigData(awsSecretsManager+configName, userOfSecretAs, true, awsSecretsManager, isSubPathNeeded, isFilePermissionNeeded, isSecret)
			data.RoleARN = ""
			conf = data
		}
	case kubernetes:
		{
			data := GetConfigData(kubernetes+configName, userOfSecretAs, false, "", isSubPathNeeded, isFilePermissionNeeded, isSecret)
			data.RoleARN = ""
			data.Data = GetDataForConfigOrSecret(isSecret)
			conf = data
		}
	case externalKubernetes:
		{
			data := GetConfigData(externalKubernetes+configName, userOfSecretAs, false, kubernetes, isSubPathNeeded, isFilePermissionNeeded, isSecret)
			data.RoleARN = ""
			conf = data
		}
	}
	configDataList = append(configDataList, conf)
	configMapDataRequestDTO.ConfigData = configDataList
	return configMapDataRequestDTO
}

func GetConfigData(configName string, userOfSecretAs string, isSecretDataNeeded bool, externalSecretType string, isSubPathNeeded bool, isFilePermissionRequired bool, isSecret bool) ConfigData {
	conf := ConfigData{}
	conf.Name = configName
	conf.Type = userOfSecretAs
	if userOfSecretAs == volume {
		conf.MountPath = "/directory-path"
		conf.SubPath = isSubPathNeeded

		if isFilePermissionRequired {
			conf.FilePermission = "0744"
		}

	}
	conf.External = true
	conf.ExternalSecretType = externalSecretType
	if isSecretDataNeeded {
		conf.SecretData = GetSecretData(isSecret)
	}
	return conf
}

func GetSecretData(isSecret bool) []SecretData {
	var secretDataList = make([]SecretData, 0)
	data := SecretData{}
	var key, name string
	if isSecret {
		key = base64.StdEncoding.EncodeToString([]byte("service/credentials"))
		name = base64.StdEncoding.EncodeToString([]byte("secret-key"))
	} else {
		key = "service/credentials"
		name = "secret-key"
	}
	data.Key = key
	data.Name = name
	secretDataList = append(secretDataList, data)
	return secretDataList
}

func GetDataForConfigOrSecret(isSecret bool) Data {
	data := Data{}
	var value1, value2 string
	if isSecret {
		value1 = base64.StdEncoding.EncodeToString([]byte("value1"))
		value2 = base64.StdEncoding.EncodeToString([]byte("value2"))
	} else {
		value1 = "value1"
		value2 = "value2"
	}
	data.Key1 = value1
	data.Key2 = value2
	return data
}

func (structConfigMapRouter StructConfigMapRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructConfigMapRouter {
	switch apiName {
	case SaveGlobalConfigmapApi:
		json.Unmarshal(response, &structConfigMapRouter.saveConfigMapResponseDTO)

	}
	return structConfigMapRouter
}

func HitSaveGlobalConfigMap(payload []byte, authToken string) SaveConfigMapResponseDTO {
	resp, err := Base.MakeApiCall(SaveGlobalConfigmapApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveGlobalConfigmapApi)
	structConfigMapRouter := StructConfigMapRouter{}
	configMapRouter := structConfigMapRouter.UnmarshalGivenResponseBody(resp.Body(), SaveGlobalConfigmapApi)
	return configMapRouter.saveConfigMapResponseDTO
}

func HitGetEnvironmentConfigMap(appId int, envId int, authToken string) SaveConfigMapResponseDTO {
	id := strconv.Itoa(appId)
	environmentId := strconv.Itoa(envId)
	apiUrl := GetEnvironmentConfigMapApiUrl + id + "/" + environmentId
	resp, err := Base.MakeApiCall(apiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetEnvironmentConfigMapApi)

	structConfigMapRouter := StructConfigMapRouter{}
	pipelineConfigRouter := structConfigMapRouter.UnmarshalGivenResponseBody(resp.Body(), SaveGlobalConfigmapApi)
	return pipelineConfigRouter.saveConfigMapResponseDTO
}

func HitSaveGlobalSecretApi(payload []byte, authToken string) SaveConfigMapResponseDTO {
	resp, err := Base.MakeApiCall(SaveGlobalSecretApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveGlobalSecretApi)
	structConfigMapRouter := StructConfigMapRouter{}
	configMapRouter := structConfigMapRouter.UnmarshalGivenResponseBody(resp.Body(), SaveGlobalConfigmapApi)
	return configMapRouter.saveConfigMapResponseDTO
}

func HitGetEnvironmentSecretApi(appId int, envId int, authToken string) SaveConfigMapResponseDTO {
	id := strconv.Itoa(appId)
	environmentId := strconv.Itoa(envId)
	apiUrl := GetEnvSecretApiUrl + id + "/" + environmentId
	resp, err := Base.MakeApiCall(apiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetEnvironmentConfigMapApi)
	structConfigMapRouter := StructConfigMapRouter{}
	pipelineConfigRouter := structConfigMapRouter.UnmarshalGivenResponseBody(resp.Body(), SaveGlobalConfigmapApi)
	return pipelineConfigRouter.saveConfigMapResponseDTO
}
