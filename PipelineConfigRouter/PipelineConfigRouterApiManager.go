package PipelineConfigRouter

import (
	"automation-suite/dockerRegRouter"
	"automation-suite/testUtils"
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
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result AppDetails    `json:"result"`
	Errors []Base.Errors `json:"errors"`
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

type GetChartReferenceResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		ChartRefs []struct {
			Id      int    `json:"id"`
			Version string `json:"version"`
			Name    string `json:"name"`
		} `json:"chartRefs"`
		LatestChartRef    int `json:"latestChartRef"`
		LatestAppChartRef int `json:"latestAppChartRef"`
	} `json:"result"`
}

type LivenessProbe struct {
	Path                string        `json:"Path"`
	Command             []interface{} `json:"command"`
	FailureThreshold    int           `json:"failureThreshold"`
	HttpHeaders         []interface{} `json:"httpHeaders"`
	InitialDelaySeconds int           `json:"initialDelaySeconds"`
	PeriodSeconds       int           `json:"periodSeconds"`
	Port                int           `json:"port"`
	Scheme              string        `json:"scheme"`
	SuccessThreshold    int           `json:"successThreshold"`
	Tcp                 bool          `json:"tcp"`
	TimeoutSeconds      int           `json:"timeoutSeconds"`
}

// BlueGreen === GetCdPipelineStrategies ==== /////////
type BlueGreen struct {
	AutoPromotionSeconds  int  `json:"autoPromotionSeconds"`
	ScaleDownDelaySeconds int  `json:"scaleDownDelaySeconds"`
	PreviewReplicaCount   int  `json:"previewReplicaCount"`
	AutoPromotionEnabled  bool `json:"autoPromotionEnabled"`
}

type Canary struct {
	MaxSurge       string `json:"maxSurge"`
	MaxUnavailable int    `json:"maxUnavailable"`
	Steps          []struct {
		SetWeight int `json:"setWeight,omitempty"`
		Pause     struct {
			Duration int `json:"duration"`
		} `json:"pause,omitempty"`
	} `json:"steps"`
}

type Rolling struct {
	MaxSurge       string `json:"maxSurge"`
	MaxUnavailable int    `json:"maxUnavailable"`
}
type GetCdPipelineStrategiesResponseDto struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Errors []Base.Errors `json:"errors"`
	Result struct {
		PipelineStrategy []struct {
			DeploymentTemplate string `json:"deploymentTemplate"`
			Config             struct {
				Deployment struct {
					Strategy struct {
						Rolling   Rolling   `json:"rolling,omitempty"`
						BlueGreen BlueGreen `json:"blueGreen,omitempty"`
						Canary    Canary    `json:"canary,omitempty"`
						Recreate  struct {
						} `json:"recreate,omitempty"`
					} `json:"strategy"`
				} `json:"deployment"`
			} `json:"config"`
			Default bool `json:"default"`
		} `json:"pipelineStrategy"`
	} `json:"result"`
}

type PipelineSuggestedCDResponseDTO struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result string        `json:"result"`
	Errors []Base.Errors `json:"errors"`
}

type EnvironmentDetailsResponseDTO struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Errors []Base.Errors `json:"errors"`
	Result []struct {
		Id                    int    `json:"id"`
		EnvironmentName       string `json:"environment_name"`
		Active                bool   `json:"active"`
		Default               bool   `json:"default"`
		Namespace             string `json:"namespace"`
		IsClusterCdActive     bool   `json:"isClusterCdActive"`
		EnvironmentIdentifier string `json:"environmentIdentifier"`
	} `json:"result"`
}

