package RequestDTOs

type CreateSSODetailsRequestDTO struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Config struct {
		Id     string `json:"id"`
		Label  string `json:"label"`
		Type   string `json:"type"`
		Name   string `json:"name"`
		Config struct {
			Issuer        string   `json:"issuer"`
			ClientID      string   `json:"clientID"`
			ClientSecret  string   `json:"clientSecret"`
			RedirectURI   string   `json:"redirectURI"`
			HostedDomains []string `json:"hostedDomains"`
		} `json:"config"`
	} `json:"config"`
	Active bool `json:"active"`
}
