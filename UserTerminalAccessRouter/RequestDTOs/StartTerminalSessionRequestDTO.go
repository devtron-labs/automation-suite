package RequestDTOs

type StartTerminalSessionRequestDTO struct {
	ClusterId int    `json:"clusterId"`
	BaseImage string `json:"baseImage"`
	ShellName string `json:"shellName"`
	NodeName  string `json:"nodeName"`
	Namespace string `json:"namespace"`
}
