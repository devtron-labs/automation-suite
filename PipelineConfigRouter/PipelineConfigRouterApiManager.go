package PipelineConfigRouter

import (
	"automation-suite/PipelineConfigRouter/RequestDTOs"
	"automation-suite/PipelineConfigRouter/ResponseDTOs"
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
						Rolling   RequestDTOs.Rolling `json:"rolling,omitempty"`
						BlueGreen BlueGreen           `json:"blueGreen,omitempty"`
						Canary    Canary              `json:"canary,omitempty"`
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
	saveDeploymentTemplateResponseDTO  SaveDeploymentTemplateResponseDTO
	getWorkflowDetails                 GetWorkflowDetails
	createWorkflowResponseDto          CreateWorkflowResponseDto
	fetchSuggestedCiPipelineName       FetchSuggestedCiPipelineName
	saveCdPipelineRequestDTO           RequestDTOs.SaveCdPipelineRequestDTO
	saveCdPipelineResponseDTO          ResponseDTOs.SaveCdPipelineResponseDTO
	deleteCdPipelineRequestDTO         RequestDTOs.DeleteCdPipelineRequestDTO
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
	slice.CheckoutPath = "./" + Base.GetRandomStringOfGivenLength(5)
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

func GetRequestPayloadForSaveDeploymentTemplate(AppId int, chartRefId int, defaultOverride DefaultAppOverride) SaveDeploymentTemplateRequestDTO {
	saveDeploymentTemplateRequestDTO := SaveDeploymentTemplateRequestDTO{}
	saveDeploymentTemplateRequestDTO.AppId = AppId
	saveDeploymentTemplateRequestDTO.ChartRefId = chartRefId
	saveDeploymentTemplateRequestDTO.ValuesOverride = defaultOverride
	saveDeploymentTemplateRequestDTO.DefaultAppOverride = defaultOverride
	return saveDeploymentTemplateRequestDTO
}

func HitSaveDeploymentTemplateApi(payload []byte, authToken string) SaveDeploymentTemplateResponseDTO {
	resp, err := Base.MakeApiCall(SaveDeploymentTemplateAPiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveDeploymentTemplateApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), SaveDeploymentTemplateApi)
	return pipelineConfigRouter.saveDeploymentTemplateResponseDTO
}

func HitSaveCdPipelineApi(payload []byte, authToken string) ResponseDTOs.SaveCdPipelineResponseDTO {
	resp, err := Base.MakeApiCall(SaveCdPipelineApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveCdPipelineApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), SaveCdPipelineApi)
	return pipelineConfigRouter.saveCdPipelineResponseDTO
}

func HitDeleteCdPipelineApi(payload []byte, authToken string) RequestDTOs.DeleteCdPipelineRequestDTO {
	resp, err := Base.MakeApiCall(DeleteCdPipelineApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, DeleteCdPipelineApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteCdPipelineApi)
	return pipelineConfigRouter.deleteCdPipelineRequestDTO
}

func GetPayloadForDeleteCdPipeline(AppId int, pipelineId int) RequestDTOs.DeleteCdPipelineRequestDTO {
	deleteRequest := RequestDTOs.DeleteCdPipelineRequestDTO{}
	deleteRequest.Pipeline.Id = pipelineId
	deleteRequest.Action = 1
	deleteRequest.AppId = AppId
	return deleteRequest
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
	case SaveDeploymentTemplateApi:
		json.Unmarshal(response, &structPipelineConfigRouter.saveDeploymentTemplateResponseDTO)
	case CreateWorkflowApi:
		json.Unmarshal(response, &structPipelineConfigRouter.createWorkflowResponseDto)
	case FetchSuggestedCiPipelineNameApi:
		json.Unmarshal(response, &structPipelineConfigRouter.fetchSuggestedCiPipelineName)
	case GetWorkflowDetailsApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getWorkflowDetails)
	case DeleteAppApi:
		json.Unmarshal(response, &structPipelineConfigRouter.deleteResponseDto)
	case SaveCdPipelineApi:
		json.Unmarshal(response, &structPipelineConfigRouter.saveCdPipelineResponseDTO)
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
	//byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(suite.createAppResponseDto.Result.Id, suite.createAppResponseDto.Result.AppName, suite.createAppResponseDto.Result.TeamId, suite.createAppResponseDto.Result.TemplateId)
	//HitDeleteAppApi(byteValueOfDeleteApp, suite.createAppResponseDto.Result.Id, suite.authToken)
	//testUtils.DeleteFile("OutputDataGetChartReferenceViaAppId")
}

