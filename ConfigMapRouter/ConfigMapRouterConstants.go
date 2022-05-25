package ConfigMapRouter

const (
	SaveGlobalConfigmapApiUrl     string = "/orchestrator/config/global/cm"
	SaveGlobalConfigmapApi        string = "SaveGlobalConfigmapApi"
	GetEnvironmentConfigMapApiUrl string = "/orchestrator/config/environment/cm/"
	GetEnvironmentConfigMapApi    string = "GetEnvironmentConfigMapApi"
	SaveEnvironmentSecretApiUrl   string = "/orchestrator/config/environment/cs"
	SaveEnvironmentSecretApi      string = "SaveEnvironmentSecretApi"
	AWSSystemManager              string = "AWSSystemManager"
	HashiCorpVault                string = "HashiCorpVault"
	AWSSecretsManager             string = "AWSSecretsManager"
	KubernetesSecret              string = "KubernetesSecret"
	ExternalKubernetesSecret      string = "ExternalKubernetesSecret"
	environment                   string = "environment"
	volume                        string = "volume"
)
