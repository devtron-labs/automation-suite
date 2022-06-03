package ResponseDTOs

type GetListResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Url    string `json:"url"`
		Active bool   `json:"active"`
		Label  string `json:"label,omitempty"`
	} `json:"result"`
}