/////////////////=== Create Workflow API ====//////////////
type CiMaterial struct {
	Source struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"source"`
	GitMaterialId   int    `json:"gitMaterialId"`
	Id              int    `json:"id"`
	GitMaterialName string `json:"gitMaterialName"`
}

type PreBuildStage struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Steps []struct {
		Id                  int      `json:"id"`
		Name                string   `json:"name"`
		Description         string   `json:"description"`
		Index               int      `json:"index"`
		StepType            string   `json:"stepType"`
		OutputDirectoryPath []string `json:"outputDirectoryPath"`
		InlineStepDetail    struct {
			ScriptType             string `json:"scriptType"`
			Script                 string `json:"script"`
			StoreScriptAt          string `json:"storeScriptAt"`
			MountDirectoryFromHost bool   `json:"mountDirectoryFromHost"`
			CommandArgsMap         []struct {
				Command string   `json:"command"`
				Args    []string `json:"args"`
			} `json:"commandArgsMap"`
			InputVariables []struct {
				Id                   int    `json:"id"`
				Name                 string `json:"name"`
				Format               string `json:"format"`
				Description          string `json:"description"`
				Value                string `json:"value"`
				VariableType         string `json:"variableType"`
				RefVariableName      string `json:"refVariableName,omitempty"`
				RefVariableStage     string `json:"refVariableStage"`
				RefVariableStepIndex int    `json:"refVariableStepIndex,omitempty"`
			} `json:"inputVariables"`
			OutputVariables []struct {
				Id               int    `json:"id"`
				Name             string `json:"name"`
				Format           string `json:"format"`
				Description      string `json:"description"`
				Value            string `json:"value"`
				VariableType     string `json:"variableType"`
				RefVariableStage string `json:"refVariableStage"`
			} `json:"outputVariables"`
			ConditionDetails []struct {
				Id                  int    `json:"id"`
				ConditionOnVariable string `json:"conditionOnVariable"`
				ConditionType       string `json:"conditionType"`
				ConditionOperator   string `json:"conditionOperator"`
				ConditionalValue    string `json:"conditionalValue"`
			} `json:"conditionDetails"`
			MountCodeToContainer     bool   `json:"mountCodeToContainer,omitempty"`
			MountCodeToContainerPath string `json:"mountCodeToContainerPath,omitempty"`
			ContainerImagePath       string `json:"containerImagePath,omitempty"`
			MountPathMap             []struct {
				FilePathOnDisk      string `json:"filePathOnDisk"`
				FilePathOnContainer string `json:"filePathOnContainer"`
			} `json:"mountPathMap,omitempty"`
			PortMap []struct {
				PortOnLocal     int `json:"portOnLocal"`
				PortOnContainer int `json:"portOnContainer"`
			} `json:"portMap,omitempty"`
			IsMountCustomScript bool `json:"isMountCustomScript,omitempty"`
		} `json:"inlineStepDetail"`
		PluginRefStepDetail interface{} `json:"pluginRefStepDetail"`
	} `json:"steps"`
}
type Step struct {
	Id                  int      `json:"id"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Index               int      `json:"index"`
	StepType            string   `json:"stepType"`
	OutputDirectoryPath []string `json:"outputDirectoryPath"`
	InlineStepDetail    struct {
		ScriptType       string           `json:"scriptType"`
		Script           string           `json:"script"`
		StoreScriptAt    string           `json:"storeScriptAt"`
		CommandArgsMap   []CommandArgsMap `json:"commandArgsMap"`
		InputVariables   []InputVariables `json:"inputVariables"`
		OutputVariables  []InputVariables `json:"outputVariables"`
		ConditionDetails []struct {
			Id                  int    `json:"id"`
			ConditionOnVariable string `json:"conditionOnVariable"`
			ConditionType       string `json:"conditionType"`
			ConditionOperator   string `json:"conditionOperator"`
			ConditionalValue    string `json:"conditionalValue"`
		} `json:"conditionDetails"`
		MountCodeToContainer     bool   `json:"mountCodeToContainer,omitempty"`
		MountCodeToContainerPath string `json:"mountCodeToContainerPath,omitempty"`
		MountDirectoryFromHost   bool   `json:"mountDirectoryFromHost"`
		ContainerImagePath       string `json:"containerImagePath,omitempty"`
		MountPathMap             []struct {
			FilePathOnDisk      string `json:"filePathOnDisk"`
			FilePathOnContainer string `json:"filePathOnContainer"`
		} `json:"mountPathMap,omitempty"`
		PortMap []struct {
			PortOnLocal     int `json:"portOnLocal"`
			PortOnContainer int `json:"portOnContainer"`
		} `json:"portMap,omitempty"`
		IsMountCustomScript bool `json:"isMountCustomScript,omitempty"`
	} `json:"inlineStepDetail"`
	PluginRefStepDetail interface{} `json:"pluginRefStepDetail"`
}
type CiPipeline struct {
	IsManual         bool              `json:"isManual"`
	DockerArgs       map[string]string `json:"dockerArgs"`
	IsExternal       bool              `json:"isExternal"`
	ParentCiPipeline int               `json:"parentCiPipeline"`
	ParentAppId      int               `json:"parentAppId"`
	ExternalCiConfig struct {
		Id         int    `json:"id"`
		WebhookUrl string `json:"webhookUrl"`
		Payload    string `json:"payload"`
		AccessKey  string `json:"accessKey"`
	} `json:"externalCiConfig"`
	CiMaterial    []CiMaterial `json:"ciMaterial"`
	Name          string       `json:"name"`
	Id            int          `json:"id"`
	Active        bool         `json:"active"`
	LinkedCount   int          `json:"linkedCount"`
	ScanEnabled   bool         `json:"scanEnabled"`
	AppWorkflowId int          `json:"appWorkflowId"`
	PreBuildStage struct {
		Id    int    `json:"id"`
		Type  string `json:"type"`
		Steps []Step `json:"steps"`
	} `json:"preBuildStage"`
	PostBuildStage struct {
		Id    int    `json:"id"`
		Type  string `json:"type"`
		Steps []Step `json:"steps"`
	} `json:"postBuildStage"`
}
type GetWorkflowDetails struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Result CiPipeline `json:"result"`
}
type ConditionDetails struct {
	Id                  int    `json:"id"`
	ConditionOnVariable string `json:"conditionOnVariable"`
	ConditionType       string `json:"conditionType"`
	ConditionOperator   string `json:"conditionOperator"`
	ConditionalValue    string `json:"conditionalValue"`
}

