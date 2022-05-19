package PipelineConfigRouter

import (
	"automation-suite/dockerRegRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type DeleteResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type CreateAppRequestDto struct {
	Id         int    `json:"id"`
	AppName    string `json:"appName"`
	TeamId     int    `json:"teamId"`
	TemplateId int    `json:"templateId"`
}
type CreateAppResponseDto struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Result CreateAppRequestDto `json:"result"`
	Errors []Base.Errors       `json:"errors"`
}

type DockerBuildConfig struct {
	DockerfilePath string `json:"dockerfilePath"`
	Args           struct {
	} `json:"args"`
	DockerfileRepository   string `json:"dockerfileRepository"`
	DockerfileRelativePath string `json:"dockerfileRelativePath"`
	GitMaterialId          int    `json:"gitMaterialId"`
}
type SaveAppCiPipelineRequestDTO struct {
	Id                interface{}       `json:"id"`
	AppId             int               `json:"appId"`
	DockerRegistry    string            `json:"dockerRegistry"`
	DockerRepository  string            `json:"dockerRepository"`
	BeforeDockerBuild []interface{}     `json:"beforeDockerBuild"`
	DockerBuildConfig DockerBuildConfig `json:"dockerBuildConfig"`
	AfterDockerBuild  []interface{}     `json:"afterDockerBuild"`
	AppName           string            `json:"appName"`
}

type CreateAppMaterialRequestDto struct {
	AppId     int            `json:"appId"`
	Materials []AppMaterials `json:"material"`
}

type AppMaterials struct {
	Url             string `json:"url"`
	Id              int    `json:"id"`
	GitProviderId   int    `json:"gitProviderId"`
	CheckoutPath    string `json:"checkoutPath"`
	FetchSubmodules bool   `json:"fetchSubmodules"`
}

type CreateAppMaterialResponseDto struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Result AppDetails `json:"result"`
}
type AppDetails struct {
	AppId    int            `json:"appId"`
	Material []AppMaterials `json:"material"`
}

type SaveAppCiPipelineResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		AppName string `json:"appName"`
		AppId   int    `json:"appId"`
	} `json:"result"`
	Errors []Base.Errors `json:"errors"`
}

type GetCiPipelineViaIdResponseDTO struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Errors []Base.Errors `json:"errors"`
	Result struct {
		Id                int               `json:"id"`
		AppId             int               `json:"appId"`
		DockerRegistry    string            `json:"dockerRegistry"`
		DockerRepository  string            `json:"dockerRepository"`
		DockerBuildConfig DockerBuildConfig `json:"dockerBuildConfig"`
		AppName           string            `json:"appName"`
		Materials         interface{}       `json:"materials"`
		ScanEnabled       bool              `json:"scanEnabled"`
	} `json:"result"`
}

type GetContainerRegistryResponseDTO struct {
	Code   int                                             `json:"code"`
	Status string                                          `json:"status"`
	Result []*dockerRegRouter.SaveDockerRegistryRequestDto `json:"result"`
}

type StructPipelineConfigRouter struct {
	saveAppCiPipelineRequestDTO     SaveAppCiPipelineRequestDTO
	createAppResponseDto            CreateAppResponseDto
	deleteResponseDto               DeleteResponseDto
	createAppMaterialRequestDto     CreateAppMaterialRequestDto
	createAppMaterialResponseDto    CreateAppMaterialResponseDto
	saveAppCiPipelineResponseDTO    SaveAppCiPipelineResponseDTO
	getCiPipelineViaIdResponseDTO   GetCiPipelineViaIdResponseDTO
	getContainerRegistryResponseDTO GetContainerRegistryResponseDTO
}

type EnvironmentConfigPipelineConfigRouter struct {
	GitHubProjectUrl       string `env:"GITHUB_URL_TO_CLONE_PROJECT" envDefault:"https://github.com/devtron-labs/sample-go-app.git"`
	DockerRegistry         string `env:"DOCKER_REGISTRY" envDefault:"erdipak"`
	DockerfilePath         string `env:"DOCKER_FILE_PATH" envDefault:"./Dockerfile"`
	DockerfileRepository   string `env:"DOCKER_FILE_REPO" envDefault:"sample-go-app"`
	DockerfileRelativePath string `env:"DOCKER_FILE_RELATIVE_PATH" envDefault:"Dockerfile"`
}

