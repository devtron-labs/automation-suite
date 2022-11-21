package RequestDTOs

type SaveDockerRegistryRequestDto struct {
	Id           string `json:"id"`
	PluginId     string `json:"pluginId"`
	RegistryUrl  string `json:"registryUrl"`
	RegistryType string `json:"registryType"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	IsDefault    bool   `json:"isDefault"`
	Connection   string `json:"connection"`
	Cert         string `json:"cert"`
	Active       bool   `json:"active"`
}
