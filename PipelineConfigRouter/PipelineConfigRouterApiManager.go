package PipelineConfigRouter

import (
	"automation-suite/PipelineConfigRouter/RequestDTOs"
	"automation-suite/PipelineConfigRouter/ResponseDTOs"
	"os"
	"strings"
	"testing"
	"time"

	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"strconv"
)

type DeleteResponseDto struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result string        `json:"result"`
	Errors []Base.Errors `json:"errors"`
}

type GetCdWorkflowsDto struct {
	Code   int                     `json:"code"`
	Status string                  `json:"status"`
	Result []CdWorkflowResponseDto `json:"result"`
}

type GetCdWorkflowsDetailsDto struct {
	Code   int                   `json:"code"`
	Status string                `json:"status"`
	Result CdWorkflowResponseDto `json:"result"`
}

type CiWorkflowDetailsDto struct {
	Code   int                  `json:"code"`
	Status string               `json:"status"`
	Result []CiWorkflowResponse `json:"result"`
}

type CiWorkflowResponse struct {
	Id                 int       `json:"id"`
	Name               string    `json:"name"`
	Status             string    `json:"status"`
	PodStatus          string    `json:"podStatus"`
	Message            string    `json:"message"`
	StartedOn          time.Time `json:"startedOn"`
	FinishedOn         time.Time `json:"finishedOn"`
	CiPipelineId       int       `json:"ciPipelineId"`
	Namespace          string    `json:"namespace"`
	LogLocation        string    `json:"logLocation"`
	BlobStorageEnabled bool      `json:"blobStorageEnabled"`
	TriggeredBy        int32     `json:"triggeredBy"`
	Artifact           string    `json:"artifact"`
	TriggeredByEmail   string    `json:"triggeredByEmail"`
	Stage              string    `json:"stage"`
	ArtifactId         int       `json:"artifactId"`
}

type CdWorkflowResponseDto struct {
	Id                 int       `json:"id"`
	CdWorkflowId       int       `json:"cd_workflow_id"`
	Name               string    `json:"name"`
	Status             string    `json:"status"`
	PodStatus          string    `json:"pod_status"`
	Message            string    `json:"message"`
	StartedOn          time.Time `json:"started_on"`
	FinishedOn         time.Time `json:"finished_on"`
	PipelineId         int       `json:"pipeline_id"`
	Namespace          string    `json:"namespace"`
	LogFilePath        string    `json:"log_file_path"`
	TriggeredBy        int32     `json:"triggered_by"`
	EmailId            string    `json:"email_id"`
	Image              string    `json:"image"`
	MaterialInfo       string    `json:"material_info,omitempty"`
	DataSource         string    `json:"data_source,omitempty"`
	CiArtifactId       int       `json:"ci_artifact_id,omitempty"`
	WorkflowType       string    `json:"workflow_type,omitempty"`
	ExecutorType       string    `json:"executor_type,omitempty"`
	BlobStorageEnabled bool      `json:"blobStorageEnabled"`
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
type SaveDockerRegistryRequestDto struct {
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
}

type GetContainerRegistryResponseDTO struct {
	Code   int                            `json:"code"`
	Status string                         `json:"status"`
	Result []SaveDockerRegistryRequestDto `json:"result"`
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
						//						Rolling   RequestDTOs.Rolling `json:"rolling,omitempty"`
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
	//saveAppCiPipelineRequestDTO        SaveAppCiPipelineRequestDTO
	createAppResponseDto         CreateAppResponseDto
	deleteResponseDto            DeleteResponseDto
	createAppMaterialRequestDto  CreateAppMaterialRequestDto
	createAppMaterialResponseDto CreateAppMaterialResponseDto
	getAppDetailsResponseDto     GetAppDetailsResponseDto
	saveAppCiPipelineResponseDTO SaveAppCiPipelineResponseDTO
	//getCiPipelineViaIdResponseDTO      GetCiPipelineViaIdResponseDTO
	getContainerRegistryResponseDTO    GetContainerRegistryResponseDTO
	getChartReferenceResponseDTO       GetChartReferenceResponseDTO
	getAppTemplateResponseDto          GetAppTemplateResponseDto
	getCdPipelineStrategiesResponseDto GetCdPipelineStrategiesResponseDto
	pipelineSuggestedCDResponseDTO     PipelineSuggestedCDResponseDTO
	environmentDetailsResponseDTO      EnvironmentDetailsResponseDTO
	saveDeploymentTemplateResponseDTO  SaveDeploymentTemplateResponseDTO
	getWorkflowDetails                 RequestDTOs.GetWorkflowDetails
	createWorkflowResponseDto          ResponseDTOs.CreateWorkflowResponseDto
	fetchSuggestedCiPipelineName       FetchSuggestedCiPipelineName
	fetchAllAppWorkflowResponseDto     FetchAllAppWorkflowResponseDto
	getAppDeploymentStatusTimelineDto  ResponseDTOs.GetAppDeploymentStatusTimelineDTO
	saveCdPipelineRequestDTO           RequestDTOs.SaveCdPipelineRequestDTO
	saveCdPipelineResponseDTO          ResponseDTOs.SaveCdPipelineResponseDTO
	deleteCdPipelineRequestDTO         RequestDTOs.DeleteCdPipelineRequestDTO
	getCdPipeResponseDTO               ResponseDTOs.GetCdPipeResponseDTO
	getWorkflowStatusResponseDTO       ResponseDTOs.GetWorkflowStatusResponseDTO
	getCiPipelineMaterialResponseDTO   ResponseDTOs.GetCiPipelineMaterialResponseDTO
	triggerCiPipelineResponseDTO       ResponseDTOs.TriggerCiPipelineResponseDTO
	updateAppMaterialResponseDTO       ResponseDTOs.UpdateAppMaterialResponseDTO
	appListForAutocompleteResponseDTO  ResponseDTOs.AppListForAutocompleteResponseDTO
	appListByTeamIdsResponseDTO        ResponseDTOs.AppListByTeamIdsResponseDTO
	fetchMaterialsResponseDTO          ResponseDTOs.FetchMaterialsResponseDTO
	getCiPipelineMinResponseDTO        ResponseDTOs.GetCiPipelineMinResponseDTO
	refreshMaterialsResponseDTO        ResponseDTOs.RefreshMaterialsResponseDTO
	saveAppCiPipelineRequestDTO        RequestDTOs.SaveAppCiPipelineRequestDTO
	getCiPipelineViaIdResponseDTO      ResponseDTOs.GetCiPipelineViaIdResponseDTO
}

/*type EnvironmentConfigPipelineConfigRouter struct {
	GitHubProjectUrl       string `env:"GITHUB_URL_TO_CLONE_PROJECT" envDefault:"https://github.com/devtron-labs/sample-go-app.git"`
	DockerRegistry         string `env:"DOCKER_REGISTRY" envDefault:"devtron-quay-dpk"`
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
}*/

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
	//pipelineConfig, _ := GetEnvironmentConfigPipelineConfigRouter()
	envConf := Base.ReadBaseEnvConfig()
	file := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)
	var slice AppMaterials
	slice.Url = file.GitHubProjectUrl
	slice.GitProviderId = gitProviderId
	slice.FetchSubmodules = fetchSubmodules
	slice.CheckoutPath = "./"
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
		Material   []Materials `json:"material"`
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

