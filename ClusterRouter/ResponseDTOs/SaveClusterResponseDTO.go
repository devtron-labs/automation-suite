package ResponseDTOs

import (
	"automation-suite/ClusterRouter/RequestDTOs"
	"automation-suite/testUtils"
)

type SaveClusterResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Result Cluster            `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}

type Cluster struct {
	Id                      int                        `json:"id"`
	ClusterName             string                     `json:"cluster_name"`
	ServerUrl               string                     `json:"server_url"`
	Active                  bool                       `json:"active"`
	Config                  RequestDTOs.Config         `json:"config"`
	PrometheusAuth          RequestDTOs.PrometheusAuth `json:"prometheusAuth"`
	DefaultClusterComponent interface{}                `json:"defaultClusterComponent"`
	AgentInstallationStage  int                        `json:"agentInstallationStage"`
	K8SVersion              string                     `json:"k8sVersion"`
}
