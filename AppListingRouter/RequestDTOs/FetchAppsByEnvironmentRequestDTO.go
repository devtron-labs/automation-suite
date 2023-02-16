package RequestDTOs

type FetchAppsByEnvironmentRequestDTO struct {
	Environments      []int    `json:"environments"`
	Statuses          []string `json:"statuses"`
	Teams             []int    `json:"teams"`
	AppNameSearch     string   `json:"appNameSearch"`
	SortOrder         string   `json:"sortOrder"`
	SortBy            string   `json:"sortBy"`
	Offset            int      `json:"offset"`
	Size              int      `json:"size"`
	DeploymentGroupId int      `json:"deploymentGroupId"`
	Namespaces        []string `json:"namespaces"` //{clusterId}_{namespace}
	AppStatuses       []string `json:"appStatuses"`
}
