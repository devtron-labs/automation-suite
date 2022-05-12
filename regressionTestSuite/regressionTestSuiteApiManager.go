package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type FetchAllGitopsConfigResponseDto struct {
	Code   int                            `json:"code"`
	Status string                         `json:"status"`
	Result []CreateGitopsConfigRequestDto `json:"result"`
}

type DeleteResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}
type InstallationScriptStruct struct {
	createAppResponseDto            CreateAppResponseDto
	createGitopsConfigResponseDto   CreateGitopsConfigResponseDto
	createTeamResponseDto           CreateTeamResponseDto
	fetchAllGitopsConfigResponseDto FetchAllGitopsConfigResponseDto
	deleteResponseDto               DeleteResponseDto
	getAutocompleteResponseDto      GetAutocompleteResponseDto
	fetchAllStageStatusResponseDto  FetchAllStageStatusResponseDto
	fetchOtherEnvResponseDto        FetchOtherEnvResponseDto
	fetchAllAppWorkflowResponseDto  FetchAllAppWorkflowResponseDto
	fetchAppGetResponseDto          FetchAppGetResponseDto
	createAppMaterialResponseDto    CreateAppMaterialResponseDto
}
type CreateGitopsConfigRequestDto struct {
	Id                   int    `json:"id"`
	Provider             string `json:"provider"`
	Username             string `json:"username"`
	Token                string `json:"token"`
	GitLabGroupId        string `json:"gitLabGroupId"`
	GitHubOrgId          string `json:"gitHubOrgId"`
	Host                 string `json:"host"`
	Active               bool   `json:"active"`
	AzureProjectName     string `json:"azureProjectName"`
	BitBucketWorkspaceId string `json:"bitBucketWorkspaceId"`
	BitBucketProjectKey  string `json:"bitBucketProjectKey"`
}

type CreateGitopsConfigResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result Result `json:"result"`
}

type Result struct {
	SuccessfulStages []string "successfulStages"
	DeleteRepoFailed bool     `json:"deleteRepoFailed"`
}

type CreateTeamRequestDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
type CreateTeamResponseDto struct {
	Code   int                  `json:"code"`
	Status string               `json:"status"`
	Result CreateTeamRequestDto `json:"result"`
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
	Errors []Error             `json:"errors"`
}
type Error struct {
	InternalMessage string `json:"internalMessage"`
	UserMessage     string `json:"userMessage"`
}
type FetchAllAppWorkflowResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []Apps `json:"result"`
}
type Apps struct {
	AppId     int        `json:"appId"`
	AppName   string     `json:"appName"`
	Workflows []Workflow `json:"workflows"`
}
type Workflow struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	AppId int    `json:"appId"`
	Tree  []Tree `json:"tree"`
}
type Tree struct {
	Id            int    `json:"id"`
	AppWorkflowId int    `json:"appWorkflowId"`
	Type          string `json:"type"`
	ComponentId   int    `json:"componentId"`
	ParentId      int    `json:"parentId"`
	ParentType    string `json:"parentType"`
}
type StageStatus struct {
	Stage     int    `json:"stage"`
	StageName string `json:"stageName"`
	Status    bool   `json:"status"`
	Required  bool   `json:"required"`
}
type FetchAllStageStatusResponseDto struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result []StageStatus `json:"result"`
}
type FetchAppGetResponseDto struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Result []AppGet `json:"result"`
}
type AppGet struct {
	Id         int         `json:"id"`
	AppName    string      `json:"appName"`
	TeamId     int         `json:"teamId"`
	TemplateId int         `json:"templateId"`
	Material   []Materials `json:material`
}
type Materials struct {
	Name            string `json:"name"`
	Url             string `json:"url"`
	Id              int    `json:"id"`
	GitProviderId   int    `json:"gitProviderId"`
	CheckoutPath    string `json:"checkoutPath"`
	FetchSubmodules bool   `json:"fetchSubmodules"`
}
type FetchOtherEnvResponseDto struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Result []OtherEnvResult `json:"result"`
}
type OtherEnvResult struct {
	EnvironmentId   int    `json:"environmentId"`
	EnvironmentName string `json:"environmentName"`
	AppMetrics      bool   `json:"appMetrics"`
	InfraMetrics    bool   `json:"infraMetrics"`
	Prod            bool   `json:"prod"`
}

type GetAutocompleteResponseDto struct {
	Code   int                    `json:"code"`
	Status string                 `json:"status"`
	Result []CreateTeamRequestDto `json:"result"`
}
type CreateAppMaterialRequestDto struct {
	AppId     int            `json:"app_id"`
	Materials []AppMaterials `json:"material"`
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
type AppMaterials struct {
	Url             string `json:"url"`
	Id              int    `json:"id"`
	GitProviderId   int    `json:"gitProviderId"`
	CheckoutPath    string `json:"checkoutPath"`
	FetchSubmodules bool   `json:"fetchSubmodules"`
}

type regressionTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *regressionTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
	fmt.Println(suite.authToken)
}

