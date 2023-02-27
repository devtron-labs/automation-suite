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

// todo need to take a call for data we are not getting config data in GetEnvSecret API

func (suite *ConfigsMapRouterTestSuite) TestClassA4GetEnvironmentSecret() {
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

	suite.Run("A=1=KubernetesSecretAsEnvVariable", func() {
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(0, configName, createAppApiResponse.Id, environment, kubernetes, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		configId = configMap.Result.Id
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData)
		assert.Equal(suite.T(), environment, environmentSecret.Result.ConfigData[index-1].Type)
		assert.Equal(suite.T(), kubernetes+configName, environmentSecret.Result.ConfigData[index-1].Name)
	})

	suite.Run("A=2=AddNewKubernetesSecretAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, kubernetes, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), environment, environmentSecret.Result.ConfigData[index].Type)
		assert.Equal(suite.T(), kubernetes+newConfigName, environmentSecret.Result.ConfigData[index].Name)
	})

	suite.Run("A=3=KubernetesSecretAsDataVolume", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, kubernetes, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), volume, environmentSecret.Result.ConfigData[index].Type)
	})

	suite.Run("A=4=KubernetesSecretAsDataVolumeHavingSubPath", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, kubernetes, true, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].SubPath)
	})

	suite.Run("A=5=KubernetesSecretAsDataVolumeHavingSubPathAndFilePermission", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, kubernetes, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	suite.Run("A=6=ExternalSecretAsEnvVariable", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, environment, externalKubernetes, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), environment, environmentSecret.Result.ConfigData[index].Type)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
	})

	suite.Run("A=7=ExternalSecretAsDataVolume", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, externalKubernetes, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), volume, environmentSecret.Result.ConfigData[index].Type)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
	})

	suite.Run("A=8=ExternalSecretAsDataVolumeHavingSubPath", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, externalKubernetes, true, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), volume, environmentSecret.Result.ConfigData[index].Type)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].SubPath)
	})

	suite.Run("A=9=ExternalSecretAsDataVolumeHavingSubPathAndFilePermission", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, externalKubernetes, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
		assert.Equal(suite.T(), "", environmentSecret.Result.ConfigData[index].Data.Key1)
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	suite.Run("B=1=AddNewExternalSecretAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, externalKubernetes, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
		assert.Equal(suite.T(), "", environmentSecret.Result.ConfigData[index].Data.Key2)
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	suite.Run("B=2=AWSSystemManagerAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, awsSystemManager, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), awsSystemManager, environmentSecret.Result.ConfigData[index].ExternalSecretType)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
	})

	suite.Run("B=3=AWSSystemManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, awsSystemManager, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), volume, environmentSecret.Result.ConfigData[index].Type)
		assert.Equal(suite.T(), awsSystemManager, environmentSecret.Result.ConfigData[index].ExternalSecretType)
	})

	suite.Run("B=4=AWSSystemManagerAsDataVolumeHavingSubPath", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, awsSystemManager, true, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), awsSystemManager, environmentSecret.Result.ConfigData[index].ExternalSecretType)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].SubPath)
	})

	suite.Run("B=5=AWSSystemManagerAsDataVolumeHavingSubPathAndFilePermission", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, awsSystemManager, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	suite.Run("B=6=AddNewAWSSystemManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, awsSystemManager, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[0].External)
		assert.Equal(suite.T(), awsSystemManager+newConfigName, environmentSecret.Result.ConfigData[index].Name)
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	suite.Run("B=7=AWSSecretsManagerAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, awsSecretsManager, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), awsSecretsManager, environmentSecret.Result.ConfigData[index].ExternalSecretType)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
	})

	suite.Run("B=8=AWSSecretsManagerAsDataVolume", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, awsSecretsManager, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), volume, environmentSecret.Result.ConfigData[index].Type)
		assert.Equal(suite.T(), awsSecretsManager, environmentSecret.Result.ConfigData[index].ExternalSecretType)
	})

	suite.Run("B=9=AWSSecretsManagerAsDataVolumeHavingSubPath", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, awsSecretsManager, true, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), awsSecretsManager, environmentSecret.Result.ConfigData[index].ExternalSecretType)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].SubPath)
	})

	suite.Run("C=1=AWSSecretsManagerAsDataVolumeHavingSubPathAndFilePermission", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, awsSecretsManager, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	suite.Run("C=2=AddNewAWSSecretsManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, awsSecretsManager, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
		assert.Equal(suite.T(), awsSecretsManager+newConfigName, environmentSecret.Result.ConfigData[index].Name)
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	suite.Run("C=3=HashiCorpVaultAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, hashiCorpVault, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), hashiCorpVault, environmentSecret.Result.ConfigData[index].ExternalSecretType)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
	})

	suite.Run("C=4=HashiCorpVaultAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, hashiCorpVault, false, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), volume, environmentSecret.Result.ConfigData[index].Type)
		assert.Equal(suite.T(), hashiCorpVault, environmentSecret.Result.ConfigData[index].ExternalSecretType)
	})

	suite.Run("C=5=HashiCorpVaultAsDataVolumeHavingSubPath", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, hashiCorpVault, true, false, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.Equal(suite.T(), hashiCorpVault, environmentSecret.Result.ConfigData[index].ExternalSecretType)
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].SubPath)
	})

	suite.Run("C=6=HashiCorpVaultAsDataVolumeHavingSubPathAndFilePermission", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, hashiCorpVault, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	suite.Run("C=7=AddNewHashiCorpVaultAsDataVolume", func() {
		configNameStr := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, configNameStr, createAppApiResponse.Id, volume, hashiCorpVault, true, true, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)
		environmentSecret := HelperRouter.HitGetEnvironmentSecretApi(createAppApiResponse.Id, 1, suite.authToken)
		index := len(environmentSecret.Result.ConfigData) - 1
		assert.True(suite.T(), environmentSecret.Result.ConfigData[index].External)
		assert.Equal(suite.T(), hashiCorpVault+configNameStr, environmentSecret.Result.ConfigData[index].Name)
		assert.Equal(suite.T(), "0744", environmentSecret.Result.ConfigData[index].FilePermission)
	})

	log.Println("=== Here we Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

}
