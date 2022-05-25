package ConfigMapRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"strconv"
)

type ConfigMapDataRequestDTO struct {
	Id         int          `json:"id"`
	AppId      int          `json:"appId"`
	ConfigData []ConfigData `json:"configData"`
}

type Data struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}
type ConfigData struct {
	Name             string      `json:"name"`
	Data             Data        `json:"data"`
	DefaultData      Data        `json:"defaultData"`
	Global           bool        `json:"global"`
	SecretData       interface{} `json:"secretData"`
	Type             string      `json:"type"`
	External         bool        `json:"external"`
	MountPath        string      `json:"mountPath"`
	DefaultMountPath string      `json:"defaultMountPath,omitempty"`
	RoleARN          string      `json:"roleARN"`
	SubPath          bool        `json:"subPath"`
	FilePermission   string      `json:"filePermission"`
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

func getRequestPayloadForSaveConfigmap(configId int, configName string, appId int, typeReq string, external bool, subPath bool, isFilePermissionRequired bool, updateData bool) ConfigMapDataRequestDTO {
	configMapDataRequestDTO := ConfigMapDataRequestDTO{}
	configMapDataRequestDTO.AppId = appId
	configMapDataRequestDTO.Id = configId
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

func HitSaveConfigMap(payload []byte, authToken string) SaveConfigMapResponseDTO {
	resp, err := Base.MakeApiCall(SaveConfigmapApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveConfigmapApi)
	structConfigMapRouter := StructConfigMapRouter{}
	configMapRouter := structConfigMapRouter.UnmarshalGivenResponseBody(resp.Body(), SaveConfigmapApi)
	return configMapRouter.saveConfigMapResponseDTO
}

func HitGetEnvironmentConfigMap(appId int, envId int, authToken string) SaveConfigMapResponseDTO {
	id := strconv.Itoa(appId)
	envirId := strconv.Itoa(envId)
	apiUrl := GetEnvironmentConfigMapApiUrl + id + "/" + envirId
	resp, err := Base.MakeApiCall(apiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetEnvironmentConfigMapApi)

	structConfigMapRouter := StructConfigMapRouter{}
	pipelineConfigRouter := structConfigMapRouter.UnmarshalGivenResponseBody(resp.Body(), SaveConfigmapApi)
	return pipelineConfigRouter.saveConfigMapResponseDTO
}

func (structConfigMapRouter StructConfigMapRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructConfigMapRouter {
	switch apiName {
	case SaveConfigmapApi:
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
