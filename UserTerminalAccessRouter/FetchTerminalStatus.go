package UserTerminalAccessRouter

import (
	"automation-suite/K8sCapacityRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *UserTerminalAccessRoutersTestSuite) TestFetchTerminalSession() {
	log.Println("Here we are fetching existing cluster List")
	var clusterId int
	var nodeName string
	var TerminalAccessId int
	ClusterList := K8sCapacityRouter.HitGetClusterListApi(suite.authToken)
	for _, cluster := range ClusterList.Result {
		if cluster.Name == "default_cluster" {
			clusterId = cluster.Id
			nodeName = cluster.NodeNames[0]
			break
		}
	}
	log.Println("Here we are starting terminal session")
	StartTerminalSessionDTO := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", nodeName, "default", "sh")
	StartTerminalSessionJson, _ := json.Marshal(StartTerminalSessionDTO)
	StartTerminalSessionApiResponse := HitStartTerminalSessionApi(StartTerminalSessionJson, suite.authToken)
	TerminalAccessId = StartTerminalSessionApiResponse.Result.TerminalAccessId

	suite.Run("A=1=FetchTerminalStatusWithValidTerminalAccessId", func() {
		queryParams := map[string]string{"terminalAccessId": strconv.Itoa(TerminalAccessId)}
		TerminalStatus := HitFetchTerminalStatusApi(queryParams, suite.authToken)
		assert.True(suite.T(), TerminalStatus.Result.UserId >= 1)
		assert.NotEmpty(suite.T(), TerminalStatus.Result.Status)
		assert.NotEmpty(suite.T(), TerminalStatus.Result.PodName)
		if TerminalStatus.Result.Status == "Starting" {
			time.Sleep(10 * time.Second)
			TerminalStatus = HitFetchTerminalStatusApi(queryParams, suite.authToken)
			assert.Equal(suite.T(), "Running", TerminalStatus.Result.Status)
			assert.NotEmpty(suite.T(), TerminalStatus.Result.UserTerminalSessionId)
		}
		assert.Equal(suite.T(), "Running", TerminalStatus.Result.Status)
	})

	suite.Run("A=2=FetchTerminalStatusWithInvalidTerminalAccessId", func() {
		queryParams := map[string]string{"terminalAccessId": strconv.Itoa(Base.GetRandomNumberOf9Digit())}
		TerminalStatus := HitFetchTerminalStatusApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), 404, TerminalStatus.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", TerminalStatus.Errors[0].UserMessage)
	})
}
