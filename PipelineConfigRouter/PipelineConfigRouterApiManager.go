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

/////////////////=== GetAppTemplateResponseDto ====//////////////

type ContainerPort struct {
	EnvoyPort        int    `json:"envoyPort"`
	IdleTimeout      string `json:"idleTimeout"`
	Name             string `json:"name"`
	Port             int    `json:"port"`
	ServicePort      int    `json:"servicePort"`
	SupportStreaming bool   `json:"supportStreaming"`
	UseHTTP2         bool   `json:"useHTTP2"`
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

type ReadinessProbe struct {
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

type GetAppTemplateResponseDto struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Errors []Base.Errors `json:"errors"`
	Result struct {
		GlobalConfig struct {
			DefaultAppOverride struct {
				ContainerPort   []ContainerPort `json:"ContainerPort"`
				EnvVariables    []interface{}   `json:"EnvVariables"`
				GracePeriod     int             `json:"GracePeriod"`
				LivenessProbe   LivenessProbe   `json:"LivenessProbe"`
				MaxSurge        int             `json:"MaxSurge"`
				MaxUnavailable  int             `json:"MaxUnavailable"`
				MinReadySeconds int             `json:"MinReadySeconds"`
				ReadinessProbe  ReadinessProbe  `json:"ReadinessProbe"`
				Spec            struct {
					Affinity struct {
						Key    interface{} `json:"Key"`
						Values string      `json:"Values"`
						Key1   string      `json:"key"`
					} `json:"Affinity"`
				} `json:"Spec"`
				Args struct {
					Enabled bool     `json:"enabled"`
					Value   []string `json:"value"`
				} `json:"args"`
				Autoscaling struct {
					MaxReplicas                       int `json:"MaxReplicas"`
					MinReplicas                       int `json:"MinReplicas"`
					TargetCPUUtilizationPercentage    int `json:"TargetCPUUtilizationPercentage"`
					TargetMemoryUtilizationPercentage int `json:"TargetMemoryUtilizationPercentage"`
					Annotations                       struct {
					} `json:"annotations"`
					Behavior struct {
					} `json:"behavior"`
					Enabled      bool          `json:"enabled"`
					ExtraMetrics []interface{} `json:"extraMetrics"`
					Labels       struct {
					} `json:"labels"`
				} `json:"autoscaling"`
				Command struct {
					Enabled bool          `json:"enabled"`
					Value   []interface{} `json:"value"`
				} `json:"command"`
				ContainerSecurityContext struct {
				} `json:"containerSecurityContext"`
				Containers        []interface{} `json:"containers"`
				DbMigrationConfig struct {
					Enabled bool `json:"enabled"`
				} `json:"dbMigrationConfig"`
				Envoyproxy struct {
					ConfigMapName string `json:"configMapName"`
					Image         string `json:"image"`
					Resources     struct {
						Limits struct {
							Cpu    string `json:"cpu"`
							Memory string `json:"memory"`
						} `json:"limits"`
						Requests struct {
							Cpu    string `json:"cpu"`
							Memory string `json:"memory"`
						} `json:"requests"`
					} `json:"resources"`
				} `json:"envoyproxy"`
				Image struct {
					PullPolicy string `json:"pullPolicy"`
				} `json:"image"`
				ImagePullSecrets []interface{} `json:"imagePullSecrets"`
				Ingress          struct {
					Annotations struct {
					} `json:"annotations"`
					ClassName string `json:"className"`
					Enabled   bool   `json:"enabled"`
					Hosts     []struct {
						Host     string   `json:"host"`
						PathType string   `json:"pathType"`
						Paths    []string `json:"paths"`
					} `json:"hosts"`
					Labels struct {
					} `json:"labels"`
					Tls []interface{} `json:"tls"`
				} `json:"ingress"`
				IngressInternal struct {
					Annotations struct {
					} `json:"annotations"`
					ClassName string `json:"className"`
					Enabled   bool   `json:"enabled"`
					Hosts     []struct {
						Host     string   `json:"host"`
						PathType string   `json:"pathType"`
						Paths    []string `json:"paths"`
					} `json:"hosts"`
					Tls []interface{} `json:"tls"`
				} `json:"ingressInternal"`
				InitContainers  []interface{} `json:"initContainers"`
				KedaAutoscaling struct {
					Advanced struct {
					} `json:"advanced"`
					AuthenticationRef struct {
					} `json:"authenticationRef"`
					Enabled                bool   `json:"enabled"`
					EnvSourceContainerName string `json:"envSourceContainerName"`
					MaxReplicaCount        int    `json:"maxReplicaCount"`
					MinReplicaCount        int    `json:"minReplicaCount"`
					TriggerAuthentication  struct {
						Enabled bool   `json:"enabled"`
						Name    string `json:"name"`
						Spec    struct {
						} `json:"spec"`
					} `json:"triggerAuthentication"`
					Triggers []interface{} `json:"triggers"`
				} `json:"kedaAutoscaling"`
				PauseForSecondsBeforeSwitchActive int `json:"pauseForSecondsBeforeSwitchActive"`
				PodAnnotations                    struct {
				} `json:"podAnnotations"`
				PodLabels struct {
				} `json:"podLabels"`
				PodSecurityContext struct {
				} `json:"podSecurityContext"`
				Prometheus struct {
					Release string `json:"release"`
				} `json:"prometheus"`
				RawYaml      []interface{} `json:"rawYaml"`
				ReplicaCount int           `json:"replicaCount"`
				Resources    struct {
					Limits struct {
						Cpu    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"limits"`
					Requests struct {
						Cpu    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"requests"`
				} `json:"resources"`
				Secret struct {
					Data struct {
					} `json:"data"`
					Enabled bool `json:"enabled"`
				} `json:"secret"`
				Server struct {
					Deployment struct {
						Image    string `json:"image"`
						ImageTag string `json:"image_tag"`
					} `json:"deployment"`
				} `json:"server"`
				Service struct {
					Annotations struct {
					} `json:"annotations"`
					LoadBalancerSourceRanges []interface{} `json:"loadBalancerSourceRanges"`
					Type                     string        `json:"type"`
				} `json:"service"`
				ServiceAccount struct {
					Annotations struct {
					} `json:"annotations"`
					Create bool   `json:"create"`
					Name   string `json:"name"`
				} `json:"serviceAccount"`
				Servicemonitor struct {
					AdditionalLabels struct {
					} `json:"additionalLabels"`
				} `json:"servicemonitor"`
				Tolerations                     []interface{} `json:"tolerations"`
				TopologySpreadConstraints       []interface{} `json:"topologySpreadConstraints"`
				VolumeMounts                    []interface{} `json:"volumeMounts"`
				Volumes                         []interface{} `json:"volumes"`
				WaitForSecondsBeforeScalingDown int           `json:"waitForSecondsBeforeScalingDown"`
			} `json:"defaultAppOverride"`
			Readme string `json:"readme"`
			Schema struct {
				Schema     string `json:"$schema"`
				Type       string `json:"type"`
				Properties struct {
					ContainerPort struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Items       struct {
							Type       string `json:"type"`
							Properties struct {
								EnvoyPort struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"envoyPort"`
								IdleTimeout struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"idleTimeout"`
								Name struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"name"`
								Port struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"port"`
								ServicePort struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"servicePort"`
								SupportStreaming struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"supportStreaming"`
								UseHTTP2 struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"useHTTP2"`
							} `json:"properties"`
						} `json:"items"`
					} `json:"ContainerPort"`
					EnvVariables struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"EnvVariables"`
					GracePeriod struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"GracePeriod"`
					LivenessProbe struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Path struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"Path"`
							Command struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"command"`
							FailureThreshold struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"failureThreshold"`
							HttpHeaders struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"httpHeaders"`
							InitialDelaySeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"initialDelaySeconds"`
							PeriodSeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"periodSeconds"`
							Port struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"port"`
							Scheme struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"scheme"`
							SuccessThreshold struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"successThreshold"`
							Tcp struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"tcp"`
							TimeoutSeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"timeoutSeconds"`
						} `json:"properties"`
					} `json:"LivenessProbe"`
					MaxSurge struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"MaxSurge"`
					MaxUnavailable struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"MaxUnavailable"`
					MinReadySeconds struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"MinReadySeconds"`
					ReadinessProbe struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Path struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"Path"`
							Command struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"command"`
							FailureThreshold struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"failureThreshold"`
							HttpHeader struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"httpHeader"`
							InitialDelaySeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"initialDelaySeconds"`
							PeriodSeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"periodSeconds"`
							Port struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"port"`
							Scheme struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"scheme"`
							SuccessThreshold struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"successThreshold"`
							Tcp struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"tcp"`
							TimeoutSeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"timeoutSeconds"`
						} `json:"properties"`
					} `json:"ReadinessProbe"`
					Spec struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Affinity struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Key struct {
										AnyOf []struct {
											Type        string `json:"type"`
											Description string `json:"description,omitempty"`
											Title       string `json:"title,omitempty"`
										} `json:"anyOf"`
									} `json:"Key"`
									Values struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"Values"`
									Key1 struct {
										Type string `json:"type"`
									} `json:"key"`
								} `json:"properties"`
							} `json:"Affinity"`
						} `json:"properties"`
					} `json:"Spec"`
					Args struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
							Value struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Items       []struct {
									Type string `json:"type"`
								} `json:"items"`
							} `json:"value"`
						} `json:"properties"`
					} `json:"args"`
					Autoscaling struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							MaxReplicas struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"MaxReplicas"`
							MinReplicas struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"MinReplicas"`
							TargetCPUUtilizationPercentage struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"TargetCPUUtilizationPercentage"`
							TargetMemoryUtilizationPercentage struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"TargetMemoryUtilizationPercentage"`
							Behavior struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"behavior"`
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
							ExtraMetrics struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"extraMetrics"`
						} `json:"properties"`
					} `json:"autoscaling"`
					Command struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
							} `json:"enabled"`
							Value struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"value"`
						} `json:"properties"`
					} `json:"command"`
					ContainerSecurityContext struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"containerSecurityContext"`
					Containers struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
					} `json:"containers"`
					DbMigrationConfig struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
						} `json:"properties"`
					} `json:"dbMigrationConfig"`
					Envoyproxy struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							ConfigMapName struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"configMapName"`
							Image struct {
								Type        string `json:"type"`
								Description string `json:"description"`
							} `json:"image"`
							Resources struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Limits struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
										Properties  struct {
											Cpu struct {
												Type        string `json:"type"`
												Format      string `json:"format"`
												Description string `json:"description"`
												Title       string `json:"title"`
											} `json:"cpu"`
											Memory struct {
												Type        string `json:"type"`
												Format      string `json:"format"`
												Description string `json:"description"`
												Title       string `json:"title"`
											} `json:"memory"`
										} `json:"properties"`
									} `json:"limits"`
									Requests struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
										Properties  struct {
											Cpu struct {
												Type        string `json:"type"`
												Format      string `json:"format"`
												Description string `json:"description"`
												Title       string `json:"title"`
											} `json:"cpu"`
											Memory struct {
												Type        string `json:"type"`
												Format      string `json:"format"`
												Description string `json:"description"`
												Title       string `json:"title"`
											} `json:"memory"`
										} `json:"properties"`
									} `json:"requests"`
								} `json:"properties"`
							} `json:"resources"`
						} `json:"properties"`
					} `json:"envoyproxy"`
					Image struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							PullPolicy struct {
								Type        string   `json:"type"`
								Description string   `json:"description"`
								Title       string   `json:"title"`
								Enum        []string `json:"enum"`
							} `json:"pullPolicy"`
						} `json:"properties"`
					} `json:"image"`
					ImagePullSecrets struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"imagePullSecrets"`
					Ingress struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Annotations struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"annotations"`
							ClassName struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Default     string `json:"default"`
							} `json:"className"`
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
							Hosts struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Items       []struct {
									Type       string `json:"type"`
									Properties struct {
										Host struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
										} `json:"host"`
										PathType struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
										} `json:"pathType"`
										Paths struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
											Items       []struct {
												Type string `json:"type"`
											} `json:"items"`
										} `json:"paths"`
									} `json:"properties"`
								} `json:"items"`
							} `json:"hosts"`
							Tls struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"tls"`
						} `json:"properties"`
					} `json:"ingress"`
					IngressInternal struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Properties  struct {
							Annotations struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"annotations"`
							ClassName struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Default     string `json:"default"`
							} `json:"className"`
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
							Hosts struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Items       []struct {
									Type       string `json:"type"`
									Properties struct {
										Host struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
										} `json:"host"`
										PathType struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
										} `json:"pathType"`
										Paths struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
											Items       []struct {
												Type string `json:"type"`
											} `json:"items"`
										} `json:"paths"`
									} `json:"properties"`
								} `json:"items"`
							} `json:"hosts"`
							Tls struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"tls"`
						} `json:"properties"`
					} `json:"ingressInternal"`
					InitContainers struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"initContainers"`
					KedaAutoscaling struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Advanced struct {
								Type string `json:"type"`
							} `json:"advanced"`
							AuthenticationRef struct {
								Type string `json:"type"`
							} `json:"authenticationRef"`
							Enabled struct {
								Type string `json:"type"`
							} `json:"enabled"`
							EnvSourceContainerName struct {
								Type string `json:"type"`
							} `json:"envSourceContainerName"`
							MaxReplicaCount struct {
								Type string `json:"type"`
							} `json:"maxReplicaCount"`
							MinReplicaCount struct {
								Type string `json:"type"`
							} `json:"minReplicaCount"`
							TriggerAuthentication struct {
								Type       string `json:"type"`
								Properties struct {
									Enabled struct {
										Type string `json:"type"`
									} `json:"enabled"`
									Name struct {
										Type string `json:"type"`
									} `json:"name"`
									Spec struct {
										Type string `json:"type"`
									} `json:"spec"`
								} `json:"properties"`
							} `json:"triggerAuthentication"`
							Triggers struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
							} `json:"triggers"`
						} `json:"properties"`
					} `json:"kedaAutoscaling"`
					PauseForSecondsBeforeSwitchActive struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"pauseForSecondsBeforeSwitchActive"`
					PodAnnotations struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"podAnnotations"`
					PodLabels struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"podLabels"`
					PodSecurityContext struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"podSecurityContext"`
					Prometheus struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Release struct {
								Type        string `json:"type"`
								Description string `json:"description"`
							} `json:"release"`
						} `json:"properties"`
					} `json:"prometheus"`
					RawYaml struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"rawYaml"`
					ReplicaCount struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"replicaCount"`
					Resources struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Limits struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Cpu struct {
										Type        string `json:"type"`
										Format      string `json:"format"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"cpu"`
									Memory struct {
										Type        string `json:"type"`
										Format      string `json:"format"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"memory"`
								} `json:"properties"`
							} `json:"limits"`
							Requests struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Cpu struct {
										Type        string `json:"type"`
										Format      string `json:"format"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"cpu"`
									Memory struct {
										Type        string `json:"type"`
										Format      string `json:"format"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"memory"`
								} `json:"properties"`
							} `json:"requests"`
						} `json:"properties"`
					} `json:"resources"`
					Secret struct {
						Type       string `json:"type"`
						Properties struct {
							Data struct {
								Type string `json:"type"`
							} `json:"data"`
							Enabled struct {
								Type string `json:"type"`
							} `json:"enabled"`
						} `json:"properties"`
					} `json:"secret"`
					Server struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Deployment struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Image struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"image"`
									ImageTag struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"image_tag"`
								} `json:"properties"`
							} `json:"deployment"`
						} `json:"properties"`
					} `json:"server"`
					Service struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Annotations struct {
								Type        string `json:"type"`
								Title       string `json:"title"`
								Description string `json:"description"`
							} `json:"annotations"`
							Type struct {
								Type        string   `json:"type"`
								Description string   `json:"description"`
								Title       string   `json:"title"`
								Enum        []string `json:"enum"`
							} `json:"type"`
						} `json:"properties"`
					} `json:"service"`
					ServiceAccount struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Annotations struct {
								Type        string `json:"type"`
								Title       string `json:"title"`
								Description string `json:"description"`
							} `json:"annotations"`
							Name struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"name"`
							Create struct {
								Type string `json:"type"`
							} `json:"create"`
						} `json:"properties"`
					} `json:"serviceAccount"`
					Servicemonitor struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							AdditionalLabels struct {
								Type string `json:"type"`
							} `json:"additionalLabels"`
						} `json:"properties"`
					} `json:"servicemonitor"`
					Tolerations struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"tolerations"`
					TopologySpreadConstraints struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"topologySpreadConstraints"`
					VolumeMounts struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"volumeMounts"`
					Volumes struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"volumes"`
					WaitForSecondsBeforeScalingDown struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"waitForSecondsBeforeScalingDown"`
				} `json:"properties"`
			} `json:"schema"`
		} `json:"globalConfig"`
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
