package ClusterRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *ClustersRouterTestSuite) TestDeleteCluster() {
	envConf := Base.ReadBaseEnvConfig()
	file := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)

	suite.Run("A=1=DeleteClusterWithValidPayload", func() {
		clusterName := "automation-cluster" + strings.ToLower(Base.GetRandomStringOfGivenLength(7))
		requestPayload := GetRequestPayloadForSaveOrDeleteCluster(0, clusterName, file.ClusterBearerToken, file.ClusterServerUrl)
		byteValueOfStruct, _ := json.Marshal(requestPayload)
		log.Println("Hitting The Save Cluster API")
		saveClusterResponse := HitSaveClusterApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), clusterName, saveClusterResponse.Result.ClusterName)
		assert.Equal(suite.T(), file.ClusterBearerToken, saveClusterResponse.Result.Config.BearerToken)
		assert.Equal(suite.T(), file.ClusterServerUrl, saveClusterResponse.Result.ServerUrl)
	})
}
