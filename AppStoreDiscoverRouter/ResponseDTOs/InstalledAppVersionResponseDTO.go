package ResponseDTOs

import Base "automation-suite/testUtils"

type InstalledAppVersionResponseDTO struct {
	Code   int                    `json:"code"`
	Status string                 `json:"status"`
	Result InstalledAppVersionDTO `json:"result"`
	Error  []Base.Errors          `json:"errors"`
}

type InstalledAppVersionDTO struct {
	Id                 int    `json:"id"`
	AppId              int    `json:"appId"`
	AppName            string `json:"appName"`
	TeamId             int    `json:"teamId"`
	EnvironmentId      int    `json:"environmentId"`
	InstalledAppId     int    `json:"installedAppId"`
	AppStoreVersion    int    `json:"appStoreVersion"`
	ValuesOverrideYaml string `json:"valuesOverrideYaml"`
	Readme             string `json:"readme"`
	ReferenceValueId   int    `json:"referenceValueId"`
	ReferenceValueKind string `json:"referenceValueKind"`
	AppStoreId         int    `json:"appStoreId"`
	AppStoreName       string `json:"appStoreName"`
	Deprecated         bool   `json:"deprecated"`
	ClusterId          int    `json:"clusterId"`
	Namespace          string `json:"namespace"`
	AppOfferingMode    string `json:"appOfferingMode"`
	GitOpsRepoName     string `json:"gitOpsRepoName"`
	GitOpsPath         string `json:"gitOpsPath"`
	GitHash            string `json:"gitHash"`
}
