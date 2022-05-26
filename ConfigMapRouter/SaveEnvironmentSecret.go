package ConfigMapRouter

import (
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *ConfigsMapRouterTestSuite) TestClassA3SaveEnvironmentSecret() {
	config, _ := PipelineConfigRouter.GetEnvironmentConfigPipelineConfigRouter()

	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result

	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := PipelineConfigRouter.GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := PipelineConfigRouter.HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)

	log.Println("=== Here we saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := PipelineConfigRouter.GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	PipelineConfigRouter.HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
	var configId int
	suite.Run("A=1=KubernetesSecretAsEnvVariable", func() {
		requestPayloadForSecret := getRequestPayloadForSecret(0, configName, createAppApiResponse.Id, environment, KubernetesSecret, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value1", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), KubernetesSecret+configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=2=AddNewKubernetesSecretAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecret(configId, newConfigName, createAppApiResponse.Id, environment, KubernetesSecret, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value2", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), KubernetesSecret+newConfigName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=3=KubernetesSecretAsDataVolume", func() {
		requestPayloadForSecret := getRequestPayloadForSecret(configId, configName, createAppApiResponse.Id, volume, KubernetesSecret, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=4=KubernetesSecretAsDataVolumeHavingSubPath", func() {
		requestPayloadForSecret := getRequestPayloadForSecret(configId, configName, createAppApiResponse.Id, volume, KubernetesSecret, true, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=5=KubernetesSecretAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSecret := getRequestPayloadForSecret(configId, configName, createAppApiResponse.Id, volume, KubernetesSecret, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("A=6=ExternalSecretAsEnvVariable", func() {
		requestPayloadForSecret := getRequestPayloadForSecret(configId, configName, createAppApiResponse.Id, environment, ExternalKubernetesSecret, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("A=7=ExternalSecretAsDataVolume", func() {
		requestPayloadForSecret := getRequestPayloadForSecret(configId, configName, createAppApiResponse.Id, volume, ExternalKubernetesSecret, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=8=ExternalSecretAsDataVolumeHavingSubPath", func() {
		requestPayloadForSecret := getRequestPayloadForSecret(configId, configName, createAppApiResponse.Id, volume, ExternalKubernetesSecret, true, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=9=ExternalSecretAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSecret := getRequestPayloadForSecret(configId, configName, createAppApiResponse.Id, volume, ExternalKubernetesSecret, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=1=AddNewExternalSecretAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecret(configId, newConfigName, createAppApiResponse.Id, volume, ExternalKubernetesSecret, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=2=AWSSystemManagerAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecret(configId, newConfigName, createAppApiResponse.Id, environment, AWSSystemManager, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), "service/credentials", configMap.Result.ConfigData[0].SecretData[0].Key)
		assert.Equal(suite.T(), AWSSystemManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("B=3=AWSSystemManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecret(configId, newConfigName, createAppApiResponse.Id, volume, AWSSystemManager, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), AWSSystemManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("B=4=AWSSystemManagerAsDataVolumeHavingSubPath", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecret(configId, newConfigName, createAppApiResponse.Id, volume, AWSSystemManager, true, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), AWSSystemManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("B=5=AWSSystemManagerAsDataVolumeHavingSubPathAndFilePermission", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecret(configId, newConfigName, createAppApiResponse.Id, volume, AWSSystemManager, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=6=AddNewAWSSystemManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecret(configId, newConfigName, createAppApiResponse.Id, volume, AWSSystemManager, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), AWSSystemManager+newConfigName, configMap.Result.ConfigData[0].Name)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
		Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
	})

}
