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
