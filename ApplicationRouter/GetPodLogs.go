package ApplicationRouter

import (
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"context"
	"fmt"
	"github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	v1alpha12 "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *ApplicationsRouterTestSuite) TestGetPodLogs() {
	createAppApiResponse, workflowResponse := PipelineConfigRouter.CreateNewAppWithCiCd(suite.authToken, false)
	ciPipelineId := workflowResponse.Result.CiPipelines[0].Id
	var ciWorkflowId string
	var container string
	var containerName string
	appId := createAppApiResponse.Id
	//log.Println("=== Here we are getting workflow status material ===")
	/*updatedWorkflowStatus := PipelineConfigRouter.HitGetWorkflowStatus(createAppApiResponse.Id, suite.authToken)
	if updatedWorkflowStatus.Result.CdWorkflowStatus[0].DeployStatus == "Not Deployed" || updatedWorkflowStatus.Code != 200 {
	*/log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(ciPipelineId, suite.authToken)
	time.Sleep(5 * time.Second)
	log.Println("=== Here we are Triggering CI/CD and verifying CI/CD Deploy Status ===")
	ciWorkflowId = triggerAndVerifyCiPipeline(createAppApiResponse, pipelineMaterial, ciPipelineId, suite)
	//}
	log.Println("=== Here we are getting ResourceTree ===")
	ResourceTreeApiResponse := HitGetResourceTreeApi(createAppApiResponse.AppName, suite.authToken)
	container = ResourceTreeApiResponse.Result.PodMetadata[0].Containers[0]
	containerName = ResourceTreeApiResponse.Result.PodMetadata[0].Name
	queryParams := make(map[string]string)
	queryParams["container"] = container
	queryParams["follow"] = "true"
	queryParams["namespace"] = "devtron-demo"
	queryParams["tailLines"] = "500"
	url := Base.CreateUrlForEventStreamsHavingQueryParam(queryParams)
	ContainersUrl := ApplicationsRouterBaseUrl + container + "-devtron-demo" + "/pods/" + containerName + "/logs?" + url

	suite.Run("A=1=GetPodLogsForValidContainer", func() {
		Base.ReadEventStreamsForSpecificApi(ContainersUrl, suite.authToken, containerName, suite.T())
	})

	// need to check artifacts for ci, pre-cd, post-cd
	pipelineId := strconv.Itoa(ciPipelineId)
	envIdString := "1" //devtron-demo env

	suite.Run("A=2=CheckForCiArtifacts", func() {
		suite.checkForCiArtifacts(pipelineId, ciWorkflowId)
	})

	suite.Run("A=3=CheckForCiLogs", func() {
		suite.checkForCiLogs(pipelineId, ciWorkflowId, 9, "STAGE:  running PRE-CI steps")
	})

	suite.Run("A=4=CheckForPreCdAndPostCdArtifactsAndLogs", func() {
		cdPipeLineResponse := PipelineConfigRouter.HitGetAppCdPipeline(strconv.Itoa(appId), suite.authToken)
		suite.checkForPreCdAndPostCdArtifactsAndLogs(strconv.Itoa(appId), envIdString, strconv.Itoa(cdPipeLineResponse.Result.Pipelines[0].Id))
	})

	// for ci, pre-cd, post-cd logs first need to delete workflow after succeed
	//PipelineConfigRouter.DeleteAppWithCiCd(suite.authToken)
}

func (suite *ApplicationsRouterTestSuite) checkForCiArtifacts(pipelineId string, ciWorkflowId string) {
	artifactsUrl := PipelineConfigRouter.GetCiPipelineBaseUrl + "/" + pipelineId + "/artifacts/" + ciWorkflowId
	suite.hitAndCheckArtifactDownload(artifactsUrl)
}

