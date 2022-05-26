package PipelineConfigRouter

type SaveDeploymentTemplateRequestDTO struct {
	AppId              int                `json:"appId"`
	ChartRefId         int                `json:"chartRefId"`
	ValuesOverride     DefaultAppOverride `json:"valuesOverride"`
	DefaultAppOverride DefaultAppOverride `json:"defaultAppOverride"`
}

type SaveDeploymentTemplateResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id                      int                `json:"id"`
		AppId                   int                `json:"appId"`
		RefChartTemplate        string             `json:"refChartTemplate"`
		RefChartTemplateVersion string             `json:"refChartTemplateVersion"`
		ChartRepositoryId       int                `json:"chartRepositoryId"`
		DefaultAppOverride      DefaultAppOverride `json:"defaultAppOverride"`
		ChartRefId              int                `json:"chartRefId"`
		Latest                  bool               `json:"latest"`
		IsAppMetricsEnabled     bool               `json:"isAppMetricsEnabled"`
		Schema                  interface{}        `json:"schema"`
		Readme                  string             `json:"readme"`
	} `json:"result"`
}
