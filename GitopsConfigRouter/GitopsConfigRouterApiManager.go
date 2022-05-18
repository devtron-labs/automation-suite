package GitopsConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
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
type StructGitopsConfigRouter struct {
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

func (structGitopsConfigRouter StructGitopsConfigRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructGitopsConfigRouter {
	switch apiName {
	case FetchAllGitopsConfigApi:
		json.Unmarshal(response, &structGitopsConfigRouter.fetchAllGitopsConfigResponseDto)
	case CreateGitopsConfigApi:
		json.Unmarshal(response, &structGitopsConfigRouter.createGitopsConfigResponseDto)
	}
	return structGitopsConfigRouter
}

func HitFetchAllGitopsConfigApi(authToken string) FetchAllGitopsConfigResponseDto {
	resp, err := Base.MakeApiCall(SaveGitopsConfigApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchAllGitopsConfigApi)

	structGitopsConfigRouter := StructGitopsConfigRouter{}
	gitopsConfigRouter := structGitopsConfigRouter.UnmarshalGivenResponseBody(resp.Body(), FetchAllGitopsConfigApi)
	return gitopsConfigRouter.fetchAllGitopsConfigResponseDto
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

	structGitopsConfigRouter := StructGitopsConfigRouter{}
	gitopsConfigRouter := structGitopsConfigRouter.UnmarshalGivenResponseBody(resp.Body(), CreateGitopsConfigApi)
	return gitopsConfigRouter.createGitopsConfigResponseDto
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
