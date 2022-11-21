package ResponseDTOs

type TemplateValuesResponseDTO struct {
	Code   int            `json:"code"`
	Status string         `json:"status"`
	Result TemplateValues `json:"result"`
}

type TemplateValues struct {
	Id                int    `json:"id"`
	AppStoreVersionId int    `json:"appStoreVersionId"`
	Name              string `json:"name"`
	Values            string `json:"values"`
	ChartVersion      string `json:"chartVersion"`
}