type StructPipelineConfigRouter struct {
	saveAppCiPipelineRequestDTO        SaveAppCiPipelineRequestDTO
	createAppResponseDto               CreateAppResponseDto
	deleteResponseDto                  DeleteResponseDto
	createAppMaterialRequestDto        CreateAppMaterialRequestDto
	createAppMaterialResponseDto       CreateAppMaterialResponseDto
	getAppDetailsResponseDto           GetAppDetailsResponseDto
	saveAppCiPipelineResponseDTO       SaveAppCiPipelineResponseDTO
	getCiPipelineViaIdResponseDTO      GetCiPipelineViaIdResponseDTO
	getContainerRegistryResponseDTO    GetContainerRegistryResponseDTO
	getChartReferenceResponseDTO       GetChartReferenceResponseDTO
	getAppTemplateResponseDto          GetAppTemplateResponseDto
	getCdPipelineStrategiesResponseDto GetCdPipelineStrategiesResponseDto
	pipelineSuggestedCDResponseDTO     PipelineSuggestedCDResponseDTO
	environmentDetailsResponseDTO      EnvironmentDetailsResponseDTO
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
func GetAppMaterialRequestDto(appId int, gitProviderId int, fetchSubmodules bool) CreateAppMaterialRequestDto {
	pipelineConfig, _ := GetEnvironmentConfigPipelineConfigRouter()
	var slice AppMaterials
	slice.Url = pipelineConfig.GitHubProjectUrl
	slice.GitProviderId = gitProviderId
	slice.FetchSubmodules = fetchSubmodules
	var createAppMaterialRequestDto CreateAppMaterialRequestDto
	createAppMaterialRequestDto.AppId = appId
	createAppMaterialRequestDto.Materials = append(createAppMaterialRequestDto.Materials, slice)
	return createAppMaterialRequestDto
}
func HitCreateAppMaterialApi(payload []byte, appId int, gitProviderId int, fetchSubmodules bool, authToken string) CreateAppMaterialResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		byteValueOfStruct, _ := json.Marshal(GetAppMaterialRequestDto(appId, gitProviderId, fetchSubmodules))
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := Base.MakeApiCall(CreateAppMaterialApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, CreateAppMaterialApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), CreateAppMaterialApi)
	return pipelineConfigRouter.createAppMaterialResponseDto
}

type GetAppDetailsResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id         int         `json:"id"`
		AppName    string      `json:"appName"`
		TeamId     int         `json:"teamId"`
		TemplateId int         `json:"templateId"`
		Material   []Materials `json:"material""`
	} `json:"result"`
	Errors []Base.Errors `json:"errors"`
}
type Materials struct {
	Name            string `json:"name"`
	Url             string `json:"url"`
	Id              int    `json:"id"`
	GitProviderId   int    `json:"gitProviderId"`
	CheckoutPath    string `json:"checkoutPath"`
	FetchSubmodules bool   `json:"fetchSubmodules"`
}

func HitGetMaterial(appId int, authToken string) GetAppDetailsResponseDto {
	id := strconv.Itoa(appId)
	resp, err := Base.MakeApiCall(GetAppDetailsApiUrl+"/"+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAppDetailsApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppDetailsApi)
	return pipelineConfigRouter.getAppDetailsResponseDto
}

type DeleteAppMaterialRequestDto struct {
	AppId     int          `json:"appId"`
	Materials AppMaterials `json:"material"`
}

