package ApplicationRouter

import (
	"automation-suite/ApplicationRouter/ResponseDTOs"
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructApplicationRouter struct {
	resourceTreeResponseDTO     ResponseDTOs.ResourceTreeResponseDTO
	managedResourcesResponseDTO ResponseDTOs.ManagedResourcesResponseDTO
	listResponseDTO             ResponseDTOs.ListResponseDTO
	terminalSessionResponseDTO  ResponseDTOs.TerminalSessionResponseDTO
	applicationResponseDTO      ResponseDTOs.ApplicationResponseDTO
}

func HitGetResourceTreeApi(appName string, authToken string) ResponseDTOs.ResourceTreeResponseDTO {
	resp, err := Base.MakeApiCall(ApplicationsRouterBaseUrl+appName+"-devtron-demo"+"/resource-tree", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetResourceTreeApi)
	structAppLabelsRouter := StructApplicationRouter{}
	applicationRepoRouter := structAppLabelsRouter.UnmarshalGivenResponseBody(resp.Body(), GetResourceTreeApi)
	return applicationRepoRouter.resourceTreeResponseDTO
}

func HitGetManagedResourcesApi(appName string, authToken string) ResponseDTOs.ManagedResourcesResponseDTO {
	resp, err := Base.MakeApiCall(ApplicationsRouterBaseUrl+appName+"-devtron-demo"+"/managed-resources", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetManagedResourcesApi)
	structAppLabelsRouter := StructApplicationRouter{}
	applicationRepoRouter := structAppLabelsRouter.UnmarshalGivenResponseBody(resp.Body(), GetManagedResourcesApi)
	return applicationRepoRouter.managedResourcesResponseDTO
}

//we are no longer using this API from UI
/*func HitGetListApi(queryParams map[string]string, authToken string) ResponseDTOs.ListResponseDTO {
	resp, err := Base.MakeApiCall(ApplicationsRouterBaseUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetListApi)
	structAppLabelsRouter := StructApplicationRouter{}
	applicationRepoRouter := structAppLabelsRouter.UnmarshalGivenResponseBody(resp.Body(), GetListApi)
	return applicationRepoRouter.listResponseDTO
}*/

func HitGetTerminalSessionApi(AppId string, EnvId string, NameSpace string, PodName string, AppName string, authToken string) ResponseDTOs.TerminalSessionResponseDTO {
	resp, err := Base.MakeApiCall(GetTerminalSessionApiUrl+AppId+"/"+EnvId+"/"+NameSpace+"/"+PodName+"/sh/"+AppName, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetTerminalSessionApi)
	structAppLabelsRouter := StructApplicationRouter{}
	applicationRepoRouter := structAppLabelsRouter.UnmarshalGivenResponseBody(resp.Body(), GetTerminalSessionApi)
	return applicationRepoRouter.terminalSessionResponseDTO
}

func HitGetApplicationApi(name string, authToken string) ResponseDTOs.ApplicationResponseDTO {
	resp, err := Base.MakeApiCall(GetApplicationApiUrl+name, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetApplicationApi)
	structAppLabelsRouter := StructApplicationRouter{}
	applicationRepoRouter := structAppLabelsRouter.UnmarshalGivenResponseBody(resp.Body(), GetApplicationApi)
	return applicationRepoRouter.applicationResponseDTO
}

func (structApplicationRouter StructApplicationRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructApplicationRouter {
	switch apiName {
	case GetResourceTreeApi:
		json.Unmarshal(response, &structApplicationRouter.resourceTreeResponseDTO)
	case GetManagedResourcesApi:
		json.Unmarshal(response, &structApplicationRouter.managedResourcesResponseDTO)
	case GetListApi:
		json.Unmarshal(response, &structApplicationRouter.listResponseDTO)
	case GetTerminalSessionApi:
		json.Unmarshal(response, &structApplicationRouter.terminalSessionResponseDTO)
	case GetApplicationApi:
		json.Unmarshal(response, &structApplicationRouter.applicationResponseDTO)
	}
	return structApplicationRouter
}

type ApplicationsRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *ApplicationsRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
func (suite *ApplicationsRouterTestSuite) AfterSuite() {
	PipelineConfigRouter.DeleteAppWithCiCd(suite.authToken)
}
