package ResponseDTO

type Tool struct {
	Id   int    `json:"id"`
	Icon string `json:"icon"`
	Name string `json:"Name"`
}
type FetchAllToolsResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []Tool `json:"result"`
}