func GetPayLoadForDeleteAppMaterialAPI(appId int, slice2 AppMaterials) []byte {
	var deleteAppMaterialRequestDto DeleteAppMaterialRequestDto
	deleteAppMaterialRequestDto.AppId = appId
	deleteAppMaterialRequestDto.Materials.Id = slice2.Id
	deleteAppMaterialRequestDto.Materials.Url = slice2.Url
	deleteAppMaterialRequestDto.Materials.GitProviderId = slice2.GitProviderId
	deleteAppMaterialRequestDto.Materials.CheckoutPath = slice2.CheckoutPath
	deleteAppMaterialRequestDto.Materials.FetchSubmodules = slice2.FetchSubmodules
	byteValueOfStruct, _ := json.Marshal(deleteAppMaterialRequestDto)
	return byteValueOfStruct
}
func HitDeleteAppMaterialApi(byteValueOfStruct []byte, authToken string) DeleteResponseDto {
	resp, err := Base.MakeApiCall(CreateAppMaterialApiUrl+"/delete", http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, DeleteAppMaterialApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteAppMaterialApi)
	return pipelineConfigRouter.deleteResponseDto
}
func GetRequestPayloadForSaveAppCiPipeline(AppId int, dockerRegistry string, dockerRepository string, dockerfilePath string, dockerfileRepository string, dockerfileRelativePath string, gitMaterialId int) SaveAppCiPipelineRequestDTO {
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

func HitGetChartReferenceViaAppId(appId string, authToken string) GetChartReferenceResponseDTO {
	resp, err := Base.MakeApiCall(GetChartReferenceViaAppIdApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetChartReferenceViaAppIdApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetChartReferenceViaAppIdApi)
	return pipelineConfigRouter.getChartReferenceResponseDTO
}

func HitGetTemplateViaAppIdAndChartRefId(appId string, chartRefId string, authToken string) GetAppTemplateResponseDto {
	resp, err := Base.MakeApiCall(GetAppTemplateViaAppIdAndChartRefIdApiUrl+appId+"/"+chartRefId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetChartReferenceViaAppIdApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppTemplateViaAppIdAndChartRefIdApi)
	return pipelineConfigRouter.getAppTemplateResponseDto
}

func HitGetCdPipelineStrategies(appId string, authToken string) GetCdPipelineStrategiesResponseDto {
	resp, err := Base.MakeApiCall(GetCdPipelineStrategiesApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetCdPipelineStrategiesApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetCdPipelineStrategiesApi)
	return pipelineConfigRouter.getCdPipelineStrategiesResponseDto
}

func HitGetPipelineSuggestedCD(appId string, authToken string) PipelineSuggestedCDResponseDTO {
	resp, err := Base.MakeApiCall(GetPipelineSuggestedCDApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetPipelineSuggestedCDApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetPipelineSuggestedCDApi)
	return pipelineConfigRouter.pipelineSuggestedCDResponseDTO
}

func HitGetAllEnvironmentDetails(queryParams map[string]string, authToken string) EnvironmentDetailsResponseDTO {
	resp, err := Base.MakeApiCall(GetAllEnvironmentDetailsApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetAllEnvironmentDetailsApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetAllEnvironmentDetailsApi)
	return pipelineConfigRouter.environmentDetailsResponseDTO
}

func (structPipelineConfigRouter StructPipelineConfigRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructPipelineConfigRouter {
	switch apiName {
	case DeleteAppMaterialApi:
		json.Unmarshal(response, &structPipelineConfigRouter.deleteResponseDto)
	case GetAppDetailsApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getAppDetailsResponseDto)
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
	case GetChartReferenceViaAppIdApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getChartReferenceResponseDTO)
	case GetAppTemplateViaAppIdAndChartRefIdApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getAppTemplateResponseDto)
	case GetCdPipelineStrategiesApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getCdPipelineStrategiesResponseDto)

	case GetPipelineSuggestedCDApi:
		json.Unmarshal(response, &structPipelineConfigRouter.pipelineSuggestedCDResponseDTO)
	case GetAllEnvironmentDetailsApi:
		json.Unmarshal(response, &structPipelineConfigRouter.environmentDetailsResponseDTO)
	}
	return structPipelineConfigRouter
}

// PipelinesConfigRouterTestSuite PipelineConfigSuite =================PipelineConfigSuite Setup =========================
type PipelinesConfigRouterTestSuite struct {
	suite.Suite
	authToken                    string
	createAppResponseDto         CreateAppResponseDto
	createAppMaterialResponseDto CreateAppMaterialResponseDto
}

// SetupSuite This method runs on first priority before starting the suite means before executing any test case of the suite
func (suite *PipelinesConfigRouterTestSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
	suite.createAppResponseDto = suite.CreateApp()
	suite.createAppMaterialResponseDto = suite.CreateAppMaterial()
}

func (suite *PipelinesConfigRouterTestSuite) CreateApp() CreateAppResponseDto {
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	createAppRequestDto := GetAppRequestDto("app"+appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, "app"+appName, 1, 0, suite.authToken)
	return createAppResponseDto
}

func (suite *PipelinesConfigRouterTestSuite) CreateAppMaterial() CreateAppMaterialResponseDto {
	createAppMaterialRequestDto := GetAppMaterialRequestDto(suite.createAppResponseDto.Result.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponseDto := HitCreateAppMaterialApi(appMaterialByteValue, suite.createAppResponseDto.Result.Id, 1, false, suite.authToken)
	return createAppMaterialResponseDto
}

func (suite *PipelinesConfigRouterTestSuite) TearDownSuite() {
	log.Println("=== Running the after suite method for deleting the data created via automation ===")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(suite.createAppResponseDto.Result.Id, suite.createAppResponseDto.Result.AppName, suite.createAppResponseDto.Result.TeamId, suite.createAppResponseDto.Result.TemplateId)
	HitDeleteAppApi(byteValueOfDeleteApp, suite.createAppResponseDto.Result.Id, suite.authToken)
	testUtils.DeleteFile("OutputDataGetChartReferenceViaAppId")
}
