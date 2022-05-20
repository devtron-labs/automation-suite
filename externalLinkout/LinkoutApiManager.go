package externalLinkout

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"net/http"
)

type CreateLinkRequestDto struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Url              string `json:"url"`
	MonitoringToolId int    `json:"monitoringToolId"`
	ClusterIds       []int  `json:"clusterIds"`
	Active           bool   `json:"active"`
}
type CreateToolRequestDto struct {
	Id   int    `json:"id"`
	Icon string `json:"icon"`
	Name string `json:"Name"`
}
type CreateLinkResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	Result struct {
		Id               int      `json:"id"`
		Name             string   `json:"name"`
		Url              string   `json:"url"`
		MonitoringToolId int      `json:"monitoringToolId"`
		ClusterIds       []string `json:"clusterIds"`
		Active           bool     `json:"active"`
	} `json:"result"`
}
type DeleteLinkResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}
type SaveToolResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors struct {
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	Result struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	} `json:"result"`
}

type FetchAllToolsResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	} `json:"result"`
}
type GetLinkByIdResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id               int      `json:"id"`
		Name             string   `json:"name"`
		Url              string   `json:"url"`
		MonitoringToolId int      `json:"monitoringToolId"`
		ClusterIds       []string `json:"clusterIds"`
		Active           bool     `json:"active"`
	} `json:"result"`
}
type LinkRouterStruct struct {
	createLinkResponseDto    CreateLinkResponseDto
	deleteLinkResponseDto    DeleteLinkResponseDto
	saveToolResponseDto      SaveToolResponseDto
	fetchAllToolsResponseDto FetchAllToolsResponseDto
	deleteToolResponseDto    DeleteToolResponseDto
	fetchAllLinkResponseDto  FetchAllLinkResponseDto
	getLinkByIdResponseDto   GetLinkByIdResponseDto
}
type DeleteToolResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}
type FetchAllLinkResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id               int      `json:"id"`
		Name             string   `json:"name"`
		Url              string   `json:"url"`
		MonitoringToolId int      `json:"monitoringToolId"`
		ClusterIds       []string `json:"clusterIds"`
		Active           bool     `json:"active"`
	} `json:"result"`
}

func GetSaveLinkRequestDto() CreateLinkRequestDto {
	var createLinkRequestDto CreateLinkRequestDto
	createLinkRequestDto.Name = Base.GetRandomStringOfGivenLength(10)
	createLinkRequestDto.Active = true
	createLinkRequestDto.MonitoringToolId = 1
	createLinkRequestDto.Url = Base.GetRandomStringOfGivenLength(20)
	return createLinkRequestDto
}

func GetSaveLinkRequestInvalidClusterIdDto() CreateLinkRequestDto {
	slice := []int{rand.Intn(89-10) + 10}
	var createLinkRequestDto CreateLinkRequestDto
	createLinkRequestDto.Name = Base.GetRandomStringOfGivenLength(10)
	createLinkRequestDto.Active = true
	createLinkRequestDto.MonitoringToolId = 1
	createLinkRequestDto.ClusterIds = slice
	createLinkRequestDto.Url = Base.GetRandomStringOfGivenLength(20)
	return createLinkRequestDto
}
func GetSaveLinkRequestOneValidOneInvalidClusterId() CreateLinkRequestDto {
	slice := []int{1, rand.Intn(89-10) + 10}
	var createLinkRequestDto CreateLinkRequestDto
	createLinkRequestDto.Name = Base.GetRandomStringOfGivenLength(10)
	createLinkRequestDto.Active = true
	createLinkRequestDto.MonitoringToolId = 1
	createLinkRequestDto.ClusterIds = slice
	createLinkRequestDto.Url = Base.GetRandomStringOfGivenLength(20)
	return createLinkRequestDto
}
func GetSaveLinkRequestInvalidMonitoringToolIdDto() CreateLinkRequestDto {
	slice := []int{rand.Intn(89-10) + 10}
	var createLinkRequestDto CreateLinkRequestDto
	createLinkRequestDto.Name = Base.GetRandomStringOfGivenLength(10)
	createLinkRequestDto.Active = true
	createLinkRequestDto.MonitoringToolId = slice[0]
	createLinkRequestDto.Url = Base.GetRandomStringOfGivenLength(20)
	return createLinkRequestDto
}

func (linkRouterStruct LinkRouterStruct) UnmarshalGivenResponseBody(response []byte, apiName string) LinkRouterStruct {
	switch apiName {
	case FetchAllLinkApi:
		json.Unmarshal(response, &linkRouterStruct.fetchAllToolsResponseDto)
	case CreateLinkApi:
		json.Unmarshal(response, &linkRouterStruct.createLinkResponseDto)
	case DeleteLinkApi:
		json.Unmarshal(response, &linkRouterStruct.deleteLinkResponseDto)

	}
	return linkRouterStruct
}
func HitCreateLinkApi(payload []byte, authToken string) CreateLinkResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		var createLinkRequestDto CreateLinkRequestDto
		createLinkRequestDto.Name = Base.GetRandomStringOfGivenLength(10)
		createLinkRequestDto.Active = false
		createLinkRequestDto.MonitoringToolId = 1
		createLinkRequestDto.Url = Base.GetRandomStringOfGivenLength(20)
		byteValueOfStruct, _ := json.Marshal(createLinkRequestDto)
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := Base.MakeApiCall(CreateLinkApiUrl, http.MethodPost, payloadOfApi, nil, "")
	Base.HandleError(err, CreateLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	linkRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), CreateLinkApi)
	return linkRouter.createLinkResponseDto
}

