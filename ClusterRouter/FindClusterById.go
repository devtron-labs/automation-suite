package ClusterRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *ClustersRouterTestSuite) TestFindById() {
	envConf := Base.ReadBaseEnvConfig()
	file := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)
	clusterName := "automation-cluster" + strings.ToLower(Base.GetRandomStringOfGivenLength(7))
	suite.Run("A=1=GetClusterByValidId", func() {
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(0, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		saveClusterResponse := HitSaveClusterApi(byteValueOfStruct, suite.authToken)
		log.Println("=== Here we are getting cluster via id ===")
		queryParams := map[string]string{"id": strconv.Itoa(saveClusterResponse.Result.Id)}
		getClusterResponse := HitGetClusterByIdApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), clusterName, getClusterResponse.Result.ClusterName)
		assert.Equal(suite.T(), file.ClusterBearerToken, getClusterResponse.Result.Config.BearerToken)
		assert.Equal(suite.T(), file.ClusterServerUrl, getClusterResponse.Result.ServerUrl)
		log.Println("=== Here we are deleting the cluster after verification ===")
		requestPayload = GetRequestPayloadForSaveOrDeleteCluster(saveClusterResponse.Result.Id, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ = json.Marshal(requestPayload)
		deleteClusterResponse := HitDeleteClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Cluster deleted successfully.", deleteClusterResponse.Result)
	})

	suite.Run("A=2=GetClusterByInvalidId", func() {
		randomId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		queryParams := map[string]string{"id": randomId}
		getClusterResponse := HitGetClusterByIdApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), 404, getClusterResponse.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", getClusterResponse.Errors[0].UserMessage)
	})

}