func HitGetApp(appId int, authToken string) GetAppDetailsResponseDto {
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
func GetRequestPayloadForSaveAppCiPipeline(AppId int, dockerRegistry string, dockerRepository string, dockerfilePath string, dockerfileRepository string, dockerfileRelativePath string, gitMaterialId int) RequestDTOs.SaveAppCiPipelineRequestDTO {
	CiBuildConfig := GetCiBuildConfig(dockerfilePath, dockerfileRepository, dockerfileRelativePath, gitMaterialId)
	saveAppCiPipelineRequestDTO := RequestDTOs.SaveAppCiPipelineRequestDTO{}
	saveAppCiPipelineRequestDTO.AppId = AppId
	saveAppCiPipelineRequestDTO.CiBuildConfig = CiBuildConfig
	saveAppCiPipelineRequestDTO.DockerRepository = dockerRepository
	saveAppCiPipelineRequestDTO.DockerRegistry = dockerRegistry
	return saveAppCiPipelineRequestDTO
}

func GetCiBuildConfig(dockerfilePath string, dockerfileRepository string, dockerfileRelativePath string, gitMaterialId int) RequestDTOs.CiBuildConfig {
	CiBuildConfig := RequestDTOs.CiBuildConfig{}
	CiBuildConfig.GitMaterialId = gitMaterialId
	CiBuildConfig.CiBuildType = "self-dockerfile-build"
	CiBuildConfig.DockerBuildConfig.DockerfileContent = ""
	CiBuildConfig.DockerBuildConfig.DockerfilePath = dockerfilePath
	CiBuildConfig.DockerBuildConfig.DockerfileRepository = dockerfileRepository
	CiBuildConfig.DockerBuildConfig.DockerfileRelativePath = dockerfileRelativePath
	//CiBuildConfig.DockerBuildConfig.TargetPlatform = "linux/arm64"
	return CiBuildConfig
}

func HitSaveAppCiPipeline(payload []byte, authToken string) SaveAppCiPipelineResponseDTO {
	resp, err := Base.MakeApiCall(SaveAppCiPipelineApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveAppCiPipelineApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), SaveAppCiPipelineApi)
	return pipelineConfigRouter.saveAppCiPipelineResponseDTO
}

func HitGetCiPipelineViaId(appId string, authToken string) ResponseDTOs.GetCiPipelineViaIdResponseDTO {
	resp, err := Base.MakeApiCall(GetCiPipelineViaIdApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetCiPipelineViaIdApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetCiPipelineViaIdApi)
	return pipelineConfigRouter.getCiPipelineViaIdResponseDTO
}