func (suite *ApplicationsRouterTestSuite) checkForCiLogs(pipelineId string, ciWorkflowId string, LogLineNumber int, logString string) {
	workflowsDownloadUrl := PipelineConfigRouter.GetCiPipelineBaseUrl + "/" + pipelineId + "/workflows"
	workflows, err := PipelineConfigRouter.FetchCiWorkflows(workflowsDownloadUrl, suite.authToken)
	assert.True(suite.T(), err == nil, err)
	workflowPodName := ""
	workflowNameSpace := ""
	workflowResponses := workflows.Result
	for _, response := range workflowResponses {
		if ciWorkflowId == strconv.Itoa(response.Id) {
			workflowPodName = response.Name
			workflowNameSpace = response.Namespace
		}
	}
	ciLogsDownloadUrlFormat := PipelineConfigRouter.GetCiPipelineBaseUrl + "/%s/workflow/%s/logs"
	ciLogsDownloadUrl := fmt.Sprintf(ciLogsDownloadUrlFormat, pipelineId, ciWorkflowId)
	suite.hitAndCheckBuildLogs(ciLogsDownloadUrl, workflowNameSpace, workflowPodName, LogLineNumber, logString)
}

func (suite *ApplicationsRouterTestSuite) hitAndCheckArtifactDownload(artifactsUrl string) {
	artifactDownloadStatusCode, err := PipelineConfigRouter.HitCiArtifactsDownloadApi(artifactsUrl, suite.authToken)
	assert.True(suite.T(), err == nil, err)
	assert.True(suite.T(), artifactDownloadStatusCode == 200, artifactDownloadStatusCode)
}

func (suite *ApplicationsRouterTestSuite) hitAndCheckBuildLogs(downloadUrl string, namespace string, wfName string, logLineIndex int, logString string) {
	PipelineConfigRouter.HitLogsDownloadApi(downloadUrl, suite.authToken, suite.T(), logLineIndex, logString)
	var wfClient v1alpha12.WorkflowInterface
	var err error
	wfClient, err = suite.getClientInstance(namespace)
	if err != nil {
		fmt.Println("error while creating wf client", err)
		return
	}
	err = wfClient.Delete(context.Background(), wfName, v1.DeleteOptions{})
	if err != nil {
		fmt.Println("error while deleting wf object ", wfName, err)
		return
	}
	PipelineConfigRouter.HitLogsDownloadApi(downloadUrl, suite.authToken, suite.T(), logLineIndex, logString)
}

