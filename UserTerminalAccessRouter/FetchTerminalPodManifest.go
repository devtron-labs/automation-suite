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

func (suite *UserTerminalAccessRoutersTestSuite) TestFetchTerminalPodManifest() {
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

	suite.Run("A=1=FetchPodManifestWithValidTerminalAccessId", func() {
		queryParams := map[string]string{"terminalAccessId": strconv.Itoa(TerminalAccessId)}
		log.Println("Here we are trying to fetch pod manifest")
		TerminalPodManifest := HitFetchTerminalPodManifestApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), 200, TerminalPodManifest.Code)
		if TerminalPodManifest.Result.Manifest.Status.Phase == "Pending" {
			time.Sleep(10 * time.Second)
			TerminalPodManifest = HitFetchTerminalPodManifestApi(queryParams, suite.authToken)
			assert.Equal(suite.T(), "Running", TerminalPodManifest.Result.Manifest.Status.Phase)
		}
		assert.Equal(suite.T(), "Running", TerminalPodManifest.Result.Manifest.Status.Phase)
	})

	suite.Run("A=2=FetchPodEventsWithInvalidTerminalAccessId", func() {
		queryParams := map[string]string{"terminalAccessId": strconv.Itoa(Base.GetRandomNumberOf9Digit())}
		TerminalPodManifest := HitFetchTerminalPodManifestApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), 500, TerminalPodManifest.Code)
		assert.Equal(suite.T(), "unable to fetch manifest", TerminalPodManifest.Errors[0].UserMessage)
	})

	suite.Run("A=3=FetchPodEventsWithInValidNode", func() {
		log.Println("Here we are starting terminal session with invalid Node Name")
		StartTerminalSessionObj := CreatePayLoadForStartTerminalSession(clusterId, "ubuntu:latest", "InvalidNodeName", "default", "sh")
		StartTerminalSessionPayload, _ := json.Marshal(StartTerminalSessionObj)
		ApiResponse := HitStartTerminalSessionApi(StartTerminalSessionPayload, suite.authToken)
		TerminalAccessId = ApiResponse.Result.TerminalAccessId
		log.Println("Here we are trying to fetch pod manifest")
		queryParams := map[string]string{"terminalAccessId": strconv.Itoa(TerminalAccessId)}
		TerminalPodManifest := HitFetchTerminalPodManifestApi(queryParams, suite.authToken)
		if TerminalPodManifest.Result.Manifest.Status.Phase == "Pending" {
			time.Sleep(10 * time.Second)
			TerminalPodManifest = HitFetchTerminalPodManifestApi(queryParams, suite.authToken)
			assert.Equal(suite.T(), "Pending", TerminalPodManifest.Result.Manifest.Status.Phase)
		}
	})
}
