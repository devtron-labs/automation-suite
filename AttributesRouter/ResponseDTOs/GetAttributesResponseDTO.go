package ResponseDTOs

type AttributesDTO struct {
	Id     int    `json:"id"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
	Active bool   `json:"active"`
	UserId int32  `json:"-"`
}

type GetAttributesResponseDTO struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result AttributesDTO `json:"result"`
}
