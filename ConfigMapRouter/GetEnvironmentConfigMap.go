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

func (suite *ConfigsMapRouterTestSuite) TestClassA2GetEnvironmentConfigMap() {
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
	var configId int

	suite.Run("A=1=EnvironmentConfigMapWithoutCreatingCM", func() {
		randomAppId := Base.GetRandomNumberOf9Digit()
		randomEnvId := Base.GetRandomNumberOf9Digit()
		envConfigResponse := HelperRouter.HitGetEnvironmentConfigMap(randomAppId, randomEnvId, suite.authToken)
		log.Println("Validating the response of GetEnvConfig API")
		assert.Empty(suite.T(), envConfigResponse.Result.ConfigData)
	})

	suite.Run("A=2=KubernetesConfigmapAsEnvVariable", func() {
		configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		log.Println("=== Here We are saving a global config map ===")
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(0, configName, createAppApiResponse.Id, environment, kubernetes, false, false, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		savedResponse := HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configId = savedResponse.Result.Id

		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		log.Println("Validating the response of GetEnvConfig API")
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value1", configMap.Result.ConfigData[0].DefaultData.Key1)
		assert.Equal(suite.T(), kubernetes+configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=3=AddNewKubernetesConfigmapAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, kubernetes, false, false, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)

		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		log.Println("Validating the response of GetEnvConfig API")
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[1].Type)
		assert.Equal(suite.T(), "value2", configMap.Result.ConfigData[1].DefaultData.Key2)
		assert.Equal(suite.T(), kubernetes+newConfigName, configMap.Result.ConfigData[1].Name)
	})

	suite.Run("A=4=KubernetesConfigmapAsDataVolume", func() {
		configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, kubernetes, false, false, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[2].Type)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[2].DefaultMountPath)
	})

	suite.Run("A=5=KubernetesConfigmapAsDataVolumeHavingSubPath", func() {
		configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, kubernetes, true, false, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[3].SubPath)
	})

	suite.Run("A=6=KubernetesConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, kubernetes, true, true, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[4].FilePermission)
	})

	suite.Run("A=7=ExternalConfigmapAsEnvVariable", func() {
		configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, environment, externalKubernetes, false, false, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[5].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[5].External)
	})

	suite.Run("A=8=ExternalConfigmapAsDataVolume", func() {
		configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, externalKubernetes, false, false, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[6].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[6].External)
	})

	suite.Run("A=9=ExternalConfigmapAsDataVolumeHavingSubPath", func() {
		configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, externalKubernetes, true, false, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[7].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[7].External)
		assert.True(suite.T(), configMap.Result.ConfigData[7].SubPath)
	})

	suite.Run("B=1=ExternalConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, externalKubernetes, true, true, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[8].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[8].Data.Key1)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[8].DefaultMountPath)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[8].FilePermission)
	})

	suite.Run("B=2=AddNewExternalConfigmapAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, externalKubernetes, true, true, false, true)
		byteRequestPayloadForSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		HelperRouter.HitSaveGlobalConfigMap(byteRequestPayloadForSecret, suite.authToken)
		configMap := HelperRouter.HitGetEnvironmentConfigMap(createAppApiResponse.Id, 1, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[9].External)
		assert.Equal(suite.T(), externalKubernetes+newConfigName, configMap.Result.ConfigData[9].Name)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[9].FilePermission)
	})
	log.Println("=== Here We are Deleting the test data created for Automation ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

}