func GetEnvironmentConfigPipelineConfigRouter() (*EnvironmentConfigPipelineConfigRouter, error) {
	cfg := &EnvironmentConfigPipelineConfigRouter{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

func GetAppRequestDto(appName string, teamId int, templateId int) CreateAppRequestDto {
	var createAppRequestDto CreateAppRequestDto
	createAppRequestDto.AppName = appName
	createAppRequestDto.TeamId = teamId
	createAppRequestDto.TemplateId = templateId
	return createAppRequestDto
}

func GetPayLoadForDeleteAppAPI(id int, appName string, teamId int, templateId int) []byte {
	var createAppRequestDto CreateAppRequestDto
	createAppRequestDto.Id = id
	createAppRequestDto.AppName = appName
	createAppRequestDto.TeamId = teamId
	createAppRequestDto.TemplateId = templateId
	byteValueOfStruct, _ := json.Marshal(createAppRequestDto)
	return byteValueOfStruct
}

func HitDeleteAppApi(byteValueOfStruct []byte, id int, authToken string) DeleteResponseDto {
	resp, err := Base.MakeApiCall(CreateAppApiUrl+"/"+strconv.Itoa(id), http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, DeleteAppApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	apiRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteAppApi)
	return apiRouter.deleteResponseDto
}

func HitCreateAppApi(payload []byte, appName string, teamId int, templateId int, authToken string) CreateAppResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		var createAppRequestDto CreateAppRequestDto
		createAppRequestDto.AppName = appName
		createAppRequestDto.TeamId = teamId
		createAppRequestDto.TemplateId = templateId
		byteValueOfStruct, _ := json.Marshal(createAppRequestDto)
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := Base.MakeApiCall(CreateAppApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, CreateAppApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), CreateAppApi)
	return pipelineConfigRouter.createAppResponseDto
}

func GetAppMaterialRequestDto(appId int, url string, gitProviderId int, fetchSubmodules bool) CreateAppMaterialRequestDto {
	var createAppMaterialRequestDto CreateAppMaterialRequestDto
	var slice AppMaterials
	slice.Url = url
	slice.GitProviderId = gitProviderId
	slice.FetchSubmodules = fetchSubmodules
	createAppMaterialRequestDto.AppId = appId
	createAppMaterialRequestDto.Materials = append(createAppMaterialRequestDto.Materials, slice)
	return createAppMaterialRequestDto
}
func HitCreateAppMaterialApi(payload []byte, appId int, url string, gitProviderId int, fetchSubmodules bool, authToken string) CreateAppMaterialResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		var createAppMaterialRequestDto CreateAppMaterialRequestDto
		var slice AppMaterials
		slice.Url = url
		slice.GitProviderId = gitProviderId
		slice.FetchSubmodules = fetchSubmodules
		createAppMaterialRequestDto.AppId = appId
		createAppMaterialRequestDto.Materials = append(createAppMaterialRequestDto.Materials, slice)
		byteValueOfStruct, _ := json.Marshal(createAppMaterialRequestDto)
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := Base.MakeApiCall(CreateAppMaterialApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, CreateAppMaterialApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	installationScriptRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), CreateAppMaterialApi)
	return installationScriptRouter.createAppMaterialResponseDto
}

func getRequestPayloadForSaveAppCiPipeline(AppId int, dockerRegistry string, dockerRepository string, dockerfilePath string, dockerfileRepository string, dockerfileRelativePath string, gitMaterialId int) SaveAppCiPipelineRequestDTO {
	saveAppCiPipelineRequestDTO := SaveAppCiPipelineRequestDTO{}
	saveAppCiPipelineRequestDTO.AppId = AppId
	saveAppCiPipelineRequestDTO.DockerRepository = dockerRepository
	saveAppCiPipelineRequestDTO.DockerRegistry = dockerRegistry
	saveAppCiPipelineRequestDTO.DockerBuildConfig.DockerfilePath = dockerfilePath
	saveAppCiPipelineRequestDTO.DockerBuildConfig.DockerfileRepository = dockerfileRepository
	saveAppCiPipelineRequestDTO.DockerBuildConfig.DockerfileRelativePath = dockerfileRelativePath
	saveAppCiPipelineRequestDTO.DockerBuildConfig.GitMaterialId = gitMaterialId
	return saveAppCiPipelineRequestDTO
}

