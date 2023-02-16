package SSOLoginRouter

import (
	"automation-suite/SSOLoginRouter/RequestDTOs"
	"automation-suite/SSOLoginRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructSSOLoginRouter struct {
	getListResponseDto         ResponseDTOs.GetListResponseDTO
	getSSODetailsResponse      ResponseDTOs.GetSSODetailsResponseDTO
	createSSODetailsRequestDto RequestDTOs.CreateSSODetailsRequestDTO
}

func (structSSOLoginRouter StructSSOLoginRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructSSOLoginRouter {
	switch apiName {
	case GetListApi:
		json.Unmarshal(response, &structSSOLoginRouter.getListResponseDto)
	case GetSsoDetailsApi:
		json.Unmarshal(response, &structSSOLoginRouter.getSSODetailsResponse)
	case GetSSOConfigByName:
		json.Unmarshal(response, &structSSOLoginRouter.getSSODetailsResponse)
	case UpdateSSODetailsApi:
		json.Unmarshal(response, &structSSOLoginRouter.getSSODetailsResponse)
	}
	return structSSOLoginRouter
}

func HitGetListApi(authToken string) ResponseDTOs.GetListResponseDTO {
	resp, err := Base.MakeApiCall(GetListApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetListApi)

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), GetListApi)
	return ssoRouter.getListResponseDto
}

func HitGetSSODetailsApi(ssoDetailsId string, authToken string) ResponseDTOs.GetSSODetailsResponseDTO {
	resp, err := Base.MakeApiCall(GetSSOConfigByNameApiUrl+"/"+ssoDetailsId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetSsoDetailsApi)

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), GetSsoDetailsApi)
	return ssoRouter.getSSODetailsResponse
}

func HitGetLoginConfigByNameApi(queryParams map[string]string, authToken string) ResponseDTOs.GetSSODetailsResponseDTO {
	resp, err := Base.MakeApiCall(GetSSOConfigByNameApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetSSOConfigByName)

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), GetSSOConfigByName)
	return ssoRouter.getSSODetailsResponse
}

func HitUpdateSSODetailsApi(byteValue []byte, authToken string) ResponseDTOs.GetSSODetailsResponseDTO {
	resp, err := Base.MakeApiCall(UpdateSSODetailsApiUrl, http.MethodPut, string(byteValue), nil, authToken)
	Base.HandleError(err, UpdateSSODetailsApi)

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), UpdateSSODetailsApi)
	return ssoRouter.getSSODetailsResponse
}

type SSOLoginTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *SSOLoginTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
