package TeamRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type SaveTeamResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	Result struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"result"`
}

type GetTeamByIdResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"result"`
}

type SaveTeamRequestDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type DeleteTeamResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type FetchAllTeamResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"result"`
}

type TeamRouterStruct struct {
	saveTeamResponseDto     SaveTeamResponseDto
	deleteTeamResponseDto   DeleteTeamResponseDto
	fetchAllTeamResponseDto FetchAllTeamResponseDto
	getTeamByIdResponseDto  GetTeamByIdResponseDto
}

func (teamRouterStruct TeamRouterStruct) UnmarshalGivenResponseBody(response []byte, apiName string) TeamRouterStruct {
	switch apiName {
	case FetchAllTeamApi:
		json.Unmarshal(response, &teamRouterStruct.fetchAllTeamResponseDto)
	case SaveTeamApi:
		json.Unmarshal(response, &teamRouterStruct.saveTeamResponseDto)
	case DeleteTeamApi:
		json.Unmarshal(response, &teamRouterStruct.deleteTeamResponseDto)
	case GetTeamByIdApi:
		json.Unmarshal(response, &teamRouterStruct.getTeamByIdResponseDto)
	}
	return teamRouterStruct
}

func HitSaveTeamApi(payload []byte, authToken string) SaveTeamResponseDto {
	var payloadOfApi string
	if payload != nil {
		payloadOfApi = string(payload)
	} else {
		saveTeamRequestDto := GetSaveTeamRequestDto()
		byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)
		payloadOfApi = string(byteValueOfStruct)
	}

	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodPost, payloadOfApi, nil, authToken)
	Base.HandleError(err, SaveTeamApi)

	teamRouterStruct := TeamRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), SaveTeamApi)
	return teamRouter.saveTeamResponseDto
}

func GetSaveTeamRequestDto() SaveTeamRequestDto {
	var saveTeamRequestDto SaveTeamRequestDto
	teamName := Base.GetRandomStringOfGivenLength(10)
	saveTeamRequestDto.Name = teamName
	saveTeamRequestDto.Active = true
	return saveTeamRequestDto
}

func HitFetchAllTeamApi(authToken string) FetchAllTeamResponseDto {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchAllTeamApi)

	teamRouterStruct := TeamRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllTeamApi)
	return teamRouter.fetchAllTeamResponseDto
}

func HitDeleteTeamApi(byteValueOfStruct []byte, authToken string) DeleteTeamResponseDto {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, DeleteTeamApi)

	teamRouterStruct := TeamRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteTeamApi)
	return teamRouter.deleteTeamResponseDto
}

func HitGetTeamByIdApi(id string, authToken string) GetTeamByIdResponseDto {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl+"/"+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetTeamByIdApi)

	teamRouterStruct := TeamRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), GetTeamByIdApi)
	return teamRouter.getTeamByIdResponseDto
}

func HitUpdateTeamApi(byteValueOfStruct []byte, authToken string) SaveTeamResponseDto {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodPut, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, UpdateTeamApi)

	teamRouterStruct := TeamRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), SaveTeamApi)
	return teamRouter.saveTeamResponseDto
}

func HitFetchForAutocompleteApi(authToken string) FetchAllTeamResponseDto {
	resp, err := Base.MakeApiCall(FetchForAutocompleteApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchForAutocompleteApi)

	teamRouterStruct := TeamRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllTeamApi)
	return teamRouter.fetchAllTeamResponseDto
}

func GetPayLoadForDeleteAPI(id int, name string, isActive bool) []byte {
	var updateTeamDto SaveTeamRequestDto
	updateTeamDto.Id = id
	updateTeamDto.Name = name
	updateTeamDto.Active = isActive
	byteValueOfStruct, _ := json.Marshal(updateTeamDto)
	return byteValueOfStruct
}

func GetUpdateTeamRequestPayload(id int, teamName string) []byte {
	var updateTeamRequestDto SaveTeamRequestDto
	updateTeamRequestDto.Name = teamName
	updateTeamRequestDto.Id = id
	updateTeamRequestDto.Active = true
	byteValueOfStruct, _ := json.Marshal(updateTeamRequestDto)
	return byteValueOfStruct
}

type TeamTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *TeamTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
