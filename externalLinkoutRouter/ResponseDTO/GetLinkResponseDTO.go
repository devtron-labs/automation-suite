package ResponseDTO

type GetLinkByIdResponseDto struct {
	Code   int                  `json:"code"`
	Status string               `json:"status"`
	Result []ExternalLinkOutDTo `json:"result"`
}

type ExternalLinkOutDTo struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Url              string   `json:"url"`
	MonitoringToolId int      `json:"monitoringToolId"`
	ClusterIds       []string `json:"clusterIds"`
	Active           bool     `json:"active"`
}