func GetPayLoadForDeleteLinkAPI(id int, name string, monitoringToolId int, url string, isActive bool) []byte {
	var updateLinkDto CreateLinkRequestDto
	updateLinkDto.Id = id
	updateLinkDto.Name = name
	updateLinkDto.MonitoringToolId = monitoringToolId
	updateLinkDto.Url = url
	updateLinkDto.Active = isActive
	byteValueOfStruct, _ := json.Marshal(updateLinkDto)
	return byteValueOfStruct
}
func HitDeleteLinkApi(byteValueOfStruct []byte, authToken string) DeleteLinkResponseDto {
	resp, err := Base.MakeApiCall(CreateLinkApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, "")
	Base.HandleError(err, DeleteLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	linkRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteLinkApi)
	return linkRouter.deleteLinkResponseDto
}

////////////////////////////   Tools

func HitCreateToolApi(payload []byte) SaveToolResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		var createToolRequestDto CreateToolRequestDto
		createToolRequestDto.Id = 1
		createToolRequestDto.Name = Base.GetRandomStringOfGivenLength(10)
		createToolRequestDto.Icon = Base.GetRandomStringOfGivenLength(10)
		byteValueOfStruct, _ := json.Marshal(createToolRequestDto)
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := Base.MakeApiCall(SaveToolApiUrl, http.MethodPost, payloadOfApi, nil, "")
	Base.HandleError(err, SaveToolApi)

	linkRouterStruct := LinkRouterStruct{}
	toolRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), SaveToolApi)
	return toolRouter.saveToolResponseDto
}

func HitFetchAllToolsApi() FetchAllToolsResponseDto {
	resp, err := Base.MakeApiCall(SaveLinkApiUrl, http.MethodGet, "", nil, "")
	Base.HandleError(err, FetchAllLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	linkRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllLinkApi)
	return linkRouter.fetchAllToolsResponseDto
}
func GetPayLoadForDeleteToolAPI(id int, name string, icon string) []byte {
	var updateToolDto CreateToolRequestDto
	updateToolDto.Id = id
	updateToolDto.Name = name
	updateToolDto.Icon = icon
	byteValueOfStruct, _ := json.Marshal(updateToolDto)
	return byteValueOfStruct
}

func HitDeleteToolApi(byteValueOfStruct []byte) DeleteToolResponseDto {
	resp, err := Base.MakeApiCall(SaveToolApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, "")
	Base.HandleError(err, DeleteToolApi)

	linkRouterStruct := LinkRouterStruct{}
	toolRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteToolApi)
	return toolRouter.deleteToolResponseDto
}

//////////////////////

func HitFetchAllLinkApi() FetchAllLinkResponseDto {
	resp, err := Base.MakeApiCall(SaveLinkApiUrl, http.MethodGet, "", nil, "")
	Base.HandleError(err, FetchAllLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	linkRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllLinkApi)
	return linkRouter.fetchAllLinkResponseDto
}
func HitFetchAllLinkByClusterIdApi(id map[string]string) FetchAllLinkResponseDto {
	resp, err := Base.MakeApiCall(SaveLinkApiUrl, http.MethodGet, "", id, "")
	Base.HandleError(err, FetchAllLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	linkRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllLinkApi)
	return linkRouter.fetchAllLinkResponseDto
}

func HitGetLinkByIdApi(id string, authToken string) GetLinkByIdResponseDto {
	resp, err := Base.MakeApiCall(SaveLinkApiUrl+"/"+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetLinkByIdApi)

	linkRouterStruct := LinkRouterStruct{}
	linkRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), GetLinkByIdApi)
	return linkRouter.getLinkByIdResponseDto
}

func GetUpdateLinkRequestPayload(id int, linkName string, monitoringToolId int) []byte {
	var updateLinkRequestDto CreateLinkRequestDto
	updateLinkRequestDto.Id = id
	updateLinkRequestDto.Name = linkName
	updateLinkRequestDto.MonitoringToolId = monitoringToolId
	updateLinkRequestDto.Active = true
	byteValueOfStruct, _ := json.Marshal(updateLinkRequestDto)
	return byteValueOfStruct
}
func GetUpdateLinkRequestPayloadInvalidMonitorigId(id int, linkName string) []byte {
	var updateLinkRequestDto CreateLinkRequestDto
	updateLinkRequestDto.Id = id
	updateLinkRequestDto.Name = linkName
	updateLinkRequestDto.MonitoringToolId = rand.Intn(89-10) + 10
	updateLinkRequestDto.Active = true
	byteValueOfStruct, _ := json.Marshal(updateLinkRequestDto)
	return byteValueOfStruct
}
func GetUpdateLinkRequestPayloadInvalidClusterId(id int, linkName string) []byte {
	slice := []int{rand.Intn(89-10) + 10}
	var updateLinkRequestDto CreateLinkRequestDto
	updateLinkRequestDto.Id = id
	updateLinkRequestDto.Name = linkName
	updateLinkRequestDto.MonitoringToolId = 1
	updateLinkRequestDto.Active = true
	updateLinkRequestDto.ClusterIds = slice
	byteValueOfStruct, _ := json.Marshal(updateLinkRequestDto)
	return byteValueOfStruct
}

func HitUpdateLinkApi(byteValueOfStruct []byte, authToken string) CreateLinkResponseDto {
	resp, err := Base.MakeApiCall(SaveLinkApiUrl, http.MethodPut, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, UpdateLinkApi)

	linkRouterStruct := LinkRouterStruct{}
	linkRouter := linkRouterStruct.UnmarshalGivenResponseBody(resp.Body(), CreateLinkApi)
	return linkRouter.createLinkResponseDto
}

type LinkOutRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *LinkOutRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
