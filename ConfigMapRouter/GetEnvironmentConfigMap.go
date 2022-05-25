package ConfigMapRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *ConfigsMapRouterTestSuite) TestClassA2GetEnvironmentConfigMap() {
	var createdAppId, configId int
	var createdAppResponse Base.CreateAppResponseDto
	configName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))

	suite.Run("A=1=EnvironmentConfigMapWithoutCreatingCM", func() {
		randomAppId := Base.GetRandomNumberOf9Digit()
		randomEnvId := Base.GetRandomNumberOf9Digit()
		envConfigResponse := HitGetEnvironmentConfigMap(randomAppId, randomEnvId, suite.authToken)
		log.Println("Validating the response of GetEnvConfig API")
		assert.Empty(suite.T(), envConfigResponse.Result.ConfigData)
	})

	suite.Run("A=2=KubernetesConfigmapAsEnvVariable", func() {
		log.Println("=== Here We are creating a new App ===")
		createdAppResponse = Base.CreateApp(suite.authToken)
		createdAppId = createdAppResponse.Result.Id

		log.Println("=== Here We are saving a config map ===")
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(0, configName, createdAppId, "environment", false, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		savedResponse := HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configId = savedResponse.Result.Id

		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		log.Println("Validating the response of GetEnvConfig API")
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value1", configMap.Result.ConfigData[0].DefaultData.Key1)
		assert.Equal(suite.T(), configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=3=AddNewKubernetesConfigmapAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, newConfigName, createdAppId, "environment", false, false, false, true)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)

		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		log.Println("Validating the response of GetEnvConfig API")
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value2", configMap.Result.ConfigData[0].DefaultData.Key2)
		assert.Equal(suite.T(), configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=4=KubernetesConfigmapAsDataVolume", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createdAppId, "volume", false, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].DefaultMountPath)
	})

	suite.Run("A=5=KubernetesConfigmapAsDataVolumeHavingSubPath", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createdAppId, "volume", false, true, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=6=KubernetesConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createdAppId, "volume", false, true, true, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("A=7=ExternalConfigmapAsEnvVariable", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createdAppId, "environment", true, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		assert.Equal(suite.T(), "environment", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("A=8=ExternalConfigmapAsDataVolume", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createdAppId, "volume", true, false, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].DefaultMountPath)
	})

	suite.Run("A=9=ExternalConfigmapAsDataVolumeHavingSubPath", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createdAppId, "volume", true, true, false, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		assert.Equal(suite.T(), "volume", configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("B=1=ExternalConfigmapAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, configName, createdAppId, "volume", true, true, true, false)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=2=AddNewExternalConfigmapAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSaveAppCiPipeline := getRequestPayloadForSaveConfigmap(configId, newConfigName, createdAppId, "volume", true, true, true, true)
		byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
		log.Println("=== Hitting the SaveConfigMap API ====")
		HitSaveConfigMap(byteValueOfSaveAppCiPipeline, suite.authToken)
		configMap := HitGetEnvironmentConfigMap(createdAppId, 1, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)

		log.Println("=== Here We are Deleting the test data created for Automation ===")
		Base.DeleteApp(createdAppId, createdAppResponse.Result.AppName, createdAppResponse.Result.TeamId, createdAppResponse.Result.TemplateId, suite.authToken)

	})
}
