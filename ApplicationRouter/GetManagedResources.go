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

func (suite *ApplicationsRouterTestSuite) TestClassGetManagedResources() {
	config, _ := PipelineConfigRouter.GetEnvironmentConfigPipelineConfigRouter()
	var configId int
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appName := createAppApiResponse.AppName
	log.Println("=== App Name is :====", createAppApiResponse.AppName)

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
