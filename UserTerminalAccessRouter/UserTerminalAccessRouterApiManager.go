package UserTerminalAccessRouter

import (
	"automation-suite/UserTerminalAccessRouter/RequestDTOs"
	"automation-suite/UserTerminalAccessRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructUserTerminalAccessRouter struct {
	startTerminalSessionResponseDTO ResponseDTOs.StartTerminalSessionResponseDTO
	terminalPodEventsResponseDTO    ResponseDTOs.TerminalPodEventsResponseDTO
	terminalPodManifestResponseDTO  ResponseDTOs.TerminalPodManifestResponseDTO
}

func CreatePayLoadForStartTerminalSession(ClusterId int, BaseImage string, NodeName string, NameSpace string, ShellName string) RequestDTOs.StartTerminalSessionRequestDTO {
	startTerminalSessionRequestDTO := RequestDTOs.StartTerminalSessionRequestDTO{}
	startTerminalSessionRequestDTO.ClusterId = ClusterId
	startTerminalSessionRequestDTO.BaseImage = BaseImage
	startTerminalSessionRequestDTO.ShellName = ShellName
	startTerminalSessionRequestDTO.NodeName = NodeName
	startTerminalSessionRequestDTO.Namespace = NameSpace
	return startTerminalSessionRequestDTO
}

func CreatePayLoadForUpdateTerminalShellSession(TerminalAccessId int, ShellName string) RequestDTOs.UpdateTerminalShellSessionRequestDTO {
	updateTerminalShellSessionRequestDTO := RequestDTOs.UpdateTerminalShellSessionRequestDTO{}
	updateTerminalShellSessionRequestDTO.ShellName = ShellName
	updateTerminalShellSessionRequestDTO.TerminalAccessId = TerminalAccessId
	return updateTerminalShellSessionRequestDTO
}

func HitStartTerminalSessionApi(payload []byte, authToken string) ResponseDTOs.StartTerminalSessionResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+StartTerminalSessionApiPath, http.MethodPut, string(payload), nil, authToken)
	Base.HandleError(err, StartTerminalSessionApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), StartTerminalSessionApi)
	return userRouter.startTerminalSessionResponseDTO
}

func HitFetchTerminalStatusApi(queryParams map[string]string, authToken string) ResponseDTOs.StartTerminalSessionResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+FetchTerminalStatusApiPath, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, FetchTerminalStatusApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), FetchTerminalStatusApi)
	return userRouter.startTerminalSessionResponseDTO
}

func HitFetchTerminalPodEventsApi(queryParams map[string]string, authToken string) ResponseDTOs.TerminalPodEventsResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+FetchTerminalPodEventsApiPath, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, FetchTerminalPodEventsApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), FetchTerminalPodEventsApi)
	return userRouter.terminalPodEventsResponseDTO
}

func HitFetchTerminalPodManifestApi(queryParams map[string]string, authToken string) ResponseDTOs.TerminalPodManifestResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+FetchTerminalPodManifestApiPath, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, FetchTerminalPodManifestApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), FetchTerminalPodManifestApi)
	return userRouter.terminalPodManifestResponseDTO
}

func HitStopTerminalSessionApi(queryParams map[string]string, authToken string) ResponseDTOs.StartTerminalSessionResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+StopTerminalSessionApiPath, http.MethodPost, "", queryParams, authToken)
	Base.HandleError(err, StopTerminalSessionApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), StopTerminalSessionApi)
	return userRouter.startTerminalSessionResponseDTO
}

func HitUpdateTerminalSessionApi(payload []byte, authToken string) ResponseDTOs.StartTerminalSessionResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+UpdateTerminalSessionApiPath, http.MethodPut, string(payload), nil, authToken)
	Base.HandleError(err, UpdateTerminalSessionApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), UpdateTerminalSessionApi)
	return userRouter.startTerminalSessionResponseDTO
}

func HitUpdateTerminalShellSessionApi(payload []byte, authToken string) ResponseDTOs.StartTerminalSessionResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+UpdateTerminalShellSessionApiPath, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, UpdateTerminalShellSessionApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), UpdateTerminalShellSessionApi)
	return userRouter.startTerminalSessionResponseDTO
}

func HitDisconnectTerminalSessionApi(queryParams map[string]string, authToken string) ResponseDTOs.StartTerminalSessionResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+DisconnectTerminalSessionApiPath, http.MethodPost, "", queryParams, authToken)
	Base.HandleError(err, DisconnectTerminalSessionApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), DisconnectTerminalSessionApi)
	return userRouter.startTerminalSessionResponseDTO
}

func HitDisconnectAllTerminalSessionAndRetryApi(payload []byte, authToken string) ResponseDTOs.StartTerminalSessionResponseDTO {
	resp, err := Base.MakeApiCall(UserTerminalRouterBaseUrl+DisconnectAllTerminalSessionAndRetryApiPath, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, DisconnectAllTerminalSessionAndRetryApi)
	structUserTerminalAccessRouter := StructUserTerminalAccessRouter{}
	userRouter := structUserTerminalAccessRouter.UnmarshalGivenResponseBody(resp.Body(), DisconnectAllTerminalSessionAndRetryApi)
	return userRouter.startTerminalSessionResponseDTO
}

func (structUserTerminalAccessRouter StructUserTerminalAccessRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructUserTerminalAccessRouter {
	switch apiName {
	case StartTerminalSessionApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.startTerminalSessionResponseDTO)
	case FetchTerminalStatusApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.startTerminalSessionResponseDTO)
	case FetchTerminalPodEventsApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.terminalPodEventsResponseDTO)
	case FetchTerminalPodManifestApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.terminalPodManifestResponseDTO)
	case StopTerminalSessionApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.startTerminalSessionResponseDTO)
	case UpdateTerminalSessionApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.startTerminalSessionResponseDTO)
	case UpdateTerminalShellSessionApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.startTerminalSessionResponseDTO)
	case DisconnectTerminalSessionApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.startTerminalSessionResponseDTO)
	case DisconnectAllTerminalSessionAndRetryApi:
		json.Unmarshal(response, &structUserTerminalAccessRouter.startTerminalSessionResponseDTO)
	}
	return structUserTerminalAccessRouter
}

type UserTerminalAccessRoutersTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *UserTerminalAccessRoutersTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
