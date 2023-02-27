package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"time"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassD3TriggerCiPipeline() {
	createAppApiResponse, workflowResponse := CreateNewAppWithCiCd(suite.authToken, false)
	time.Sleep(2 * time.Second)
	log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := HitGetCiPipelineMaterial(workflowResponse.Result.CiPipelines[0].Id, suite.authToken)

	//here we are hitting GetWorkFlow API 2 time one just after the triggerCiPipeline and one after 4 minutes of triggering
	suite.Run("A=1=TriggerCiPipelineWithValidPayload", func() {
		payloadForTriggerCiPipeline := CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, workflowResponse.Result.CiPipelines[0].Id, pipelineMaterial.Result[0].Id, true)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		triggerCiPipelineResponse := HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
		if triggerCiPipelineResponse.Result.AuthStatus != "allowed for all pipelines" {
			time.Sleep(5 * time.Second)
			triggerCiPipelineResponse = HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
			assert.Equal(suite.T(), "allowed for all pipelines", triggerCiPipelineResponse.Result.AuthStatus)
			assert.NotNil(suite.T(), triggerCiPipelineResponse.Result.ApiResponse)
		}
		time.Sleep(5 * time.Second)
		log.Println("=== Here we are getting workflow after triggering ===")
		workflowStatus := HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
		if workflowStatus.Result.CiWorkflowStatus[0].CiStatus == "Starting" {
			time.Sleep(5 * time.Second)
			workflowStatus = HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
			assert.Equal(suite.T(), "Running", workflowStatus.Result.CiWorkflowStatus[0].CiStatus)
		} else {
			assert.Equal(suite.T(), "Running", workflowStatus.Result.CiWorkflowStatus[0].CiStatus)
		}
		log.Println("=== Here we are getting workflow and verifying the status after triggering via poll function ===")
		assert.True(suite.T(), PollForGettingCdDeployStatusAfterTrigger(createAppApiResponse.Id, suite.authToken))
		updatedWorkflowStatus := HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
		assert.Equal(suite.T(), "Succeeded", updatedWorkflowStatus.Result.CiWorkflowStatus[0].CiStatus)
		assert.Equal(suite.T(), "Healthy", updatedWorkflowStatus.Result.CdWorkflowStatus[0].DeployStatus)
	})

	suite.Run("A=2=TriggerCiPipelineWithInvalidateCacheAsFalse", func() {
		payloadForTriggerCiPipeline := CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, workflowResponse.Result.CiPipelines[0].Id, pipelineMaterial.Result[0].Id, true)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		triggerCiPipelineResponse := HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
		if triggerCiPipelineResponse.Result.AuthStatus != "allowed for all pipelines" {
			time.Sleep(2 * time.Second)
			triggerCiPipelineResponse = HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
			time.Sleep(5 * time.Second)
			assert.Equal(suite.T(), "allowed for all pipelines", triggerCiPipelineResponse.Result.AuthStatus)
			assert.NotNil(suite.T(), triggerCiPipelineResponse.Result.ApiResponse)
		}
	})

	suite.Run("A=3=TriggerCiPipelineWithInvalidPipelineId", func() {
		invalidPipeLineId := Base.GetRandomNumberOf9Digit()
		payloadForTriggerCiPipeline := CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, invalidPipeLineId, pipelineMaterial.Result[0].Id, true)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		triggerCiPipelineResponse := HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", triggerCiPipelineResponse.Errors[0].UserMessage)
	})

	suite.Run("A=4=TriggerCiPipelineWithInvalidPipelineId", func() {
		invalidMaterialId := Base.GetRandomNumberOf9Digit()
		payloadForTriggerCiPipeline := CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, workflowResponse.Result.CiPipelines[0].Id, invalidMaterialId, true)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		triggerCiPipelineResponse := HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", triggerCiPipelineResponse.Errors[0].InternalMessage)
	})

	DeleteAppWithCiCd(suite.authToken)
}

func PollForGettingCdDeployStatusAfterTrigger(id int, authToken string) bool {
	count := 0
	for {
		updatedWorkflowStatus := HitGetWorkflowStatus(id, authToken)
		deploymentStatus := updatedWorkflowStatus.Result.CdWorkflowStatus[0].DeployStatus
		time.Sleep(1 * time.Second)
		count = count + 1
		if deploymentStatus == "Healthy" || count >= 600 {
			break
		}
	}
	return true
}
