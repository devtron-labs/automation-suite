package RequestDTOs

type ChartRepoRequestDTO struct {
	Id           int    `json:"id,omitempty" validate:"number"`
	Name         string `json:"name,omitempty" validate:"required"`
	Url          string `json:"url,omitempty"`
	UserName     string `json:"userName,omitempty"`
	Password     string `json:"password,omitempty"`
	SshKey       string `json:"sshKey,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	AuthMode     string `json:"authMode,omitempty" validate:"required"`
	Active       bool   `json:"active"`
	Default      bool   `json:"default"`
	UserId       int32  `json:"-"`
	CustomErrMsg string `json:"customErrMsg"`
	ActualErrMsg string `json:"actualErrMsg"`
}
