package PipelineConfigRouter

import (
	"automation-suite/HelperRouter"
	"automation-suite/PipelineConfigRouter/ResponseDTOs"
	dtos "automation-suite/PipelineConfigRouter/ResponseDTOs"
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/tidwall/sjson"
	"log"
	"strconv"
	"time"
)

var (
	createAppApiResponsePtr *testUtils.CreateAppRequestDto
	workflowResponsePtr     *dtos.CreateWorkflowResponseDto
	savePipelineResponsePtr *ResponseDTOs.SaveCdPipelineResponseDTO
)

const SuccessCode = 200

func CreateNewAppWithCiCd(authToken string) (testUtils.CreateAppRequestDto, dtos.CreateWorkflowResponseDto) {
	if createAppApiResponsePtr != nil && workflowResponsePtr != nil {
		return *createAppApiResponsePtr, *workflowResponsePtr
	}
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	var configId int
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := testUtils.CreateApp(authToken).Result
	log.Println("=== App Name is :====>", createAppApiResponse.AppName)

	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, authToken)

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, "test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	time.Sleep(2 * time.Second)
	getChartReferenceResponse := HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching Template using getAppTemplateAPI ===")
	getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), authToken)

	log.Println("=== Here we are fetching DefaultAppOverride from template response ===")
	defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

	log.Println("=== Here we are creating payload for SaveTemplate API ===")
	saveDeploymentTemplate := GetRequestPayloadForSaveDeploymentTemplate(createAppApiResponse.Id, latestChartRef, defaultAppOverride)
	byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)
	jsonOfSaveDeploymentTemp := string(byteValueOfSaveDeploymentTemplate)
	jsonWithMicroserviceToleration, _ := sjson.Set(jsonOfSaveDeploymentTemp, "valuesOverride.tolerations.0", map[string]interface{}{"effect": "NoSchedule", "key": "microservice", "operator": "Equal", "value": "true"})
	finalJson, _ := sjson.Set(jsonWithMicroserviceToleration, "valuesOverride.tolerations.1", map[string]interface{}{"effect": "NoSchedule", "key": "kubernetes.azure.com/scalesetpriority", "operator": "Equal", "value": "spot"})
	updatedByteValueOfSaveDeploymentTemplate := []byte(finalJson)

	log.Println("=== Here we are hitting SaveTemplate API ===")
	HitSaveDeploymentTemplateApi(updatedByteValueOfSaveDeploymentTemplate, authToken)

	log.Println("=== Here we are saving Global Configmap ===")
	requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(0, "-config1", createAppApiResponse.Id, "environment", "kubernetes", false, false, false, false)
	byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
	globalConfigMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, authToken)
	configId = globalConfigMap.Result.Id

	log.Println("=== Here we are saving Global Secret ===")
	requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, "-secret1", createAppApiResponse.Id, "environment", "kubernetes", false, false, true, false)
	byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
	HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, authToken)

	log.Println("=== Here we are saving workflow with Pre/Post CI ===")
	workflowResponse := HitCreateWorkflowApiWithFullPayload(createAppApiResponse.Id, authToken)
	preStageScript, _ := testUtils.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
	postStageScript, _ := testUtils.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")

	log.Println("=== Here we are saving CD pipeline ===")
	payload := GetRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.Result.AppWorkflowId, 1, workflowResponse.Result.CiPipelines[0].Id, workflowResponse.Result.CiPipelines[0].ParentCiPipeline, "AUTOMATIC", string(preStageScript), string(postStageScript), "AUTOMATIC")
	bytePayload, _ := json.Marshal(payload)
	savePipelineResponse := HitSaveCdPipelineApi(bytePayload, authToken)

	createAppApiResponsePtr = &createAppApiResponse
	workflowResponsePtr = &workflowResponse
	savePipelineResponsePtr = &savePipelineResponse
	return createAppApiResponse, workflowResponse
	//clean created data

}

func DeleteAppWithCiCd(authToken string) bool {
	if createAppApiResponsePtr == nil || workflowResponsePtr == nil || savePipelineResponsePtr == nil {
		log.Println("=== No app is present ===")
		return true
	}
	log.Println("=== Here we are Deleting the CD pipeline ===")
	deletePipelinePayload := GetPayloadForDeleteCdPipeline(createAppApiResponsePtr.Id, savePipelineResponsePtr.Result.Pipelines[0].Id)
	deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
	HitForceDeleteCdPipelineApi(deletePipelineByteCode, authToken)

	log.Println("=== Here we are Deleting the CI pipeline ===")
	DeleteCiPipeline(createAppApiResponsePtr.Id, workflowResponsePtr.Result.CiPipelines[0].Id, authToken)
	log.Println("=== Here we are Deleting CI Workflow ===")
	HitDeleteWorkflowApi(createAppApiResponsePtr.Id, workflowResponsePtr.Result.AppWorkflowId, authToken)
	log.Println("=== Here we are Deleting the app after all verifications ===")
	deleteResponse := testUtils.DeleteApp(createAppApiResponsePtr.Id, createAppApiResponsePtr.AppName, createAppApiResponsePtr.TeamId, createAppApiResponsePtr.TemplateId, authToken)
	return deleteResponse.Code == SuccessCode
}
