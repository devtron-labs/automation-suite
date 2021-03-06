package regressionTestSuite

const (
	CreateAppApi            string = "CreateAppApi"
	CreateGitopsConfigApi   string = "CreateGitopsConfigApi"
	SaveGitopsConfigApiUrl  string = "/orchestrator/gitops/config"
	CreateTeamApi           string = "CreateTeamApi"
	SaveTeamApiUrl          string = "/orchestrator/team"
	FetchAllGitopsConfigApi string = "FetchAllLinkApi"
	GetAutocompleteApiUrl   string = "/orchestrator/team"
	FetchAllAutocompleteApi string = "FetchAllAutocompleteApi"

	// createApp_test_urls
	SaveAppApiUrl string = "/orchestrator/app"

	GetStageStatusApiUrl   string = "/orchestrator/app/stage/status"
	FetchAllStageStatusApi string = "FetchAllStageStatusApi"
	GetOtherEnvApiUrl      string = "/orchestrator/app/other-env"
	FetchOtherEnvApi       string = "FetchOtherEnvApi"
	GetAppWorkflowApiUrl   string = "/orchestrator/app/app-wf/"
	FetchAllAppWorkflowApi string = "FetchAllAppWorkflowApi"
	GetAppGetApiUrl        string = "/orchestrator/app/get"
	FetchAppGetApi         string = "FetchAppGetApi"
	SaveAppMaterialApiUrl  string = "/orchestrator/app/material"
	CreateAppMaterialApi   string = "CreateAppMaterialApi"
)
