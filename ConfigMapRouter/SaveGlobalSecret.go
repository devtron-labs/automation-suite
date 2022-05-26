package ConfigMapRouter

import (
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *ConfigsMapRouterTestSuite) TestClassA3SaveGlobalSecret() {
	config, _ := PipelineConfigRouter.GetEnvironmentConfigPipelineConfigRouter()
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
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(0, configName, createAppApiResponse.Id, environment, Kubernetes, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value1", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), Kubernetes+configName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=2=AddNewKubernetesSecretAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, Kubernetes, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		configId = configMap.Result.Id
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "value2", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), Kubernetes+newConfigName, configMap.Result.ConfigData[0].Name)
	})

	suite.Run("A=3=KubernetesSecretAsDataVolume", func() {
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, Kubernetes, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=4=KubernetesSecretAsDataVolumeHavingSubPath", func() {
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, Kubernetes, true, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=5=KubernetesSecretAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, Kubernetes, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("A=6=ExternalSecretAsEnvVariable", func() {
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, environment, ExternalKubernetes, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), environment, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("A=7=ExternalSecretAsDataVolume", func() {
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, ExternalKubernetes, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("A=8=ExternalSecretAsDataVolumeHavingSubPath", func() {
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, ExternalKubernetes, true, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("A=9=ExternalSecretAsDataVolumeHavingSubPathAndFilePermission", func() {
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, configName, createAppApiResponse.Id, volume, ExternalKubernetes, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key1)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=1=AddNewExternalSecretAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, ExternalKubernetes, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "", configMap.Result.ConfigData[0].Data.Key2)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=2=AWSSystemManagerAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, AWSSystemManager, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), "service/credentials", configMap.Result.ConfigData[0].SecretData[0].Key)
		assert.Equal(suite.T(), AWSSystemManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("B=3=AWSSystemManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, AWSSystemManager, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), AWSSystemManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("B=4=AWSSystemManagerAsDataVolumeHavingSubPath", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, AWSSystemManager, true, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), AWSSystemManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("B=5=AWSSystemManagerAsDataVolumeHavingSubPathAndFilePermission", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, AWSSystemManager, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=6=AddNewAWSSystemManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, AWSSystemManager, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), AWSSystemManager+newConfigName, configMap.Result.ConfigData[0].Name)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("B=7=AWSSecretsManagerAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, AWSSecretsManager, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), "service/credentials", configMap.Result.ConfigData[0].SecretData[0].Key)
		assert.Equal(suite.T(), AWSSecretsManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("B=8=AWSSecretsManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, AWSSecretsManager, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), AWSSecretsManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("B=9=AWSSecretsManagerAsDataVolumeHavingSubPath", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, AWSSecretsManager, true, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), AWSSecretsManager, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("C=1=AWSSecretsManagerAsDataVolumeHavingSubPathAndFilePermission", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, AWSSecretsManager, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("C=2=AddNewAWSSecretsManagerAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, AWSSecretsManager, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), AWSSecretsManager+newConfigName, configMap.Result.ConfigData[0].Name)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("C=3=HashiCorpVaultAsEnvVariable", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, environment, HashiCorpVault, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), "service/credentials", configMap.Result.ConfigData[0].SecretData[0].Key)
		assert.Equal(suite.T(), HashiCorpVault, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
	})

	suite.Run("C=4=HashiCorpVaultAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, HashiCorpVault, false, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), volume, configMap.Result.ConfigData[0].Type)
		assert.Equal(suite.T(), HashiCorpVault, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.Equal(suite.T(), "/directory-path", configMap.Result.ConfigData[0].MountPath)
	})

	suite.Run("C=5=HashiCorpVaultAsDataVolumeHavingSubPath", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, HashiCorpVault, true, false)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.Equal(suite.T(), HashiCorpVault, configMap.Result.ConfigData[0].ExternalSecretType)
		assert.True(suite.T(), configMap.Result.ConfigData[0].SubPath)
	})

	suite.Run("C=6=HashiCorpVaultAsDataVolumeHavingSubPathAndFilePermission", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, HashiCorpVault, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	suite.Run("C=7=AddNewHashiCorpVaultAsDataVolume", func() {
		newConfigName := strings.ToLower(Base.GetRandomStringOfGivenLength(6))
		requestPayloadForSecret := getRequestPayloadForSecretOrConfig(configId, newConfigName, createAppApiResponse.Id, volume, HashiCorpVault, true, true)
		byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
		log.Println("=== Hitting the SaveGlobalSecret API ====")
		configMap := HitSaveEnvironmentSecret(byteValueOfSecret, suite.authToken)
		assert.True(suite.T(), configMap.Result.ConfigData[0].External)
		assert.Equal(suite.T(), HashiCorpVault+newConfigName, configMap.Result.ConfigData[0].Name)
		assert.Equal(suite.T(), "0744", configMap.Result.ConfigData[0].FilePermission)
	})

	log.Println("=== Here we Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

}
