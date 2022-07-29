package ClusterRouter

import (
	"automation-suite/ClusterRouter/RequestDTOs"
	"automation-suite/ClusterRouter/ResponseDTOs"

	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
)

type StructClusterRouter struct {
	saveClusterResponseDTO   ResponseDTOs.SaveClusterResponseDTO
	deleteClusterResponseDTO ResponseDTOs.DeleteClusterResponseDTO
	clusterByIdResponseDTO   ResponseDTOs.ClusterByIdResponseDTO
	findAllForAutocomplete   ResponseDTOs.FindAllForAutocomplete
}

func HitSaveClusterApi(payload []byte, authToken string) ResponseDTOs.SaveClusterResponseDTO {
	resp, err := Base.MakeApiCall(SaveClusterApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, SaveClusterApi)
	structClusterRouter := StructClusterRouter{}
	apiRouter := structClusterRouter.UnmarshalGivenResponseBody(resp.Body(), SaveClusterApi)
	return apiRouter.saveClusterResponseDTO
}

func HitDeleteClusterApi(payload []byte, authToken string) ResponseDTOs.DeleteClusterResponseDTO {
	resp, err := Base.MakeApiCall(SaveClusterApiUrl, http.MethodDelete, string(payload), nil, authToken)
	Base.HandleError(err, DeleteClusterApi)
	structClusterRouter := StructClusterRouter{}
	apiRouter := structClusterRouter.UnmarshalGivenResponseBody(resp.Body(), DeleteClusterApi)
	return apiRouter.deleteClusterResponseDTO
}

func HitGetClusterByIdApi(queryParams map[string]string, authToken string) ResponseDTOs.ClusterByIdResponseDTO {
	resp, err := Base.MakeApiCall(SaveClusterApiUrl, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetClusterById)
	structClusterRouter := StructClusterRouter{}
	apiRouter := structClusterRouter.UnmarshalGivenResponseBody(resp.Body(), GetClusterById)
	return apiRouter.clusterByIdResponseDTO
}

func HitFindAllClusterForAutocomplete(authToken string) ResponseDTOs.FindAllForAutocomplete {
	resp, err := Base.MakeApiCall(SaveClusterApiUrl+"/autocomplete", http.MethodGet, "", nil, authToken)
	Base.HandleError(err, FindAllForAutocompleteApi)
	structClusterRouter := StructClusterRouter{}
	apiRouter := structClusterRouter.UnmarshalGivenResponseBody(resp.Body(), FindAllForAutocompleteApi)
	return apiRouter.findAllForAutocomplete
}

func (structClusterRouter StructClusterRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructClusterRouter {
	switch apiName {
	case SaveClusterApi:
		json.Unmarshal(response, &structClusterRouter.saveClusterResponseDTO)
	case DeleteClusterApi:
		json.Unmarshal(response, &structClusterRouter.deleteClusterResponseDTO)
	case GetClusterById:
		json.Unmarshal(response, &structClusterRouter.clusterByIdResponseDTO)
	case FindAllForAutocompleteApi:
		json.Unmarshal(response, &structClusterRouter.findAllForAutocomplete)
	}
	return structClusterRouter
}

func GetRequestPayloadForSaveOrDeleteCluster(clusterId int, clusterName string, bearerToken string, serverUrl string) RequestDTOs.SaveClusterRequestDTO {
	var saveClusterRequestDto RequestDTOs.SaveClusterRequestDTO
	saveClusterRequestDto.ClusterName = clusterName
	saveClusterRequestDto.Config.BearerToken = bearerToken
	saveClusterRequestDto.ServerUrl = serverUrl
	saveClusterRequestDto.Active = true
	if clusterId != 0 {
		saveClusterRequestDto.Id = clusterId
	}
	return saveClusterRequestDto
}

type ClustersRouterTestSuite struct {
	suite.Suite
	authToken string
}

// SetupSuite This method runs on first priority before starting the suite means before executing any test case of the suite
func (suite *ClustersRouterTestSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
}
