package AppStoreDiscoverRouter

import (
	"automation-suite/AppStoreDiscoverRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructAppStoreDiscoverRouter struct {
	discoverAppApiResponse ResponseDTOs.DiscoverAppApiResponse
}

func HitDiscoverAppApi(queryParams map[string]string, authToken string) ResponseDTOs.DiscoverAppApiResponse {
	resp, err := Base.MakeApiCall(DiscoverAppApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, DiscoverAppApi)
	structAppStoreDiscoverRouter := StructAppStoreDiscoverRouter{}
	appStoreDiscoverRouter := structAppStoreDiscoverRouter.UnmarshalGivenResponseBody(resp.Body(), DiscoverAppApi)
	return appStoreDiscoverRouter.discoverAppApiResponse
}

func (structAppStoreDiscoverRouter StructAppStoreDiscoverRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppStoreDiscoverRouter {
	switch apiName {
	case DiscoverAppApi:
		json.Unmarshal(response, &structAppStoreDiscoverRouter.discoverAppApiResponse)
	}
	return structAppStoreDiscoverRouter
}

type AppStoreDiscoverTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *AppStoreDiscoverTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
