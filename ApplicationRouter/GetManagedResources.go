package ApplicationRouter

import (
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *ApplicationsRouterTestSuite) TestClassGetManagedResources() {
	createAppApiResponse, workflowResponse := PipelineConfigRouter.CreateNewAppWithCiCd(suite.authToken, false)
	log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.Result.CiPipelines[0].Id, suite.authToken)

	appName := createAppApiResponse.AppName
	log.Println("=== Here we are getting workflow status material ===")
	updatedWorkflowStatus := PipelineConfigRouter.HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
	if updatedWorkflowStatus.Result.CdWorkflowStatus[0].DeployStatus == "Not Deployed" || updatedWorkflowStatus.Code != 200 {
		log.Println("=== Here we are Triggering CI/CD and verifying CI/CD Deploy Status ===")
		triggerAndVerifyCiPipeline(createAppApiResponse, pipelineMaterial, workflowResponse.Result.CiPipelines[0].Id, suite)
	}

	suite.Run("A=1=GetManagedResourcesWithInvalidAppName", func() {
		randomAppName := Base.GetRandomStringOfGivenLength(8)
		managedResourceApiResponse := HitGetManagedResourcesApi(randomAppName, suite.authToken)
		assert.True(suite.T(), strings.Contains(managedResourceApiResponse.Errors[0].InternalMessage, "[{rpc error: code = NotFound desc = error getting application: application.argoproj.io"))
	})

	suite.Run("A=2=GetManagedResourcesWithValidAppName", func() {
		managedResourceApiResponse := HitGetManagedResourcesApi(createAppApiResponse.AppName, suite.authToken)
		assert.Equal(suite.T(), "argoproj.io", managedResourceApiResponse.Result.Items[3].Group)
		assert.Equal(suite.T(), "Rollout", managedResourceApiResponse.Result.Items[3].Kind)
		assert.True(suite.T(), strings.Contains(managedResourceApiResponse.Result.Items[3].Name, appName))
		assert.True(suite.T(), strings.Contains(managedResourceApiResponse.Result.Items[2].Name, appName))
		assert.Equal(suite.T(), "ConfigMap", managedResourceApiResponse.Result.Items[0].Kind)
		assert.Equal(suite.T(), "Secret", managedResourceApiResponse.Result.Items[1].Kind)
		for i := 0; i < 3; i++ {
			assert.NotNil(suite.T(), managedResourceApiResponse.Result.Items[i].LiveState)
			assert.NotNil(suite.T(), managedResourceApiResponse.Result.Items[i].TargetState)
			assert.NotNil(suite.T(), managedResourceApiResponse.Result.Items[i].NormalizedLiveState)
			assert.NotNil(suite.T(), managedResourceApiResponse.Result.Items[i].PredictedLiveState)
			assert.Equal(suite.T(), "devtron-demo", managedResourceApiResponse.Result.Items[i].Namespace)
		}
	})

	//PipelineConfigRouter.DeleteAppWithCiCd(suite.authToken)
}
