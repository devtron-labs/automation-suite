package PipelineConfigRouter

import (
	"automation-suite/HelperRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"log"
	"strconv"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassD3TriggerCiPipeline() {
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

	requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(0, "-config1", createAppApiResponse.Id, "environment", "kubernetes", false, false, false)
	byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
	globalConfigMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
	configId = globalConfigMap.Result.Id

	log.Println("=== Here we are saving Global Secret ===")
	requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, "-secret1", createAppApiResponse.Id, "environment", "kubernetes", false, false, true)
	byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
	HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)

	log.Println("=== Here we are saving workflow with Pre/Post CI ===")
	workflowResponse := HitCreateWorkflowApiWithFullPayload(createAppApiResponse.Id, suite.authToken).Result

	preStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
	postStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")

	log.Println("=== Here we are saving CD pipeline ===")
	payload := getRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, Automatic, string(preStageScript), string(postStageScript), Automatic)
	bytePayload, _ := json.Marshal(payload)
	savePipelineResponse := HitSaveCdPipelineApi(bytePayload, suite.authToken)

	log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := HitGetCiPipelineMaterial(workflowResponse.CiPipelines[0].Id, suite.authToken)

	suite.Run("A=1=GetWorkflowStatusWithoutTriggering", func() {
		payloadForTriggerCiPipeline := createPayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, workflowResponse.CiPipelines[0].Id, pipelineMaterial.Result[0].Id)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		triggerCiPipelineResponse := HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
		log.Println(triggerCiPipelineResponse)
	})

	/*suite.Run("A=2=GetWorkflowStatusWithInvalidAppId", func() {
		randomAppId := Base.GetRandomNumberOf9Digit()
		workflowStatus := HitGetWorkflowStatus(randomAppId, suite.authToken)
		assert.Nil(suite.T(), workflowStatus.Result.CiWorkflowStatus)
		assert.Nil(suite.T(), workflowStatus.Result.CdWorkflowStatus)
	})*/

	log.Println("=== Here we are Deleting the CD pipeline ===")
	deletePipelinePayload := GetPayloadForDeleteCdPipeline(createAppApiResponse.Id, savePipelineResponse.Result.Pipelines[0].Id)
	deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
	HitForceDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)

	log.Println("=== Here we are Deleting the CI pipeline ===")
	DeleteWorkflow(createAppApiResponse.Id, workflowResponse.CiPipelines[0].Id, suite.authToken)

	log.Println("=== Here we are Deleting the app after all verifications ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
