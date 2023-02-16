package AppListingRouter

import (
	"automation-suite/PipelineConfigRouter"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"time"
)

func (suite *AppsListingRouterTestSuite) TestFetchAppsByEnvironment() {

	suite.Run("A=1=FetchAppsByEnvironmentWithDefaultFilter", func() {
		log.Println("Here we are creating request payload")
		Envs := []int{}
		Teams := []int{}
		Namespaces := []string{}
		AppStatuses := []string{}
		requestDTOForApiFetchAppsByEnvironment := GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
		payloadForApiFetchAppsByEnvironment, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)
		log.Println("Here we are fetching the app list without creating a new one")
		fetchOtherEnvResponseDto := HitApiFetchAppsByEnvironment(payloadForApiFetchAppsByEnvironment, suite.authToken)
		appCountBeforeCreationNewOne := fetchOtherEnvResponseDto.Result.AppCount
		log.Println("Here we are creating a new app with CI/CD")
		createAppApiResponse, workflowResponse := PipelineConfigRouter.CreateNewAppWithCiCd(suite.authToken)
		time.Sleep(2 * time.Second)
		log.Println("=== Here we are getting pipeline material ===")
		pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.Result.CiPipelines[0].Id, suite.authToken)
		payloadForTriggerCiPipeline := PipelineConfigRouter.CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, workflowResponse.Result.CiPipelines[0].Id, pipelineMaterial.Result[0].Id, true)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		triggerCiPipelineResponse := PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
		if triggerCiPipelineResponse.Result.AuthStatus != "allowed for all pipelines" {
			time.Sleep(5 * time.Second)
			triggerCiPipelineResponse = PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
			assert.Equal(suite.T(), "allowed for all pipelines", triggerCiPipelineResponse.Result.AuthStatus)
			assert.NotNil(suite.T(), triggerCiPipelineResponse.Result.ApiResponse)
		}
		time.Sleep(5 * time.Second)
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
		assert.True(suite.T(), PipelineConfigRouter.PollForGettingCdDeployStatusAfterTrigger(createAppApiResponse.Id, suite.authToken))
		updatedWorkflowStatus := PipelineConfigRouter.HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
		assert.Equal(suite.T(), "Succeeded", updatedWorkflowStatus.Result.CiWorkflowStatus[0].CiStatus)
		assert.Equal(suite.T(), "Succeeded", updatedWorkflowStatus.Result.CdWorkflowStatus[0].DeployStatus)
		log.Println("Here we are fetching the app list after creating a new app")
		fetchOtherEnvResponseDto = HitApiFetchAppsByEnvironment(payloadForApiFetchAppsByEnvironment, suite.authToken)
		appCountAfterCreationNewOne := fetchOtherEnvResponseDto.Result.AppCount
		assert.Equal(suite.T(), appCountBeforeCreationNewOne+1, appCountAfterCreationNewOne)
		assert.Equal(suite.T(), 200, fetchOtherEnvResponseDto.Code)
		lastIndexOfApps := len(fetchOtherEnvResponseDto.Result.AppContainers)
		assert.Equal(suite.T(), createAppApiResponse.Id, fetchOtherEnvResponseDto.Result.AppContainers[lastIndexOfApps-1].AppId)
		assert.Equal(suite.T(), 1, fetchOtherEnvResponseDto.Result.AppContainers[lastIndexOfApps-1].ProjectId)
		assert.Equal(suite.T(), 1, fetchOtherEnvResponseDto.Result.AppContainers[lastIndexOfApps-1].Environments[0].EnvironmentId)
		assert.Equal(suite.T(), "default_cluster", fetchOtherEnvResponseDto.Result.AppContainers[lastIndexOfApps-1].Environments[0].ClusterName)
		assert.Equal(suite.T(), "devtron-demo", fetchOtherEnvResponseDto.Result.AppContainers[lastIndexOfApps-1].Environments[0].Namespace)
		assert.Equal(suite.T(), "devtron-demo", fetchOtherEnvResponseDto.Result.AppContainers[lastIndexOfApps-1].Environments[0].EnvironmentName)
		assert.Equal(suite.T(), "Healthy", fetchOtherEnvResponseDto.Result.AppContainers[lastIndexOfApps-1].Environments[0].AppStatus)
	})

	suite.Run("A=2=FetchAppsByEnvironmentWithAppStatusFilter", func() {
		log.Println("Here we are creating request payload")
		Envs := []int{}
		Teams := []int{}
		Namespaces := []string{}
		AppStatuses := []string{"Healthy"}
		requestDTOForApiFetchAppsByEnvironment := GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
		payloadForApiFetchAppsByEnvironment, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)
		log.Println("Here we are fetching the app list without creating a new one")
		appsList := HitApiFetchAppsByEnvironment(payloadForApiFetchAppsByEnvironment, suite.authToken)
		lastIndexOfList := len(appsList.Result.AppContainers)
		assert.Equal(suite.T(), "Healthy", appsList.Result.AppContainers[0].Environments[0].AppStatus)
		assert.Equal(suite.T(), "Healthy", appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].AppStatus)
	})

	suite.Run("A=3=FetchAppsByEnvironmentWithEnvFilter", func() {
		log.Println("Here we are creating request payload")
		Envs := []int{1}
		Teams := []int{}
		Namespaces := []string{}
		AppStatuses := []string{}
		requestDTOForApiFetchAppsByEnvironment := GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
		payloadForApiFetchAppsByEnvironment, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)
		log.Println("Here we are fetching the app list without creating a new one")
		appsList := HitApiFetchAppsByEnvironment(payloadForApiFetchAppsByEnvironment, suite.authToken)
		lastIndexOfList := len(appsList.Result.AppContainers)
		assert.Equal(suite.T(), 1, appsList.Result.AppContainers[0].Environments[0].EnvironmentId)
		assert.Equal(suite.T(), 1, appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].EnvironmentId)
	})

	suite.Run("A=4=FetchAppsByEnvironmentWithTeamsFilter", func() {
		log.Println("Here we are creating request payload")
		Envs := []int{}
		Teams := []int{1}
		Namespaces := []string{}
		AppStatuses := []string{}
		requestDTOForApiFetchAppsByEnvironment := GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
		payloadForApiFetchAppsByEnvironment, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)
		log.Println("Here we are fetching the app list without creating a new one")
		appsList := HitApiFetchAppsByEnvironment(payloadForApiFetchAppsByEnvironment, suite.authToken)
		lastIndexOfList := len(appsList.Result.AppContainers)
		assert.Equal(suite.T(), 1, appsList.Result.AppContainers[0].Environments[0].TeamId)
		assert.Equal(suite.T(), 1, appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].TeamId)
	})

	suite.Run("A=5=FetchAppsByEnvironmentWithClusterFilter", func() {
		log.Println("Here we are creating request payload")
		Envs := []int{}
		Teams := []int{}
		Namespaces := []string{"1"}
		AppStatuses := []string{}
		requestDTOForApiFetchAppsByEnvironment := GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
		payloadForApiFetchAppsByEnvironment, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)
		log.Println("Here we are fetching the app list without creating a new one")
		appsList := HitApiFetchAppsByEnvironment(payloadForApiFetchAppsByEnvironment, suite.authToken)
		lastIndexOfList := len(appsList.Result.AppContainers)
		assert.Equal(suite.T(), "default_cluster", appsList.Result.AppContainers[0].Environments[0].ClusterName)
		assert.Equal(suite.T(), "default_cluster", appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].ClusterName)
	})

	suite.Run("A=6=FetchAppsByEnvironmentWithClusterFilter", func() {
		log.Println("Here we are creating request payload")
		Envs := []int{}
		Teams := []int{}
		Namespaces := []string{"1_devtron-demo"}
		AppStatuses := []string{}
		requestDTOForApiFetchAppsByEnvironment := GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
		payloadForApiFetchAppsByEnvironment, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)
		log.Println("Here we are fetching the app list without creating a new one")
		appsList := HitApiFetchAppsByEnvironment(payloadForApiFetchAppsByEnvironment, suite.authToken)
		lastIndexOfList := len(appsList.Result.AppContainers)
		assert.Equal(suite.T(), "devtron-demo", appsList.Result.AppContainers[0].Environments[0].Namespace)
		assert.Equal(suite.T(), "devtron-demo", appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].Namespace)
	})

	suite.Run("A=7=FetchAppsByEnvironmentWithAllAvailableFilters", func() {
		log.Println("Here we are creating request payload")
		Envs := []int{1}
		Teams := []int{1}
		Namespaces := []string{"1_devtron-demo"}
		AppStatuses := []string{"Healthy"}
		requestDTOForApiFetchAppsByEnvironment := GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
		payloadForApiFetchAppsByEnvironment, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)
		log.Println("Here we are fetching the app list without creating a new one")
		appsList := HitApiFetchAppsByEnvironment(payloadForApiFetchAppsByEnvironment, suite.authToken)
		lastIndexOfList := len(appsList.Result.AppContainers)
		assert.Equal(suite.T(), "devtron-demo", appsList.Result.AppContainers[0].Environments[0].Namespace)
		assert.Equal(suite.T(), "devtron-demo", appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].Namespace)
		assert.Equal(suite.T(), "default_cluster", appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].ClusterName)
		assert.Equal(suite.T(), 1, appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].TeamId)
		assert.Equal(suite.T(), "Healthy", appsList.Result.AppContainers[lastIndexOfList-1].Environments[0].AppStatus)
	})

	PipelineConfigRouter.DeleteAppWithCiCd(suite.authToken)
}