// todo need to add logic to pick Host and Token from Env Variable
func (suite *ApplicationsRouterTestSuite) getClientInstance(namespace string) (v1alpha12.WorkflowInterface, error) {

	envConfig := Base.ReadBaseEnvConfig()
	baseCredentials := Base.ReadAnyJsonFile(envConfig.BaseCredentialsFile)
	config := &rest.Config{
		Host:        baseCredentials.BaseServerUrl,
		BearerToken: baseCredentials.BearerToken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	clientSet, err := versioned.NewForConfig(config)
	if err != nil {
		fmt.Println("err", err)
	}

	wfClient := clientSet.ArgoprojV1alpha1().Workflows(namespace) // create the workflow client
	return wfClient, nil
}

func (suite *ApplicationsRouterTestSuite) checkForPreCdAndPostCdArtifactsAndLogs(appId string, envId string, pipelineId string) {
	cdArtifactsUrl := PipelineConfigRouter.GetAppCdPipelineApiUrl + "workflow/history/" + appId + "/" + envId + "/" + pipelineId
	workflows, err := PipelineConfigRouter.FetchCdPipelineWorkflows(cdArtifactsUrl, suite.authToken)
	assert.True(suite.T(), err == nil, err)
	assert.True(suite.T(), workflows.Code == 200, workflows.Code)
	assert.True(suite.T(), workflows.Status == "OK", workflows.Status)

	workflowResponse := workflows.Result
	latestPreCdWorkflowRunnerId := 0
	latestPostCdWorkflowRunnerId := 0
	postWorkflowStatus := ""
	preCdWorkflowName := ""
	postCdWorkflowName := ""
	preCdNamespace := "devtron-ci"
	postCdNameSpace := "devtron-cd"
	for _, workflow := range workflowResponse {
		if workflow.WorkflowType == "POST" && latestPostCdWorkflowRunnerId == 0 {
			latestPostCdWorkflowRunnerId = workflow.Id
			postWorkflowStatus = workflow.PodStatus
			postCdNameSpace = workflow.Namespace
			postCdWorkflowName = workflow.Name
		}
		if workflow.WorkflowType == "PRE" && latestPreCdWorkflowRunnerId == 0 {
			latestPreCdWorkflowRunnerId = workflow.Id
			preCdNamespace = workflow.Namespace
			preCdWorkflowName = workflow.Name
		}
	}
	postCdRunnerId := strconv.Itoa(latestPostCdWorkflowRunnerId)
	if postWorkflowStatus != "Succeeded" {
		healthy := suite.PollForPostCdTillHealthy(appId, envId, pipelineId, postCdRunnerId)
		assert.True(suite.T(), healthy)
	}
	artifactDownloadFormat := "workflow/download/%s/%s/%s/%s"
	preCdWorkflowRunnerId := strconv.Itoa(latestPreCdWorkflowRunnerId)
	preCdArtifactDownloadUrl := PipelineConfigRouter.GetAppCdPipelineApiUrl + fmt.Sprintf(artifactDownloadFormat, appId, envId, pipelineId, preCdWorkflowRunnerId)
	suite.hitAndCheckArtifactDownload(preCdArtifactDownloadUrl)

	postCdArtifactDownloadUrl := PipelineConfigRouter.GetAppCdPipelineApiUrl + fmt.Sprintf(artifactDownloadFormat, appId, envId, pipelineId, postCdRunnerId)
	suite.hitAndCheckArtifactDownload(postCdArtifactDownloadUrl)

	logsDownloadUrlFormat := "workflow/logs/%s/%s/%s/%s"
	logsDownloadUrl := PipelineConfigRouter.GetAppCdPipelineApiUrl + fmt.Sprintf(logsDownloadUrlFormat, appId, envId, pipelineId, preCdWorkflowRunnerId)
	lastIndex := strings.LastIndex(preCdWorkflowName, "-")
	preCdWorkflowName = preCdWorkflowName[0:lastIndex]
	suite.hitAndCheckBuildLogs(logsDownloadUrl, preCdNamespace, preCdWorkflowName, 10, "Login Succeeded")
	logsDownloadUrl = PipelineConfigRouter.GetAppCdPipelineApiUrl + fmt.Sprintf(logsDownloadUrlFormat, appId, envId, pipelineId, postCdRunnerId)
	lastIndex = strings.LastIndex(postCdWorkflowName, "-")
	postCdWorkflowName = postCdWorkflowName[0:lastIndex]
	suite.hitAndCheckBuildLogs(logsDownloadUrl, postCdNameSpace, postCdWorkflowName, 38, "END_OF_STREAM")
}

func (suite *ApplicationsRouterTestSuite) PollForPostCdTillHealthy(appId string, envId string, pipelineId string, workflowRunnerId string) bool {
	workflowStatusUrlFormat := "workflow/trigger-info/%s/%s/%s/%s"
	cdWorkflowRunnerStatusUrl := PipelineConfigRouter.GetAppCdPipelineApiUrl + fmt.Sprintf(workflowStatusUrlFormat, appId, envId, pipelineId, workflowRunnerId)
	counter := 0
	for true {
		counter++
		if counter > 100 {
			break
		}
		runnerDetails, err := PipelineConfigRouter.FetchCdWorkflowRunnerDetails(cdWorkflowRunnerStatusUrl, suite.authToken)
		if err != nil {
			fmt.Println(err)
			return false
		}
		workflowResponseDto := runnerDetails.Result
		if workflowResponseDto.PodStatus == "Succeeded" {
			return true
		}
		time.Sleep(5 * time.Second)
	}
	return false
}

//todo need to add test cases for further cases if required like for invalid container
