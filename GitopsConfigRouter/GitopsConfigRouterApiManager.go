package GitopsConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"log"
	"net/http"

	"automation-suite/GitopsConfigRouter/RequestDTOs"
	"automation-suite/GitopsConfigRouter/ResponseDTOs"

	"github.com/stretchr/testify/suite"
)

type StructGitopsConfigRouter struct {
	createGitopsConfigResponseDto   ResponseDTOs.CreateGitopsConfigResponseDto
	fetchAllGitopsConfigResponseDto ResponseDTOs.FetchAllGitopsConfigResponseDto
	checkGitopsExistsResponse       ResponseDTOs.CheckGitopsExistsResponse
	updateGitopsConfigResponseDto   ResponseDTOs.UpdateGitopsConfigResponseDto
}

func HitGitopsConfigured(authToken string) ResponseDTOs.CheckGitopsExistsResponse {
	resp, err := Base.MakeApiCall(CheckGitopsConfigExistsApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, CheckGitopsConfigExistsApi)

	structGitopsConfigRouter := StructGitopsConfigRouter{}
	gitopsConfigRouter := structGitopsConfigRouter.UnmarshalGivenResponseBody(resp.Body(), CheckGitopsConfigExistsApi)
	return gitopsConfigRouter.checkGitopsExistsResponse
}

func (structGitopsConfigRouter StructGitopsConfigRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructGitopsConfigRouter {
	switch apiName {
	case FetchAllGitopsConfigApi:
		json.Unmarshal(response, &structGitopsConfigRouter.fetchAllGitopsConfigResponseDto)
	case CreateGitopsConfigApi:
		json.Unmarshal(response, &structGitopsConfigRouter.createGitopsConfigResponseDto)
	case CheckGitopsConfigExistsApi:
		json.Unmarshal(response, &structGitopsConfigRouter.checkGitopsExistsResponse)
	case UpdateGitopsConfigApi:
		json.Unmarshal(response, &structGitopsConfigRouter.updateGitopsConfigResponseDto)

	}
	return structGitopsConfigRouter
}

func HitFetchAllGitopsConfigApi(authToken string) ResponseDTOs.FetchAllGitopsConfigResponseDto {
	resp, err := Base.MakeApiCall(SaveGitopsConfigApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchAllGitopsConfigApi)

	structGitopsConfigRouter := StructGitopsConfigRouter{}
	gitopsConfigRouter := structGitopsConfigRouter.UnmarshalGivenResponseBody(resp.Body(), FetchAllGitopsConfigApi)
	return gitopsConfigRouter.fetchAllGitopsConfigResponseDto
}

func GetGitopsConfigRequestDto(provider string, username string, host string, token string, githuborgid string) RequestDTOs.CreateGitopsConfigRequestDto {
	var createGitopsConfigRequestDto RequestDTOs.CreateGitopsConfigRequestDto
	createGitopsConfigRequestDto.Provider = provider
	createGitopsConfigRequestDto.Username = username
	createGitopsConfigRequestDto.Host = host
	createGitopsConfigRequestDto.Token = token
	createGitopsConfigRequestDto.GitHubOrgId = githuborgid
	createGitopsConfigRequestDto.Active = true
	return createGitopsConfigRequestDto
}
func HitCreateGitopsConfigApi(payload []byte, provider string, username string, host string, token string, githuborgid string, authToken string) ResponseDTOs.CreateGitopsConfigResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		var createGitopsConfigRequestDto RequestDTOs.CreateGitopsConfigRequestDto
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

func UpdateGitops(authToken string) RequestDTOs.CreateGitopsConfigRequestDto {
	var createGitopsConfigRequestDto RequestDTOs.CreateGitopsConfigRequestDto
	fetchAllLinkResponseDto := HitFetchAllGitopsConfigApi(authToken)

	log.Println("Checking which is true")
	for _, createGitopsConfigRequestDto = range fetchAllLinkResponseDto.Result {
		if createGitopsConfigRequestDto.Active {
			createGitopsConfigRequestDto.Active = false
			byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)
			log.Println("Updating gitops to false")
			HitUpdateGitopsConfigApi(byteValueOfCreateGitopsConfig, authToken)
			createGitopsConfigRequestDto.Active = true
			return createGitopsConfigRequestDto
		}
	}
	return createGitopsConfigRequestDto
}

func HitUpdateGitopsConfigApi(payload []byte, authToken string) ResponseDTOs.UpdateGitopsConfigResponseDto {
	resp, err := Base.MakeApiCall(SaveGitopsConfigApiUrl, http.MethodPut, string(payload), nil, authToken)
	Base.HandleError(err, UpdateGitopsConfigApi)

	structGitopsConfigRouter := StructGitopsConfigRouter{}
	gitopsConfigRouter := structGitopsConfigRouter.UnmarshalGivenResponseBody(resp.Body(), UpdateGitopsConfigApi)
	return gitopsConfigRouter.updateGitopsConfigResponseDto
}

/*
type GitopsConfig struct {
	Provider    string `env:"PROVIDER" envDefault:"GITHUB"`
	Username    string `env:"GIT_USERNAME" envDefault:"deepak-devtron"`
	Host        string `env:"HOST" envDefault:"https://github.com/"`
	Token       string `env:"GIT_TOKEN" envDefault:"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"`
	GitHubOrgId string `env:"GITHUB_ORG_ID" envDefault:"Deepak-Deepak-Org"`
	Url         string `env:"URL" envDefault:""`
}*/

/*func GetGitopsConfig() (*GitopsConfig, error) {
	cfg := &GitopsConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from ChartRepoRouterConfig")
	}
	return cfg, err
}*/

type GitOpsRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *GitOpsRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
