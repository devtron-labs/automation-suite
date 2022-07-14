package ResponseDTOs

type GetApplicationValuesListResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Values []struct {
			Values []struct {
				Id           int    `json:"id"`
				Name         string `json:"name"`
				ChartVersion string `json:"chartVersion"`
			} `json:"values"`
			Kind string `json:"kind"`
		} `json:"values"`
	} `json:"result"`
}
