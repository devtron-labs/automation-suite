package RequestDTOs

type RbacDevtronUserApiDTO struct {
	ApiTokenName            string `json:"apiTokenName"`
	ApiToken                string `json:"apiToken"`
	ExpectedResponseCode    int    `json:"ExpectedResponseCode"`
	ExpectedAppCount        int    `json:"expectedAppCount"`
	ExpectedAppName         string `json:"expectedAppName"`
	ExpectedTeamName        string `json:"expectedTeamName"`
	ExpectedEnvironmentName string `json:"expectedEnvironmentName"`
}