func HitGetContainerRegistry(appId string, authToken string) GetContainerRegistryResponseDTO {
	resp, err := Base.MakeApiCall(PipelineRouterBaseApiUrl+appId+"/autocomplete/docker", http.MethodGet, "", nil, authToken)
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

func HitGetAllEnvironmentDetails(queryParams map[string]string, authToken string) EnvironmentDetailsResponseDTO {
	resp, err := Base.MakeApiCall(GetEnvAutocompleteApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetEnvAutocompleteApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetEnvAutocompleteApi)
	return pipelineConfigRouter.environmentDetailsResponseDTO
}

func GetRequestPayloadForSaveDeploymentTemplate(AppId int, chartRefId int, defaultOverride DefaultAppOverride) SaveDeploymentTemplateRequestDTO {
	saveDeploymentTemplateRequestDTO := SaveDeploymentTemplateRequestDTO{}
	saveDeploymentTemplateRequestDTO.AppId = AppId
	saveDeploymentTemplateRequestDTO.ChartRefId = chartRefId
	saveDeploymentTemplateRequestDTO.ValuesOverride = defaultOverride
	saveDeploymentTemplateRequestDTO.ValuesOverride.Spec.Affinity.Key = "node"
	saveDeploymentTemplateRequestDTO.ValuesOverride.Spec.Affinity.Values = "devtron"
	saveDeploymentTemplateRequestDTO.DefaultAppOverride = defaultOverride
	saveDeploymentTemplateRequestDTO.CurrentViewEditor = "ADVANCED"
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

func HitForceDeleteCdPipelineApi(payload []byte, authToken string) RequestDTOs.DeleteCdPipelineRequestDTO {
	resp, err := Base.MakeApiCall(ForceDeleteCdPipelineApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, DeleteCdPipelineApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteCdPipelineApi)
	return pipelineConfigRouter.deleteCdPipelineRequestDTO
}

func HitLogsDownloadApi(url string, authToken string, t *testing.T, indexOfLog int, logString string) {
	Base.ReadEventStreamsForSpecificApiAndVerifyResult(url, authToken, t, indexOfLog, logString)
}

func FetchCiWorkflows(url string, authToken string) (*CiWorkflowDetailsDto, error) {
	queryParams := make(map[string]string)
	queryParams["offset"] = "0"
	queryParams["size"] = "5"
	apiCall, err := Base.MakeApiCall(url, http.MethodGet, "", queryParams, authToken)
	if err != nil {
		return nil, err
	}
	ciWorkflowDetailsDto := &CiWorkflowDetailsDto{}
	err = json.Unmarshal(apiCall.Body(), ciWorkflowDetailsDto)
	return ciWorkflowDetailsDto, err

}

func HitCiArtifactsDownloadApi(url string, authToken string) (int, error) {
	apiCall, err := Base.MakeApiCall(url, http.MethodGet, "", nil, authToken)
	if err != nil {
		return 0, err
	}
	statusCode := apiCall.StatusCode()
	return statusCode, err
}

func FetchCdWorkflowRunnerDetails(url string, authToken string) (*GetCdWorkflowsDetailsDto, error) {
	apiCall, err := Base.MakeApiCall(url, http.MethodGet, "", nil, authToken)
	if err != nil {
		return nil, err
	}
	cdWorkflowDetails := &GetCdWorkflowsDetailsDto{}
	err = json.Unmarshal(apiCall.Body(), cdWorkflowDetails)
	return cdWorkflowDetails, err
}

func FetchCdPipelineWorkflows(url string, authToken string) (*GetCdWorkflowsDto, error) {
	queryParams := make(map[string]string)
	queryParams["offset"] = "0"
	queryParams["size"] = "5"
	apiCall, err := Base.MakeApiCall(url, http.MethodGet, "", queryParams, authToken)
	if err != nil {
		return nil, err
	}
	getCdWorkflowsDto := &GetCdWorkflowsDto{}
	err = json.Unmarshal(apiCall.Body(), getCdWorkflowsDto)
	return getCdWorkflowsDto, err
}

func GetPayloadForDeleteCdPipeline(AppId int, pipelineId int) RequestDTOs.DeleteCdPipelineRequestDTO {
	deleteRequest := RequestDTOs.DeleteCdPipelineRequestDTO{}
	deleteRequest.Pipeline.Id = pipelineId
	deleteRequest.Action = 1
	deleteRequest.AppId = AppId
	return deleteRequest
}

func HitGetAppCdPipeline(appId string, authToken string) ResponseDTOs.GetCdPipeResponseDTO {
	resp, err := Base.MakeApiCall(GetAppCdPipelineApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAppCdPipelineApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppCdPipelineApi)
	return pipelineConfigRouter.getCdPipeResponseDTO
}

func HitGetWorkflowStatus(appId int, authToken string) ResponseDTOs.GetWorkflowStatusResponseDTO {
	id := strconv.Itoa(appId)
	resp, err := Base.MakeApiCall(GetWorkflowStatusApiUrl+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetWorkflowStatusApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetWorkflowStatusApi)
	return pipelineConfigRouter.getWorkflowStatusResponseDTO
}

func HitGetCiPipelineMaterial(ciPipelineId int, authToken string) ResponseDTOs.GetCiPipelineMaterialResponseDTO {
	pipelineId := strconv.Itoa(ciPipelineId)
	resp, err := Base.MakeApiCall(GetCiPipelineViaIdApiUrl+pipelineId+"/material", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetCiPipelineMaterialApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetCiPipelineMaterialApi)
	return pipelineConfigRouter.getCiPipelineMaterialResponseDTO
}

func HitTriggerCiPipelineApi(payload []byte, authToken string) ResponseDTOs.TriggerCiPipelineResponseDTO {
	resp, err := Base.MakeApiCall(TriggerCiPipelineApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, TriggerCiPipelineApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), TriggerCiPipelineApi)
	return pipelineConfigRouter.triggerCiPipelineResponseDTO
}

func CreatePayloadForTriggerCiPipeline(commit string, PipelineId int, ciPipelineMaterialId int, invalidateCache bool) RequestDTOs.TriggerCiPipelineRequestDTO {
	var listOfCiPipelineMaterials []RequestDTOs.CiPipelineMaterials
	listOfCiPipelineMaterials = append(listOfCiPipelineMaterials, getCiPipelineMaterials(commit, ciPipelineMaterialId))
	TriggerCiPipelineRequestDTO := RequestDTOs.TriggerCiPipelineRequestDTO{}
	TriggerCiPipelineRequestDTO.PipelineId = PipelineId
	TriggerCiPipelineRequestDTO.CiPipelineMaterials = listOfCiPipelineMaterials
	TriggerCiPipelineRequestDTO.InvalidateCache = invalidateCache
	return TriggerCiPipelineRequestDTO
}

func getCiPipelineMaterials(commit string, ciPipelineMaterialId int) RequestDTOs.CiPipelineMaterials {
	gitCommit := RequestDTOs.GitCommit{}
	gitCommit.Commit = commit
	CiPipelineMaterial := RequestDTOs.CiPipelineMaterials{}
	CiPipelineMaterial.Id = ciPipelineMaterialId
	CiPipelineMaterial.GitCommit = gitCommit
	return CiPipelineMaterial
}

func GetPayloadForUpdateAppMaterial(appId int, url string, id int, GitProvidedId int, CheckoutPath string, FetchSubmodules bool) RequestDTOs.UpdateAppMaterialRequestDTO {
	requestDTO := RequestDTOs.UpdateAppMaterialRequestDTO{}
	requestDTO.AppId = appId
	requestDTO.Material.Url = url
	requestDTO.Material.CheckoutPath = CheckoutPath
	requestDTO.Material.Id = id
	requestDTO.Material.GitProviderId = GitProvidedId
	requestDTO.Material.FetchSubmodules = FetchSubmodules
	return requestDTO
}
func HitUpdateAppMaterialApi(payload []byte, authToken string) ResponseDTOs.UpdateAppMaterialResponseDTO {
	resp, err := Base.MakeApiCall(CreateAppMaterialApiUrl, http.MethodPut, string(payload), nil, authToken)
	Base.HandleError(err, UpdateAppMaterial)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), UpdateAppMaterial)
	return pipelineConfigRouter.updateAppMaterialResponseDTO
}

