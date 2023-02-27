package RequestDTOs

type FetchAppsByEnvironmentRequestDTO struct {
	Environments      []int    `json:"environments"`
	Statuses          []string `json:"statuses"`
	Teams             []int    `json:"teams"`
	DeploymentGroupId int      `json:"deploymentGroupId"`
	Namespaces        []string `json:"namespaces"` //{clusterId}_{namespace}
	AppStatuses       []string `json:"appStatuses"`
	AppNameSearch string `json:"appNameSearch"`
	SortBy        string `json:"sortBy"`
	SortOrder     string `json:"sortOrder"`
	Offset        int    `json:"offset"`
	HOffset       int    `json:"hOffset"`
	Size          int    `json:"size"`
}
