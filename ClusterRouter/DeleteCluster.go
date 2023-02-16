package ClusterRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *ClustersRouterTestSuite) TestDeleteCluster() {
	envConf := Base.ReadBaseEnvConfig()
	file := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)
	clusterName := "automation-cluster" + strings.ToLower(Base.GetRandomStringOfGivenLength(7))

	suite.Run("A=1=DeleteClusterWithValidPayload", func() {
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(0, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		saveClusterResponse := HitSaveClusterApi(byteValueOfStruct, suite.authToken)
		log.Println("=== Here we are getting cluster via id ===")
		queryParams := map[string]string{"id": strconv.Itoa(saveClusterResponse.Result.Id)}
		getClusterResponse := HitGetClusterByIdApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), clusterName, getClusterResponse.Result.ClusterName)
		log.Println("=== Here we are deleting the cluster after verification ===")
		requestPayload = GetRequestPayloadForSaveOrDeleteCluster(saveClusterResponse.Result.Id, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ = json.Marshal(requestPayload)
		deleteClusterResponse := HitDeleteClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Cluster deleted successfully.", deleteClusterResponse.Result)
		log.Println("=== Here we are trying to get the cluster after Deletion ===")
		getClusterResponse = HitGetClusterByIdApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), 404, getClusterResponse.Code)
	})

	suite.Run("A=2=DeleteClusterWithInvalidIdInPayload", func() {
		log.Println("=== Here we are deleting the cluster after verification ===")
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(Base.GetRandomNumberOf9Digit(), clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		deleteClusterResponse := HitDeleteClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", deleteClusterResponse.Errors[0].UserMessage)
	})

	/*suite.Run("A=3=DeleteClusterWithInvalidIdTokenAndServerUrl", func() {
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(0, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		saveClusterResponse := HitSaveClusterApi(byteValueOfStruct, suite.authToken)
		log.Println("=== Here we are deleting the cluster after verification ===")
		requestPayload = GetRequestPayloadForSaveOrDeleteCluster(saveClusterResponse.Result.Id, clusterName, "file.ClusterBearerToken", file.ClusterServerUrl)
		byteValueOfStruct, _ = json.Marshal(requestPayload)
		deleteClusterResponse := HitDeleteClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", deleteClusterResponse.Errors[0].UserMessage)

		log.Println("=== Here we are deleting the cluster after verification ===")
		requestPayload = GetRequestPayloadForSaveOrDeleteCluster(saveClusterResponse.Result.Id, clusterName, file.ClusterBearerToken, "file.ClusterServerUrl")
		byteValueOfStruct, _ = json.Marshal(requestPayload)
		deleteClusterResponse = HitDeleteClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", deleteClusterResponse.Errors[0].UserMessage)
	})*/
}

//todo there is no validation for Bearer Token and Server Url While we are passing in deleteCluster request payload
