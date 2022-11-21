package ApplicationRouter

import (
	"automation-suite/HelperRouter"
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	v1alpha12 "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/sjson"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *ApplicationsRouterTestSuite) TestGetPodLogs() {
	config, _ := PipelineConfigRouter.GetEnvironmentConfigPipelineConfigRouter()
	var configId int
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appName := createAppApiResponse.AppName
	log.Println("=== App Name is :====", appName)
	log.Println("=== Here we are creating App Material ===")
	appId := createAppApiResponse.Id
	createAppMaterialRequestDto := PipelineConfigRouter.GetAppMaterialRequestDto(appId, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := PipelineConfigRouter.HitCreateAppMaterialApi(appMaterialByteValue, appId, 1, false, suite.authToken)

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := PipelineConfigRouter.GetRequestPayloadForSaveAppCiPipeline(appId, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	PipelineConfigRouter.HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	time.Sleep(2 * time.Second)
	appIdString := strconv.Itoa(appId)
	getChartReferenceResponse := PipelineConfigRouter.HitGetChartReferenceViaAppId(appIdString, suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching Template using getAppTemplateAPI ===")
	getTemplateResponse := PipelineConfigRouter.HitGetTemplateViaAppIdAndChartRefId(appIdString, strconv.Itoa(latestChartRef), suite.authToken)

	log.Println("=== Here we are fetching DefaultAppOverride from template response ===")
	defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

	log.Println("=== Here we are creating payload for SaveTemplate API ===")
	saveDeploymentTemplate := PipelineConfigRouter.GetRequestPayloadForSaveDeploymentTemplate(appId, latestChartRef, defaultAppOverride)
	byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)
	jsonOfSaveDeploymentTemp := string(byteValueOfSaveDeploymentTemplate)
	jsonWithMicroserviceToleration, _ := sjson.Set(jsonOfSaveDeploymentTemp, "valuesOverride.tolerations.0", map[string]interface{}{"effect": "NoSchedule", "key": "microservice", "operator": "Equal", "value": "true"})
	finalJson, _ := sjson.Set(jsonWithMicroserviceToleration, "valuesOverride.tolerations.1", map[string]interface{}{"effect": "NoSchedule", "key": "kubernetes.azure.com/scalesetpriority", "operator": "Equal", "value": "spot"})
	updatedByteValueOfSaveDeploymentTemplate := []byte(finalJson)

	log.Println("=== Here we are hitting SaveTemplate API ===")
	PipelineConfigRouter.HitSaveDeploymentTemplateApi(updatedByteValueOfSaveDeploymentTemplate, suite.authToken)

	log.Println("=== Here we are saving Global Configmap ===")
	requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(0, "-config1", appId, "environment", "kubernetes", false, false, false, false)
	byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
	globalConfigMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
	configId = globalConfigMap.Result.Id

	log.Println("=== Here we are saving Global Secret ===")
	requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, "-secret1", appId, "environment", "kubernetes", false, false, true, false)
	byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
	HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)

	log.Println("=== Here we are saving workflow with Pre/Post CI ===")
	workflowResponse := PipelineConfigRouter.HitCreateWorkflowApiWithFullPayload(appId, suite.authToken).Result

	preStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
	postStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")

	log.Println("=== Here we are saving CD pipeline ===")
	ciPipelineId := workflowResponse.CiPipelines[0].Id
	payload := PipelineConfigRouter.GetRequestPayloadForSaveCdPipelineApi(appId, workflowResponse.AppWorkflowId, 1, ciPipelineId, workflowResponse.CiPipelines[0].ParentCiPipeline, "AUTOMATIC", string(preStageScript), string(postStageScript), "AUTOMATIC")
	bytePayload, _ := json.Marshal(payload)
	savePipelineResponse := PipelineConfigRouter.HitSaveCdPipelineApi(bytePayload, suite.authToken)
	time.Sleep(2 * time.Second)

	log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(ciPipelineId, suite.authToken)

	log.Println("=== Here we are Triggering CI/CD and verifying CI/CD Deploy Status ===")
	ciWorkflowId := triggerAndVerifyCiPipeline(createAppApiResponse, pipelineMaterial, ciPipelineId, suite)

	log.Println("=== Here we are getting ResourceTree ===")
	ResourceTreeApiResponse := HitGetResourceTreeApi(createAppApiResponse.AppName, suite.authToken)
	container := ResourceTreeApiResponse.Result.PodMetadata[0].Containers[0]
	containerName := ResourceTreeApiResponse.Result.PodMetadata[0].Name
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
		suite.checkForPreCdAndPostCdArtifactsAndLogs(appIdString, envIdString, strconv.Itoa(cdPipeLineResponse.Result.Pipelines[0].Id))
	})

	// for ci, pre-cd, post-cd logs first need to delete workflow after succeed

	log.Println("=== Here we are Deleting the CD pipeline ===")
	deletePipelinePayload := PipelineConfigRouter.GetPayloadForDeleteCdPipeline(appId, savePipelineResponse.Result.Pipelines[0].Id)
	deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
	PipelineConfigRouter.HitForceDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)

	log.Println("=== Here we are Deleting the CI pipeline ===")
	PipelineConfigRouter.DeleteCiPipeline(appId, ciPipelineId, suite.authToken)
	log.Println("=== Here we are Deleting CI Workflow ===")
	PipelineConfigRouter.HitDeleteWorkflowApi(appId, workflowResponse.AppWorkflowId, suite.authToken)
	log.Println("=== Here we are Deleting the app after all verifications ===")
	Base.DeleteApp(appId, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
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
	suite.hitAndCheckBuildLogs(logsDownloadUrl, postCdNameSpace, postCdWorkflowName, 40, "END_OF_STREAM")
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
