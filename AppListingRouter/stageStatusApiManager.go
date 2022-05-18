package AppListingRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type AppListingRouter struct {
	suite.Suite
	authToken string
}

func (suite *AppListingRouter) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
	fmt.Println(suite.authToken)
}

type InstallationScriptStruct struct {
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
}

func (installationScriptStruct InstallationScriptStruct) UnmarshalGivenResponseBody(response []byte, apiName string) InstallationScriptStruct {
	switch apiName {
	case FetchAllStageStatusApi:
		json.Unmarshal(response, &installationScriptStruct.fetchAllStageStatusResponseDto)
	}
	return installationScriptStruct
}

func FetchAllStageStatus(id map[string]string, authToken string) FetchAllStageStatusResponseDto {
	resp, err := Base.MakeApiCall(GetStageStatusApiUrl, http.MethodGet, "", id, authToken)
	Base.HandleError(err, FetchAllStageStatusApi)

	installationScriptStruct := InstallationScriptStruct{}
	apiRouter := installationScriptStruct.UnmarshalGivenResponseBody(resp.Body(), "FetchAllStageStatus")
	return apiRouter.fetchAllStageStatusResponseDto
}
