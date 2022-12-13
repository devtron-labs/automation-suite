package ApplicationRouter

import (
	"automation-suite/PipelineConfigRouter"
	PipelineConfigRouterResponseDTOs "automation-suite/PipelineConfigRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"time"
)

func (suite *ApplicationsRouterTestSuite) TestClassGetResourceTree() {
	createAppApiResponse, workflowResponse := Base.CreateNewAppWithCiCd(suite.authToken)
	time.Sleep(2 * time.Second)
	log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.Result.CiPipelines[0].Id, suite.authToken)
	log.Println("=== Here we are Triggering CI/CD and verifying CI/CD Deploy Status ===")
	time.Sleep(10 * time.Second)
	triggerAndVerifyCiPipeline(createAppApiResponse, pipelineMaterial, workflowResponse.Result.CiPipelines[0].Id, suite)

	suite.Run("A=1=GetResourceTreeWithValidAppName", func() {
		ResourceTreeApiResponse := HitGetResourceTreeApi(createAppApiResponse.AppName, suite.authToken)
		assert.NotNil(suite.T(), ResourceTreeApiResponse.Result, suite.authToken)
		assert.Equal(suite.T(), 8, len(ResourceTreeApiResponse.Result.Nodes))
		assert.Equal(suite.T(), "ConfigMap", ResourceTreeApiResponse.Result.Nodes[0].Kind)
		assert.Equal(suite.T(), "EndpointSlice", ResourceTreeApiResponse.Result.Nodes[7].Kind)
		assert.Equal(suite.T(), "linux", ResourceTreeApiResponse.Result.Hosts[0].SystemInfo.OperatingSystem)
		assert.Equal(suite.T(), "Healthy", ResourceTreeApiResponse.Result.Status)
	})

	suite.Run("A=2=GetResourceTreeWithInvalidAppName", func() {
		randomAppName := Base.GetRandomStringOfGivenLength(8)
		ResourceTreeApiResponse := HitGetResourceTreeApi(randomAppName, suite.authToken)
		assert.True(suite.T(), strings.Contains(ResourceTreeApiResponse.Errors[0].InternalMessage, "[{rpc error: code = NotFound desc = error getting application by name: application.argoproj.io"))
	})

	Base.DeleteAppWithCiCd(suite.authToken)
}

func triggerAndVerifyCiPipeline(createAppApiResponse Base.CreateAppRequestDto, pipelineMaterial PipelineConfigRouterResponseDTOs.GetCiPipelineMaterialResponseDTO, CiPipelineID int, suite *ApplicationsRouterTestSuite) string {
	ciTriggerWorkflowId := ""
	payloadForTriggerCiPipeline := PipelineConfigRouter.CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, CiPipelineID, pipelineMaterial.Result[0].Id, true)
	bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
	triggerCiPipelineResponse := PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
	if triggerCiPipelineResponse.Result.AuthStatus != "allowed for all pipelines" {
		time.Sleep(2 * time.Second)
		triggerCiPipelineResponse = PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "allowed for all pipelines", triggerCiPipelineResponse.Result.AuthStatus)
		assert.NotNil(suite.T(), triggerCiPipelineResponse.Result.ApiResponse)
	}
	ciTriggerWorkflowId = triggerCiPipelineResponse.Result.ApiResponse
	time.Sleep(10 * time.Second)
	log.Println("=== Here we are getting workflow after triggering ===")
	workflowStatus := PipelineConfigRouter.HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
	if workflowStatus.Result.CiWorkflowStatus[0].CiStatus == "Starting" {
		time.Sleep(5 * time.Second)
		workflowStatus = PipelineConfigRouter.HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
		assert.Equal(suite.T(), "Running", workflowStatus.Result.CiWorkflowStatus[0].CiStatus)
	} else {
		assert.Equal(suite.T(), "Running", workflowStatus.Result.CiWorkflowStatus[0].CiStatus)
	}
	log.Println("=== Here we are getting workflow and verifying the status after triggering via poll function ===")
	assert.True(suite.T(), PollForGettingCdDeployStatusAfterTrigger(createAppApiResponse.Id, suite.authToken))
	updatedWorkflowStatus := PipelineConfigRouter.HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
	assert.Equal(suite.T(), "Succeeded", updatedWorkflowStatus.Result.CiWorkflowStatus[0].CiStatus)
	assert.Equal(suite.T(), "Healthy", updatedWorkflowStatus.Result.CdWorkflowStatus[0].DeployStatus)
	return ciTriggerWorkflowId
}

func PollForGettingCdDeployStatusAfterTrigger(id int, authToken string) bool {
	count := 0
	for {
		updatedWorkflowStatus := PipelineConfigRouter.HitGetWorkflowStatus(id, authToken)
		deploymentStatus := updatedWorkflowStatus.Result.CdWorkflowStatus[0].DeployStatus
		time.Sleep(1 * time.Second)
		count = count + 1
		if deploymentStatus == "Healthy" || count >= 800 {
			break
		}
	}
	return true
}
