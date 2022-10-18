package ResponseDTOs

type AppStoreChartByNameResponseDTO struct {
	Code   int                   `json:"code"`
	Status string                `json:"status"`
	Result []AppStoreChartByName `json:"result"`
}

type AppStoreChartByName struct {
	AppStoreApplicationVersionId int    `json:"appStoreApplicationVersionId"`
	ChartId                      int    `json:"chartId"`
	ChartName                    string `json:"chartName"`
	ChartRepoId                  int    `json:"chartRepoId"`
	ChartRepoName                string `json:"chartRepoName"`
	Version                      string `json:"version"`
	Deprecated                   bool   `json:"deprecated"`
}
