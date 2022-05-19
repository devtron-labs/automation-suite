package PipelineConfigRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelineConfigSuite) TestSaveAppCiPipelineWithValidPayload() {
	log.Println("=== getting Test Data for Hitting the SaveAppCiPipeline API ====")
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	createAppApiResponse := suite.createAppResponseDto.Result
	materialId := suite.createAppMaterialResponseDto.Result.Material[0].Id
	appName := createAppApiResponse.AppName
	requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/"+appName, config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, materialId)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	log.Println("=== Hitting the SaveAppCiPipeline API ====")
	saveAppCiPipelineResponse := HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)
	assert.Equal(suite.T(), saveAppCiPipelineResponse.Result.AppName, createAppApiResponse.AppName)
}

func (suite *PipelineConfigSuite) TestSaveAppCiPipelineWithInValidAppId() {
	log.Println("=== getting Test Data for Hitting the SaveAppCiPipeline API ====")
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	createAppApiResponse := suite.createAppResponseDto.Result
	createAppMaterialResponse := suite.createAppMaterialResponseDto.Result
	appName := createAppApiResponse.AppName
	appId := testUtils.GetRandomNumberOf9Digit()
	requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveAppCiPipeline(appId, config.DockerRegistry, config.DockerRegistry+"/"+appName, config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	log.Println("=== Hitting the SaveAppCiPipeline API with Invalid AppId ====")
	saveAppCiPipelineResponse := HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)
	assert.Equal(suite.T(), saveAppCiPipelineResponse.Errors[0].UserMessage, "pg: no rows in result set")
}

func (suite *PipelineConfigSuite) TestSaveAppCiPipelineWithInValidMaterialId() {
	log.Println("=== getting Test Data for Hitting the SaveAppCiPipeline API ====")
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	createAppApiResponse := suite.createAppResponseDto.Result
	appName := createAppApiResponse.AppName
	invalidMaterialId := testUtils.GetRandomNumberOf9Digit()
	requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/"+appName, config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, invalidMaterialId)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	log.Println("=== Hitting the SaveAppCiPipeline API with Invalid Material Id ====")
	saveAppCiPipelineResponse := HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)
	assert.Equal(suite.T(), saveAppCiPipelineResponse.Errors[0].UserMessage, "ERROR #23503 insert or update on table \"ci_template\" violates foreign key constraint \"ci_template_git_material_id_fkey\"")
}

func (suite *PipelineConfigSuite) TestSaveAppCiPipelineWithInValidDockerfileRepository() {
	log.Println("=== getting Test Data for Hitting the SaveAppCiPipeline API ====")
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	createAppApiResponse := suite.createAppResponseDto.Result
	materialId := suite.createAppMaterialResponseDto.Result.Material[0].Id
	appName := createAppApiResponse.AppName
	requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry+"invalid", config.DockerRegistry+"/"+appName, config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, materialId)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	log.Println("=== Hitting the SaveAppCiPipeline API with Invalid Docker file Repository ====")
	saveAppCiPipelineResponse := HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)
	assert.Equal(suite.T(), saveAppCiPipelineResponse.Errors[0].UserMessage, "pg: no rows in result set")
}
