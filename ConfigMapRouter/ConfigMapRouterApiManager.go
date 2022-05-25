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

func getRequestPayloadForSaveSecretOrConfigmap(configId int, configName string, appId int, typeReq string, external bool, subPath bool, isFilePermissionRequired bool, updateData bool) ConfigMapAndSecretDataRequestDTO {
	configMapDataRequestDTO := ConfigMapAndSecretDataRequestDTO{}
	configMapDataRequestDTO.AppId = appId
	configMapDataRequestDTO.Id = configId
	configMapDataRequestDTO.EnvironmentId = 1
	var configDataList = make([]ConfigData, 0)
	conf := ConfigData{}
	conf.Name = configName
	conf.Type = typeReq

	if typeReq == "volume" {
		conf.MountPath = "/directory-path"
		conf.SubPath = subPath

		if isFilePermissionRequired {
			conf.FilePermission = "0744"
		}
	}
	conf.External = external
	if !external {
		conf.Data.Key1 = "value1"
		if updateData {
			conf.Data.Key2 = "value2"
		}
	}
	if external && subPath {
		conf.Data.Key1 = ""
		if updateData {
			conf.Data.Key2 = ""
		}
	}
	configDataList = append(configDataList, conf)
	configMapDataRequestDTO.ConfigData = configDataList
	return configMapDataRequestDTO
}

func getRequestPayloadForSecret(configId int, configName string, appId int, userOfSecretAs string, externalType string) ConfigMapAndSecretDataRequestDTO {
	configMapDataRequestDTO := ConfigMapAndSecretDataRequestDTO{}
	var configDataList = make([]ConfigData, 0)
	conf := ConfigData{}
	configMapDataRequestDTO.AppId = appId
	configMapDataRequestDTO.Id = configId
	configMapDataRequestDTO.EnvironmentId = 1
	switch externalType {
	case AWSSystemManager:
		{
			data := GetConfigData(configName, userOfSecretAs, true, AWSSystemManager)
			data.RoleARN = "RoleARNAdmin"
			conf = data
		}
	case HashiCorpVault:
		{
			data := GetConfigData(configName, userOfSecretAs, true, HashiCorpVault)
			data.RoleARN = ""
			conf = data
		}
	case AWSSecretsManager:
		{
			data := GetConfigData(configName, userOfSecretAs, true, AWSSecretsManager)
			data.RoleARN = ""
			conf = data
		}
	case KubernetesSecret:
		{
			data := GetConfigData(configName, userOfSecretAs, false, "")
			data.RoleARN = ""
			data.Data = GetDataForConfigOrSecret()
			conf = data
		}
	case ExternalKubernetesSecret:
		{
			data := GetConfigData(configName, userOfSecretAs, true, KubernetesSecret)
			data.RoleARN = ""
			conf = data
		}
	}
	configDataList = append(configDataList, conf)
	configMapDataRequestDTO.ConfigData = configDataList
	return configMapDataRequestDTO
}

func GetConfigData(configName string, userOfSecretAs string, isSecretDataNeeded bool, externalSecretType string) ConfigData {
	conf := ConfigData{}
	conf.Name = configName
	conf.Type = userOfSecretAs
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
	resp, err := Base.MakeApiCall(SaveEnvironmentSecretApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveEnvironmentSecretApi)
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
	authToken            string
	createAppResponseDto Base.CreateAppResponseDto
	deleteResponseDto    Base.DeleteResponseDto
}

func (suite *ConfigsMapRouterTestSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
	suite.createAppResponseDto = Base.CreateApp(suite.authToken)
}

func (suite *ConfigsMapRouterTestSuite) TearDownSuite() {
	log.Println("=== Running the after suite method for deleting the data created via automation ===")
	createAppResponse := suite.createAppResponseDto
	suite.deleteResponseDto = Base.DeleteApp(createAppResponse.Result.Id, createAppResponse.Result.AppName, createAppResponse.Result.TeamId, createAppResponse.Result.TemplateId, suite.authToken)
}
