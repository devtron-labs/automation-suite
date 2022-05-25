package ConfigMapRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

// TestClassC6SaveConfigMap todo need to take call for some changes once product will final that config file should delete after app deletion or not
// TestClassC6SaveConfigMap todo once product will final that we can add config file before deployment template of not
func (suite *ConfigsMapRouterTestSuite) TestClassC6SaveConfigMap() {
	createAppApiResponse := suite.createAppResponseDto.Result
	configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
	var configId int
	suite.Run("A=1=KubernetesConfigmapAsEnvVariable", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(0, configName, createAppApiResponse.Id, "environment", false, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value1", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=2=NewKubernetesConfigmapAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, newConfigName, createAppApiResponse.Id, "environment", false, false, false, true)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value2", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), newConfigName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=3=KubernetesConfigmapAsDataVolume", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createAppApiResponse.Id, "volume", false, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=4=KubernetesConfigmapAsDataVolumeHavingSubPath", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createAppApiResponse.Id, "volume", false, true, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=5=KubernetesConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createAppApiResponse.Id, "volume", false, true, true, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("A=6=ExternalConfigmapAsEnvVariable", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createAppApiResponse.Id, "environment", true, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("A=7=ExternalConfigmapAsDataVolume", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createAppApiResponse.Id, "volume", true, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=8=ExternalConfigmapAsDataVolumeHavingSubPath", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createAppApiResponse.Id, "volume", true, true, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=9=ExternalConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createAppApiResponse.Id, "volume", true, true, true, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveAppCiPipeline API ====")
		configMap := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})
}
