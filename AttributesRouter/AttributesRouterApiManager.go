package AttributesRouter

import (
	"automation-suite/AttributesRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructAttributesRouter struct {
	getAttributesRespDto ResponseDTOs.GetAttributesResponseDTO
}

func HitGetAttributesApi(queryParams map[string]string, authToken string) ResponseDTOs.GetAttributesResponseDTO {
	resp, err := Base.MakeApiCall(GetAttributesApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetAttributesApi)
	structAttributesRouter := StructAttributesRouter{}
	chartRepoRouter := structAttributesRouter.UnmarshalGivenResponseBody(resp.Body(), GetAttributesApi)
	return chartRepoRouter.getAttributesRespDto
}

func HitAddAttributesApi(payloadOfApi []byte, authToken string) ResponseDTOs.GetAttributesResponseDTO {
	resp, err := Base.MakeApiCall(AddAttributesApiUrl, http.MethodPost, string(payloadOfApi), nil, authToken)
	Base.HandleError(err, AddAttributesApi)
	structAttributesRouter := StructAttributesRouter{}
	chartRepoRouter := structAttributesRouter.UnmarshalGivenResponseBody(resp.Body(), GetAttributesApi)
	return chartRepoRouter.getAttributesRespDto
}

func GetPayloadForAddAttributes(value string) ResponseDTOs.AttributesDTO {
	var attributesDTO ResponseDTOs.AttributesDTO
	attributesDTO.Key = "url"
	attributesDTO.Value = value
	attributesDTO.Active = true
	return attributesDTO
}

func (structAttributesRouter StructAttributesRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAttributesRouter {
	switch apiName {
	case GetAttributesApi:
		json.Unmarshal(response, &structAttributesRouter.getAttributesRespDto)
	}
	return structAttributesRouter
}

type AttributeRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *AttributeRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
