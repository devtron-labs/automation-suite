package ResponseDTOs

type DeleteApiTokenResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Success bool `json:"success"`
	} `json:"result"`
}
