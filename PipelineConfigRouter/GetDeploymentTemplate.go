package PipelineConfigRouter

import (
	"automation-suite/testUtils"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassA8GetAppTemplate() {
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	log.Println("=== Here we are creating an App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	log.Println("=== Here we are printing AppName ===>", createAppApiResponse.AppName)
	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	time.Sleep(2 * time.Second)
	getChartReferenceResponse := HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	suite.Run("A=1=GetTemplateViaValidArgs", func() {
		getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), suite.authToken)
		assert.NotNil(suite.T(), getTemplateResponse.Result.GlobalConfig.DefaultAppOverride)
	})

	suite.Run("A=2=GetTemplateViaInvalidChartRefId", func() {
		invalidChartRefId := strconv.Itoa(testUtils.GetRandomNumberOf9Digit())
		getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), invalidChartRefId, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", getTemplateResponse.Errors[0].UserMessage)
	})

	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}

//todo need to add one more case for invalid AppId as well once dev will fix the issue for invalid app-id
