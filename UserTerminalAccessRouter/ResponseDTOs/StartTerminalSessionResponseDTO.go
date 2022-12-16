package ResponseDTOs

import "automation-suite/testUtils"

type StartTerminalSessionResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Errors []testUtils.Errors `json:"errors"`
	Result Result             `json:"result"`
}

type Result struct {
	UserTerminalSessionId string `json:"userTerminalSessionId"`
	UserId                int    `json:"userId"`
	TerminalAccessId      int    `json:"terminalAccessId"`
	Status                string `json:"status"`
	PodName               string `json:"podName"`
}
