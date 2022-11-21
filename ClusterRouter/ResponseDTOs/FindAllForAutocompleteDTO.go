package ResponseDTOs

type FindAllForAutocomplete struct {
	Code   int       `json:"code"`
	Status string    `json:"status"`
	Result []Cluster `json:"result"`
}
