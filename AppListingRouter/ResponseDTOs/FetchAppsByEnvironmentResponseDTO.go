package ResponseDTOs

import "time"

type FetchAppsByEnvironmentResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		AppContainers   []AppContainers `json:"appContainers"`
		AppCount        int             `json:"appCount"`
		DeploymentGroup DeploymentGroup `json:"deploymentGroup"`
	} `json:"result"`
}

type AppContainers struct {
	AppId        int            `json:"appId"`
	AppName      string         `json:"appName"`
	ProjectId    int            `json:"projectId"`
	Environments []Environments `json:"environments"`
}

type Environments struct {
	AppId           int         `json:"appId"`
	AppName         string      `json:"appName"`
	EnvironmentId   int         `json:"environmentId"`
	EnvironmentName string      `json:"environmentName"`
	Namespace       string      `json:"namespace"`
	ClusterName     string      `json:"clusterName"`
	Status          string      `json:"status"`
	AppStatus       string      `json:"appStatus"`
	CdStageStatus   *string     `json:"cdStageStatus"`
	PreStageStatus  interface{} `json:"preStageStatus"`
	PostStageStatus interface{} `json:"postStageStatus"`
	Default         bool        `json:"default"`
	Deleted         bool        `json:"deleted"`
	MaterialInfo    []struct {
		Author       string    `json:"author"`
		Branch       string    `json:"branch"`
		Message      string    `json:"message"`
		ModifiedTime time.Time `json:"modifiedTime"`
		Revision     string    `json:"revision"`
		Url          string    `json:"url"`
		WebhookData  string    `json:"webhookData"`
	} `json:"materialInfo"`
	CiArtifactId     int    `json:"ciArtifactId"`
	TeamId           int    `json:"teamId"`
	TeamName         string `json:"teamName"`
	LastDeployedTime string `json:"lastDeployedTime,omitempty"`
	DataSource       string `json:"dataSource,omitempty"`
}

type DeploymentGroup struct {
	Id             int         `json:"id"`
	Name           string      `json:"name"`
	AppCount       int         `json:"appCount"`
	NoOfApps       string      `json:"noOfApps"`
	EnvironmentId  int         `json:"environmentId"`
	CiPipelineId   int         `json:"ciPipelineId"`
	CiMaterialDTOs interface{} `json:"ciMaterialDTOs"`
}
