package K8sCapacityRouter

import (
	"automation-suite/K8sCapacityRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructK8sCapacityRouter struct {
	getClusterListResponseDTO ResponseDTOs.GetClusterListResponseDTO
}

func HitGetClusterListApi(authToken string) ResponseDTOs.GetClusterListResponseDTO {
	resp, err := Base.MakeApiCall(K8sCapacityRoutersBaseUrl+GetClusterListApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, GetClusterListApi)
	structK8sCapacityRouter := StructK8sCapacityRouter{}
	gitopsConfigRouter := structK8sCapacityRouter.UnmarshalGivenResponseBody(resp.Body(), GetClusterListApi)
	return gitopsConfigRouter.getClusterListResponseDTO
}

func (structK8sCapacityRouter StructK8sCapacityRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructK8sCapacityRouter {
	switch apiName {
	case GetClusterListApi:
		json.Unmarshal(response, &structK8sCapacityRouter.getClusterListResponseDTO)
	}
	return structK8sCapacityRouter
}

type K8sCapacityRoutersTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *K8sCapacityRoutersTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
