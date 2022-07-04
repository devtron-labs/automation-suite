package ChartRepositoryRouter

import (
	"automation-suite/ChartRepositoryRouter/RequestDTOs"
	"automation-suite/ChartRepositoryRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructChartRepoRouter struct {
	getChartRepoListResponseDto   ResponseDTOs.GetChartRepoListResponseDTO
	createChartRepoResponseDto    ResponseDTOs.CreateChartRepoResponseDTO
	deleteChartRepoResponseDto    ResponseDTOs.DeleteChartRepoResponseDTO
	triggerChartSyncManualRespDto ResponseDTOs.TriggerChartSyncManualRespDTo
}

func HitCreateChartRepoApi(payload []byte, authToken string) ResponseDTOs.CreateChartRepoResponseDTO {
	resp, err := Base.MakeApiCall(CreateChartRepoApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, CreateChartRepo)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), CreateChartRepo)
	return chartRepoRouter.createChartRepoResponseDto
}

func HitUpdateChartRepoApi(payload []byte, authToken string) ResponseDTOs.CreateChartRepoResponseDTO {
	resp, err := Base.MakeApiCall(UpdateChartRepoUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, UpdateChartRepo)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), CreateChartRepo)
	return chartRepoRouter.createChartRepoResponseDto
}

func HitGetChartRepoList(authToken string) ResponseDTOs.GetChartRepoListResponseDTO {
	resp, err := Base.MakeApiCall(GetChartRepoListApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetChartRepoListApi)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), GetChartRepoListApi)
	return chartRepoRouter.getChartRepoListResponseDto
}

func HitGetChartRepoViaId(authToken string, id string) ResponseDTOs.CreateChartRepoResponseDTO {
	resp, err := Base.MakeApiCall(DeleteChartRepoApiUrl+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetChartRepoById)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), CreateChartRepo)
	return chartRepoRouter.createChartRepoResponseDto
}

func HitDeleteChartRepo(payload []byte, authToken string) ResponseDTOs.DeleteChartRepoResponseDTO {
	resp, err := Base.MakeApiCall(DeleteChartRepoApiUrl, http.MethodDelete, string(payload), nil, authToken)
	Base.HandleError(err, DeleteChartRepoApi)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteChartRepoApi)
	return chartRepoRouter.deleteChartRepoResponseDto
}

func HitValidateChartRepo(payload string, authToken string) ResponseDTOs.CreateChartRepoResponseDTO {
	resp, err := Base.MakeApiCall(ValidateChartRepoApiUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, ValidateChartRepoApi)
	structChartRepoRouter := StructChartRepoRouter{}
	chartRepoRouter := structChartRepoRouter.UnmarshalGivenResponseBody(resp.Body(), CreateChartRepo)
	return chartRepoRouter.createChartRepoResponseDto
}

func HitTriggerChartSyncManualApi(authToken string) ResponseDTOs.TriggerChartSyncManualRespDTo {
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

func createChartRepoRequestPayload(authenticateType string, repoId int, RepoName string, repoUrl string, AccessToken string, Active bool) RequestDTOs.ChartRepoRequestDTO {
	chartRepoRequestDto := RequestDTOs.ChartRepoRequestDTO{}
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
