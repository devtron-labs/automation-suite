package RequestDTOs

type SaveClusterRequestDTO struct {
	Id             int            `json:"id"`
	ClusterName    string         `json:"cluster_name"`
	Config         Config         `json:"config"`
	Active         bool           `json:"active"`
	PrometheusUrl  string         `json:"prometheus_url"`
	PrometheusAuth PrometheusAuth `json:"prometheusAuth"`
	ServerUrl      string         `json:"server_url"`
}

type Config struct {
	BearerToken string `json:"bearer_token"`
}

type PrometheusAuth struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
