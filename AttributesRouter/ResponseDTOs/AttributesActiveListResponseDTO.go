package ResponseDTOs

type AttributesActiveListResponseDTO struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Result []AttributesDTO `json:"result"`
}
