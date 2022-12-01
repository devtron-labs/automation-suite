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

func GetSaveLinkRequestDtoList(noOfLinkRequired int, identifiers string, monitoringToolId int, IsEditable bool, ParentObjectType string, Url string, ChildObjType string, ClusterId int) []RequestDTO.CreateLinkRequestDto1 {
	var listOfCreateLinkRequestDTO []RequestDTO.CreateLinkRequestDto1

	for i := 0; i < noOfLinkRequired; i++ {
		var createLinkRequestDto RequestDTO.CreateLinkRequestDto1
		identifiers := GetIdentifierObject(identifiers, ChildObjType, ClusterId)
		createLinkRequestDto.MonitoringToolId = monitoringToolId
		createLinkRequestDto.Name = "automation" + strconv.Itoa(i+1) + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		createLinkRequestDto.Description = "This is description for testing purpose only"
		createLinkRequestDto.Identifiers = identifiers
		createLinkRequestDto.IsEditable = IsEditable
		createLinkRequestDto.Type = ParentObjectType
		createLinkRequestDto.Url = Url + strconv.Itoa(i+1)
		listOfCreateLinkRequestDTO = append(listOfCreateLinkRequestDTO, createLinkRequestDto)
	}
	return listOfCreateLinkRequestDTO
}

func GetIdentifierObject(ChildIdentifier string, ChildObjType string, ClusterId int) []RequestDTO.Identifiers {
	var listOfIdentifiers []RequestDTO.Identifiers
	var identifier RequestDTO.Identifiers
	identifier.Identifier = ChildIdentifier
	identifier.Type = ChildObjType
	identifier.ClusterId = ClusterId
	listOfIdentifiers = append(listOfIdentifiers, identifier)
	return listOfIdentifiers
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
