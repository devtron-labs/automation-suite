package OrchestratorServerRouter

import (
	"automation-suite/OrchestratorServerRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructOrchestratorServerRouter struct {
	getOrchestratorResponse ResponseDTOs.GetOrchestratorResponse
}

func HitGetOrchestratorServerApi(authToken string) ResponseDTOs.GetOrchestratorResponse {
	resp, err := Base.MakeApiCall(GetOrchestratorServerApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetOrchestratorServerApi)

	structOrchestratorServerRouter := StructOrchestratorServerRouter{}
	orchestratorServerRouter := structOrchestratorServerRouter.UnmarshalGivenResponseBody(resp.Body(), GetOrchestratorServerApi)
	return orchestratorServerRouter.getOrchestratorResponse
}

func (structOrchestratorServerRouter StructOrchestratorServerRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructOrchestratorServerRouter {
	switch apiName {
	case GetOrchestratorServerApi:
		json.Unmarshal(response, &structOrchestratorServerRouter.getOrchestratorResponse)
	}
	return structOrchestratorServerRouter
}

type ServerRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *ServerRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
