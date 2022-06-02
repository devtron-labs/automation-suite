package AppListingRouter

import (
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

type StructAppListingRouter struct {
	fetchAllStageStatusResponseDto FetchAllStageStatusResponseDto
	createAppResponseDto           CreateAppResponseDto
	createAppMaterialResponseDto   CreateAppMaterialResponseDto
	deleteResponseDto              DeleteResponseDto
	fetchOtherEnvResponseDto       FetchOtherEnvResponseDto
}
type FetchAllStageStatusResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Stage     int    `json:"stage"`
		StageName string `json:"stageName"`
		Status    bool   `json:"status"`
		Required  bool   `json:"required"`
	} `json:"result"`
	Errors []Base.Errors `json:"errors"`
}

func (structAppListingRouter StructAppListingRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppListingRouter {
	switch apiName {
	case FetchAllStageStatusApi:
		json.Unmarshal(response, &structAppListingRouter.fetchAllStageStatusResponseDto)

	case CreateAppApi:
		json.Unmarshal(response, &structAppListingRouter.createAppResponseDto)

	case CreateAppMaterialApi:
		json.Unmarshal(response, &structAppListingRouter.createAppMaterialResponseDto)
	case DeleteAppApi:
		json.Unmarshal(response, &structAppListingRouter.deleteResponseDto)
	case FetchOtherEnvApi:
		json.Unmarshal(response, &structAppListingRouter.fetchOtherEnvResponseDto)
	}

	return structAppListingRouter
}

func FetchAllStageStatus(id int, authToken string) FetchAllStageStatusResponseDto {
	AppId := map[string]string{
		"app-id": strconv.Itoa(id),
	}
	resp, err := Base.MakeApiCall(GetStageStatusApiUrl, http.MethodGet, "", AppId, authToken)
	Base.HandleError(err, FetchAllStageStatusApi)

	structAppListingRouter := StructAppListingRouter{}
	apiRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), FetchAllStageStatusApi)
	return apiRouter.fetchAllStageStatusResponseDto
}

type AppsListingRouterTestSuite struct {
	suite.Suite
	authToken                    string
	createAppResponseDto         CreateAppResponseDto
	createAppMaterialResponseDto CreateAppMaterialResponseDto
}

//func (suite *AppsListingRouterTestSuite) SetupSuite() {
//	suite.authToken = Base.GetAuthToken()
//}

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

func GetAppRequestDto(appName string, teamId int, templateId int) CreateAppRequestDto {
	var createAppRequestDto CreateAppRequestDto
	createAppRequestDto.AppName = appName
	createAppRequestDto.TeamId = teamId
	createAppRequestDto.TemplateId = templateId
	return createAppRequestDto
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

	structAppListingRouter := StructAppListingRouter{}
	pipelineConfigRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), CreateAppApi)
	return pipelineConfigRouter.createAppResponseDto
}

type CreateAppMaterialRequestDto struct {
	AppId     int            `json:"appId"`
	Materials []AppMaterials `json:"material"`
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

	structAppListingRouter := StructAppListingRouter{}
	pipelineConfigRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), CreateAppMaterialApi)
	return pipelineConfigRouter.createAppMaterialResponseDto
}

// SetupSuite This method runs on first priority before starting the suite means before executing any test case of the suite
func (suite *AppsListingRouterTestSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
	suite.createAppResponseDto = suite.CreateApp()
	suite.createAppMaterialResponseDto = suite.CreateAppMaterial()
}

func (suite *AppsListingRouterTestSuite) CreateApp() CreateAppResponseDto {
	appName := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	createAppRequestDto := GetAppRequestDto("app"+appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, "app"+appName, 1, 0, suite.authToken)
	return createAppResponseDto
}

func (suite *AppsListingRouterTestSuite) CreateAppMaterial() CreateAppMaterialResponseDto {
	createAppMaterialRequestDto := GetAppMaterialRequestDto(suite.createAppResponseDto.Result.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponseDto := HitCreateAppMaterialApi(appMaterialByteValue, suite.createAppResponseDto.Result.Id, 1, false, suite.authToken)
	return createAppMaterialResponseDto
}

func (suite *AppsListingRouterTestSuite) TearDownSuite() {
	log.Println("=== Running the after suite method for deleting the data created via automation ===")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(suite.createAppResponseDto.Result.Id, suite.createAppResponseDto.Result.AppName, suite.createAppResponseDto.Result.TeamId, suite.createAppResponseDto.Result.TemplateId)
	HitDeleteAppApi(byteValueOfDeleteApp, suite.createAppResponseDto.Result.Id, suite.authToken)
	//Base.DeleteFile("OutputDataGetChartReferenceViaAppId")
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

	structAppListingRouter := StructAppListingRouter{}
	apiRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteAppApi)
	return apiRouter.deleteResponseDto
}

type DeleteResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type FetchOtherEnvResponseDto struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Result []OtherEnvResult `json:"result"`
	Errors []Base.Errors    `json:"errors"`
}
type OtherEnvResult struct {
	EnvironmentId   int    `json:"environmentId"`
	EnvironmentName string `json:"environmentName"`
	AppMetrics      bool   `json:"appMetrics"`
	InfraMetrics    bool   `json:"infraMetrics"`
	Prod            bool   `json:"prod"`
}

func FetchOtherEnv(id int, authToken string) FetchOtherEnvResponseDto {
	AppId := map[string]string{
		"app-id": strconv.Itoa(id),
	}
	resp, err := Base.MakeApiCall(GetOtherEnvApiUrl, http.MethodGet, "", AppId, authToken)
	Base.HandleError(err, FetchOtherEnvApi)

	structAppListingRouter := StructAppListingRouter{}
	apiRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), FetchOtherEnvApi)
	return apiRouter.fetchOtherEnvResponseDto
}
