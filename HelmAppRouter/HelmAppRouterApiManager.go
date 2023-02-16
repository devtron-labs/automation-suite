package HelmAppRouter

import (
	"automation-suite/HelmAppRouter/RequestDTOs"
	"automation-suite/HelmAppRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructHelmAppRouter struct {
	deploymentHistoryResponseDto      ResponseDTOs.DeploymentHistoryResponseDTO
	rollbackApplicationApiResponseDto ResponseDTOs.RollbackApplicationApiResponseDTO
	applicationUpdateRequestDto       RequestDTOs.ApplicationUpdateRequestDTO
	releaseInfoApiResponseDto         ResponseDTOs.ReleaseInfoApiResponseDTO
	hibernateApiResponseDto           ResponseDTOs.HibernateApiResponseDTO
	getApplicationDetailResponseDto   ResponseDTOs.GetApplicationDetailResponseDTO
}

func HitGetDeploymentHistoryById(queryParams map[string]string, authToken string) ResponseDTOs.DeploymentHistoryResponseDTO {
	resp, err := Base.MakeApiCall(GetDeploymentHistoryApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetDeploymentHistory)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), GetDeploymentHistory)
	return helmAppRouter.deploymentHistoryResponseDto
}

func HitRollbackApplicationApi(payload string, authToken string) ResponseDTOs.RollbackApplicationApiResponseDTO {
	resp, err := Base.MakeApiCall(RollbackApplicationApiUrl, http.MethodPut, payload, nil, authToken)
	Base.HandleError(err, RollbackApplication)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), RollbackApplication)
	return helmAppRouter.rollbackApplicationApiResponseDto
}

func HitGetReleaseInfoApi(queryParams map[string]string, authToken string) ResponseDTOs.ReleaseInfoApiResponseDTO {
	resp, err := Base.MakeApiCall(GetReleaseInfoApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetReleaseInfoApi)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), GetReleaseInfoApi)
	return helmAppRouter.releaseInfoApiResponseDto
}

func HitApplicationUpdateApi(queryParams map[string]string, authToken string) ResponseDTOs.DeploymentHistoryResponseDTO {
	resp, err := Base.MakeApiCall(ApplicationUpdateApiUrl, http.MethodPut, "", queryParams, authToken)
	Base.HandleError(err, ApplicationUpdate)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), ApplicationUpdate)
	return helmAppRouter.deploymentHistoryResponseDto
}

func HitHibernateWorkloadApi(payload string, authToken string) ResponseDTOs.HibernateApiResponseDTO {
	resp, err := Base.MakeApiCall(HibernateWorkLoadsApiUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, Hibernate)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), Hibernate)
	return helmAppRouter.hibernateApiResponseDto
}

func HitUnHibernateWorkloadApi(payload string, authToken string) ResponseDTOs.HibernateApiResponseDTO {
	resp, err := Base.MakeApiCall(UnHibernateWorkLoadsApiUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, Unhibernate)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), Hibernate)
	return helmAppRouter.hibernateApiResponseDto
}

func HitGetApplicationDetailApi(queryParams map[string]string, authToken string) ResponseDTOs.GetApplicationDetailResponseDTO {
	resp, err := Base.MakeApiCall(GetApplicationDetailUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetApplicationDetail)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), GetApplicationDetail)
	return helmAppRouter.getApplicationDetailResponseDto
}

func (structHelmAppRouter StructHelmAppRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructHelmAppRouter {
	switch apiName {
	case GetDeploymentHistory:
		json.Unmarshal(response, &structHelmAppRouter.deploymentHistoryResponseDto)
	case RollbackApplication:
		json.Unmarshal(response, &structHelmAppRouter.rollbackApplicationApiResponseDto)
	case GetReleaseInfoApi:
		json.Unmarshal(response, &structHelmAppRouter.releaseInfoApiResponseDto)
	case ApplicationUpdate:
		json.Unmarshal(response, &structHelmAppRouter.rollbackApplicationApiResponseDto)
	case Hibernate:
		json.Unmarshal(response, &structHelmAppRouter.hibernateApiResponseDto)
	case GetApplicationDetail:
		json.Unmarshal(response, &structHelmAppRouter.getApplicationDetailResponseDto)
	}
	return structHelmAppRouter
}

func GetRollbackAppApiRequestDto(HAppId string, version int) RequestDTOs.RollbackApplicationApiRequestDto {
	rollbackApplicationApiRequestDto := RequestDTOs.RollbackApplicationApiRequestDto{}
	rollbackApplicationApiRequestDto.HAppId = HAppId
	rollbackApplicationApiRequestDto.Version = version
	return rollbackApplicationApiRequestDto
}

type EnvironmentConfigHelmApp struct {
	HAppId                  string `env:"H_APP_ID" envDefault:"1|default|envoy-deepak-testing-v1"`
	ResourceNameToHibernate string `env:"RESOURCE_NAME_TO_HIBERNATE" envDefault:"envoy-deepak-testing-v1"`
}

func GetEnvironmentConfigForHelmApp() (*EnvironmentConfigHelmApp, error) {
	cfg := &EnvironmentConfigHelmApp{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

//Installed envoy already for test data having single Resource of deployment Kind
func createRequestPayloadForHibernateApi(appId string, kind string, name string, version string, group string, namespace string) RequestDTOs.HibernateApiRequestDTO {
	hibernateApiRequestDto := RequestDTOs.HibernateApiRequestDTO{}
	hibernateApiRequestDto.AppId = appId
	hibernateApiRequestDto.Resources = getResources(kind, name, version, group, namespace)
	return hibernateApiRequestDto
}

func getResources(kind string, name string, version string, group string, namespace string) []RequestDTOs.Resource {
	Resources := []RequestDTOs.Resource{}
	resource := RequestDTOs.Resource{}
	resource.Kind = kind
	resource.Name = name
	resource.Version = version
	resource.Group = group
	resource.Namespace = namespace
	Resources = append(Resources, resource)
	return Resources
}

type HelmAppTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *HelmAppTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