func HitGetAppListForAutocomplete(authToken string) ResponseDTOs.AppListForAutocompleteResponseDTO {
	resp, err := Base.MakeApiCall(GetAppListForAutocompleteApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAppListForAutocompleteApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppListForAutocompleteApi)
	return pipelineConfigRouter.appListForAutocompleteResponseDTO
}

func HitGetAppListByTeamIds(queryParams map[string]string, authToken string) ResponseDTOs.AppListByTeamIdsResponseDTO {
	resp, err := Base.MakeApiCall(GetAppListByTeamIdsApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetAppListByTeamIdsApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppListByTeamIdsApi)
	return pipelineConfigRouter.appListByTeamIdsResponseDTO
}

func HitFindAppsByTeamId(teamId string, authToken string) ResponseDTOs.AppListForAutocompleteResponseDTO {
	resp, err := Base.MakeApiCall(FindAppsByTeamIdApiUrl+teamId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAppListForAutocompleteApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppListForAutocompleteApi)
	return pipelineConfigRouter.appListForAutocompleteResponseDTO
}

func HitFindAppsByTeamName(teamName string, authToken string) ResponseDTOs.AppListForAutocompleteResponseDTO {
	resp, err := Base.MakeApiCall(FindAppsByTeamNameApiUrl+teamName, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAppListForAutocompleteApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppListForAutocompleteApi)
	return pipelineConfigRouter.appListForAutocompleteResponseDTO
}

func HitFetchMaterialsApi(pipelineId string, authToken string) ResponseDTOs.FetchMaterialsResponseDTO {
	resp, err := Base.MakeApiCall(FetchMaterialsApiUrl+pipelineId+"/material", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchMaterialsApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), FetchMaterialsApi)
	return pipelineConfigRouter.fetchMaterialsResponseDTO
}

func HitGetCiPipelineMin(appId string, authToken string) ResponseDTOs.GetCiPipelineMinResponseDTO {
	resp, err := Base.MakeApiCall(PipelineRouterBaseApiUrl+appId+"/ci-pipeline/min", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetCiPipelineMinApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetCiPipelineMinApi)
	return pipelineConfigRouter.getCiPipelineMinResponseDTO
}

