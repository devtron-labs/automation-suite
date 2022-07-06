package ResponseDTOs

type GetOrchestratorResponse struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result ServerInfoDto `json:"result"`
}

type ServerInfoDto struct {
	CurrentVersion   string `json:"currentVersion"`
	Status           string `json:"status"`
	ReleaseName      string `json:"releaseName"`
	InstallationType string `json:"installationType"`
}
