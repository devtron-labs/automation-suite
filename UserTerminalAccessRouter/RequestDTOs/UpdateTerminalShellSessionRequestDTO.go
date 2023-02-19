package RequestDTOs

type UpdateTerminalShellSessionRequestDTO struct {
	ShellName        string `json:"shellName"`
	TerminalAccessId int    `json:"terminalAccessId"`
}
