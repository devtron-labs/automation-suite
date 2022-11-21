package ResponseDTOs

type SaveTemplateValuesResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id                int    `json:"id"`
		AppStoreVersionId int    `json:"appStoreVersionId"`
		Name              string `json:"name"`
		Values            string `json:"values"`
	} `json:"result"`
}
