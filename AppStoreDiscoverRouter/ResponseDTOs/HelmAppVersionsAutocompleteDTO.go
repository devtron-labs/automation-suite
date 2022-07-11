package ResponseDTOs

type HelmAppVersionsDTO struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Result []HelmAppVersions `json:"result"`
}

type HelmAppVersions struct {
	Version string `json:"version"`
	Id      int    `json:"id"`
}
