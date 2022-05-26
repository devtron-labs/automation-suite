package ConfigMapRouter

const (
	SaveGlobalConfigmapApiUrl     string = "/orchestrator/config/global/cm"
	SaveGlobalConfigmapApi        string = "SaveGlobalConfigmapApi"
	GetEnvironmentConfigMapApiUrl string = "/orchestrator/config/environment/cm/"
	GetEnvironmentConfigMapApi    string = "GetEnvironmentConfigMapApi"
	SaveGlobalSecretApiUrl        string = "/orchestrator/config/global/cs"
	SaveGlobalSecretApi           string = "SaveGlobalSecretApi"
	AWSSystemManager              string = "AWSSystemManager"
	HashiCorpVault                string = "HashiCorpVault"
	AWSSecretsManager             string = "AWSSecretsManager"
	Kubernetes                    string = "Kubernetes"
	ExternalKubernetes            string = "ExternalKubernetes"
	environment                   string = "environment"
	volume                        string = "volume"
)
