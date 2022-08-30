package ResponseDTOs

import "automation-suite/testUtils"

type TerminalSessionResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Result TerminalSession    `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}

type TerminalSession struct {
	Op        string `json:"Op"`
	Data      string `json:"Data"`
	SessionID string `json:"SessionID"`
	Rows      int    `json:"Rows"`
	Cols      int    `json:"Cols"`
}
