package RequestDTOs

type CreateEnvironmentRequestDTO struct {
	Id                    int    `json:"id,omitempty" validate:"number"`
	Environment           string `json:"environment_name,omitempty" validate:"required,max=50"`
	ClusterId             int    `json:"cluster_id,omitempty" validate:"number,required"`
	ClusterName           string `json:"cluster_name,omitempty"`
	Active                bool   `json:"active"`
	Default               bool   `json:"default"`
	PrometheusEndpoint    string `json:"prometheus_endpoint,omitempty"`
	Namespace             string `json:"namespace,omitempty" validate:"name-space-component,max=50"`
	CdArgoSetup           bool   `json:"isClusterCdActive"`
	EnvironmentIdentifier string `json:"environmentIdentifier"`
}