func HitSaveAppCiPipeline(payload []byte, authToken string) SaveAppCiPipelineResponseDTO {
	resp, err := Base.MakeApiCall(SaveAppCiPipelineApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveAppCiPipelineApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), SaveAppCiPipelineApi)
	return pipelineConfigRouter.saveAppCiPipelineResponseDTO
}

func HitGetCiPipelineViaId(appId string, authToken string) GetCiPipelineViaIdResponseDTO {
	resp, err := Base.MakeApiCall(GetCiPipelineViaIdApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetCiPipelineViaIdApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetCiPipelineViaIdApi)
	return pipelineConfigRouter.getCiPipelineViaIdResponseDTO
}

func HitGetContainerRegistry(appId string, authToken string) GetContainerRegistryResponseDTO {
	resp, err := Base.MakeApiCall(GetContainerRegistryApiUrl+appId+"/autocomplete/docker", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetContainerRegistryApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetContainerRegistryApi)
	return pipelineConfigRouter.getContainerRegistryResponseDTO
}

func (structPipelineConfigRouter StructPipelineConfigRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructPipelineConfigRouter {
	switch apiName {
	case CreateAppApi:
		json.Unmarshal(response, &structPipelineConfigRouter.createAppResponseDto)
	case CreateAppMaterialApi:
		json.Unmarshal(response, &structPipelineConfigRouter.createAppMaterialResponseDto)
	case SaveAppCiPipelineApi:
		json.Unmarshal(response, &structPipelineConfigRouter.saveAppCiPipelineResponseDTO)
	case GetCiPipelineViaIdApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getCiPipelineViaIdResponseDTO)
	case GetContainerRegistryApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getContainerRegistryResponseDTO)
	}
	return structPipelineConfigRouter
}

// PipelineConfigSuite =================PipelineConfigSuite Setup =========================
type PipelineConfigSuite struct {
	suite.Suite
	authToken                    string
	createAppResponseDto         CreateAppResponseDto
	createAppMaterialResponseDto CreateAppMaterialResponseDto
}

// SetupSuite This method runs on first priority before starting the suite means before executing any test case of the suite
func (suite *PipelineConfigSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
	suite.createAppResponseDto = suite.CreateApp()
	suite.createAppMaterialResponseDto = suite.CreateAppMaterial()
}

func (suite *PipelineConfigSuite) CreateApp() CreateAppResponseDto {
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	createAppRequestDto := GetAppRequestDto("app"+appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, "app"+appName, 1, 0, suite.authToken)
	return createAppResponseDto
}

func (suite *PipelineConfigSuite) CreateAppMaterial() CreateAppMaterialResponseDto {
	configPipelineConfigRouter, _ := GetEnvironmentConfigPipelineConfigRouter()
	createAppMaterialRequestDto := GetAppMaterialRequestDto(suite.createAppResponseDto.Result.Id, configPipelineConfigRouter.GitHubProjectUrl, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponseDto := HitCreateAppMaterialApi(appMaterialByteValue, suite.createAppResponseDto.Result.Id, configPipelineConfigRouter.GitHubProjectUrl, 1, false, suite.authToken)
	return createAppMaterialResponseDto
}

func (suite *PipelineConfigSuite) TearDownSuite() {
	log.Println("=== Running the after suite method for deleting the data created via automation ===")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(suite.createAppResponseDto.Result.Id, suite.createAppResponseDto.Result.AppName, suite.createAppResponseDto.Result.TeamId, suite.createAppResponseDto.Result.TemplateId)
	HitDeleteAppApi(byteValueOfDeleteApp, suite.createAppResponseDto.Result.Id, suite.authToken)
}
