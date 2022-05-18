package AppListingRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type AppListingRouter struct {
	suite.Suite
	authToken string
}

func (suite *AppListingRouter) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}

type StructAppListingRouter struct {
	fetchAllStageStatusResponseDto FetchAllStageStatusResponseDto
}
type FetchAllStageStatusResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Stage     int    `json:"stage"`
		StageName string `json:"stageName"`
		Status    bool   `json:"status"`
		Required  bool   `json:"required"`
	} `json:"result"`
	Errors []Base.Errors `json:"errors"`
}

func (structAppListingRouter StructAppListingRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppListingRouter {
	switch apiName {
	case FetchAllStageStatusApi:
		json.Unmarshal(response, &structAppListingRouter.fetchAllStageStatusResponseDto)
	}
	return structAppListingRouter
}

func FetchAllStageStatus(id string, authToken string) FetchAllStageStatusResponseDto {
	AppId := map[string]string{
		"id": id,
	}
	resp, err := Base.MakeApiCall(GetStageStatusApiUrl, http.MethodGet, "", AppId, authToken)
	Base.HandleError(err, FetchAllStageStatusApi)

	structAppListingRouter := StructAppListingRouter{}
	apiRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), FetchAllStageStatusApi)
	return apiRouter.fetchAllStageStatusResponseDto
}
