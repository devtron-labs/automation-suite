package UserTerminalAccessRouter

import (
	"automation-suite/K8sCapacityRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func (suite *UserTerminalAccessRoutersTestSuite) TestStartTerminalSession() {
	var clusterId int
	var nodeName string
	ClusterList := K8sCapacityRouter.HitGetClusterListApi(suite.authToken)
	for _, cluster := range ClusterList.Result {
		if cluster.Name == "default_cluster" {
			clusterId = cluster.Id
			nodeName = cluster.NodeNames[0]
			break
		}
	}

	suite.Run("A=1=VerifyStartTerminalSessionApiWithValidPayload", func() {
		StartTerminalSessionDTO := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", nodeName, "default", "sh")
		StartTerminalSessionJson, _ := json.Marshal(StartTerminalSessionDTO)
		StartTerminalSessionApiResponse := HitStartTerminalSessionApi(StartTerminalSessionJson, suite.authToken)
		assert.True(suite.T(), StartTerminalSessionApiResponse.Result.UserId >= 1)
		assert.True(suite.T(), StartTerminalSessionApiResponse.Result.TerminalAccessId >= 1)
		assert.NotEmpty(suite.T(), StartTerminalSessionApiResponse.Result.PodName)
	})

	suite.Run("A=2=VerifyStartTerminalSessionApiWithInvalidClusterId", func() {
		randomClusterId := Base.GetRandomNumberOf9Digit()
		StartTerminalSessionDTO := CreatePayLoadForStartTerminalSession(randomClusterId, "ubuntu:latest", nodeName, "default", "sh")
		StartTerminalSessionJson, _ := json.Marshal(StartTerminalSessionDTO)
		StartTerminalSessionApiResponse := HitStartTerminalSessionApi(StartTerminalSessionJson, suite.authToken)
		assert.Equal(suite.T(), "ERROR #23503 insert or update on table \"user_terminal_access_data\" violates foreign key constraint \"user_terminal_access_data_cluster_id_fkey\"", StartTerminalSessionApiResponse.Errors[0].UserMessage)
	})

	suite.Run("A=3=VerifyStartTerminalSessionApiWithInvalidShellName", func() {
		StartTerminalSessionDTO := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", nodeName, "default", "invalidShellName")
		StartTerminalSessionJson, _ := json.Marshal(StartTerminalSessionDTO)
		StartTerminalSessionApiResponse := HitStartTerminalSessionApi(StartTerminalSessionJson, suite.authToken)
		assert.Equal(suite.T(), "ERROR #23503 insert or update on table \"user_terminal_access_data\" violates foreign key constraint \"user_terminal_access_data_cluster_id_fkey\"", StartTerminalSessionApiResponse.Errors[0].UserMessage)
	})

	suite.Run("A=4=VerifyStartTerminalSessionApiWithInvalidNodeName", func() {
		StartTerminalSessionDTO := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", "invalidNodeName", "default", "bash")
		StartTerminalSessionJson, _ := json.Marshal(StartTerminalSessionDTO)
		StartTerminalSessionApiResponse := HitStartTerminalSessionApi(StartTerminalSessionJson, suite.authToken)
		assert.Equal(suite.T(), "ERROR #23503 insert or update on table \"user_terminal_access_data\" violates foreign key constraint \"user_terminal_access_data_cluster_id_fkey\"", StartTerminalSessionApiResponse.Errors[0].UserMessage)
	})

	suite.Run("A=5=VerifyStartTerminalSessionApiWithInvalidNamespace", func() {
		StartTerminalSessionDTO := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", nodeName, "InvalidNameSpace", "sh")
		StartTerminalSessionJson, _ := json.Marshal(StartTerminalSessionDTO)
		StartTerminalSessionApiResponse := HitStartTerminalSessionApi(StartTerminalSessionJson, suite.authToken)
		assert.Equal(suite.T(), "namespaces \"InvalidNameSpace\" not found", StartTerminalSessionApiResponse.Errors[0].UserMessage)
	})

	//todo need to add assertion with the help of getEventApi
	suite.Run("A=6=VerifyStartTerminalSessionApiWithInvalidBaseImage", func() {
		StartTerminalSessionDTO := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", nodeName, "default", "invalidShellName")
		StartTerminalSessionJson, _ := json.Marshal(StartTerminalSessionDTO)
		StartTerminalSessionApiResponse := HitStartTerminalSessionApi(StartTerminalSessionJson, suite.authToken)
		assert.Equal(suite.T(), "ERROR #23503 insert or update on table \"user_terminal_access_data\" violates foreign key constraint \"user_terminal_access_data_cluster_id_fkey\"", StartTerminalSessionApiResponse.Errors[0].UserMessage)
	})
}
