package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type ChartRepoRequestDto struct {
	Id           int      `json:"id,omitempty" validate:"number"`
	Name         string   `json:"name,omitempty" validate:"required"`
	Url          string   `json:"url,omitempty"`
	UserName     string   `json:"userName,omitempty"`
	Password     string   `json:"password,omitempty"`
	SshKey       string   `json:"sshKey,omitempty"`
	AccessToken  string   `json:"accessToken,omitempty"`
	AuthMode     AuthMode `json:"authMode,omitempty" validate:"required"`
	Active       bool     `json:"active"`
	Default      bool     `json:"default"`
	UserId       int32    `json:"-"`
	CustomErrMsg string   `json:"customErrMsg"`
	ActualErrMsg string   `json:"actualErrMsg"`
}

type CreateChartRepoResponseDto struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Result ChartRepoRequestDto `json:"result"`
	Errors Errors              `json:"errors"`
}

type GetChartRepoListResponseDto struct {
	Code   int                   `json:"code"`
	Status string                `json:"status"`
	Result []ChartRepoRequestDto `json:"result"`
}

type DeleteChartRepoResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
	Errors Errors `json:"errors"`
}
type Errors []struct {
	Code            string `json:"code"`
	InternalMessage string `json:"internalMessage"`
	UserMessage     string `json:"userMessage"`
}

type TriggerChartSyncManualRespDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Status string `json:"status"`
	} `json:"result"`
}

type StructChartRepoRouter struct {
	getChartRepoListResponseDto   GetChartRepoListResponseDto
	createChartRepoResponseDto    CreateChartRepoResponseDto
	deleteChartRepoResponseDto    DeleteChartRepoResponseDto
	triggerChartSyncManualRespDto TriggerChartSyncManualRespDto
}

func HitCreateChartRepoApi(payload string, authToken string) CreateChartRepoResponseDto {
	resp, err := Base.MakeApiCall(CreateChartRepoApiUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, CreateChartRepo)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), CreateChartRepo)
	return chartRepoRouter.createChartRepoResponseDto
}

func HitUpdateChartRepoApi(payload string, authToken string) CreateChartRepoResponseDto {
	resp, err := Base.MakeApiCall(UpdateChartRepoUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, UpdateChartRepo)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), CreateChartRepo)
	return chartRepoRouter.createChartRepoResponseDto
}

func HitGetChartRepoList(authToken string) GetChartRepoListResponseDto {
	resp, err := Base.MakeApiCall(GetChartRepoListApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetChartRepoListApi)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), GetChartRepoListApi)
	return chartRepoRouter.getChartRepoListResponseDto
}

func HitGetChartRepoViaId(authToken string, id string) CreateChartRepoResponseDto {
	resp, err := Base.MakeApiCall(DeleteChartRepoApiUrl+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetChartRepoById)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), CreateChartRepo)
	return chartRepoRouter.createChartRepoResponseDto
}

func HitDeleteChartRepo(payload string, authToken string) DeleteChartRepoResponseDto {
	resp, err := Base.MakeApiCall(DeleteChartRepoApiUrl, http.MethodDelete, payload, nil, authToken)
	Base.HandleError(err, DeleteChartRepoApi)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteChartRepoApi)
	return chartRepoRouter.deleteChartRepoResponseDto
}

func HitValidateChartRepo(payload string, authToken string) CreateChartRepoResponseDto {
	resp, err := Base.MakeApiCall(ValidateChartRepoApiUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, ValidateChartRepoApi)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), CreateChartRepo)
	return chartRepoRouter.createChartRepoResponseDto
}

func HitTriggerChartSyncManualApi(authToken string) TriggerChartSyncManualRespDto {
	resp, err := Base.MakeApiCall(TriggerChartSyncManualApiUrl, http.MethodPost, "", nil, authToken)
	Base.HandleError(err, TriggerChartSyncManualApi)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), TriggerChartSyncManualApi)
	return chartRepoRouter.triggerChartSyncManualRespDto
}

func (structChartRepoRouter StructChartRepoRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructChartRepoRouter {
	switch apiName {
	case GetChartRepoListApi:
		json.Unmarshal(response, &structChartRepoRouter.getChartRepoListResponseDto)
	case CreateChartRepo:
		json.Unmarshal(response, &structChartRepoRouter.createChartRepoResponseDto)
	case DeleteChartRepoApi:
		json.Unmarshal(response, &structChartRepoRouter.deleteChartRepoResponseDto)
	case TriggerChartSyncManualApi:
		json.Unmarshal(response, &structChartRepoRouter.triggerChartSyncManualRespDto)
	}
	return structChartRepoRouter
}

func createChartRepoRequestPayload(authenticateType AuthMode, repoId int, RepoName string, repoUrl string, AccessToken string, Active bool) ChartRepoRequestDto {
	chartRepoRequestDto := ChartRepoRequestDto{}
	switch authenticateType {

	case AUTH_MODE_ANONYMOUS:
		var repositoryId int
		var repositoryName string
		if repoId != 0 {
			repositoryId = repoId
		}
		if RepoName != "" {
			repositoryName = RepoName
		}
		chartRepoRequestDto.Id = repositoryId
		chartRepoRequestDto.Name = repositoryName
		chartRepoRequestDto.Url = repoUrl
		chartRepoRequestDto.Active = Active
		chartRepoRequestDto.AuthMode = AUTH_MODE_ANONYMOUS

	case AUTH_MODE_ACCESS_TOKEN:
		var repositoryId int
		var repositoryName string
		if repoId != 0 {
			repositoryId = repoId
		}
		if RepoName != "" {
			repositoryName = RepoName
		}
		if AccessToken != "" {
			chartRepoRequestDto.AccessToken = AccessToken
		}
		chartRepoRequestDto.Id = repositoryId
		chartRepoRequestDto.Name = repositoryName
		chartRepoRequestDto.Url = repoUrl
		chartRepoRequestDto.Active = Active
		chartRepoRequestDto.AuthMode = AUTH_MODE_ACCESS_TOKEN

	default:
		chartRepoRequestDto.Id = repoId
		chartRepoRequestDto.Name = RepoName
		chartRepoRequestDto.Url = repoUrl
		chartRepoRequestDto.Active = Active
		chartRepoRequestDto.AuthMode = ""
	}
	return chartRepoRequestDto
}

type ChartRepoRouterConfig struct {
	ChartRepoUrl     string `env:"CHART_REPO_URL" envDefault:"https://deepak-devtron.github.io/helm-chart/"`
	ChartAccessToken string `env:"CHART_ACCESS_TOKEN" envDefault:""`
}

func GetChartRepoRouterConfig() (*ChartRepoRouterConfig, error) {
	cfg := &ChartRepoRouterConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from ChartRepoRouterConfig")
	}
	return cfg, err
}

type ChartRepoTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *ChartRepoTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
