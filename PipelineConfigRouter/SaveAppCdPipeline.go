package PipelineConfigRouter

import (
	"automation-suite/ConfigMapRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassA4GetEnvironmentSecret() {
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

	log.Println("=== Here we are hitting SaveTemplate API ===")
	HitSaveDeploymentTemplateApi(byteValueOfSaveDeploymentTemplate, suite.authToken)

	log.Println("=== Here we are saving Global Configmap ===")
	configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
	requestPayloadForConfigMap := ConfigMapRouter.GetRequestPayloadForSecretOrConfig(0, configName, createAppApiResponse.Id, "environment", "kubernetes", false, false, false)
	byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
	ConfigMapRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)

	log.Println("=== Here we are saving Global Secret ===")
	requestPayloadForSecret := ConfigMapRouter.GetRequestPayloadForSecretOrConfig(0, configName, createAppApiResponse.Id, "environment", "kubernetes", false, false, true)
	byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
	ConfigMapRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)

	log.Println("=== Here we are saving workflow with Pre/Post CI ===")

	log.Println("=== Here we Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

}
