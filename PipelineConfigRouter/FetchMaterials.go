package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassFetchMaterials() {
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

	suite.Run("A=1=FetchMaterialsWithValidCiPipelineId", func() {
		appMaterials := HitFetchMaterialsApi(strconv.Itoa(createWorkflowResponseDto.Result.CiPipelines[0].Id), suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), "https://github.com/devtron-labs/sample-go-app.git", appMaterials.Result[0].GitMaterialUrl)
		assert.Equal(suite.T(), "SOURCE_TYPE_BRANCH_FIXED", appMaterials.Result[0].Type)
		assert.Equal(suite.T(), "main", appMaterials.Result[0].Value)
		assert.NotNil(suite.T(), appMaterials.Result[0].History)
	})

	suite.Run("A=2=FetchMaterialsWithValidCiPipelineId", func() {
		randomAppId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		appMaterials := HitFetchMaterialsApi(randomAppId, suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), "pg: no rows in result set", appMaterials.Errors[0].UserMessage)
		assert.Equal(suite.T(), 404, appMaterials.Code)
	})

	log.Println("=== Here we are Deleting the CI pipeline ===")
	DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
	log.Println("=== Here we are Deleting CI Workflow ===")
	HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	log.Println("=== Here we are Deleting app after verification of flow ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
