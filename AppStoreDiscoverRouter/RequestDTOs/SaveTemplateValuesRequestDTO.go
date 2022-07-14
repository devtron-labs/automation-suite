package RequestDTOs

type SaveTemplateValuesRequestDTO struct {
	Name              string `json:"name"`
	AppStoreVersionId int    `json:"appStoreVersionId"`
	Values            string `json:"values"`
}
