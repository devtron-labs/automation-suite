package ResponseDTOs

type DeleteClusterResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}
