package SSOLoginRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type GetListResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Url    string `json:"url"`
		Active bool   `json:"active"`
		Label  string `json:"label,omitempty"`
	} `json:"result"`
}

type GetSSODetailsResponse struct {
	Code                       int                         `json:"code"`
	Status                     string                      `json:"status"`
	CreateSSODetailsRequestDto *CreateSSODetailsRequestDto `json:"result"`
}

type CreateSSODetailsRequestDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Config struct {
		Id     string `json:"id"`
		Label  string `json:"label"`
		Type   string `json:"type"`
		Name   string `json:"name"`
		Config struct {
			Issuer        string   `json:"issuer"`
			ClientID      string   `json:"clientID"`
			ClientSecret  string   `json:"clientSecret"`
			RedirectURI   string   `json:"redirectURI"`
			HostedDomains []string `json:"hostedDomains"`
		} `json:"config"`
	} `json:"config"`
	Active bool `json:"active"`
}

type StructSSOLoginRouter struct {
	getListResponseDto         GetListResponseDto
	getSSODetailsResponse      GetSSODetailsResponse
	createSSODetailsRequestDto CreateSSODetailsRequestDto
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

func HitGetListApi(authToken string) GetListResponseDto {
	resp, err := Base.MakeApiCall(GetListApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetListApi)

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), GetListApi)
	return ssoRouter.getListResponseDto
}

func HitGetSSODetailsApi(ssoDetailsId string, authToken string) GetSSODetailsResponse {
	resp, err := Base.MakeApiCall(GetSSOConfigByNameApiUrl+"/"+ssoDetailsId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetSsoDetailsApi)

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), GetSsoDetailsApi)
	return ssoRouter.getSSODetailsResponse
}

func HitGetLoginConfigByNameApi(queryParams map[string]string, authToken string) GetSSODetailsResponse {
	resp, err := Base.MakeApiCall(GetSSOConfigByNameApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetSSOConfigByName)

	structSSOLoginRouter := StructSSOLoginRouter{}
	ssoRouter := structSSOLoginRouter.UnmarshalGivenResponseBody(resp.Body(), GetSSOConfigByName)
	return ssoRouter.getSSODetailsResponse
}

func HitUpdateSSODetailsApi(byteValue []byte, authToken string) GetSSODetailsResponse {
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
