package ApplicationRouter

import (
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *ApplicationsRouterTestSuite) TestClassGetTerminalSession() {
	var container string
	createAppApiResponse, workflowResponse := PipelineConfigRouter.CreateNewAppWithCiCd(suite.authToken)
	time.Sleep(2 * time.Second)

	log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.Result.CiPipelines[0].Id, suite.authToken)

	log.Println("=== Here we are getting workflow status material ===")
	updatedWorkflowStatus := PipelineConfigRouter.HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)

	if updatedWorkflowStatus.Result.CdWorkflowStatus[0].DeployStatus == "Not Deployed" || updatedWorkflowStatus.Code != 200 {
		log.Println("=== Here we are Triggering CI/CD and verifying CI/CD Deploy Status ===")
		triggerAndVerifyCiPipeline(createAppApiResponse, pipelineMaterial, workflowResponse.Result.CiPipelines[0].Id, suite)

		log.Println("=== Here we are getting ResourceTree ===")
		ResourceTreeApiResponse := HitGetResourceTreeApi(createAppApiResponse.AppName, suite.authToken)
		container = ResourceTreeApiResponse.Result.PodMetadata[0].Containers[0]
	} else {
		log.Println("=== Here we are getting ResourceTree ===")
		ResourceTreeApiResponse := HitGetResourceTreeApi(createAppApiResponse.AppName, suite.authToken)
		container = ResourceTreeApiResponse.Result.PodMetadata[0].Containers[0]
	}

	suite.Run("A=1=GetTerminalSessionWithValidArgs", func() {
		TerminalSessionApiResponse := HitGetTerminalSessionApi(strconv.Itoa(createAppApiResponse.Id), "1", "devtron-demo", container, createAppApiResponse.AppName, suite.authToken)
		assert.NotEmpty(suite.T(), TerminalSessionApiResponse.Result.SessionID)
	})

	suite.Run("A=2=GetTerminalSessionWithInvalidEnvId", func() {
		invalidEnvId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		TerminalSessionApiResponse := HitGetTerminalSessionApi(strconv.Itoa(createAppApiResponse.Id), invalidEnvId, "devtron-demo", container, createAppApiResponse.AppName, suite.authToken)
		assert.Equal(suite.T(), TerminalSessionApiResponse.Errors[0].InternalMessage, "[{pg: no rows in result set}]", suite.authToken)
	})

	// we are not using this API , this is happening as during communication with ArgoCD fastTimeOut comes in picture that is 10 sec only
	/*suite.Run("A=3=GetApplicationViaValidName", func() {
		getApplicationApiResponse := HitGetApplicationApi(createAppApiResponse.AppName+"-devtron-demo", suite.authToken)
		assert.NotEmpty(suite.T(), getApplicationApiResponse.Result.Metadata.Uid)
		assert.Equal(suite.T(), getApplicationApiResponse.Result.Status.Health.Status, "Healthy")
		assert.True(suite.T(), strings.Contains(getApplicationApiResponse.Result.Spec.Source.RepoURL, createAppApiResponse.AppName))
		assert.Equal(suite.T(), getApplicationApiResponse.Result.Metadata.ManagedFields[0].Manager, "argocd-application-controller")
	})*/

	suite.Run("A=4=GetTerminalSessionWithInvalidName", func() {
		invalidName := Base.GetRandomStringOfGivenLength(8)
		getApplicationApiResponse := HitGetApplicationApi(invalidName, suite.authToken)
		assert.True(suite.T(), strings.Contains(getApplicationApiResponse.Errors[0].InternalMessage, "[{rpc error: code = NotFound desc = error getting application: applications.argoproj.io"))
	})

	//PipelineConfigRouter.DeleteAppWithCiCd(suite.authToken)
}

//todo will add other test cases once Devs will handle the validations