func HitRefreshMaterialsApi(gitMaterialId string, authToken string) ResponseDTOs.RefreshMaterialsResponseDTO {
	resp, err := Base.MakeApiCall(RefreshMaterialsApiUrl+gitMaterialId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, RefreshMaterialsApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), RefreshMaterialsApi)
	return pipelineConfigRouter.refreshMaterialsResponseDTO
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
	case GetPipelineSuggestedCICDApi:
		json.Unmarshal(response, &structPipelineConfigRouter.pipelineSuggestedCDResponseDTO)
	case GetEnvAutocompleteApi:
		json.Unmarshal(response, &structPipelineConfigRouter.environmentDetailsResponseDTO)
	case SaveDeploymentTemplateApi:
		json.Unmarshal(response, &structPipelineConfigRouter.saveDeploymentTemplateResponseDTO)
	case PatchCiPipelinesApi:
		json.Unmarshal(response, &structPipelineConfigRouter.createWorkflowResponseDto)
	case FetchSuggestedCiPipelineNameApi:
		json.Unmarshal(response, &structPipelineConfigRouter.fetchSuggestedCiPipelineName)
	case GetWorkflowDetailsApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getWorkflowDetails)
	case DeleteAppApi:
		json.Unmarshal(response, &structPipelineConfigRouter.deleteResponseDto)
	case FetchAllAppWorkflowApi:
		json.Unmarshal(response, &structPipelineConfigRouter.fetchAllAppWorkflowResponseDto)
	case SaveCdPipelineApi:
		json.Unmarshal(response, &structPipelineConfigRouter.saveCdPipelineResponseDTO)
	case GetAppCdPipelineApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getCdPipeResponseDTO)
	case GetWorkflowStatusApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getWorkflowStatusResponseDTO)
	case GetCiPipelineMaterialApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getCiPipelineMaterialResponseDTO)
	case TriggerCiPipelineApi:
		json.Unmarshal(response, &structPipelineConfigRouter.triggerCiPipelineResponseDTO)
	case UpdateAppMaterial:
		json.Unmarshal(response, &structPipelineConfigRouter.updateAppMaterialResponseDTO)
	case GetAppListForAutocompleteApi:
		json.Unmarshal(response, &structPipelineConfigRouter.appListForAutocompleteResponseDTO)
	case GetAppListByTeamIdsApi:
		json.Unmarshal(response, &structPipelineConfigRouter.appListByTeamIdsResponseDTO)
	case FetchMaterialsApi:
		json.Unmarshal(response, &structPipelineConfigRouter.fetchMaterialsResponseDTO)
	case GetCiPipelineMinApi:
		json.Unmarshal(response, &structPipelineConfigRouter.getCiPipelineMinResponseDTO)
	case RefreshMaterialsApi:
		json.Unmarshal(response, &structPipelineConfigRouter.refreshMaterialsResponseDTO)
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

func (suite *PipelinesConfigRouterTestSuite) CreateAppMaterial() CreateAppMaterialResponseDto {
	createAppMaterialRequestDto := GetAppMaterialRequestDto(suite.createAppResponseDto.Result.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponseDto := HitCreateAppMaterialApi(appMaterialByteValue, suite.createAppResponseDto.Result.Id, 1, false, suite.authToken)
	return createAppMaterialResponseDto
}

func (suite *PipelinesConfigRouterTestSuite) CreateApp() CreateAppResponseDto {
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	createAppRequestDto := GetAppRequestDto("app"+appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, "app"+appName, 1, 0, suite.authToken)
	return createAppResponseDto
}

func (suite *PipelinesConfigRouterTestSuite) TearDownSuite() {
	log.Println("=== Running the after suite method for deleting the data created via automation ===")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(suite.createAppResponseDto.Result.Id, suite.createAppResponseDto.Result.AppName, suite.createAppResponseDto.Result.TeamId, suite.createAppResponseDto.Result.TemplateId)
	HitDeleteAppApi(byteValueOfDeleteApp, suite.createAppResponseDto.Result.Id, suite.authToken)
	//Base.DeleteFile("OutputDataGetChartReferenceViaAppId")
}

/////////////////=== Create Workflow API ====//////////////

func HitPatchCiPipelinesApi(payload []byte, authToken string) ResponseDTOs.CreateWorkflowResponseDto {
	resp, err := Base.MakeApiCall(PatchCiPipelinesApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, PatchCiPipelinesApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), PatchCiPipelinesApi)
	return pipelineConfigRouter.createWorkflowResponseDto
}
func workflowTypeProvider(temp string) string {
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
func getRequestPayloadForCreateWorkflow(forDelete bool, wfTypeId string, appId int, wfId int) RequestDTOs.CreateWorkflowRequestDto {
	var createWorkflowRequestDto RequestDTOs.CreateWorkflowRequestDto

	if forDelete == true {
		createWorkflowRequestDto.AppWorkflowId = wfId
		createWorkflowRequestDto.AppId = appId
		createWorkflowRequestDto.Action = 2
		return createWorkflowRequestDto
	}
	wfTypeStr := workflowTypeProvider(wfTypeId)
	var CiMaterial RequestDTOs.CiMaterial
	CiMaterial.Source.Type = wfTypeStr
	createWorkflowRequestDto.CiPipeline.Active = true
	createWorkflowRequestDto.AppId = appId
	createWorkflowRequestDto.CiPipeline.CiMaterial = append(createWorkflowRequestDto.CiPipeline.CiMaterial, CiMaterial)
	return createWorkflowRequestDto
}
func HitGetWorkflowDetailsApi(appId int, wfId int, authToken string) RequestDTOs.GetWorkflowDetails {
	resp, err := Base.MakeApiCall(GetCiPipelineViaIdApiUrl+strconv.Itoa(appId)+"/"+strconv.Itoa(wfId), http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetWorkflowDetailsApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetWorkflowDetailsApi)
	return pipelineConfigRouter.getWorkflowDetails
}
func HitDeleteWorkflowApi(appId int, wfId int, authToken string) DeleteResponseDto {
	resp, err := Base.MakeApiCall(GetWorkflowApiUrl+strconv.Itoa(appId)+"/"+strconv.Itoa(wfId), http.MethodDelete, "", nil, authToken)
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

type PipelineSuggestedCDResponseDTO struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result string        `json:"result"`
	Errors []Base.Errors `json:"errors"`
}

func HitGetPipelineSuggestedCiCd(pipelineType string, appId int, authToken string) PipelineSuggestedCDResponseDTO {
	resp, err := Base.MakeApiCall(GetPipelineSuggestedCICDApiUrl+pipelineType+"/"+strconv.Itoa(appId), http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetPipelineSuggestedCICDApi)
	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), GetPipelineSuggestedCICDApi)
	return pipelineConfigRouter.pipelineSuggestedCDResponseDTO
}

