package PipelineConfigRouter

import (
	"automation-suite/testUtils"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassC8GetWorkflows() {
	envConf := Base.ReadBaseEnvConfig()
	config := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)

	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appId := createAppApiResponse.Id

	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)

	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	log.Println("=== Hitting the SaveAppCiPipeline API ====")
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("Fetching suggested ci pipeline name ")
	fetchSuggestedCiPipelineName := HitGetPipelineSuggestedCiCd("ci", appId, suite.authToken)
	log.Println("Fetching gitMaterialId ")
	fetchAppGetResponseDto := HitGetApp(appId, suite.authToken)

	log.Println("Retrieving request payload from file")
	createWorkflowRequestDto := getRequestPayloadForCreateWorkflow(false, "1", appId, 0)

	createWorkflowRequestDto.CiPipeline.CiMaterial[0].GitMaterialId = fetchAppGetResponseDto.Result.Material[0].Id
	createWorkflowRequestDto.CiPipeline.Name = fetchSuggestedCiPipelineName.Result
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].Source.Value = strings.ToLower(testUtils.GetRandomStringOfGivenLength(10))
	createWorkflowResponseDto := HitCreateWorkflowApiWithFullPayload(appId, suite.authToken)

	suite.Run("A=1=FetchAllAppWorkflowWithValidAppId", func() {
		fetchAllAppWorkflowResponseDto := FetchAllAppWorkflow(appId, suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), 200, fetchAllAppWorkflowResponseDto.Code)
		assert.Equal(suite.T(), appId, fetchAllAppWorkflowResponseDto.Result.AppId)
	})
	suite.Run("A=2=FetchAllAppWorkflowWithInvalidAppId", func() {
		fetchAllAppWorkflowResponseDto := FetchAllAppWorkflow(testUtils.GetRandomNumberOf9Digit(), suite.authToken)
		log.Println("Validating the response of FetchAllLink API")
		assert.Equal(suite.T(), "pg: no rows in result set", fetchAllAppWorkflowResponseDto.Error[0].UserMessage)
	})

	log.Println("=== Here we are Deleting the CI pipeline ===")
	DeleteCiPipeline(createAppApiResponse.Id, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
	log.Println("=== Here we are Deleting CI Workflow ===")
	HitDeleteWorkflowApi(createAppApiResponse.Id, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	log.Println("=== Here we Deleting the Test data created after verification ===")
	Base.DeleteApp(appId, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
