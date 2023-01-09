package RequestDTOs

type RbacHelmUserApiDTO struct {
	ApiTokenName          string `json:"apiTokenName"`
	ApiToken              string `json:"apiToken"`
	ExpectedResponseCode  int    `json:"ExpectedResponseCode"`
	ExpectedAppCount      int    `json:"expectedAppCount"`
	ExpectedAppName       string `json:"expectedAppName"`
	ExpectedTeamId        int    `json:"expectedTeamId"`
	ExpectedEnvironmentId int    `json:"expectedEnvironmentId"`
}
