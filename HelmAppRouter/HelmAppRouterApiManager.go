package HelmAppRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type DeploymentHistoryResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		InstalledAppInfo  InstalledAppInfo    `json:"installedAppInfo"`
		DeploymentHistory []DeploymentHistory `json:"deploymentHistory"`
	} `json:"result"`
}

type RollbackApplicationApiResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Success bool `json:"success"`
	} `json:"result"`
	Errors []Errors `json:"errors"`
}

type ReleaseInfoApiResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		InstalledAppInfo InstalledAppInfo `json:"installedAppInfo"`
		ReleaseInfo      ReleaseInfo      `json:"releaseInfo"`
	} `json:"result"`
}

type HibernateApiResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		ErrorMessage string       `json:"errorMessage"`
		Success      bool         `json:"success"`
		TargetObject TargetObject `json:"targetObject"`
	} `json:"result"`
	Errors []Errors `json:"errors"`
}

type GetApplicationDetailResponseDto struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Errors []Errors `json:"errors"`
	Result struct {
		AppDetail AppDetail `json:"appDetail"`
	} `json:"result"`
}

type InstalledAppInfo struct {
	AppId                 int    `json:"appId"`
	InstalledAppId        int    `json:"installedAppId"`
	InstalledAppVersionId int    `json:"installedAppVersionId"`
	AppStoreChartId       int    `json:"appStoreChartId"`
	EnvironmentName       string `json:"environmentName"`
	AppOfferingMode       string `json:"appOfferingMode"`
	ClusterId             int    `json:"clusterId"`
	EnvironmentId         int    `json:"environmentId"`
}

type DeploymentHistory struct {
	ChartMetadata ChartMetadata `json:"chartMetadata"`
	DockerImages  []string      `json:"dockerImages"`
	Version       int           `json:"version"`
	DeployedAt    DeployedAt    `json:"deployedAt"`
}

type RollbackApplicationApiRequestDto struct {
	HAppId  string `json:"hAppId"`
	Version int    `json:"version"`
}

type ApplicationUpdateRequestDto struct {
	Id                 int    `json:"id"`
	ReferenceValueId   int    `json:"referenceValueId"`
	ReferenceValueKind string `json:"referenceValueKind"`
	ValuesOverrideYaml string `json:"valuesOverrideYaml"`
	InstalledAppId     int    `json:"installedAppId"`
	AppStoreVersion    int    `json:"appStoreVersion"`
}

type ReleaseInfo struct {
	DeployedAppDetail DeployedAppDetail `json:"deployedAppDetail"`
	DefaultValues     string            `json:"defaultValues"`
	OverrideValues    string            `json:"overrideValues"`
	MergedValues      string            `json:"mergedValues"`
	Readme            string            `json:"readme"`
}

type DeployedAppDetail struct {
	AppId             string             `json:"appId"`
	AppName           string             `json:"appName"`
	ChartName         string             `json:"chartName"`
	EnvironmentDetail EnvironmentDetails `json:"environmentDetail"`
	LastDeployed      DeployedAt         `json:"LastDeployed"`
	ChartVersion      string             `json:"chartVersion"`
}
type HibernateApiRequestDto struct {
	AppId     string     `json:"appId"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Version   string `json:"version"`
	Namespace string `json:"namespace"`
}

type Errors struct {
	Code            string `json:"code"`
	InternalMessage string `json:"internalMessage"`
	UserMessage     string `json:"userMessage"`
}
type TargetObject struct {
	Group     string `json:"group"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Version   string `json:"version"`
}

type AppDetail struct {
	ApplicationStatus  string             `json:"applicationStatus"`
	ReleaseStatus      ReleaseStatus      `json:"releaseStatus"`
	ChartMetadata      ChartMetadata      `json:"chartMetadata"`
	EnvironmentDetails EnvironmentDetails `json:"environmentDetails"`
}

