package ResponseDTOs

import "time"

type DeploymentOfInstalledAppResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		AppStoreApplicationName      string    `json:"appStoreApplicationName"`
		ChartName                    string    `json:"chartName"`
		Icon                         string    `json:"icon"`
		Status                       string    `json:"status"`
		AppName                      string    `json:"appName"`
		InstalledAppVersionId        int       `json:"installedAppVersionId"`
		AppStoreApplicationVersionId int       `json:"appStoreApplicationVersionId"`
		EnvironmentName              string    `json:"environmentName"`
		DeployedAt                   time.Time `json:"deployedAt"`
		DeployedBy                   string    `json:"deployedBy"`
		InstalledAppId               int       `json:"installedAppId"`
		Readme                       string    `json:"readme"`
		EnvironmentId                int       `json:"environmentId"`
		Deprecated                   bool      `json:"deprecated"`
		AppOfferingMode              string    `json:"appOfferingMode"`
		ClusterId                    int       `json:"clusterId"`
		Namespace                    string    `json:"namespace"`
	} `json:"result"`
}
