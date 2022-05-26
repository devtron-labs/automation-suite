package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassA9SaveDeploymentTemplate() {

	suite.Run("A=1=SaveDeploymentTemplateWithDefaultAppOverride", func() {
		config, _ := GetEnvironmentConfigPipelineConfigRouter()

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

		log.Println("=== Here we hitting 7 verifying SaveTemplate API ===")
		saveDeploymentTemplateResponse := HitSaveDeploymentTemplateApi(byteValueOfSaveDeploymentTemplate, suite.authToken)
		assert.Equal(suite.T(), createAppApiResponse.Id, saveDeploymentTemplateResponse.Result.AppId)
		assert.Equal(suite.T(), latestChartRef, saveDeploymentTemplateResponse.Result.ChartRefId)
		assert.Equal(suite.T(), defaultAppOverride, saveDeploymentTemplateResponse.Result.DefaultAppOverride)

		log.Println("=== Here we Deleting the Test data created after verification ===")
		Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

	})

	suite.Run("A=2=SaveDeploymentTemplateWithSomeChangesInDefaultOverride", func() {
		config, _ := GetEnvironmentConfigPipelineConfigRouter()

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

		log.Println("=== Here we are changing some values in default override ===")
		saveDeploymentTemplate.ValuesOverride.GracePeriod = 50
		saveDeploymentTemplate.ValuesOverride.Resources.Limits.Memory = "100Mi"
		byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)

		log.Println("=== Here we hitting 7 verifying SaveTemplate API ===")
		saveDeploymentTemplateResponse := HitSaveDeploymentTemplateApi(byteValueOfSaveDeploymentTemplate, suite.authToken)
		assert.Equal(suite.T(), createAppApiResponse.Id, saveDeploymentTemplateResponse.Result.AppId)
		assert.Equal(suite.T(), 50, saveDeploymentTemplateResponse.Result.DefaultAppOverride.GracePeriod)
		assert.Equal(suite.T(), "100Mi", saveDeploymentTemplateResponse.Result.DefaultAppOverride.Resources.Limits.Memory)
		assert.Equal(suite.T(), latestChartRef, saveDeploymentTemplateResponse.Result.ChartRefId)

		log.Println("=== Here we Deleting the Test data created after verification ===")
		Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

	})

}
