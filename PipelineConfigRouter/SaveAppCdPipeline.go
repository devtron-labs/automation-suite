package PipelineConfigRouter

import (
	"automation-suite/ConfigMapRouter"
	"automation-suite/HelperRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassA5SaveAppCdPipeline() {
	preStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
	postStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")
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
	requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(0, "config1", createAppApiResponse.Id, "environment", "kubernetes", false, false, false)
	configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
	requestPayloadForConfigMap := ConfigMapRouter.GetRequestPayloadForSecretOrConfig(0, configName, createAppApiResponse.Id, "environment", "kubernetes", false, false, false)
	byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
	globalConfigMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
	configId = globalConfigMap.Result.Id

	log.Println("=== Here we are saving Global Secret ===")
	requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, "secret1", createAppApiResponse.Id, "environment", "kubernetes", false, false, true)
	byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
	HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
	ConfigMapRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)

	log.Println("=== Here we are saving workflow with Pre/Post CI ===")
	workflowResponse := HitCreateWorkflowApiWithFullPayload(createAppApiResponse.Id, suite.authToken).Result

	suite.Run("A=1=AutomaticStrategyWithGlobalSecretAndConfigMap", func() {
		payload := getRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, Automatic, string(preStageScript), string(postStageScript), Automatic)
		bytePayload, _ := json.Marshal(payload)
		savePipelineResponse := HitSaveCdPipelineApi(bytePayload, suite.authToken)
		deletePipelinePayload := GetPayloadForDeleteCdPipeline(createAppApiResponse.Id, savePipelineResponse.Result.Pipelines[0].CiPipelineId)
		deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
		HitDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)
	})

	log.Println("=== Here we Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

}
