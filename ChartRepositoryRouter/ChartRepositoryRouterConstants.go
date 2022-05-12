package ChartRepositoryRouter

type AuthMode string

const (
	GetChartRepoListApiUrl       string   = "/orchestrator/chart-repo/list"
	GetChartRepoListApi          string   = "GetChartRepoListApiUrl"
	CreateChartRepoApiUrl        string   = "/orchestrator/chart-repo/create"
	CreateChartRepo              string   = "CreateChartRepo"
	DeleteChartRepoApiUrl        string   = "/orchestrator/chart-repo/"
	DeleteChartRepoApi           string   = "DeleteChartRepoApi"
	GetChartRepoById             string   = "GetChartRepoById"
	UpdateChartRepoUrl           string   = "/orchestrator/chart-repo/update"
	UpdateChartRepo              string   = "UpdateChartRepo"
	ValidateChartRepoApiUrl      string   = "/orchestrator/chart-repo/validate"
	ValidateChartRepoApi         string   = "validateChartRepoApi"
	TriggerChartSyncManualApiUrl string   = "/orchestrator/chart-repo/sync-charts"
	TriggerChartSyncManualApi    string   = "TriggerChartSyncManualApi"
	AUTH_MODE_USERNAME_PASSWORD  AuthMode = "USERNAME_PASSWORD"
	AUTH_MODE_SSH                AuthMode = "SSH"
	AUTH_MODE_ACCESS_TOKEN       AuthMode = "ACCESS_TOKEN"
	AUTH_MODE_ANONYMOUS          AuthMode = "ANONYMOUS"
)
