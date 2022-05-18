package GitopsConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type GitopsConfigRouter struct {
	suite.Suite
	authToken string
}

func (suite *GitopsConfigRouter) SetupSuite() {
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
	createGitopsConfigResponseDto   CreateGitopsConfigResponseDto
	fetchAllGitopsConfigResponseDto FetchAllGitopsConfigResponseDto
	deleteResponseDto               DeleteResponseDto
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
	Result struct {
		SuccessfulStages []string "successfulStages"
		StageErrorMap    struct {
			ErrorInConnectingWithGITHUB string `json:"error in connecting with GITHUB"`
		} `json:"stageErrorMap"`
		DeleteRepoFailed bool `json:"deleteRepoFailed"`
	} `json:"result"`
}

func (installationScriptStruct InstallationScriptStruct) UnmarshalGivenResponseBody(response []byte, apiName string) InstallationScriptStruct {
	switch apiName {
	case FetchAllGitopsConfigApi:
		json.Unmarshal(response, &installationScriptStruct.fetchAllGitopsConfigResponseDto)
	case CreateGitopsConfigApi:
		json.Unmarshal(response, &installationScriptStruct.createGitopsConfigResponseDto)
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

type GitopsConfig struct {
	Provider    string `env:"PROVIDER" envDefault:""`
	Username    string `env:"USERNAME" envDefault:""`
	Host        string `env:"HOST" envDefault:""`
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
