package ResponseDTOs

type ReleaseInfoApiResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		InstalledAppInfo InstalledAppInfo `json:"installedAppInfo"`
		ReleaseInfo      ReleaseInfo      `json:"releaseInfo"`
	} `json:"result"`
}
type ReleaseInfo struct {
	DeployedAppDetail DeployedAppDetail `json:"deployedAppDetail"`
	DefaultValues     string            `json:"defaultValues"`
	OverrideValues    string            `json:"overrideValues"`
	MergedValues      string            `json:"mergedValues"`
	Readme            string            `json:"readme"`
}

type DeployedAppDetail struct {
	AppId             string             `json:"appId"`
	AppName           string             `json:"appName"`
	ChartName         string             `json:"chartName"`
	EnvironmentDetail EnvironmentDetails `json:"environmentDetail"`
	LastDeployed      DeployedAt         `json:"LastDeployed"`
	ChartVersion      string             `json:"chartVersion"`
}

type EnvironmentDetails struct {
	ClusterName string `json:"clusterName"`
	ClusterId   int    `json:"clusterId"`
	Namespace   string `json:"namespace"`
}