func getConditionDetails(id int) []RequestDTOs.ConditionDetails {
	var conditionDetails RequestDTOs.ConditionDetails
	conditionDetails.ConditionType = "TRIGGER"
	switch id {
	case 1:
		conditionDetails.ConditionOperator = "=="
	case 2:
		conditionDetails.ConditionOperator = "!="
	case 3:
		conditionDetails.ConditionOperator = ">"
	case 4:
		conditionDetails.ConditionOperator = "<"
	case 5:
		conditionDetails.ConditionOperator = "<="

	case 6:
		conditionDetails.ConditionOperator = ">="
	}
	var input []RequestDTOs.ConditionDetails
	input = append(input, conditionDetails)
	return input
}

func getPreBuildStepRequestPayloadDto(index int, scriptType string) RequestDTOs.Step {
	var step RequestDTOs.Step
	step.Name = "TASK" + Base.GetRandomStringOfGivenLength(5)
	step.Description = "This is Random Description of Step for testing purpose"
	step.Index = index
	step.StepType = "INLINE"
	step.OutputDirectoryPath = append(step.OutputDirectoryPath, "/test/output")
	step.InlineStepDetail = getInlineStepDetails(scriptType)
	return step
}

func getInlineStepDetails(scriptType string) RequestDTOs.InlineStepDetail {
	inlineStepDetail := RequestDTOs.InlineStepDetail{}
	if scriptType == "SHELL" {
		inlineStepDetail.ScriptType = scriptType
		inlineStepDetail.Script = "#!/bin/sh \nset -eo pipefail \n#set -v  ## uncomment this to debug the script \n\necho \"Here I am Printing value of VarString\"\necho $VarString\n\necho \"Here I am Printing value of VarBool\"\necho $VarBool\n\necho \"Here I am Printing value of VarNumber\"\necho $VarNumber\n\necho \"Here I am Printing value of VarDate\"\necho $VarDate\n\necho \"Here I am Printing value of VarDockerImage\"\necho $VarDockerImage\n\necho \"Here I am exporting Var1\"\nexport Var1=$Var1\n\necho \"Here I am exporting Var1\"\nexport VarBool=$VarBool\n\necho \"Here I am exporting Var1\"\nexport VarNumber=$VarNumber\n\necho \"Here I am exporting Var1\"\nexport VarDate=$VarDate\n\necho \"Here I am exporting Var1\"\nexport VarDockerImage=$VarDockerImage"

	}
	if scriptType == "CONTAINER_IMAGE" {
		inlineStepDetail.ScriptType = "CONTAINER_IMAGE"
		inlineStepDetail.Script = "#!/bin/sh \nset -eo pipefail \n#set -v  ## uncomment this to debug the script \n\necho \"Here I am Printing value of VarString\"\necho $VarString\n\necho \"Here I am Printing value of VarBool\"\necho $VarBool\n\necho \"Here I am Printing value of VarNumber\"\necho $VarNumber\n\necho \"Here I am Printing value of VarDate\"\necho $VarDate\n\necho \"Here I am Printing value of VarDockerImage\"\necho $VarDockerImage\n\necho \"Here I am exporting Var1\"\nexport Var1=$Var1\n\necho \"Here I am exporting Var1\"\nexport VarBool=$VarBool\n\necho \"Here I am exporting Var1\"\nexport VarNumber=$VarNumber\n\necho \"Here I am exporting Var1\"\nexport VarDate=$VarDate\n\necho \"Here I am exporting Var1\"\nexport VarDockerImage=$VarDockerImage"
		inlineStepDetail.StoreScriptAt = "mounted/script.sh"
		commandArgsMap := getCommandArgsMap()
		inlineStepDetail.CommandArgsMap = append(inlineStepDetail.CommandArgsMap, commandArgsMap)
		inlineStepDetail.MountCodeToContainerPath = "/sourcecode"
		inlineStepDetail.MountDirectoryFromHost = true
		inlineStepDetail.ContainerImagePath = "alpine:latest"
		inlineStepDetail.IsMountCustomScript = true
		inlineStepDetail.MountCodeToContainer = true
		portmap := getPortMap()
		mountPathMap := getMountPathMap()
		inlineStepDetail.PortMap = append(inlineStepDetail.PortMap, portmap)
		inlineStepDetail.MountPathMap = append(inlineStepDetail.MountPathMap, mountPathMap)
	}
	for i := 1; i < 6; i++ {
		inputVariables := inputVariablesSelector(i)
		inlineStepDetail.InputVariables = append(inlineStepDetail.InputVariables, inputVariables)
		inlineStepDetail.OutputVariables = append(inlineStepDetail.OutputVariables, inputVariables)
		conditionDetails := getConditionDetailsForGivenConditionType(i, "TRIGGER")
		inlineStepDetail.ConditionDetails = append(inlineStepDetail.ConditionDetails, conditionDetails)
		conditionDetails = getConditionDetailsForGivenConditionType(i, "PASS")
		inlineStepDetail.ConditionDetails = append(inlineStepDetail.ConditionDetails, conditionDetails)
	}
	return inlineStepDetail
}
func getCommandArgsMap() RequestDTOs.CommandArgsMap {
	var commandArgsMap RequestDTOs.CommandArgsMap
	commandArgsMap.Command = "sh"
	args := []string{"/mounted/test.sh"}
	commandArgsMap.Args = args
	return commandArgsMap
}

