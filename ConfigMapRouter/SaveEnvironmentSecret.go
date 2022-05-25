package ConfigMapRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *ConfigsMapRouterTestSuite) TestClassA3SaveEnvironmentSecret() {
	createAppApiResponse := suite.createAppResponseDto.Result
	configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
	var configId int
	suite.Run("A=1=KubernetesSecretAsEnvVariable", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(0, configName, createAppApiResponse.Id, environment, false, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value1", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=2=AddNewKubernetesSecretAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, newConfigName, createAppApiResponse.Id, environment, false, false, false, true)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value2", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), newConfigName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=3=KubernetesSecretAsDataVolume", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, false, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=4=KubernetesSecretAsDataVolumeHavingSubPath", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, false, true, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=5=KubernetesSecretAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, false, true, true, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("A=6=ExternalSecretAsEnvVariable", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, environment, true, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("A=7=ExternalSecretAsDataVolume", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, true, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=8=ExternalSecretAsDataVolumeHavingSubPath", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, true, true, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=9=ExternalSecretAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, true, true, true, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=1=AddNewExternalSecretAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, newConfigName, createAppApiResponse.Id, volume, true, true, true, true)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveEnvSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})
}
