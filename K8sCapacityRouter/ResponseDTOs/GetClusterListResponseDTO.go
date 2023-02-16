package ResponseDTOs

type GetClusterListResponseDTO struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Result []Result `json:"result"`
}

type Result struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	NodeCount  int      `json:"nodeCount,omitempty"`
	NodeNames  []string `json:"nodeNames"`
	NodeErrors *struct {
	} `json:"nodeErrors"`
	NodeK8SVersions []string `json:"nodeK8sVersions"`
	ServerVersion   string   `json:"serverVersion,omitempty"`
	Cpu             *struct {
		Capacity string `json:"capacity"`
	} `json:"cpu"`
	Memory *struct {
		Capacity string `json:"capacity"`
	} `json:"memory"`
	ErrorInNodeListing string `json:"errorInNodeListing,omitempty"`
}
