package ConfigMapRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

// TestClassA1SaveGlobalConfigMap todo need to take call for some changes once product will final that config file should delete after app deletion or not
// TestClassA1SaveGlobalConfigMap todo once product will final that we can add config file before deployment template of not
func (suite *ConfigsMapRouterTestSuite) TestClassA1SaveGlobalConfigMap() {
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
	var configId int
	suite.Run("A=1=KubernetesConfigmapAsEnvVariable", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(0, configName, createAppApiResponse.Id, environment, false, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value1", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=2=AddNewKubernetesConfigmapAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, newConfigName, createAppApiResponse.Id, environment, false, false, false, true)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value2", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), newConfigName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=3=KubernetesConfigmapAsDataVolume", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, false, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=4=KubernetesConfigmapAsDataVolumeHavingSubPath", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, false, true, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=5=KubernetesConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, false, true, true, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("A=6=ExternalConfigmapAsEnvVariable", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, environment, true, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("A=7=ExternalConfigmapAsDataVolume", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, true, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=8=ExternalConfigmapAsDataVolumeHavingSubPath", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, true, true, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=9=ExternalConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, configName, createAppApiResponse.Id, volume, true, true, true, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=1=AddNewExternalConfigmapAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveSecretOrConfigmap(configId, newConfigName, createAppApiResponse.Id, volume, true, true, true, true)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveGlobalConfigMap API ====")
		configMap := HitSaveGlobalConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})
}
