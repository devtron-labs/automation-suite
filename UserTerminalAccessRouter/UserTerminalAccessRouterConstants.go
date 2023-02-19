package UserTerminalAccessRouter

const (
	StartTerminalSessionApi                     string = "StartTerminalSessionApi"
	StartTerminalSessionApiPath                 string = "/start"
	UserTerminalRouterBaseUrl                   string = "/orchestrator/user/terminal"
	FetchTerminalStatusApi                      string = "FetchTerminalStatusApi"
	FetchTerminalStatusApiPath                  string = "/get"
	FetchTerminalPodEventsApi                   string = "FetchTerminalPodEventsApi"
	FetchTerminalPodEventsApiPath               string = "/pod/events"
	FetchTerminalPodManifestApi                 string = "FetchTerminalPodManifestApi"
	FetchTerminalPodManifestApiPath             string = "/pod/manifest"
	StopTerminalSessionApi                      string = "StopTerminalSessionApi"
	StopTerminalSessionApiPath                  string = "/stop"
	UpdateTerminalSessionApi                    string = "UpdateTerminalSessionApi"
	UpdateTerminalSessionApiPath                string = "/update"
	UpdateTerminalShellSessionApi               string = "UpdateTerminalShellSessionApi"
	UpdateTerminalShellSessionApiPath           string = "/update/shell"
	DisconnectTerminalSessionApi                string = "DisconnectTerminalSessionApi"
	DisconnectTerminalSessionApiPath            string = "/disconnect"
	DisconnectAllTerminalSessionAndRetryApi     string = "DisconnectAllTerminalSessionAndRetryApi"
	DisconnectAllTerminalSessionAndRetryApiPath string = "/disconnectAndRetry"
)
