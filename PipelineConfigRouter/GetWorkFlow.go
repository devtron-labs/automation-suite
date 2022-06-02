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
	appId := suite.createAppResponseDto.Result.Id
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	createAppApiResponse := suite.createAppResponseDto.Result
	createAppMaterialResponse := suite.createAppMaterialResponseDto.Result

	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	log.Println("=== Hitting the SaveAppCiPipeline API ====")
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("Fetching suggested ci pipeline name ")
	fetchSuggestedCiPipelineName := HitGetPipelineSuggestedCiCd("ci", appId, suite.authToken)
	log.Println("Fetching gitMaterialId ")
	fetchAppGetResponseDto := HitGetMaterial(appId, suite.authToken)

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

	wfId := createWorkflowResponseDto.Result.AppWorkflowId
	DeleteWorkflow(appId, wfId, suite.authToken)
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
