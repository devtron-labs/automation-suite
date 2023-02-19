package PipelineConfigRouter

import (
	"automation-suite/HelperRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassB7SaveCdPipeline() {
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	var configId int
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result

	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	getChartReferenceResponse := HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching DefaultAppOverride using getAppTemplateAPI ===")
	getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), suite.authToken)

	log.Println("=== Here we are fetching DefaultAppOverride using getAppTemplateAPI ===")
	defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

	log.Println("=== Here we are creating payload for SaveTemplate API ===")
	saveDeploymentTemplate := GetRequestPayloadForSaveDeploymentTemplate(createAppApiResponse.Id, latestChartRef, defaultAppOverride)
	byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)

	log.Println("=== Here we are hitting SaveTemplate API ===")
	HitSaveDeploymentTemplateApi(byteValueOfSaveDeploymentTemplate, suite.authToken)

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
	workflowResponse := HitCreateWorkflowApiWithFullPayload(createAppApiResponse.Id, suite.authToken).Result
	preStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
	postStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")

	suite.Run("A=1=AutomaticStrategyWithGlobalSecretAndConfigMap", func() {
		payload := GetRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, Automatic, string(preStageScript), string(postStageScript), Automatic)
		bytePayload, _ := json.Marshal(payload)
		savePipelineResponse := HitSaveCdPipelineApi(bytePayload, suite.authToken)
		time.Sleep(3 * time.Second)
		assert.Equal(suite.T(), Automatic, savePipelineResponse.Result.Pipelines[0].TriggerType)
		assert.Equal(suite.T(), payload.Pipelines[0].Strategies, savePipelineResponse.Result.Pipelines[0].Strategies)
		assert.Equal(suite.T(), payload.Pipelines[0].PostStage.Config, savePipelineResponse.Result.Pipelines[0].PostStage.Config)
		assert.Equal(suite.T(), payload.Pipelines[0].PostStageConfigMapSecretNames, savePipelineResponse.Result.Pipelines[0].PostStageConfigMapSecretNames)
		deletePipelinePayload := GetPayloadForDeleteCdPipeline(createAppApiResponse.Id, savePipelineResponse.Result.Pipelines[0].Id)
		deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
		HitForceDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)
	})

	suite.Run("A=2=ManualStrategyWithAutomaticPipelineTriggerType", func() {
		payload := GetRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, Automatic, string(preStageScript), string(postStageScript), Manual)
		bytePayload, _ := json.Marshal(payload)
		savePipelineResponse := HitSaveCdPipelineApi(bytePayload, suite.authToken)
		time.Sleep(2 * time.Second)
		assert.Equal(suite.T(), Manual, savePipelineResponse.Result.Pipelines[0].TriggerType)
		assert.Equal(suite.T(), payload.Pipelines[0].Strategies, savePipelineResponse.Result.Pipelines[0].Strategies)
		assert.Equal(suite.T(), payload.Pipelines[0].PostStage.Config, savePipelineResponse.Result.Pipelines[0].PostStage.Config)
		assert.Equal(suite.T(), payload.Pipelines[0].PostStageConfigMapSecretNames, savePipelineResponse.Result.Pipelines[0].PostStageConfigMapSecretNames)

		deletePipelinePayload := GetPayloadForDeleteCdPipeline(createAppApiResponse.Id, savePipelineResponse.Result.Pipelines[0].Id)
		deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
		HitForceDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)
	})

	suite.Run("A=3=ManualStrategyWithManualPipelineTriggerType", func() {
		payload := GetRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, Manual, string(preStageScript), string(postStageScript), Manual)
		bytePayload, _ := json.Marshal(payload)
		savePipelineResponse := HitSaveCdPipelineApi(bytePayload, suite.authToken)
		time.Sleep(2 * time.Second)
		assert.Equal(suite.T(), Manual, savePipelineResponse.Result.Pipelines[0].TriggerType)
		assert.Equal(suite.T(), Manual, savePipelineResponse.Result.Pipelines[0].PostStage.TriggerType)
		assert.Equal(suite.T(), Manual, savePipelineResponse.Result.Pipelines[0].PreStage.TriggerType)

		deletePipelinePayload := GetPayloadForDeleteCdPipeline(createAppApiResponse.Id, savePipelineResponse.Result.Pipelines[0].Id)
		deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
		HitForceDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)
	})

	suite.Run("A=4=SaveAppCdWithInvalidAppID", func() {
		payload := GetRequestPayloadForSaveCdPipelineApi(Base.GetRandomNumberOf9Digit(), workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, Manual, string(preStageScript), string(postStageScript), Manual)
		bytePayload, _ := json.Marshal(payload)
		savePipelineResponse := HitSaveCdPipelineApi(bytePayload, suite.authToken)
		time.Sleep(2 * time.Second)
		assert.Equal(suite.T(), "pg: no rows in result set", savePipelineResponse.Errors[0].UserMessage)
	})
	log.Println("=== Here we are Deleting the CI pipeline ===")
	DeleteCiPipeline(createAppApiResponse.Id, workflowResponse.CiPipelines[0].Id, suite.authToken)
	log.Println("=== Here we are Deleting CI Workflow ===")
	HitDeleteWorkflowApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, suite.authToken)
	log.Println("=== Here we Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}

//todo There is no handling of parentPipelineId and WorkflowId ,need to add test cases once issue fixed from dev side
