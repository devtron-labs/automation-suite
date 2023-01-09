package ResponseDTOs

type DeleteEnvResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}
