package RequestDTOs

type InstallAppRequestDTO struct {
	TeamId             int         `json:"teamId"`
	ReferenceValueId   int         `json:"referenceValueId"`
	ReferenceValueKind string      `json:"referenceValueKind"`
	EnvironmentId      int         `json:"environmentId"`
	Namespace          string      `json:"namespace"`
	AppStoreVersion    int         `json:"appStoreVersion"`
	ValuesOverride     interface{} `json:"valuesOverride"`
	ValuesOverrideYaml string      `json:"valuesOverrideYaml"`
	AppName            string      `json:"appName"`
}
