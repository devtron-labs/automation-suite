package ResponseDTOs

type HelmEnvAutocompleteResponseDTO struct {
	Code   int       `json:"code"`
	Status string    `json:"status"`
	Result []HelmEnv `json:"result"`
}

type HelmEnv struct {
	ClusterId    int           `json:"clusterId"`
	ClusterName  string        `json:"clusterName"`
	Environments []Environment `json:"environments"`
}

type Environment struct {
	EnvironmentId         int    `json:"environmentId"`
	EnvironmentName       string `json:"environmentName"`
	Namespace             string `json:"namespace"`
	EnvironmentIdentifier string `json:"environmentIdentifier"`
}
