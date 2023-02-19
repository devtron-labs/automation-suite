package RequestDTOs

type RollbackApplicationApiRequestDto struct {
	HAppId  string `json:"hAppId"`
	Version int    `json:"version"`
}