type ReleaseStatus struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type ChartMetadata struct {
	ChartName    string   `json:"chartName"`
	ChartVersion string   `json:"chartVersion"`
	Home         string   `json:"home"`
	Sources      []string `json:"sources"`
	Description  string   `json:"description"`
}

type EnvironmentDetails struct {
	ClusterName string `json:"clusterName"`
	ClusterId   int    `json:"clusterId"`
	Namespace   string `json:"namespace"`
}

type DeployedAt struct {
	Seconds int `json:"seconds"`
	Nanos   int `json:"nanos"`
}

type StructHelmAppRouter struct {
	deploymentHistoryResponseDto      DeploymentHistoryResponseDto
	rollbackApplicationApiResponseDto RollbackApplicationApiResponseDto
	applicationUpdateRequestDto       ApplicationUpdateRequestDto
	releaseInfoApiResponseDto         ReleaseInfoApiResponseDto
	hibernateApiResponseDto           HibernateApiResponseDto
	getApplicationDetailResponseDto   GetApplicationDetailResponseDto
}

func HitGetDeploymentHistoryById(queryParams map[string]string, authToken string) DeploymentHistoryResponseDto {
	resp, err := Base.MakeApiCall(GetDeploymentHistoryApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetDeploymentHistory)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), GetDeploymentHistory)
	return helmAppRouter.deploymentHistoryResponseDto
}

func HitRollbackApplicationApi(payload string, authToken string) RollbackApplicationApiResponseDto {
	resp, err := Base.MakeApiCall(RollbackApplicationApiUrl, http.MethodPut, payload, nil, authToken)
	Base.HandleError(err, RollbackApplication)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), RollbackApplication)
	return helmAppRouter.rollbackApplicationApiResponseDto
}

func HitGetReleaseInfoApi(queryParams map[string]string, authToken string) ReleaseInfoApiResponseDto {
	resp, err := Base.MakeApiCall(GetReleaseInfoApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetReleaseInfoApi)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), GetReleaseInfoApi)
	return helmAppRouter.releaseInfoApiResponseDto
}

func HitApplicationUpdateApi(queryParams map[string]string, authToken string) DeploymentHistoryResponseDto {
	resp, err := Base.MakeApiCall(ApplicationUpdateApiUrl, http.MethodPut, "", queryParams, authToken)
	Base.HandleError(err, ApplicationUpdate)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), ApplicationUpdate)
	return helmAppRouter.deploymentHistoryResponseDto
}

func HitHibernateWorkloadApi(payload string, authToken string) HibernateApiResponseDto {
	resp, err := Base.MakeApiCall(HibernateWorkLoadsApiUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, Hibernate)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), Hibernate)
	return helmAppRouter.hibernateApiResponseDto
}

func HitUnHibernateWorkloadApi(payload string, authToken string) HibernateApiResponseDto {
	resp, err := Base.MakeApiCall(UnHibernateWorkLoadsApiUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, Unhibernate)

	structHelmAppRouter := StructHelmAppRouter{}
	helmAppRouter := structHelmAppRouter.UnmarshalGivenResponseBody(resp.Body(), Hibernate)
	return helmAppRouter.hibernateApiResponseDto
}

func HitGetApplicationDetailApi(queryParams map[string]string, authToken string) GetApplicationDetailResponseDto {
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

func GetRollbackAppApiRequestDto(HAppId string, version int) RollbackApplicationApiRequestDto {
	rollbackApplicationApiRequestDto := RollbackApplicationApiRequestDto{}
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
func createRequestPayloadForHibernateApi(appId string, kind string, name string, version string, group string, namespace string) HibernateApiRequestDto {
	hibernateApiRequestDto := HibernateApiRequestDto{}
	hibernateApiRequestDto.AppId = appId
	hibernateApiRequestDto.Resources = getResources(kind, name, version, group, namespace)
	return hibernateApiRequestDto
}

func getResources(kind string, name string, version string, group string, namespace string) []Resource {
	Resources := []Resource{}
	resource := Resource{}
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
