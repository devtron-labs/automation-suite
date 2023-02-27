package RequestDTOs

import "automation-suite/PipelineConfigRouter"

type RbacDevtronUserApiDTO struct {
	ApiTokenName            string `json:"apiTokenName"`
	ApiToken                string `json:"apiToken"`
	ExpectedResponseCode    int    `json:"ExpectedResponseCode"`
	ExpectedAppCount        int    `json:"expectedAppCount"`
	ExpectedAppName         string `json:"expectedAppName"`
	ExpectedTeamName        string `json:"expectedTeamName"`
	ExpectedEnvironmentName string `json:"expectedEnvironmentName"`
}

type RbacDevtronDeletion struct {
	ProjectPayload []byte
	EnvPayLoad     []byte
	DevtronPayload PipelineConfigRouter.CreateAppResponseDto
	RoleGroupId    int
}
type RbacApiTokenDeletion struct {
	ApiTokenId int
	UserId     int
}
