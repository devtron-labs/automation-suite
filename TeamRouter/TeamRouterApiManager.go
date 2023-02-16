package TeamRouter

import (
	"automation-suite/TeamRouter/RequestDTOs"
	"automation-suite/TeamRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type TeamsRouterStruct struct {
	saveTeamResponseDto     ResponseDTOs.SaveTeamResponseDTO
	deleteTeamResponseDto   ResponseDTOs.DeleteTeamResponseDto
	fetchAllTeamResponseDto ResponseDTOs.FetchAllTeamResponseDTO
	getTeamByIdResponseDto  ResponseDTOs.GetTeamByIdResponseDTO
}

func (teamRouterStruct TeamsRouterStruct) UnmarshalGivenResponseBody(response []byte, apiName string) TeamsRouterStruct {
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

func HitSaveTeamApi(payload []byte, authToken string) ResponseDTOs.SaveTeamResponseDTO {
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

	teamRouterStruct := TeamsRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), SaveTeamApi)
	return teamRouter.saveTeamResponseDto
}

func GetSaveTeamRequestDto() RequestDTOs.SaveTeamRequestDto {
	var saveTeamRequestDto RequestDTOs.SaveTeamRequestDto
	teamName := Base.GetRandomStringOfGivenLength(10)
	saveTeamRequestDto.Name = teamName
	saveTeamRequestDto.Active = true
	return saveTeamRequestDto
}

func HitFetchAllTeamApi(authToken string) ResponseDTOs.FetchAllTeamResponseDTO {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchAllTeamApi)

	teamRouterStruct := TeamsRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllTeamApi)
	return teamRouter.fetchAllTeamResponseDto
}

func HitDeleteTeamApi(byteValueOfStruct []byte, authToken string) ResponseDTOs.DeleteTeamResponseDto {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodDelete, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, DeleteTeamApi)

	teamRouterStruct := TeamsRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), DeleteTeamApi)
	return teamRouter.deleteTeamResponseDto
}

func HitGetTeamByIdApi(id string, authToken string) ResponseDTOs.GetTeamByIdResponseDTO {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl+"/"+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetTeamByIdApi)

	teamRouterStruct := TeamsRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), GetTeamByIdApi)
	return teamRouter.getTeamByIdResponseDto
}

func HitUpdateTeamApi(byteValueOfStruct []byte, authToken string) ResponseDTOs.SaveTeamResponseDTO {
	resp, err := Base.MakeApiCall(SaveTeamApiUrl, http.MethodPut, string(byteValueOfStruct), nil, authToken)
	Base.HandleError(err, UpdateTeamApi)

	teamRouterStruct := TeamsRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), SaveTeamApi)
	return teamRouter.saveTeamResponseDto
}

func HitFetchForAutocompleteApi(authToken string) ResponseDTOs.FetchAllTeamResponseDTO {
	resp, err := Base.MakeApiCall(FetchTeamAutocompleteApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FetchForAutocompleteApi)

	teamRouterStruct := TeamsRouterStruct{}
	teamRouter := teamRouterStruct.UnmarshalGivenResponseBody(resp.Body(), FetchAllTeamApi)
	return teamRouter.fetchAllTeamResponseDto
}

func GetPayLoadForDeleteAPI(id int, name string, isActive bool) []byte {
	var updateTeamDto RequestDTOs.SaveTeamRequestDto
	updateTeamDto.Id = id
	updateTeamDto.Name = name
	updateTeamDto.Active = isActive
	byteValueOfStruct, _ := json.Marshal(updateTeamDto)
	return byteValueOfStruct
}

func GetUpdateTeamRequestPayload(id int, teamName string) []byte {
	var updateTeamRequestDto RequestDTOs.SaveTeamRequestDto
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
