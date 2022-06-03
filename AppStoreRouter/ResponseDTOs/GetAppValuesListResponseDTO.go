package ResponseDTOs

type GetApplicationValuesListResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Values []struct {
			Values []Values `json:"values"`
			Kind   string   `json:"kind"`
		} `json:"values"`
	} `json:"result"`
}

type Values struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	ChartVersion      string `json:"chartVersion"`
	AppStoreVersionId int    `json:"appStoreVersionId,omitempty"`
	EnvironmentName   string `json:"environmentName,omitempty"`
}
