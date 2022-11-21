package RequestDTO

type CreateLinkRequestDto struct {
	Id               int    `json:"id"`
	Active           bool   `json:"active"`
	MonitoringToolId int    `json:"monitoringToolId"`
	Name             string `json:"name"`
	ClusterIds       []int  `json:"clusterIds"`
	Url              string `json:"url"`
}
