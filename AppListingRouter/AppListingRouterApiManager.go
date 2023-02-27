package AppListingRouter

import (
	"automation-suite/AppListingRouter/RequestDTOs"
	"automation-suite/AppListingRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"strconv"
)

type StructAppListingRouter struct {
	fetchAllStageStatusResponseDto    ResponseDTOs.FetchAllStageStatusResponseDTO
	fetchOtherEnvResponseDto          ResponseDTOs.FetchOtherEnvResponseDTO
	fetchAppsByEnvironmentResponseDTO ResponseDTOs.FetchAppsByEnvironmentResponseDTO
}

func FetchAllStageStatus(id int, authToken string) ResponseDTOs.FetchAllStageStatusResponseDTO {
	AppId := map[string]string{
		"app-id": strconv.Itoa(id),
	}
	resp, err := Base.MakeApiCall(GetStageStatusApiUrl, http.MethodGet, "", AppId, authToken)
	Base.HandleError(err, FetchAllStageStatusApi)

	structAppListingRouter := StructAppListingRouter{}
	apiRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), FetchAllStageStatusApi)
	return apiRouter.fetchAllStageStatusResponseDto
}

func FetchOtherEnv(id int, authToken string) ResponseDTOs.FetchOtherEnvResponseDTO {
	AppId := map[string]string{
		"app-id": strconv.Itoa(id),
	}
	resp, err := Base.MakeApiCall(GetOtherEnvApiUrl, http.MethodGet, "", AppId, authToken)
	Base.HandleError(err, FetchOtherEnvApi)

	structAppListingRouter := StructAppListingRouter{}
	apiRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), FetchOtherEnvApi)
	return apiRouter.fetchOtherEnvResponseDto
}

func GetPayloadForApiFetchAppsByEnvironment(Environments []int, Teams []int, Namespaces []string, AppNameSearch string, AppStatuses []string, SortOrder string, Offset int, HOffset int, Size int) RequestDTOs.FetchAppsByEnvironmentRequestDTO {
	FetchAppsByEnvironmentRequestDTO := RequestDTOs.FetchAppsByEnvironmentRequestDTO{}
	FetchAppsByEnvironmentRequestDTO.Environments = Environments
	FetchAppsByEnvironmentRequestDTO.Teams = Teams
	FetchAppsByEnvironmentRequestDTO.Namespaces = Namespaces
	FetchAppsByEnvironmentRequestDTO.AppNameSearch = AppNameSearch
	FetchAppsByEnvironmentRequestDTO.AppStatuses = AppStatuses
	FetchAppsByEnvironmentRequestDTO.SortBy = "appNameSort"
	FetchAppsByEnvironmentRequestDTO.SortOrder = SortOrder
	FetchAppsByEnvironmentRequestDTO.Offset = Offset
	FetchAppsByEnvironmentRequestDTO.Size = Size
	return FetchAppsByEnvironmentRequestDTO
}

func HitApiFetchAppsByEnvironment(payload []byte, authToken string) ResponseDTOs.FetchAppsByEnvironmentResponseDTO {
	resp, err := Base.MakeApiCall(AppListingRoutersBaseUrl+FetchAppsByEnvironmentUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, FetchAppsByEnvironment)
	structAppListingRouter := StructAppListingRouter{}
	apiRouter := structAppListingRouter.UnmarshalGivenResponseBody(resp.Body(), FetchAppsByEnvironment)
	return apiRouter.fetchAppsByEnvironmentResponseDTO
}

func (structAppListingRouter StructAppListingRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppListingRouter {
	switch apiName {
	case FetchAllStageStatusApi:
		json.Unmarshal(response, &structAppListingRouter.fetchAllStageStatusResponseDto)
	case FetchOtherEnvApi:
		json.Unmarshal(response, &structAppListingRouter.fetchOtherEnvResponseDto)
	case FetchAppsByEnvironment:
		json.Unmarshal(response, &structAppListingRouter.fetchAppsByEnvironmentResponseDTO)
	}
	return structAppListingRouter
}

type AppsListingRouterTestSuite struct {
	suite.Suite
	authToken string
}

// SetupSuite This method runs on first priority before starting the suite means before executing any test case of the suite
func (suite *AppsListingRouterTestSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
}
