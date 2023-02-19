package ResponseDTOs

type InstalledAppResponseDTO struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Result InstalledAppDTO `json:"result"`
}

type InstalledAppDTO struct {
	Id                    int    `json:"id"`
	AppId                 int    `json:"appId"`
	AppName               string `json:"appName"`
	TeamId                int    `json:"teamId"`
	EnvironmentId         int    `json:"environmentId"`
	InstalledAppId        int    `json:"installedAppId"`
	InstalledAppVersionId int    `json:"installedAppVersionId"`
	AppStoreVersion       int    `json:"appStoreVersion"`
	ValuesOverrideYaml    string `json:"valuesOverrideYaml"`
	ReferenceValueId      int    `json:"referenceValueId"`
	ReferenceValueKind    string `json:"referenceValueKind"`
	AppStoreId            int    `json:"appStoreId"`
	AppStoreName          string `json:"appStoreName"`
	Deprecated            bool   `json:"deprecated"`
	ClusterId             int    `json:"clusterId"`
	Namespace             string `json:"namespace"`
	AppOfferingMode       string `json:"appOfferingMode"`
	GitOpsRepoName        string `json:"gitOpsRepoName"`
	GitOpsPath            string `json:"gitOpsPath"`
	GitHash               string `json:"gitHash"`
}
