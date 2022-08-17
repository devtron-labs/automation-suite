package ApplicationRouter

import (
	"automation-suite/ApplicationRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructApplicationRouter struct {
	resourceTreeResponseDTO ResponseDTOs.ResourceTreeResponseDTO
}

func HitGetResourceTreeApi(appName string, authToken string) ResponseDTOs.ResourceTreeResponseDTO {
	resp, err := Base.MakeApiCall(GetResourceTreeApiBaseUrl+appName+"-devtron-demo"+"/resource-tree", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetResourceTreeApi)
	structAppLabelsRouter := StructApplicationRouter{}
	applicationRepoRouter := structAppLabelsRouter.UnmarshalGivenResponseBody(resp.Body(), GetResourceTreeApi)
	return applicationRepoRouter.resourceTreeResponseDTO
}

func (structApplicationRouter StructApplicationRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructApplicationRouter {
	switch apiName {
	case GetResourceTreeApi:
		json.Unmarshal(response, &structApplicationRouter.resourceTreeResponseDTO)
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
