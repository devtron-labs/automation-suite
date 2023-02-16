package RequestDTOs

/*type SaveDockerRegistryRequestDto struct {
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
}*/

type SaveDockerRegistryRequestDTO struct {
	Id           string    `json:"id"`
	IpsConfig    IpsConfig `json:"ipsConfig"`
	IsDefault    bool      `json:"isDefault"`
	Password     string    `json:"password"`
	PluginId     string    `json:"pluginId"`
	RegistryType string    `json:"registryType"`
	RegistryUrl  string    `json:"registryUrl"`
	Username     string    `json:"username"`
}

type IpsConfig struct {
	AppliedClusterIdsCsv string `json:"appliedClusterIdsCsv"`
	CredentialType       string `json:"credentialType"`
	CredentialValue      string `json:"credentialValue"`
	Id                   int    `json:"id"`
	IgnoredClusterIdsCsv string `json:"ignoredClusterIdsCsv"`
}