func getPortMap() RequestDTOs.PortMap {
	var portMap RequestDTOs.PortMap
	portMap.PortOnContainer = 8080
	portMap.PortOnLocal = 9000
	return portMap
}

func getMountPathMap() RequestDTOs.MountPathMap {
	var mountPathMap RequestDTOs.MountPathMap
	mountPathMap.FilePathOnDisk = "./"
	mountPathMap.FilePathOnContainer = "./"
	return mountPathMap
}
func getConditionDetailsForGivenConditionType(id int, ConditionType string) RequestDTOs.ConditionDetails {
	var conditionDetails RequestDTOs.ConditionDetails
	conditionDetails.ConditionType = ConditionType
	switch id {
	case 1:
		conditionDetails.ConditionOperator = "=="
		conditionDetails.ConditionOnVariable = "VarNumber"
		conditionDetails.ConditionalValue = "1"
	case 2:
		conditionDetails.ConditionOperator = "!="
		conditionDetails.ConditionOnVariable = "VarNumber"
		conditionDetails.ConditionalValue = "3"
	case 3:
		conditionDetails.ConditionOperator = ">"
		conditionDetails.ConditionOnVariable = "VarNumber"
		conditionDetails.ConditionalValue = "0"
	case 4:
		conditionDetails.ConditionOperator = "<"
		conditionDetails.ConditionOnVariable = "VarNumber"
		conditionDetails.ConditionalValue = "5"
	case 5:
		conditionDetails.ConditionOperator = "<="
		conditionDetails.ConditionOnVariable = "VarNumber"
		conditionDetails.ConditionalValue = "3"

	case 6:
		conditionDetails.ConditionOperator = ">="
		conditionDetails.ConditionOnVariable = "VarNumber"
		conditionDetails.ConditionalValue = "0"
	}
	return conditionDetails
}

func inputVariablesSelector(inputType int) RequestDTOs.InputVariables {
	var inputVariable RequestDTOs.InputVariables
	switch inputType {
	case 1:
		inputVariable.Format = "STRING"
		inputVariable.Name = "VarString"
		inputVariable.Value = "Deepak"
		inputVariable.VariableType = "NEW"
		inputVariable.Description = "This is description of Variable ==>" + inputVariable.Name
	case 2:
		inputVariable.Format = "BOOL"
		inputVariable.Value = "true"
		inputVariable.VariableType = "NEW"
		inputVariable.Name = "VarBool"
		inputVariable.Description = "This is description of Variable ==>" + inputVariable.Name
	case 3:
		inputVariable.Format = "NUMBER"
		inputVariable.Value = "1"
		inputVariable.VariableType = "NEW"
		inputVariable.Name = "VarNumber"
		inputVariable.Description = "This is description of Variable ==>" + inputVariable.Name
	case 4:
		inputVariable.Format = "DATE"
		inputVariable.Value = "2006-01-02"
		inputVariable.VariableType = "NEW"
		inputVariable.Name = "VarDate"
		inputVariable.Description = "This is description of Variable ==>" + inputVariable.Name
	case 5:
		inputVariable.Format = "STRING"
		inputVariable.VariableType = "GLOBAL"
		inputVariable.RefVariableName = "DOCKER_IMAGE_TAG"
		inputVariable.Name = "VarDockerImage"
		inputVariable.Description = "This is description of Variable ==>" + inputVariable.Name
	}
	return inputVariable
}

