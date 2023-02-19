package ClusterRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *ClustersRouterTestSuite) TestSaveCluster() {
	envConf := Base.ReadBaseEnvConfig()
	file := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)
	clusterName := "automation-cluster" + strings.ToLower(Base.GetRandomStringOfGivenLength(7))
	suite.Run("A=1=SaveClusterWithValidCredentials", func() {
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(0, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		log.Println("=== Hitting The Save Cluster API ===")
		saveClusterResponse := HitSaveClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), clusterName, saveClusterResponse.Result.ClusterName)
		assert.Equal(suite.T(), file.ClusterBearerToken, saveClusterResponse.Result.Config.BearerToken)
		assert.Equal(suite.T(), file.ClusterServerUrl, saveClusterResponse.Result.ServerUrl)
		queryParams := map[string]string{"id": strconv.Itoa(saveClusterResponse.Result.Id)}
		getClusterResponse := HitGetClusterByIdApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), clusterName, getClusterResponse.Result.ClusterName)
		log.Println("=== Here we are deleting the cluster after verification ===")
		requestPayload = GetRequestPayloadForSaveOrDeleteCluster(saveClusterResponse.Result.Id, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ = json.Marshal(requestPayload)
		deleteClusterResponse := HitDeleteClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Cluster deleted successfully.", deleteClusterResponse.Result)
	})

	suite.Run("A=2=SaveClusterWithInvalidServerUrl", func() {
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(0, clusterName, file.ClusterBearerToken, "https://invalid-url.com")
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		log.Println("=== Hitting The Save Cluster API ===")
		saveClusterResponse := HitSaveClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), 500, saveClusterResponse.Code)
		assert.Equal(suite.T(), "unable to parse the server version: invalid character '\u003c' looking for beginning of value", saveClusterResponse.Errors[0].UserMessage)
	})

	suite.Run("A=3=SaveClusterWithInvalidBearerToken", func() {
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(0, clusterName, "file.ClusterBearerToken", file.ClusterServerUrl)
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		log.Println("=== Hitting The Save Cluster API ===")
		saveClusterResponse := HitSaveClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), 500, saveClusterResponse.Code)
		assert.Equal(suite.T(), "the server has asked for the client to provide credentials", saveClusterResponse.Errors[0].UserMessage)
	})

}

//todo need to remove status code check once dev will fix it
//todo need to add other advance test cases as well with Prometheus Details and metrics
