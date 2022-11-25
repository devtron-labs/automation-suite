package ResponseDTOs

type CheckAppExistsResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Name   string `json:"name"`
		Exists bool   `json:"exists"`
	} `json:"result"`
}
