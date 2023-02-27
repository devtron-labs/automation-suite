package ConfigMapRouter

import (
	"automation-suite/HelperRouter"
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

// TestClassA1SaveGlobalConfigMap todo need to take call for some changes once product will final that config file should delete after app deletion or not
// TestClassA1SaveGlobalConfigMap todo once product will final that we can add config file before deployment template of not
func (suite *ConfigsMapRouterTestSuite) TestClassA1SaveGlobalConfigMap() {
	envConf := Base.ReadBaseEnvConfig()
	config := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result

	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := PipelineConfigRouter.GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := PipelineConfigRouter.HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := PipelineConfigRouter.GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	PipelineConfigRouter.HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	getChartReferenceResponse := PipelineConfigRouter.HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching DefaultAppOverride using getAppTemplateAPI ===")
	getTemplateResponse := PipelineConfigRouter.HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), suite.authToken)

	log.Println("=== Here we are fetching DefaultAppOverride using getAppTemplateAPI ===")
	defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

	log.Println("=== Here we are creating payload for SaveTemplate API ===")
	saveDeploymentTemplate := PipelineConfigRouter.GetRequestPayloadForSaveDeploymentTemplate(createAppApiResponse.Id, latestChartRef, defaultAppOverride)
	byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)

	log.Println("=== Here we are hitting SaveTemplate API ===")
	PipelineConfigRouter.HitSaveDeploymentTemplateApi(byteValueOfSaveDeploymentTemplate, suite.authToken)
	configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
	var configId int

	suite.Run("A=1=KubernetesConfigmapAsEnvVariable", func() {
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(0, configName, createAppApiResponse.Id, environment, kubernetes, false, false, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value1", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), kubernetes+configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=2=AddNewKubernetesConfigmapAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, kubernetes, false, false, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value2", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), kubernetes+newConfigName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=3=KubernetesConfigmapAsDataVolume", func() {
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, kubernetes, false, false, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=4=KubernetesConfigmapAsDataVolumeHavingSubPath", func() {
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, kubernetes, true, false, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=5=KubernetesConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, kubernetes, true, true, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("A=6=ExternalConfigmapAsEnvVariable", func() {
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, environment, externalKubernetes, false, false, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("A=7=ExternalConfigmapAsDataVolume", func() {
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, externalKubernetes, false, false, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=8=ExternalConfigmapAsDataVolumeHavingSubPath", func() {
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, externalKubernetes, true, false, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=9=ExternalConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, externalKubernetes, true, true, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=1=AddNewExternalConfigmapAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, externalKubernetes, true, true, false, true)
		byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), externalKubernetes+newConfigName, configMap.Result.ConfigData[0].Name)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	log.Println("=== Here we Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

}
