package ConfigMapRouter

const (
	SaveGlobalConfigmapApiUrl     string = "/orchestrator/config/global/cm"
	SaveGlobalConfigmapApi        string = "SaveGlobalConfigmapApi"
	GetEnvironmentConfigMapApiUrl string = "/orchestrator/config/environment/cm/"
	GetEnvironmentConfigMapApi    string = "GetEnvironmentConfigMapApi"
	SaveGlobalSecretApiUrl        string = "/orchestrator/config/global/cs"
	SaveGlobalSecretApi           string = "SaveGlobalSecretApi"
	awsSystemManager              string = "awsSystemManager"
	hashiCorpVault                string = "hashiCorpVault"
	awsSecretsManager             string = "awsSecretsManager"
	kubernetes                    string = "kubernetes"
	externalKubernetes            string = "externalKubernetes"
	environment                   string = "environment"
	volume                        string = "volume"
	GetEnvSecretApiUrl            string = "/orchestrator/config/environment/cs/"
	GetEnvSecretApi               string = "GetEnvSecretApi"
)
