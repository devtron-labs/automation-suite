package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassGetCiPipelineMin() {
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appId := createAppApiResponse.Id
	log.Println("=== AppName is===>", createAppApiResponse.AppName)
	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := HitCreateAppMaterialApi(appMaterialByteValue, appId, 1, false, suite.authToken).Result

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(appId, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	time.Sleep(2 * time.Second)
	getChartReferenceResponse := HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching Template using getAppTemplateAPI ===")
	getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), suite.authToken)

	log.Println("=== Here we are fetching DefaultAppOverride from template response ===")
	defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

	log.Println("=== Here we are creating payload for SaveTemplate API ===")
	saveDeploymentTemplate := GetRequestPayloadForSaveDeploymentTemplate(createAppApiResponse.Id, latestChartRef, defaultAppOverride)
	byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)

	log.Println("=== Here we are hitting SaveTemplate API ===")
	HitSaveDeploymentTemplateApi(byteValueOfSaveDeploymentTemplate, suite.authToken)

	log.Println("Fetching suggested ci pipeline name ")
	fetchSuggestedCiPipelineName := HitGetPipelineSuggestedCiCd("ci", appId, suite.authToken)

	log.Println("Fetching gitMaterialId")
	fetchAppGetResponseDto := HitGetApp(appId, suite.authToken)

	log.Println("Retrieving request payload for creating workflow from file")
	createWorkflowRequestDto := getRequestPayloadForCreateWorkflow(false, "1", appId, 0)
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].GitMaterialId = fetchAppGetResponseDto.Result.Material[0].Id
	createWorkflowRequestDto.CiPipeline.Name = fetchSuggestedCiPipelineName.Result
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].Source.Value = "main"

	log.Println("Here we are Patching CiPipelines")
	byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
	createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

	suite.Run("A=1=GetCiPipelineWithValidAppId", func() {
		ciPipelineMin := HitGetCiPipelineMin(strconv.Itoa(appId), suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), "NORMAL", ciPipelineMin.Result[0].PipelineType)
		assert.Equal(suite.T(), false, ciPipelineMin.Result[0].ScanEnabled)
		assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].Name, ciPipelineMin.Result[0].Name)
		assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].Id, ciPipelineMin.Result[0].Id)

		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)

	})

	suite.Run("A=2=GetCiPipelineWithEnabledScan", func() {
		log.Println("Retrieving request payload for creating workflow from file")
		createWorkflowRequestDto = getRequestPayloadForCreateWorkflow(false, "1", appId, 0)
		createWorkflowRequestDto.CiPipeline.CiMaterial[0].GitMaterialId = fetchAppGetResponseDto.Result.Material[0].Id
		createWorkflowRequestDto.CiPipeline.Name = fetchSuggestedCiPipelineName.Result
		createWorkflowRequestDto.CiPipeline.CiMaterial[0].Source.Value = "main"
		createWorkflowRequestDto.CiPipeline.ScanEnabled = true

		log.Println("Here we are Patching CiPipelines")
		byteValueOfCreateWorkflow, _ = json.Marshal(createWorkflowRequestDto)
		createWorkflowResponseDto = HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		ciPipelineMin := HitGetCiPipelineMin(strconv.Itoa(appId), suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), "NORMAL", ciPipelineMin.Result[0].PipelineType)
		assert.Equal(suite.T(), true, ciPipelineMin.Result[0].ScanEnabled)
		assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].Name, ciPipelineMin.Result[0].Name)
		assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].Id, ciPipelineMin.Result[0].Id)

		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	suite.Run("A=3=GetCiPipelineWithInvalidAppId", func() {
		randomAppId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		ciPipelineMin := HitGetCiPipelineMin(randomAppId, suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), "no ci pipeline found", ciPipelineMin.Errors[0].UserMessage)
	})

	log.Println("=== Here we are Deleting app after verification of flow ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
