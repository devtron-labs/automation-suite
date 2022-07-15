package externalLinkoutRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"automation-suite/externalLinkoutRouter/RequestDTO"
	"automation-suite/externalLinkoutRouter/ResponseDTO"

	"github.com/stretchr/testify/suite"
)

type LinkRouterStruct struct {
	createLinkResponseDto    ResponseDTO.CreateLinkResponseDto
	fetchAllToolsResponseDto ResponseDTO.FetchAllToolsResponseDto
	getLinkByIdResponseDto   ResponseDTO.GetLinkByIdResponseDto
}

func GetSaveLinkRequestDto(monitoringToolId int, slice []int) []RequestDTO.CreateLinkRequestDto {
	var createLinkRequestDto RequestDTO.CreateLinkRequestDto
	createLinkRequestDto.Name = "automated-" + strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	createLinkRequestDto.Active = true
	createLinkRequestDto.MonitoringToolId = monitoringToolId
	createLinkRequestDto.Url = "http://www." + strings.ToLower(Base.GetRandomStringOfGivenLength(5)) + ".com/{namespace}/{appName}/details/{appId}/env/{envId}/details/"
	createLinkRequestDto.ClusterIds = append(createLinkRequestDto.ClusterIds, slice...)
	var createLinkRequestDto2 []RequestDTO.CreateLinkRequestDto
	createLinkRequestDto2 = append(createLinkRequestDto2, createLinkRequestDto)
	return createLinkRequestDto2
}
func GetUpdateLinkRequestPayload(id int, linkName string, monitoringToolId int, url string) RequestDTO.CreateLinkRequestDto {
	var updateLinkRequestDto RequestDTO.CreateLinkRequestDto
	updateLinkRequestDto.Id = id
	updateLinkRequestDto.Name = linkName
	updateLinkRequestDto.MonitoringToolId = monitoringToolId
	updateLinkRequestDto.Url = url
	updateLinkRequestDto.Active = true
	return updateLinkRequestDto
}
func (linkRouterStruct LinkRouterStruct) UnmarshalGivenResponseBody(response []byte, apiName string) LinkRouterStruct {
	switch apiName {
	case FetchAllLinkApi:
		json.Unmarshal(response, &linkRouterStruct.getLinkByIdResponseDto)
	case CreateLinkApi:
		json.Unmarshal(response, &linkRouterStruct.createLinkResponseDto)
	case SaveToolApi:
		json.Unmarshal(response, &linkRouterStruct.fetchAllToolsResponseDto)
	case UpdateLinkApi:
		json.Unmarshal(response, &linkRouterStruct.createLinkResponseDto)

	}

	return linkRouterStruct
}
func HitCreateLinkApi(payload []byte, authToken string) ResponseDTO.CreateLinkResponseDto {
	resp, err := Base.MakeApiCall(SaveExternalLink, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, CreateLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	externalLinkOutRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), CreateLinkApi)
	return externalLinkOutRouter.createLinkResponseDto
}

func HitDeleteLinkApi(id int, authToken string) ResponseDTO.CreateLinkResponseDto {
	resp, err := Base.MakeApiCall(SaveExternalLink+"?id="+strconv.Itoa(id), http.MethodDelete, "", nil, authToken)
	Base.HandleError(err, CreateLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	externalLinkOutRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), CreateLinkApi)
	return externalLinkOutRouter.createLinkResponseDto
}

func HitFetchAllToolsApi(authToken string) ResponseDTO.FetchAllToolsResponseDto {
	resp, err := Base.MakeApiCall(SaveToolApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, SaveToolApi)

	linkRouterStruct := LinkRouterStruct{}
	externalLinkOutRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), SaveToolApi)
	return externalLinkOutRouter.fetchAllToolsResponseDto
}

func HitFetchAllLinkApi(authToken string) ResponseDTO.GetLinkByIdResponseDto {
	resp, err := Base.MakeApiCall(SaveExternalLink, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchAllLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	externalLinkOutRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllLinkApi)
	return externalLinkOutRouter.getLinkByIdResponseDto
}
func HitFetchLinksByClusterIdApi(queryParams map[string]string, authToken string) ResponseDTO.GetLinkByIdResponseDto {
	resp, err := Base.MakeApiCall(SaveExternalLink, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, FetchAllLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	externalLinkOutRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllLinkApi)
	return externalLinkOutRouter.getLinkByIdResponseDto
}

func HitGetLinkByIdApi(id string, authToken string) ResponseDTO.GetLinkByIdResponseDto {
	resp, err := Base.MakeApiCall(SaveExternalLink+"/"+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetLinkByIdApi)

	linkRouterStruct := LinkRouterStruct{}
	externalLinkOutRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), GetLinkByIdApi)
	return externalLinkOutRouter.getLinkByIdResponseDto
}

func HitUpdateLinkApi(byteValueOfStruct []byte, authToken string) ResponseDTO.CreateLinkResponseDto {
	resp, err := Base.MakeApiCall(SaveExternalLink, http.MethodPut, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, UpdateLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	externalLinkOutRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), UpdateLinkApi)
	return externalLinkOutRouter.createLinkResponseDto
}

type LinkOutRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *LinkOutRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
