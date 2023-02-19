package ResponseDTOs

type AppListForAutocompleteResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"result"`
}
