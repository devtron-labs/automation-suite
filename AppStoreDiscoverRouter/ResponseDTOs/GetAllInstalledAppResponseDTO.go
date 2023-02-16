package ResponseDTOs

import "time"

type GetAllInstalledAppResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		ApplicationType string      `json:"applicationType"`
		ClusterIds      interface{} `json:"clusterIds"`
		HelmApps        []HelmApps  `json:"helmApps"`
	} `json:"result"`
}

type HelmApps struct {
	AppId             string            `json:"appId"`
	AppName           string            `json:"appName"`
	ChartAvatar       string            `json:"chartAvatar"`
	ChartName         string            `json:"chartName"`
	EnvironmentDetail EnvironmentDetail `json:"environmentDetail"`
	LastDeployedAt    time.Time         `json:"lastDeployedAt"`
	ProjectId         int               `json:"projectId"`
}

type EnvironmentDetail struct {
	ClusterId       int    `json:"clusterId"`
	ClusterName     string `json:"clusterName"`
	EnvironmentId   int    `json:"environmentId"`
	EnvironmentName string `json:"environmentName"`
	Namespace       string `json:"namespace"`
}