func HitCreateWorkflowApiWithFullPayload(appId int, authToken string) ResponseDTOs.CreateWorkflowResponseDto {
	var createWorkflowRequestDto RequestDTOs.CreateWorkflowRequestDto
	configFile, err := os.Open("../testdata/PipeLineConfigRouter/CreateWorkflow/CreateWorkflowPreAndPostBuildRequestPayload.json")
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&createWorkflowRequestDto); err != nil {
		panic(err)
	}

	expectedPayload := getRequestPayloadForCreateWorkflow(false, "1", appId, 0)

	createWorkflowRequestDto.AppId = appId
	createWorkflowRequestDto.CiPipeline.Active = expectedPayload.CiPipeline.Active
	fetchSuggestedCiPipelineName := HitGetPipelineSuggestedCiCd("ci", appId, authToken)
	createWorkflowRequestDto.CiPipeline.Name = fetchSuggestedCiPipelineName.Result

	fetchAppGetResponseDto := HitGetApp(appId, authToken)

	branchValue := "main"

	for _, j := range fetchAppGetResponseDto.Result.Material {

		var CiMaterial RequestDTOs.CiMaterial
		CiMaterial.GitMaterialId = j.Id
		CiMaterial.Source.Type = expectedPayload.CiPipeline.CiMaterial[0].Source.Type
		CiMaterial.Source.Value = branchValue
		createWorkflowRequestDto.CiPipeline.CiMaterial = append(createWorkflowRequestDto.CiPipeline.CiMaterial, CiMaterial)

	}
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].GitMaterialId = fetchAppGetResponseDto.Result.Material[0].Id

	byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
	log.Println("Hitting the Create Workflow Api with valid payload")
	createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, authToken)
	return createWorkflowResponseDto
}

func DeleteCiPipeline(appId int, ciPipelineId int, authToken string) {
	getWorkflowDetailsResponseDto := HitGetWorkflowDetailsApi(appId, ciPipelineId, authToken)
	log.Println("Getting data for delete ci-pipeline")
	deleteCiPipelineRequestDto := getRequestPayloadForCreateWorkflow(true, "1", appId, ciPipelineId)
	deleteCiPipelineRequestDto.CiPipeline = getWorkflowDetailsResponseDto.Result
	log.Println("Removing the data created via ci-pipeline API")
	byteValueOfDeleteCiPipeline, _ := json.Marshal(deleteCiPipelineRequestDto)
	log.Println("Hitting the Create Workflow Api with action=2 for delete ci-pipeline")
	HitPatchCiPipelinesApi(byteValueOfDeleteCiPipeline, authToken)
	log.Println("Deleting workflow")
	return
}

func GetRequestPayloadForSaveCdPipelineApi(appId int, AppWorkflowId int, EnvironmentId int, CiPipelineId int, ParentPipelineId int, strategy string, prescript string, postscript string, pipelineTriggerType string) RequestDTOs.SaveCdPipelineRequestDTO {
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
	preStageConfigMapSecretNames.ConfigMaps = []string{"kubernetes-config1"}
	preStageConfigMapSecretNames.Secrets = []string{"kubernetes-secret1"}
	return preStageConfigMapSecretNames
}

func getPostStageConfigMapSecretNames() RequestDTOs.StageConfigMapSecretNames {
	postStageConfigMapSecretNames := RequestDTOs.StageConfigMapSecretNames{}
	postStageConfigMapSecretNames.ConfigMaps = []string{"kubernetes-config1"}
	postStageConfigMapSecretNames.Secrets = []string{"kubernetes-secret1"}
	return postStageConfigMapSecretNames
}

type FetchAllAppWorkflowResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		AppId     int    `json:"appId"`
		AppName   string `json:"appName"`
		Workflows []struct {
			Id    int    `json:"id"`
			Name  string `json:"name"`
			AppId int    `json:"appId"`
			Tree  []struct {
				Id            int    `json:"id"`
				AppWorkflowId int    `json:"appWorkflowId"`
				Type          string `json:"type"`
				ComponentId   int    `json:"componentId"`
				ParentId      int    `json:"parentId"`
				ParentType    string `json:"parentType"`
			} `json:"tree"`
		} `json:"workflows"`
	} `json:"result"`
	Error []Base.Errors `json:"errors"`
}

func FetchAllAppWorkflow(id int, authToken string) FetchAllAppWorkflowResponseDto {
	resp, err := Base.MakeApiCall(GetWorkflowApiUrl+strconv.Itoa(id), http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchAllAppWorkflowApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), FetchAllAppWorkflowApi)
	return pipelineConfigRouter.fetchAllAppWorkflowResponseDto
}

func GetAppDeploymentStatusTimeline(appId int, envId int, authToken string) ResponseDTOs.GetAppDeploymentStatusTimelineDTO {
	resp, err := Base.MakeApiCall(GetAppDeploymentStatusTimelineApiUrl+strconv.Itoa(appId)+"/"+strconv.Itoa(envId), http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchAllAppWorkflowApi)

	structPipelineConfigRouter := StructPipelineConfigRouter{}
	pipelineConfigRouter := structPipelineConfigRouter.UnmarshalGivenResponseBody(resp.Body(), FetchAllAppWorkflowApi)
	return pipelineConfigRouter.getAppDeploymentStatusTimelineDto
}
