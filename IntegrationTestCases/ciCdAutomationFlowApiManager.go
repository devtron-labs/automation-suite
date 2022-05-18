package IntegrationTestCases

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
	"strconv"
)

type CiCdAutomationFlow struct {
	suite.Suite
	authToken string
}

func (suite *CiCdAutomationFlow) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
	fmt.Println(suite.authToken)
}

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
	saveDockerRegistryResponseDto   SaveDockerRegistryResponseDto
	deleteDockerRegistryResponse    DeleteDockerRegistryResponse
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
	StageErrorMap    struct {
		ErrorInConnectingWithGITHUB string `json:"error in connecting with GITHUB"`
	} `json:"stageErrorMap"`
	DeleteRepoFailed bool `json:"deleteRepoFailed"`
}

type CreateTeamRequestDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
type CreateTeamResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	Result struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"result"`
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
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id         int    `json:"id"`
		AppName    string `json:"appName"`
		TeamId     int    `json:"teamId"`
		TemplateId int    `json:"templateId"`
		Material   []struct {
			Name            string `json:"name"`
			Url             string `json:"url"`
			Id              int    `json:"id"`
			GitProviderId   int    `json:"gitProviderId"`
			CheckoutPath    string `json:"checkoutPath"`
			FetchSubmodules bool   `json:"fetchSubmodules"`
		} `json:"material"`
	} `json:"result"`
}

type FetchOtherEnvResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		EnvironmentId   int    `json:"environmentId"`
		EnvironmentName string `json:"environmentName"`
		AppMetrics      bool   `json:"appMetrics"`
		InfraMetrics    bool   `json:"infraMetrics"`
		Prod            bool   `json:"prod"`
	} `json:"result"`
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

func (installationScriptStruct InstallationScriptStruct) UnmarshalGivenResponseBody(response []byte, apiName string) InstallationScriptStruct {
	switch apiName {

	case DeleteAppApi:
		json.Unmarshal(response, &installationScriptStruct.deleteResponseDto)
	case FetchAllAutocompleteApi:
		json.Unmarshal(response, &installationScriptStruct.getAutocompleteResponseDto)
	case DeleteDockerRegistry:
		json.Unmarshal(response, &installationScriptStruct.deleteDockerRegistryResponse)
	case SaveDockerRegistryApi:
		json.Unmarshal(response, &installationScriptStruct.saveDockerRegistryResponseDto)
	case FetchAllGitopsConfigApi:
		json.Unmarshal(response, &installationScriptStruct.fetchAllGitopsConfigResponseDto)
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
	case DeleteAppMaterialApi:
		json.Unmarshal(response, &installationScriptStruct.deleteResponseDto)
	}
	return installationScriptStruct
}

func HitFetchAllGitopsConfigApi(authToken string) FetchAllGitopsConfigResponseDto {
	resp, err := Base.MakeApiCall(SaveGitopsConfigApiUrl, http.MethodGet, "", nil, authToken)
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
	Base.HandleError(err, CreateGitopsConfigApi)

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), CreateGitopsConfigApi)
	return installationScriptRouter.createGitopsConfigResponseDto
}

type SaveTeamRequestDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type DeleteTeamResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}

func GetSaveTeamRequestDto() SaveTeamRequestDto {
	var saveTeamRequestDto SaveTeamRequestDto
	teamName := Base.GetRandomStringOfGivenLength(10)
	saveTeamRequestDto.Name = teamName
	saveTeamRequestDto.Active = true
	return saveTeamRequestDto
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
	Base.HandleError(err, CreateTeamApi)

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), CreateTeamApi)
	return installationScriptRouter.createTeamResponseDto
}
func GetPayLoadForDeleteTeamAPI(id int, name string, active bool) []byte {
	var createTeamRequestDto CreateTeamRequestDto
	createTeamRequestDto.Id = id
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
	resp, err := Base.MakeApiCall(SaveAppApiUrl+"/"+strconv.Itoa(id), http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, DeleteAppApi)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteAppApi)
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
	Base.HandleError(err, CreateAppMaterialApi)

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), CreateAppMaterialApi)
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
	Base.HandleError(err, DeleteAppMaterialApi)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteAppMaterialApi)
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
	Host        string `env:"HOST" envDefault:"""`
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

type DockerRegistry struct {
	Id           string `env:"ID" envDefault:""`
	PluginId     string `env:"PLUGINID" envDefault:""`
	RegistryType string `env:"REGISTRYTYPE" envDefault:""`
	RegistryUrl  string `env:"REGISTRYURL" envDefault:""`
	Username     string `env:"USERNAME" envDefault:""`
	Password     string `env:"PASSWORD" envDefault:""`
}

func GetDockerRegistry() (*DockerRegistry, error) {
	cfg := &DockerRegistry{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from ChartRepoRouterConfig")
	}
	return cfg, err
}

type SaveDockerRegistryRequestDto struct {
	Id           string `json:"id"`
	PluginId     string `json:"pluginId"`
	RegistryType string `json:"registryType"`
	IsDefault    bool   `json:"isDefault"`
	RegistryUrl  string `json:"registryUrl"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func GetDockerRegistryRequestDto(isRepeat bool, id string, pluginId string, regType string, regUrl string, isDefault bool, username string, password string) SaveDockerRegistryRequestDto {
	if isRepeat == false {
		dockerRegistry, _ := GetDockerRegistry()
		var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
		saveDockerRegistryRequestDto.Id = Base.GetRandomStringOfGivenLength(10)
		saveDockerRegistryRequestDto.PluginId = dockerRegistry.PluginId
		saveDockerRegistryRequestDto.RegistryType = dockerRegistry.RegistryType
		saveDockerRegistryRequestDto.RegistryUrl = dockerRegistry.RegistryUrl
		saveDockerRegistryRequestDto.IsDefault = false
		saveDockerRegistryRequestDto.Username = dockerRegistry.Username
		saveDockerRegistryRequestDto.Password = dockerRegistry.Password
		return saveDockerRegistryRequestDto
	}

	var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
	saveDockerRegistryRequestDto.Id = id
	saveDockerRegistryRequestDto.PluginId = pluginId
	saveDockerRegistryRequestDto.RegistryType = regType
	saveDockerRegistryRequestDto.RegistryUrl = regUrl
	saveDockerRegistryRequestDto.IsDefault = isDefault
	saveDockerRegistryRequestDto.Username = username
	saveDockerRegistryRequestDto.Password = password
	return saveDockerRegistryRequestDto
}

type SaveDockerRegistryResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
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
	} `json:"result"`

	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
}

func HitSaveDockerRegistryApi(isRepeat bool, payload []byte, id string, pluginId string, regUrl string, regType string, username string, password string, isDefault bool, authToken string) SaveDockerRegistryResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		if isRepeat == false {
			dockerRegistry, _ := GetDockerRegistry()
			var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
			saveDockerRegistryRequestDto.Id = dockerRegistry.Id
			saveDockerRegistryRequestDto.PluginId = dockerRegistry.PluginId
			saveDockerRegistryRequestDto.RegistryType = dockerRegistry.RegistryType
			saveDockerRegistryRequestDto.RegistryUrl = dockerRegistry.RegistryUrl
			saveDockerRegistryRequestDto.IsDefault = false
			saveDockerRegistryRequestDto.Username = dockerRegistry.Username
			saveDockerRegistryRequestDto.Password = dockerRegistry.Password
			byteValueOfStruct, _ := json.Marshal(saveDockerRegistryRequestDto)
			payloadOfApi = string(byteValueOfStruct)
		} else {
			var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
			saveDockerRegistryRequestDto.Id = id
			saveDockerRegistryRequestDto.PluginId = pluginId
			saveDockerRegistryRequestDto.RegistryType = regType
			saveDockerRegistryRequestDto.RegistryUrl = regUrl
			saveDockerRegistryRequestDto.IsDefault = isDefault
			saveDockerRegistryRequestDto.Username = username
			saveDockerRegistryRequestDto.Password = password
			byteValueOfStruct, _ := json.Marshal(saveDockerRegistryRequestDto)
			payloadOfApi = string(byteValueOfStruct)
		}
	}

	resp, err := Base.MakeApiCall(SaveDockerRegistryApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, SaveDockerRegistryApi)

	installationScriptStruct := InstallationScriptStruct{}
	installationScriptRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), SaveDockerRegistryApi)
	return installationScriptRouter.saveDockerRegistryResponseDto
}
func GetPayLoadForDeleteDockerRegistryAPI(id string, pluginId string, regUrl string, regType string, username string, password string, isDefault bool) []byte {
	var saveDockerRegistryRequestDto SaveDockerRegistryRequestDto
	saveDockerRegistryRequestDto.Id = id
	saveDockerRegistryRequestDto.PluginId = pluginId
	saveDockerRegistryRequestDto.RegistryUrl = regUrl
	saveDockerRegistryRequestDto.RegistryType = regType
	saveDockerRegistryRequestDto.Username = username
	saveDockerRegistryRequestDto.Password = password
	saveDockerRegistryRequestDto.IsDefault = isDefault
	byteValueOfStruct, _ := json.Marshal(saveDockerRegistryRequestDto)
	return byteValueOfStruct
}

type DeleteDockerRegistryResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}

func HitDeleteDockerRegistryApi(byteValueOfStruct []byte, authToken string) DeleteDockerRegistryResponse {
	resp, err := Base.MakeApiCall(SaveDockerRegistryApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, DeleteDockerRegistry)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteDockerRegistry)
	return apiRouter.deleteDockerRegistryResponse
}
