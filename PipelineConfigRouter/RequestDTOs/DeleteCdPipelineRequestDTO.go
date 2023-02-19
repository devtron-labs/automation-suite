package RequestDTOs

type DeleteCdPipelineRequestDTO struct {
	Action   int `json:"action"`
	AppId    int `json:"appId"`
	Pipeline struct {
		Id int `json:"id"`
	} `json:"pipeline"`
}
