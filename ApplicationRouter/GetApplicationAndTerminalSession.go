package ApplicationRouter

import (
	"automation-suite/HelperRouter"
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/sjson"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *ApplicationsRouterTestSuite) TestClassGetTerminalSession() {
	config, _ := PipelineConfigRouter.GetEnvironmentConfigPipelineConfigRouter()
	var configId int
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	log.Println("=== App Name is :====>", createAppApiResponse.AppName)

	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := PipelineConfigRouter.GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := PipelineConfigRouter.HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := PipelineConfigRouter.GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	PipelineConfigRouter.HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	time.Sleep(2 * time.Second)
	getChartReferenceResponse := PipelineConfigRouter.HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching Template using getAppTemplateAPI ===")
	getTemplateResponse := PipelineConfigRouter.HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), suite.authToken)

	log.Println("=== Here we are fetching DefaultAppOverride from template response ===")
	defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

	log.Println("=== Here we are creating payload for SaveTemplate API ===")
	saveDeploymentTemplate := PipelineConfigRouter.GetRequestPayloadForSaveDeploymentTemplate(createAppApiResponse.Id, latestChartRef, defaultAppOverride)
	byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)
	jsonOfSaveDeploymentTemp := string(byteValueOfSaveDeploymentTemplate)
	jsonWithMicroserviceToleration, _ := sjson.Set(jsonOfSaveDeploymentTemp, "valuesOverride.tolerations.0", map[string]interface{}{"effect": "NoSchedule", "key": "microservice", "operator": "Equal", "value": "true"})
	finalJson, _ := sjson.Set(jsonWithMicroserviceToleration, "valuesOverride.tolerations.1", map[string]interface{}{"effect": "NoSchedule", "key": "kubernetes.azure.com/scalesetpriority", "operator": "Equal", "value": "spot"})
	updatedByteValueOfSaveDeploymentTemplate := []byte(finalJson)

	log.Println("=== Here we are hitting SaveTemplate API ===")
	PipelineConfigRouter.HitSaveDeploymentTemplateApi(updatedByteValueOfSaveDeploymentTemplate, suite.authToken)

	log.Println("=== Here we are saving Global Configmap ===")
	requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(0, "-config1", createAppApiResponse.Id, "environment", "kubernetes", false, false, false, false)
	byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
	globalConfigMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
	configId = globalConfigMap.Result.Id

	log.Println("=== Here we are saving Global Secret ===")
	requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, "-secret1", createAppApiResponse.Id, "environment", "kubernetes", false, false, true, false)
	byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
	HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)

	log.Println("=== Here we are saving workflow with Pre/Post CI ===")
	workflowResponse := PipelineConfigRouter.HitCreateWorkflowApiWithFullPayload(createAppApiResponse.Id, suite.authToken).Result

	preStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
	postStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")

	log.Println("=== Here we are saving CD pipeline ===")
	payload := PipelineConfigRouter.GetRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, "AUTOMATIC", string(preStageScript), string(postStageScript), "AUTOMATIC")
	bytePayload, _ := json.Marshal(payload)
	savePipelineResponse := PipelineConfigRouter.HitSaveCdPipelineApi(bytePayload, suite.authToken)
	time.Sleep(2 * time.Second)

	log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.CiPipelines[0].Id, suite.authToken)

	log.Println("=== Here we are Triggering CI/CD and verifying CI/CD Deploy Status ===")
	triggerAndVerifyCiPipeline(createAppApiResponse, pipelineMaterial, workflowResponse.CiPipelines[0].Id, suite)

	log.Println("=== Here we are getting ResourceTree ===")
	ResourceTreeApiResponse := HitGetResourceTreeApi(createAppApiResponse.AppName, suite.authToken)
	container := ResourceTreeApiResponse.Result.PodMetadata[0].Containers[0]

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

	log.Println("=== Here we are Deleting the CD pipeline ===")
	deletePipelinePayload := PipelineConfigRouter.GetPayloadForDeleteCdPipeline(createAppApiResponse.Id, savePipelineResponse.Result.Pipelines[0].Id)
	deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
	PipelineConfigRouter.HitForceDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)

	log.Println("=== Here we are Deleting the CI pipeline ===")
	PipelineConfigRouter.DeleteCiPipeline(createAppApiResponse.Id, workflowResponse.CiPipelines[0].Id, suite.authToken)
	log.Println("=== Here we are Deleting CI Workflow ===")
	PipelineConfigRouter.HitDeleteWorkflowApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, suite.authToken)
	log.Println("=== Here we are Deleting the app after all verifications ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}

//todo will add other test cases once Devs will handle the validations
