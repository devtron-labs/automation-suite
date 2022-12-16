package UserTerminalAccessRouter

import (
	"automation-suite/K8sCapacityRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTerminalAccessRoutersTestSuite) TestUpdateTerminalShellSessionApi() {
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
	log.Println("Here we are hitting fetch terminal session API")
	queryParams := map[string]string{"terminalAccessId": strconv.Itoa(TerminalAccessId)}
	TerminalStatus := HitFetchTerminalStatusApi(queryParams, suite.authToken)

	suite.Run("A=1=UpdateTerminalShellSessionStatusWithValidTerminalAccessId", func() {
		UpdateTerminalShellSessionRequestDto := CreatePayLoadForUpdateTerminalShellSession(TerminalAccessId, "bash")
		UpdateTerminalShellSessionRequestPayload, _ := json.Marshal(UpdateTerminalShellSessionRequestDto)
		UpdatedTerminalShellSession := HitUpdateTerminalShellSessionApi(UpdateTerminalShellSessionRequestPayload, suite.authToken)
		assert.Equal(suite.T(), TerminalStatus.Result.UserTerminalSessionId, UpdatedTerminalShellSession.Result.UserTerminalSessionId)
		assert.Equal(suite.T(), TerminalAccessId, UpdatedTerminalShellSession.Result.TerminalAccessId)
		assert.Equal(suite.T(), TerminalStatus.Result.PodName, UpdatedTerminalShellSession.Result.PodName)
	})

	suite.Run("A=2=UpdateTerminalShellSessionStatusWithInvalidTerminalAccessId", func() {
		UpdateTerminalShellSessionRequestDto := CreatePayLoadForUpdateTerminalShellSession(Base.GetRandomNumberOf9Digit(), "bash")
		UpdateTerminalShellSessionRequestPayload, _ := json.Marshal(UpdateTerminalShellSessionRequestDto)
		UpdatedTerminalShellSession := HitUpdateTerminalShellSessionApi(UpdateTerminalShellSessionRequestPayload, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", UpdatedTerminalShellSession.Errors[0].UserMessage)
	})
}