func (installationScriptStruct InstallationScriptStruct) UnmarshalGivenResponseBody(response []byte, apiName string) InstallationScriptStruct {
	switch apiName {

	case CreateAppApi:
		json.Unmarshal(response, &installationScriptStruct.createAppResponseDto)
	case CreateGitopsConfigApi:
		json.Unmarshal(response, &installationScriptStruct.createGitopsConfigResponseDto)
	case CreateTeamApi:
		json.Unmarshal(response, &installationScriptStruct.createTeamResponseDto)
	case CreateAppMaterialApi:
		json.Unmarshal(response, &installationScriptStruct.createAppMaterialResponseDto)
	case FetchAllStageStatusApi:
		json.Unmarshal(response, &installationScriptStruct.fetchAllStageStatusResponseDto)
	}
	return installationScriptStruct
}

func HitFetchAllGitopsConfigApi() FetchAllGitopsConfigResponseDto {
	resp, err := Base.MakeApiCall(SaveGitopsConfigApiUrl, http.MethodGet, "", nil, "")
	Base.HandleError(err, FetchAllGitopsConfigApi)

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllGitopsConfigApi)
	return installationScriptRouter.fetchAllGitopsConfigResponseDto
}

func GetPayLoadForDeleteGitopsConfigAPI(id int, provider string, username string, host string, token string) []byte {
	var deleteGitopsConfigRequestDto CreateGitopsConfigRequestDto
	deleteGitopsConfigRequestDto.Id = id
	deleteGitopsConfigRequestDto.Provider = provider
	deleteGitopsConfigRequestDto.Username = username
	deleteGitopsConfigRequestDto.Host = host
	deleteGitopsConfigRequestDto.Token = token
	byteValueOfStruct, _ := json.Marshal(deleteGitopsConfigRequestDto)
	return byteValueOfStruct
}

func HitDeleteLinkApi(byteValueOfStruct []byte, authToken string) DeleteResponseDto {
	resp, err := Base.MakeApiCall(SaveGitopsConfigApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, "DeleteLinkApi")

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "DeleteLink")
	return apiRouter.deleteResponseDto
}

func GetGitopsConfigRequestDto(provider string, username string, host string, token string, githuborgid string) CreateGitopsConfigRequestDto {
	var createGitopsConfigRequestDto CreateGitopsConfigRequestDto
	createGitopsConfigRequestDto.Provider = provider
	createGitopsConfigRequestDto.Username = username
	createGitopsConfigRequestDto.Host = host
	createGitopsConfigRequestDto.Token = token
	createGitopsConfigRequestDto.GitHubOrgId = githuborgid
	createGitopsConfigRequestDto.Active = true
	return createGitopsConfigRequestDto
}
func HitCreateGitopsConfigApi(payload []byte, provider string, username string, host string, token string, githuborgid string, authToken string) CreateGitopsConfigResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		var createGitopsConfigRequestDto CreateGitopsConfigRequestDto
		createGitopsConfigRequestDto.Provider = provider
		createGitopsConfigRequestDto.Username = username
		createGitopsConfigRequestDto.Host = host
		createGitopsConfigRequestDto.Token = token
		createGitopsConfigRequestDto.GitHubOrgId = githuborgid
		createGitopsConfigRequestDto.Active = true
		byteValueOfStruct, _ := json.Marshal(createGitopsConfigRequestDto)
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := Base.MakeApiCall(SaveGitopsConfigApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, "CreateGitopsConfigApi")

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "CreateGitopsConfig")
	return installationScriptRouter.createGitopsConfigResponseDto
}

func GetTeamRequestDto(name string, active bool) CreateTeamRequestDto {
	var createTeamRequestDto CreateTeamRequestDto
	createTeamRequestDto.Name = name
	createTeamRequestDto.Active = active
	return createTeamRequestDto
}
func HitCreateTeamApi(payload []byte, name string, active bool, authToken string) CreateTeamResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		var createTeamRequestDto CreateTeamRequestDto
		createTeamRequestDto.Name = name
		createTeamRequestDto.Active = active
		byteValueOfStruct, _ := json.Marshal(createTeamRequestDto)
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, "CreateTeamApi")

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "CreateTeam")
	return installationScriptRouter.createTeamResponseDto
}
func GetPayLoadForDeleteTeamAPI(name string, active bool) []byte {
	var createTeamRequestDto CreateTeamRequestDto
	createTeamRequestDto.Name = name
	createTeamRequestDto.Active = active
	byteValueOfStruct, _ := json.Marshal(createTeamRequestDto)
	return byteValueOfStruct
}

