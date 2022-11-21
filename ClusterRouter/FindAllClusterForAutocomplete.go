package ClusterRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *ClustersRouterTestSuite) TestFindAllClusterForAutocomplete() {
	envConf := Base.ReadBaseEnvConfig()
	file := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)
	clusterName := "automation-cluster" + strings.ToLower(Base.GetRandomStringOfGivenLength(7))
	suite.Run("A=1=GetClusterByValidId", func() {
		log.Println("=== Here we are getting no of clusters via autocomplete API before adding new  ===")
		getAllClustersList := HitFindAllClusterForAutocomplete(suite.authToken)
		noOfCluster := len(getAllClustersList.Result)
		log.Println("=== Here we are saving new cluster")
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(0, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		saveClusterResponse := HitSaveClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), clusterName, saveClusterResponse.Result.ClusterName)
		log.Println("=== Here we are getting no of clusters via autocomplete API after adding new  ===")
		getAllClustersList = HitFindAllClusterForAutocomplete(suite.authToken)
		assert.Equal(suite.T(), noOfCluster+1, len(getAllClustersList.Result))
		log.Println("=== Here we are deleting the cluster after verification ===")
		requestPayload = GetRequestPayloadForSaveOrDeleteCluster(saveClusterResponse.Result.Id, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ = json.Marshal(requestPayload)
		deleteClusterResponse := HitDeleteClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Cluster deleted successfully.", deleteClusterResponse.Result)
	})
}
