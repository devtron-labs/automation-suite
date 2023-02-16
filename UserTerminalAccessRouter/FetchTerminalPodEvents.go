package UserTerminalAccessRouter

import (
	"automation-suite/K8sCapacityRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTerminalAccessRoutersTestSuite) TestFetchTerminalPodEvents() {
	log.Println("Here we are fetching existing cluster List")
	var clusterId int
	var nodeName string
	var TerminalAccessId int
	ClusterList := K8sCapacityRouter.HitGetClusterListApi(suite.authToken)
	for _, cluster := range ClusterList.Result {
		if cluster.Name == "default_cluster" {
			clusterId = cluster.Id
			nodeName = cluster.NodeNames[1]
			break
		}
	}
	log.Println("Here we are starting terminal session")
	StartTerminalSessionDTO := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", nodeName, "default", "sh")
	StartTerminalSessionJson, _ := json.Marshal(StartTerminalSessionDTO)
	StartTerminalSessionApiResponse := HitStartTerminalSessionApi(StartTerminalSessionJson, suite.authToken)
	TerminalAccessId = StartTerminalSessionApiResponse.Result.TerminalAccessId

	suite.Run("A=1=FetchPodEventsWithValidTerminalAccessId", func() {
		var isExpectedEventPresent bool
		queryParams := map[string]string{"terminalAccessId": strconv.Itoa(TerminalAccessId)}
		TerminalPodEvents := HitFetchTerminalPodEventsApi(queryParams, suite.authToken)
		for _, PodEventItems := range TerminalPodEvents.Result.Events.Items {
			if PodEventItems.Reason == "Scheduled" {
				assert.Contains(suite.T(), PodEventItems.Message, "Successfully assigned")
				isExpectedEventPresent = true
			}
		}
		assert.True(suite.T(), isExpectedEventPresent)
	})

	suite.Run("A=2=FetchPodEventsWithInvalidTerminalAccessId", func() {
		queryParams := map[string]string{"terminalAccessId": strconv.Itoa(Base.GetRandomNumberOf9Digit())}
		TerminalPodEvents := HitFetchTerminalPodEventsApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), 500, TerminalPodEvents.Code)
		assert.Equal(suite.T(), "unable to fetch pod event", TerminalPodEvents.Errors[0].UserMessage)
	})

	suite.Run("A=3=FetchPodEventsWithInValidNode", func() {
		log.Println("Here we are starting terminal session with invalid Node Name")
		StartTerminalSessionObj := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", "InvalidNodeName", "default", "sh")
		StartTerminalSessionPayload, _ := json.Marshal(StartTerminalSessionObj)
		ApiResponse := HitStartTerminalSessionApi(StartTerminalSessionPayload, suite.authToken)
		TerminalAccessId = ApiResponse.Result.TerminalAccessId
		log.Println("Here we are trying to fetch Events for Pod")
		var isExpectedEventPresent bool
		queryParams := map[string]string{"terminalAccessId": strconv.Itoa(TerminalAccessId)}
		TerminalPodEvents := HitFetchTerminalPodEventsApi(queryParams, suite.authToken)
		for _, PodEventItems := range TerminalPodEvents.Result.Events.Items {
			if PodEventItems.Reason == "FailedScheduling" {
				assert.Contains(suite.T(), PodEventItems.Message, "pod didn't tolerate")
				isExpectedEventPresent = true
			}
		}
		assert.True(suite.T(), isExpectedEventPresent)
	})
}
