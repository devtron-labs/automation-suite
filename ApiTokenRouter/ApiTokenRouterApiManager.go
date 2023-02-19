package ApiTokenRouter

import (
	"automation-suite/ApiTokenRouter/RequestDTOs"
	"automation-suite/ApiTokenRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
	"time"
)

type StructApiTokenRouter struct {
	createApiTokenResponseDTO  ResponseDTOs.CreateApiTokenResponseDTO
	getAllApiTokensResponseDTO ResponseDTOs.GetAllApiTokensResponseDTO
	deleteApiTokenResponseDTO  ResponseDTOs.DeleteApiTokenResponseDTO
}

func HitCreateApiTokenApi(payload string, authToken string) ResponseDTOs.CreateApiTokenResponseDTO {
	resp, err := Base.MakeApiCall(ApiTokenRoutersBaseUrl, http.MethodPost, payload, nil, authToken)
	Base.HandleError(err, CreateApiToken)
	structApiTokenRouter := StructApiTokenRouter{}
	apiTokenRepoRouter := structApiTokenRouter.UnmarshalGivenResponseBody(resp.Body(), CreateApiToken)
	return apiTokenRepoRouter.createApiTokenResponseDTO
}

func GetPayLoadForCreateApiToken() RequestDTOs.CreateApiTokenRequestDTO {
	createApiTokenRequestDTO := RequestDTOs.CreateApiTokenRequestDTO{}
	createApiTokenRequestDTO.ExpireAtInMs = getFutureTimestamp(7)
	createApiTokenRequestDTO.Name = "Token" + Base.GetRandomStringOfGivenLength(5)
	createApiTokenRequestDTO.Description = "This is sample Description for Testing via Automation"
	return createApiTokenRequestDTO
}

func getFutureTimestamp(futureDaysFromNow time.Duration) int64 {
	now := time.Now()
	daysFromNow := time.Hour * 24 * futureDaysFromNow
	diff := now.Add(daysFromNow)
	timestamp := diff.Unix() * 1000
	return timestamp
}

func HitGetAllApiTokens(authToken string) ResponseDTOs.GetAllApiTokensResponseDTO {
	resp, err := Base.MakeApiCall(ApiTokenRoutersBaseUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAllApiTokens)
	structApiTokenRouter := StructApiTokenRouter{}
	apiTokenRepoRouter := structApiTokenRouter.UnmarshalGivenResponseBody(resp.Body(), GetAllApiTokens)
	return apiTokenRepoRouter.getAllApiTokensResponseDTO
}

func HitDeleteApiToken(tokenId string, authToken string) ResponseDTOs.DeleteApiTokenResponseDTO {
	resp, err := Base.MakeApiCall(ApiTokenRoutersBaseUrl+"/"+tokenId, http.MethodDelete, "", nil, authToken)
	Base.HandleError(err, DeleteApiToken)
	structApiTokenRouter := StructApiTokenRouter{}
	apiTokenRepoRouter := structApiTokenRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteApiToken)
	return apiTokenRepoRouter.deleteApiTokenResponseDTO
}

func HitUpdateApiToken(tokenId string, authToken string) ResponseDTOs.CreateApiTokenResponseDTO {
	resp, err := Base.MakeApiCall(ApiTokenRoutersBaseUrl+"/"+tokenId, http.MethodPut, "", nil, authToken)
	Base.HandleError(err, UpdateApiToken)
	structApiTokenRouter := StructApiTokenRouter{}
	apiTokenRepoRouter := structApiTokenRouter.UnmarshalGivenResponseBody(resp.Body(), CreateApiToken)
	return apiTokenRepoRouter.createApiTokenResponseDTO
}

func (structApiTokenRouter StructApiTokenRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructApiTokenRouter {
	switch apiName {
	case CreateApiToken:
		json.Unmarshal(response, &structApiTokenRouter.createApiTokenResponseDTO)
	case GetAllApiTokens:
		json.Unmarshal(response, &structApiTokenRouter.getAllApiTokensResponseDTO)
	case DeleteApiToken:
		json.Unmarshal(response, &structApiTokenRouter.deleteApiTokenResponseDTO)
	}
	return structApiTokenRouter
}

type ApiTokenRoutersTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *ApiTokenRoutersTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
