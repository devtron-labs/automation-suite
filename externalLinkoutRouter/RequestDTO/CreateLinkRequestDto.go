package RequestDTO

type CreateLinkRequestDto struct {
	Id               int    `json:"id"`
	Active           bool   `json:"active"`
	MonitoringToolId int    `json:"monitoringToolId"`
	Name             string `json:"name"`
	ClusterIds       []int  `json:"clusterIds"`
	Url              string `json:"url"`
}

type CreateLinkRequestDto1 struct {
	Id               int           `json:"id"`
	MonitoringToolId int           `json:"monitoringToolId"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Type             string        `json:"type"`
	Identifiers      []Identifiers `json:"identifiers"`
	Url              string        `json:"url"`
	IsEditable       bool          `json:"isEditable"`
}

type Identifiers struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	ClusterId  int    `json:"clusterId"`
}
