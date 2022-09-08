package RequestDTOs

type CreateApiTokenRequestDTO struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ExpireAtInMs int64  `json:"expireAtInMs"`
}
