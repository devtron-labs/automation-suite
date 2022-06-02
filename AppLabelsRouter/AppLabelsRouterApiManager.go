package AppLabelsRouter

import (
	"automation-suite/AppLabelsRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructAppLabelsRouter struct {
	appMetaInfoResponseDto ResponseDTOs.AppMetaInfoResponseDto
}

func HitGetAppMetaInfoByIdApi(appId string, authToken string) ResponseDTOs.AppMetaInfoResponseDto {
	resp, err := Base.MakeApiCall(GetAppMetaInfoByIdApiUrl+appId, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetAppMetaInfoByIdApi)
	structAppLabelsRouter := StructAppLabelsRouter{}
	appLabelRepoRouter := structAppLabelsRouter.UnmarshalGivenResponseBody(resp.Body(), GetAppMetaInfoByIdApi)
	return appLabelRepoRouter.appMetaInfoResponseDto
}

func (structAppLabelsRouter StructAppLabelsRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructAppLabelsRouter {
	switch apiName {
	case GetAppMetaInfoByIdApi:
		json.Unmarshal(response, &structAppLabelsRouter.appMetaInfoResponseDto)
	}
	return structAppLabelsRouter
}

type AppLabelRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *AppLabelRouterTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