// Payload with added missing fields
type CreateWorkflowRequestDto struct {
	AppId         int        `json:"appId"`
	AppWorkflowId int        `json:"appWorkflowId"`
	Action        int        `json:"action"`
	CiPipeline    CiPipeline `json:"ciPipeline"`
}
type CreateWorkflowResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id                int    `json:"id"`
		AppId             int    `json:"appId"`
		DockerRegistry    string `json:"dockerRegistry"`
		DockerRepository  string `json:"dockerRepository"`
		DockerBuildConfig struct {
			GitMaterialId          int    `json:"gitMaterialId"`
			DockerfileRelativePath string `json:"dockerfileRelativePath"`
		} `json:"dockerBuildConfig"`
		CiPipelines []struct {
			IsManual         bool              `json:"isManual"`
			DockerArgs       map[string]string `json:"dockerArgs"`
			IsExternal       bool              `json:"isExternal"`
			ParentCiPipeline int               `json:"parentCiPipeline"`
			ParentAppId      int               `json:"parentAppId"`
			ExternalCiConfig struct {
				Id         int    `json:"id"`
				WebhookUrl string `json:"webhookUrl"`
				Payload    string `json:"payload"`
				AccessKey  string `json:"accessKey"`
			} `json:"externalCiConfig"`
			CiMaterial []struct {
				Source struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"source"`
				GitMaterialId   int    `json:"gitMaterialId"`
				Id              int    `json:"id"`
				GitMaterialName string `json:"gitMaterialName"`
			} `json:"ciMaterial"`

			Name          string `json:"name"`
			Id            int    `json:"id"`
			Active        bool   `json:"active"`
			LinkedCount   int    `json:"linkedCount"`
			ScanEnabled   bool   `json:"scanEnabled"`
			PreBuildStage struct {
				Id    int    `json:"id"`
				Type  string `json:"type"`
				Steps []struct {
					Id                  int      `json:"id"`
					Name                string   `json:"name"`
					Description         string   `json:"description"`
					Index               int      `json:"index"`
					StepType            string   `json:"stepType"`
					OutputDirectoryPath []string `json:"outputDirectoryPath"`
					InlineStepDetail    struct {
						ScriptType             string `json:"scriptType"`
						Script                 string `json:"script"`
						StoreScriptAt          string `json:"storeScriptAt"`
						MountDirectoryFromHost bool   `json:"mountDirectoryFromHost"`
						CommandArgsMap         []struct {
							Command string   `json:"command"`
							Args    []string `json:"args"`
						} `json:"commandArgsMap"`
						InputVariables []struct {
							Id                   int    `json:"id"`
							Name                 string `json:"name"`
							Format               string `json:"format"`
							Description          string `json:"description"`
							Value                string `json:"value"`
							VariableType         string `json:"variableType"`
							RefVariableName      string `json:"refVariableName,omitempty"`
							RefVariableStage     string `json:"refVariableStage"`
							RefVariableStepIndex int    `json:"refVariableStepIndex,omitempty"`
						} `json:"inputVariables"`
						OutputVariables []struct {
							Id               int    `json:"id"`
							Name             string `json:"name"`
							Format           string `json:"format"`
							Description      string `json:"description"`
							Value            string `json:"value"`
							VariableType     string `json:"variableType"`
							RefVariableStage string `json:"refVariableStage"`
						} `json:"outputVariables"`
						ConditionDetails []struct {
							Id                  int    `json:"id"`
							ConditionOnVariable string `json:"conditionOnVariable"`
							ConditionType       string `json:"conditionType"`
							ConditionOperator   string `json:"conditionOperator"`
							ConditionalValue    string `json:"conditionalValue"`
						} `json:"conditionDetails"`
						MountCodeToContainer     bool   `json:"mountCodeToContainer,omitempty"`
						MountCodeToContainerPath string `json:"mountCodeToContainerPath,omitempty"`
						ContainerImagePath       string `json:"containerImagePath,omitempty"`
						MountPathMap             []struct {
							FilePathOnDisk      string `json:"filePathOnDisk"`
							FilePathOnContainer string `json:"filePathOnContainer"`
						} `json:"mountPathMap,omitempty"`
						PortMap []struct {
							PortOnLocal     int `json:"portOnLocal"`
							PortOnContainer int `json:"portOnContainer"`
						} `json:"portMap,omitempty"`
						IsMountCustomScript bool `json:"isMountCustomScript,omitempty"`
					} `json:"inlineStepDetail"`
					PluginRefStepDetail interface{} `json:"pluginRefStepDetail"`
				} `json:"steps"`
			} `json:"preBuildStage"`
			PostBuildStage struct {
				Id    int    `json:"id"`
				Type  string `json:"type"`
				Steps []struct {
					Id                  int      `json:"id"`
					Name                string   `json:"name"`
					Description         string   `json:"description"`
					Index               int      `json:"index"`
					StepType            string   `json:"stepType"`
					OutputDirectoryPath []string `json:"outputDirectoryPath"`
					InlineStepDetail    struct {
						ScriptType             string `json:"scriptType"`
						Script                 string `json:"script"`
						StoreScriptAt          string `json:"storeScriptAt"`
						MountDirectoryFromHost bool   `json:"mountDirectoryFromHost"`
						CommandArgsMap         []struct {
							Command string   `json:"command"`
							Args    []string `json:"args"`
						} `json:"commandArgsMap"`
						InputVariables []struct {
							Id                   int    `json:"id"`
							Name                 string `json:"name"`
							Format               string `json:"format"`
							Description          string `json:"description"`
							Value                string `json:"value"`
							VariableType         string `json:"variableType"`
							RefVariableName      string `json:"refVariableName,omitempty"`
							RefVariableStage     string `json:"refVariableStage"`
							RefVariableStepIndex int    `json:"refVariableStepIndex,omitempty"`
						} `json:"inputVariables"`
						OutputVariables []struct {
							Id               int    `json:"id"`
							Name             string `json:"name"`
							Format           string `json:"format"`
							Description      string `json:"description"`
							Value            string `json:"value"`
							VariableType     string `json:"variableType"`
							RefVariableStage string `json:"refVariableStage"`
						} `json:"outputVariables"`
						ConditionDetails []struct {
							Id                  int    `json:"id"`
							ConditionOnVariable string `json:"conditionOnVariable"`
							ConditionType       string `json:"conditionType"`
							ConditionOperator   string `json:"conditionOperator"`
							ConditionalValue    string `json:"conditionalValue"`
						} `json:"conditionDetails"`
						MountCodeToContainer     bool   `json:"mountCodeToContainer,omitempty"`
						MountCodeToContainerPath string `json:"mountCodeToContainerPath,omitempty"`
						ContainerImagePath       string `json:"containerImagePath,omitempty"`
						MountPathMap             []struct {
							FilePathOnDisk      string `json:"filePathOnDisk"`
							FilePathOnContainer string `json:"filePathOnContainer"`
						} `json:"mountPathMap,omitempty"`
						PortMap []struct {
							PortOnLocal     int `json:"portOnLocal"`
							PortOnContainer int `json:"portOnContainer"`
						} `json:"portMap,omitempty"`
						IsMountCustomScript bool `json:"isMountCustomScript,omitempty"`
					} `json:"inlineStepDetail"`
					PluginRefStepDetail interface{} `json:"pluginRefStepDetail"`
				} `json:"steps"`
			} `json:"postBuildStage"`
		} `json:"ciPipelines"`
		AppName   string `json:"appName"`
		Materials []struct {
			GitMaterialId int    `json:"gitMaterialId"`
			MaterialName  string `json:"materialName"`
		} `json:"materials"`
		AppWorkflowId int  `json:"appWorkflowId"`
		ScanEnabled   bool `json:"scanEnabled"`
	} `json:"result"`
}

func HitCreateWorkflowApi(payload []byte, authToken string) CreateWorkflowResponseDto {
	resp, err := Base.MakeApiCall(CreateWorkflowApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, CreateWorkflowApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), CreateWorkflowApi)
	return pipelineConfigRouter.createWorkflowResponseDto
}
func worflowTypeProvider(temp string) string {
	var str string
	switch temp {
	case "1":
		str = "SOURCE_TYPE_BRANCH_FIXED"
	case "2":
		str = "WEBHOOK"
	case "3":
		str = "WEBHOOK"
	}
	return str
}
func getRequestPayloadForCreateWorkflow(forDelete bool, wfTypeId string, appId int, wfId int) CreateWorkflowRequestDto {
	var createWorkflowRequestDto CreateWorkflowRequestDto

	if forDelete == true {
		createWorkflowRequestDto.AppWorkflowId = wfId
		createWorkflowRequestDto.AppId = appId
		createWorkflowRequestDto.Action = 2
		return createWorkflowRequestDto
	}
	wfTypeStr := worflowTypeProvider(wfTypeId)
	var CiMaterial CiMaterial
	CiMaterial.Source.Type = wfTypeStr
	createWorkflowRequestDto.CiPipeline.Active = true
	createWorkflowRequestDto.AppId = appId
	createWorkflowRequestDto.CiPipeline.CiMaterial = append(createWorkflowRequestDto.CiPipeline.CiMaterial, CiMaterial)
	return createWorkflowRequestDto
}
func HitGetWorkflowGetailsApi(appId int, wfId int, authToken string) GetWorkflowDetails {
	resp, err := Base.MakeApiCall(GetCiPipelineViaIdApiUrl+strconv.Itoa(appId)+"/"+strconv.Itoa(wfId), http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetWorkflowDetailsApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetWorkflowDetailsApi)
	return pipelineConfigRouter.getWorkflowDetails
}
func HitDeleteWorkflowApi(appId int, wfId int, authToken string) DeleteResponseDto {
	resp, err := Base.MakeApiCall(DeleteWorkflowApiUrl+strconv.Itoa(appId)+"/"+strconv.Itoa(wfId), http.MethodDelete, "", nil, authToken)
	Base.HandleError(err, DeleteAppApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteAppApi)
	return pipelineConfigRouter.deleteResponseDto
}

type FetchSuggestedCiPipelineName struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}

func HitFetchSuggestedCiPipelineName(appId int, authToken string) FetchSuggestedCiPipelineName {
	resp, err := Base.MakeApiCall(FetchSuggestedCiPipelineNameApiUrl+strconv.Itoa(appId), http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchSuggestedCiPipelineNameApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), FetchSuggestedCiPipelineNameApi)
	return pipelineConfigRouter.fetchSuggestedCiPipelineName
}

type InputVariables struct {
	Id                        int    `json:"id"`
	Name                      string `json:"name"`
	Format                    string `json:"format"`
	Description               string `json:"description"`
	IsExposed                 bool   `json:"isExposed"`
	AllowEmptyValue           bool   `json:"allowEmptyValue"`
	Value                     string `json:"value"`
	VariableType              string `json:"variableType"`
	VariableStepIndexInPlugin int    `json:"variableStepIndexInPlugin"`
	RefVariableStage          string `json:"refVariableStage"`
	RefVariableName           string `json:"refVariableName"`
}

func inputVariablesSelector(inputType int) []InputVariables {
	var inputVariable InputVariables
	switch inputType {
	case 1:
		inputVariable.Format = "STRING"
		inputVariable.Value = Base.GetRandomStringOfGivenLength(5)
		inputVariable.VariableType = "NEW"
		break
	case 2:
		inputVariable.Format = "BOOL"
		inputVariable.Value = "true"
		inputVariable.VariableType = "NEW"
		break
	case 3:
		inputVariable.Format = "NUMBER"
		inputVariable.Value = strconv.Itoa(Base.GetRandomNumberOf9Digit())
		inputVariable.VariableType = "NEW"
		break
	case 4:
		inputVariable.Format = "DATE"
		inputVariable.Value = "2006-01-02"
		inputVariable.VariableType = "NEW"
		break
	case 5:
		inputVariable.Format = "STRING"
		inputVariable.VariableType = "GLOBAL"
		inputVariable.RefVariableName = "DOCKER_IMAGE_TAG"
		break
	}
	inputVariable.Name = Base.GetRandomStringOfGivenLength(5) + "_" + inputVariable.Format
	inputVariable.Description = inputVariable.Name + "_Desc_" + Base.GetRandomStringOfGivenLength(10)
	var input []InputVariables
	input = append(input, inputVariable)
	return input
}
func getConditionDetails(id int) []ConditionDetails {
	var conditionDetails ConditionDetails
	conditionDetails.ConditionType = "TRIGGER"
	switch id {
	case 1:
		conditionDetails.ConditionOperator = "=="
		break
	case 2:
		conditionDetails.ConditionOperator = "!="
		break
	case 3:
		conditionDetails.ConditionOperator = ">"
		break
	case 4:
		conditionDetails.ConditionOperator = "<"
		break
	case 5:
		conditionDetails.ConditionOperator = "<="
		break
	case 6:
		conditionDetails.ConditionOperator = ">="
		break
	}
	var input []ConditionDetails
	input = append(input, conditionDetails)
	return input
}

type CommandArgsMap struct {
	Command string `json:"command"`
	Args    []Args `json:"args"`
}
type Args struct {
	Arg string `json:"args"`
}

func getPreBuildStepRequestPayloadDto(scriptType int) []Step {
	var step Step
	step.Name = Base.GetRandomStringOfGivenLength(10)
	step.Description = Base.GetRandomStringOfGivenLength(20)
	step.StepType = "INLINE"
	switch scriptType {
	case 1:
		{
			step.InlineStepDetail.ScriptType = "SHELL"

			outputVariables := inputVariablesSelector(1)
			step.InlineStepDetail.OutputVariables = append(step.InlineStepDetail.OutputVariables, outputVariables[0])

			step.OutputDirectoryPath = append(step.OutputDirectoryPath, "./"+Base.GetRandomStringOfGivenLength(5))
			break
		}
	case 2:
		{
			step.InlineStepDetail.ScriptType = "CONTAINER_IMAGE"
			step.InlineStepDetail.ContainerImagePath = "alpine:latest"
			/*
				var arg Args
				arg.Arg = "/" + Base.GetRandomStringOfGivenLength(5) + ".sh"
				var args []Args
				args = append(args, arg)
				var commandArgsMap CommandArgsMap
				commandArgsMap.Command = "sh"
				commandArgsMap.Args = args

				var commandArgsMap2 []CommandArgsMap
				commandArgsMap2 = append(commandArgsMap2, commandArgsMap)
				step.InlineStepDetail.CommandArgsMap = append(step.InlineStepDetail.CommandArgsMap, commandArgsMap2[0])

			*/
			break
		}

	}
	step.InlineStepDetail.Script = "#!/bin/sh \nset -eo pipefail \n#set -v  ## uncomment this to debug the script \n"
	i := 0
	log.Println("Adding tasks")
	for i = 1; i < 6; i++ {
		inputVariables := inputVariablesSelector(i)
		conditionDetails := getConditionDetails(i)
		step.InlineStepDetail.InputVariables = append(step.InlineStepDetail.InputVariables, inputVariables[0])
		step.InlineStepDetail.ConditionDetails = append(step.InlineStepDetail.ConditionDetails, conditionDetails[0])
		step.InlineStepDetail.ConditionDetails[0].ConditionOnVariable = step.InlineStepDetail.InputVariables[0].Name
		step.InlineStepDetail.ConditionDetails[0].ConditionalValue = Base.GetRandomStringOfGivenLength(4)
	}
	step.InlineStepDetail.MountCodeToContainer = false
	step.InlineStepDetail.MountDirectoryFromHost = false
	var steps []Step
	steps = append(steps, step)
	return steps
}

func HitCreateWorkflowApiWithFullPayload(appId int, authToken string) CreateWorkflowResponseDto {
	createWorkflowRequestDto := getRequestPayloadForCreateWorkflow(false, "1", appId, 0)
	key := Base.GetRandomStringOfGivenLength(10)
	createWorkflowRequestDto.CiPipeline.DockerArgs = make(map[string]string)
	createWorkflowRequestDto.CiPipeline.DockerArgs[key] = Base.GetRandomStringOfGivenLength(10)
	fetchSuggestedCiPipelineName := HitFetchSuggestedCiPipelineName(appId, authToken)
	createWorkflowRequestDto.CiPipeline.Name = fetchSuggestedCiPipelineName.Result
	fetchAppGetResponseDto := HitGetMaterial(appId, authToken)
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].GitMaterialId = fetchAppGetResponseDto.Result.Material[0].Id
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].Source.Value = strings.ToLower(Base.GetRandomStringOfGivenLength(10))

	i := 0
	for i = 1; i < 3; i++ {
		preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i)
		postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i)

		createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload[0])
		createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload[0])
	}

	byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
	log.Println("Hitting the Create Workflow Api with valid payload")
	createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, authToken)
	return createWorkflowResponseDto
}

func DeleteWorkflow(appId int, wfId int, authToken string) {
	getWorkflowDetailsResponseDto := HitGetWorkflowGetailsApi(appId, wfId, authToken)
	log.Println("Validating get workflow details api")
	log.Println("Getting data for delete ci-pipeline")

	deleteCiPipelineRequestDto := getRequestPayloadForCreateWorkflow(true, "1", appId, wfId)
	deleteCiPipelineRequestDto.CiPipeline = getWorkflowDetailsResponseDto.Result
	log.Println("Removing the data created via ci-pipeline API")
	byteValueOfDeleteCiPipeline, _ := json.Marshal(deleteCiPipelineRequestDto)

	log.Println("Hitting the Create Workflow Api with action=2 for delete ci-pipeline")
	//deleteCiPipelineResponseDto := HitCreateWorkflowApi(byteValueOfDeleteCiPipeline, authToken)
	HitCreateWorkflowApi(byteValueOfDeleteCiPipeline, authToken)

	log.Println("Validating delete ci-pipeline")
	//assert.Equal(suite.T(), deleteCiPipelineResponseDto.Result.AppId, appId)

	log.Println("Deleting workflow")
	//respOfDeleteWorkflowApi := HitDeleteWorkflowApi(appId, wfId, authToken)
	HitDeleteWorkflowApi(appId, wfId, authToken)
	//assert.Equal(suite.T(), 200, respOfDeleteWorkflowApi.Code)
	return
}

func getRequestPayloadForSaveCdPipelineApi(appId int, AppWorkflowId int, EnvironmentId int, CiPipelineId int, ParentPipelineId int, strategy string, prescript string, postscript string, pipelineTriggerType string) RequestDTOs.SaveCdPipelineRequestDTO {
	CdPipelineRequestDTO := RequestDTOs.SaveCdPipelineRequestDTO{}
	CdPipelineRequestDTO.AppId = appId
	CdPipelineRequestDTO.Pipelines = getPipeLines(AppWorkflowId, EnvironmentId, CiPipelineId, ParentPipelineId, strategy, prescript, postscript, pipelineTriggerType)
	return CdPipelineRequestDTO
}

func getPipeLines(AppWorkflowId int, EnvironmentId int, CiPipelineId int, ParentPipelineId int, strategy string, prescript string, postscript string, pipelineTriggerType string) []RequestDTOs.Pipeline {
	var pipelines []RequestDTOs.Pipeline
	pipeline := RequestDTOs.Pipeline{}
	pipeline.AppWorkflowId = AppWorkflowId
	pipeline.EnvironmentId = EnvironmentId
	pipeline.DeploymentTemplate = "ROLLING"
	pipeline.CiPipelineId = CiPipelineId
	pipeline.TriggerType = pipelineTriggerType
	pipeline.Name = "cd-pipeline"
	pipeline.Strategies = getStrategies()
	pipeline.Namespace = "devtron-demo"
	pipeline.PreStage = getPreStage(strategy, prescript)
	pipeline.PostStage = getPostStage(strategy, postscript)
	pipeline.PreStageConfigMapSecretNames = getPreStageConfigMapSecretNames()
	pipeline.PostStageConfigMapSecretNames = getPostStageConfigMapSecretNames()
	pipeline.ParentPipelineId = ParentPipelineId
	pipeline.ParentPipelineType = "CI_PIPELINE"
	pipelines = append(pipelines, pipeline)
	return pipelines
}

func getStrategies() []RequestDTOs.Strategies {
	var strategies []RequestDTOs.Strategies
	strategy := RequestDTOs.Strategies{}
	strategy.Config.Deployment.Strategy.Rolling = getRolling()
	strategy.Default = true
	strategy.DeploymentTemplate = "ROLLING"
	strategies = append(strategies, strategy)
	return strategies
}

func getRolling() RequestDTOs.Rolling {
	rolling := RequestDTOs.Rolling{}
	rolling.MaxSurge = "25%"
	rolling.MaxUnavailable = 1
	return rolling
}

func getPreStage(triggerType string, script string) RequestDTOs.Stage {
	preStage := RequestDTOs.Stage{}
	preStage.TriggerType = triggerType
	preStage.Config = script
	preStage.Name = "Pre-Deployment"
	preStage.Switch = "config"
	preStage.IsCollapse = false
	return preStage
}

func getPostStage(triggerType string, script string) RequestDTOs.Stage {
	postStage := RequestDTOs.Stage{}
	postStage.TriggerType = triggerType
	postStage.Config = script
	postStage.Name = "Post-Deployment"
	postStage.Switch = "config"
	postStage.IsCollapse = false
	return postStage
}

func getPreStageConfigMapSecretNames() RequestDTOs.StageConfigMapSecretNames {
	preStageConfigMapSecretNames := RequestDTOs.StageConfigMapSecretNames{}
	preStageConfigMapSecretNames.ConfigMaps = []string{"config1"}
	preStageConfigMapSecretNames.Secrets = []string{"secret1"}
	return preStageConfigMapSecretNames
}

func getPostStageConfigMapSecretNames() RequestDTOs.StageConfigMapSecretNames {
	postStageConfigMapSecretNames := RequestDTOs.StageConfigMapSecretNames{}
	postStageConfigMapSecretNames.ConfigMaps = []string{"config1"}
	postStageConfigMapSecretNames.Secrets = []string{"secret1"}
	return postStageConfigMapSecretNames
}
