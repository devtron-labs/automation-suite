package ConfigMapRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"log"
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

func getRequestPayloadForSecretOrConfig(configId int, configName string, appId int, userOfSecretAs string, externalType string, isSubPathNeeded bool, isFilePermissionNeeded bool) ConfigMapAndSecretDataRequestDTO {
	configMapDataRequestDTO := ConfigMapAndSecretDataRequestDTO{}
	var configDataList = make([]ConfigData, 0)
	conf := ConfigData{}
	configMapDataRequestDTO.AppId = appId
	configMapDataRequestDTO.Id = configId
	configMapDataRequestDTO.EnvironmentId = 1
	switch externalType {
	case AWSSystemManager:
		{
			data := GetConfigData(AWSSystemManager+configName, userOfSecretAs, true, AWSSystemManager, isSubPathNeeded, isFilePermissionNeeded)
			data.RoleARN = "RoleARNAdmin"
			conf = data
		}
	case HashiCorpVault:
		{
			data := GetConfigData(HashiCorpVault+configName, userOfSecretAs, true, HashiCorpVault, isSubPathNeeded, isFilePermissionNeeded)
			data.RoleARN = ""
			conf = data
		}
	case AWSSecretsManager:
		{
			data := GetConfigData(AWSSecretsManager+configName, userOfSecretAs, true, AWSSecretsManager, isSubPathNeeded, isFilePermissionNeeded)
			data.RoleARN = ""
			conf = data
		}
	case Kubernetes:
		{
			data := GetConfigData(Kubernetes+configName, userOfSecretAs, false, "", isSubPathNeeded, isFilePermissionNeeded)
			data.RoleARN = ""
			data.Data = GetDataForConfigOrSecret()
			conf = data
		}
	case ExternalKubernetes:
		{
			data := GetConfigData(ExternalKubernetes+configName, userOfSecretAs, false, Kubernetes, isSubPathNeeded, isFilePermissionNeeded)
			data.RoleARN = ""
			conf = data
		}
	}
	configDataList = append(configDataList, conf)
	configMapDataRequestDTO.ConfigData = configDataList
	return configMapDataRequestDTO
}

func GetConfigData(configName string, userOfSecretAs string, isSecretDataNeeded bool, externalSecretType string, isSubPathNeeded bool, isFilePermissionRequired bool) ConfigData {
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
		conf.SecretData = GetSecretData()
	}
	return conf
}

func GetSecretData() []SecretData {
	var secretDataList = make([]SecretData, 0)
	data := SecretData{}
	data.Key = "service/credentials"
	data.Name = "secret-key"
	data.IsBinary = true
	data.Property = "property-name"
	secretDataList = append(secretDataList, data)
	return secretDataList
}

func GetDataForConfigOrSecret() Data {
	data := Data{}
	data.Key1 = "value1"
	data.Key2 = "value2"
	return data
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

func HitSaveEnvironmentSecret(payload []byte, authToken string) SaveConfigMapResponseDTO {
	resp, err := Base.MakeApiCall(SaveGlobalSecretApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveGlobalSecretApi)
	structConfigMapRouter := StructConfigMapRouter{}
	configMapRouter := structConfigMapRouter.UnmarshalGivenResponseBody(resp.Body(), SaveGlobalConfigmapApi)
	return configMapRouter.saveConfigMapResponseDTO
}

func (structConfigMapRouter StructConfigMapRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructConfigMapRouter {
	switch apiName {
	case SaveGlobalConfigmapApi:
		json.Unmarshal(response, &structConfigMapRouter.saveConfigMapResponseDTO)

	}
	return structConfigMapRouter
}

// ConfigsMapRouterTestSuite =================PipelineConfigSuite Setup =========================
type ConfigsMapRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *ConfigsMapRouterTestSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
}