func HitDeleteTeamApi(byteValueOfStruct []byte, authToken string) DeleteResponseDto {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, "DeleteLinkApi")

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "DeleteLink")
	return apiRouter.deleteResponseDto
}

func HitFetchAllTeamApi(authToken string) GetAutocompleteResponseDto {
	resp, err := Base.MakeApiCall(GetAutocompleteApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchAllAutocompleteApi)

	installationScriptStruct := InstallationScriptStruct{}
	fetchAllRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllAutocompleteApi)
	return fetchAllRouter.getAutocompleteResponseDto
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

	resp, err := Base.MakeApiCall(SaveAppApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, CreateAppApi)

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), CreateAppApi)
	return installationScriptRouter.createAppResponseDto
}
func GetPayLoadForDeleteAppAPI(appName string, teamId int, templateId int) []byte {
	var createAppRequestDto CreateAppRequestDto
	createAppRequestDto.AppName = appName
	createAppRequestDto.TeamId = teamId
	createAppRequestDto.TemplateId = templateId
	byteValueOfStruct, _ := json.Marshal(createAppRequestDto)
	return byteValueOfStruct
}
func HitDeleteAppApi(byteValueOfStruct []byte, authToken string) DeleteResponseDto {
	resp, err := Base.MakeApiCall(SaveAppApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, "DeleteAppApi")

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "DeleteApp")
	return apiRouter.deleteResponseDto
}

func FetchAllStageStatus(id map[string]string, authToken string) FetchAllStageStatusResponseDto {
	resp, err := Base.MakeApiCall(GetStageStatusApiUrl, http.MethodGet, "", id, authToken)
	Base.HandleError(err, FetchAllStageStatusApi)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "FetchAllStageStatus")
	return apiRouter.fetchAllStageStatusResponseDto
}

func FetchAllAppWorkflow(id map[string]string, authToken string) FetchAllAppWorkflowResponseDto {
	resp, err := Base.MakeApiCall(GetAppWorkflowApiUrl, http.MethodGet, "", id, authToken)
	Base.HandleError(err, FetchAllAppWorkflowApi)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "FetchAllAppWorkflow")
	return apiRouter.fetchAllAppWorkflowResponseDto
}

func FetchAppGet(id map[string]string, authToken string) FetchAppGetResponseDto {
	resp, err := Base.MakeApiCall(GetAppGetApiUrl, http.MethodGet, "", id, authToken)
	Base.HandleError(err, FetchAppGetApi)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "FetchAppGet")
	return apiRouter.fetchAppGetResponseDto
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

	resp, err := Base.MakeApiCall(SaveAppMaterialApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, "CreateAppMaterialApi")

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "CreateAppMaterial")
	return installationScriptRouter.createAppMaterialResponseDto
}
func GetPayLoadForDeleteAppMaterialAPI(appId int, slice2 AppMaterials) []byte {
	var createAppMaterialRequestDto CreateAppMaterialRequestDto
	var slice AppMaterials
	slice.Url = slice2.Url
	slice.GitProviderId = slice2.GitProviderId
	slice.FetchSubmodules = slice2.FetchSubmodules
	createAppMaterialRequestDto.AppId = appId
	createAppMaterialRequestDto.Materials = append(createAppMaterialRequestDto.Materials, slice)
	byteValueOfStruct, _ := json.Marshal(createAppMaterialRequestDto)
	return byteValueOfStruct
}
func HitDeleteAppMaterialApi(byteValueOfStruct []byte, authToken string) DeleteResponseDto {
	resp, err := Base.MakeApiCall(SaveAppMaterialApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, "DeleteAppMaterialApi")

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "DeleteAppMaterial")
	return apiRouter.deleteResponseDto
}

func FetchOtherEnv(id map[string]string, authToken string) FetchOtherEnvResponseDto {
	resp, err := Base.MakeApiCall(GetOtherEnvApiUrl, http.MethodGet, "", id, authToken)
	Base.HandleError(err, FetchOtherEnvApi)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "FetchOtherEnv")
	return apiRouter.fetchOtherEnvResponseDto
}

type GitopsConfig struct {
	Provider    string `env:"PROVIDER" envDefault:""`
	Username    string `env:"USERNAME" envDefault:""`
	Host        string `env:"HOST" envDefault:"https://github.com/"`
	Token       string `env:"TOKEN" envDefault:""`
	GitHubOrgId string `env:"GITHUB_ORG_ID" envDefault:""`
	Url         string `env:"URL" envDefault:""`
}

func GetGitopsConfig() (*GitopsConfig, error) {
	cfg := &GitopsConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from ChartRepoRouterConfig")
	}
	return cfg, err
}
