package ResponseDTOs

type DeploymentHistoryResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		InstalledAppInfo  InstalledAppInfo    `json:"installedAppInfo"`
		DeploymentHistory []DeploymentHistory `json:"deploymentHistory"`
	} `json:"result"`
}

type InstalledAppInfo struct {
	AppId                 int    `json:"appId"`
	InstalledAppId        int    `json:"installedAppId"`
	InstalledAppVersionId int    `json:"installedAppVersionId"`
	AppStoreChartId       int    `json:"appStoreChartId"`
	EnvironmentName       string `json:"environmentName"`
	AppOfferingMode       string `json:"appOfferingMode"`
	ClusterId             int    `json:"clusterId"`
	EnvironmentId         int    `json:"environmentId"`
}

type DeploymentHistory struct {
	ChartMetadata ChartMetadata `json:"chartMetadata"`
	DockerImages  []string      `json:"dockerImages"`
	Version       int           `json:"version"`
	DeployedAt    DeployedAt    `json:"deployedAt"`
}

type ChartMetadata struct {
	ChartName    string   `json:"chartName"`
	ChartVersion string   `json:"chartVersion"`
	Home         string   `json:"home"`
	Sources      []string `json:"sources"`
	Description  string   `json:"description"`
}

type DeployedAt struct {
	Seconds int `json:"seconds"`
	Nanos   int `json:"nanos"`
}
