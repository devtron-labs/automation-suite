package ResponseDTOs

type ListResponseDTO struct {
	Result struct {
		Metadata struct {
			ResourceVersion string `json:"resourceVersion"`
		} `json:"metadata"`
		Items interface{} `json:"items"`
	} `json:"result"`
}
