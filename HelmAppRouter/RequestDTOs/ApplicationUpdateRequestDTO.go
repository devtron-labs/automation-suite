package RequestDTOs

type ApplicationUpdateRequestDTO struct {
	Id                 int    `json:"id"`
	ReferenceValueId   int    `json:"referenceValueId"`
	ReferenceValueKind string `json:"referenceValueKind"`
	ValuesOverrideYaml string `json:"valuesOverrideYaml"`
	InstalledAppId     int    `json:"installedAppId"`
	AppStoreVersion    int    `json:"appStoreVersion"`
}
