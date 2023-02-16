package CommonRouter

import (
	"automation-suite/CommonRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
)

type StructCommonRouter struct {
	globalChecklistResponseDTO ResponseDTOs.GlobalChecklistResponseDTO
}

func HitGlobalChecklistApi(authToken string) ResponseDTOs.GlobalChecklistResponseDTO {
	resp, err := Base.MakeApiCall(GlobalChecklistApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GlobalChecklistApi)
	structCommonRouter := StructCommonRouter{}
	apiRouter := structCommonRouter.UnmarshalGivenResponseBody(resp.Body(), GlobalChecklistApi)
	return apiRouter.globalChecklistResponseDTO
}
func (structCommonRouter StructCommonRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructCommonRouter {
	switch apiName {
	case GlobalChecklistApi:
		json.Unmarshal(response, &structCommonRouter.globalChecklistResponseDTO)
	}
	return structCommonRouter
}

type BaseCommonRouterTestSuite struct {
	suite.Suite
	authToken string
}

// SetupSuite This method runs on first priority before starting the suite means before executing any test case of the suite
func (suite *BaseCommonRouterTestSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
}
