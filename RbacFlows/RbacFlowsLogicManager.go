package RbacFlows

import (
	"automation-suite/EnvironmentRouter"
	EnvironmentRouterResponseDTOs "automation-suite/EnvironmentRouter/ResponseDTOs"
	"automation-suite/PipelineConfigRouter"
	"automation-suite/TeamRouter"
	TeamRouterResponseDTOs "automation-suite/TeamRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"time"
)

type UserRole struct {
	Id      int32  `json:"id" validate:"number"`
	EmailId string `json:"email_id" validate:"email"`
	Role    string `json:"role"`
}

type UserInfo struct {
	Id           int32        `json:"id" validate:"number"`
	EmailId      string       `json:"email_id" validate:"required"`
	Roles        []string     `json:"roles,omitempty"`
	AccessToken  string       `json:"access_token,omitempty"`
	UserType     string       `json:"-"`
	LastUsedAt   time.Time    `json:"-"`
	LastUsedByIp string       `json:"-"`
	Exist        bool         `json:"-"`
	UserId       int32        `json:"-"` // created or modified user id
	RoleFilters  []RoleFilter `json:"roleFilters"`
	Status       string       `json:"status,omitempty"`
	Groups       []string     `json:"groups"`
	SuperAdmin   bool         `json:"superAdmin,notnull"`
}

type RoleGroup struct {
	Id          int32        `json:"id" validate:"number"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	RoleFilters []RoleFilter `json:"roleFilters"`
	Status      string       `json:"status,omitempty"`
	UserId      int32        `json:"-"` // created or modified user id
}

type RoleFilter struct {
	Entity      string `json:"entity"`
	Team        string `json:"team"`
	EntityName  string `json:"entityName"`
	Environment string `json:"environment"`
	Action      string `json:"action"`
	AccessType  string `json:"accessType"`
}

type Role struct {
	Id   int    `json:"id" validate:"number"`
	Role string `json:"role" validate:"required"`
}

type RoleData struct {
	Id          int    `json:"id" validate:"number"`
	Role        string `json:"role" validate:"required"`
	Entity      string `json:"entity"`
	Team        string `json:"team"`
	EntityName  string `json:"entityName"`
	Environment string `json:"environment"`
	Action      string `json:"action"`
	AccessType  string `json:"accessType"`
}

type SSOLoginDto struct {
	Id     int32           `json:"id"`
	Name   string          `json:"name,omitempty"`
	Label  string          `json:"label,omitempty"`
	Url    string          `json:"url,omitempty"`
	Config json.RawMessage `json:"config,omitempty"`
	Active bool            `json:"active"`
	UserId int32           `json:"-"`
}

func CreateProject(payload []byte, authToken string) TeamRouterResponseDTOs.SaveTeamResponseDTO {
	response := TeamRouter.HitSaveTeamApi(payload, authToken)
	return response
}

func DeleteProject(payload []byte, authToken string) TeamRouterResponseDTOs.DeleteTeamResponseDto {
	response := TeamRouter.HitDeleteTeamApi(payload, authToken)
	return response
}

func CreateEnv(payload []byte, authToken string) EnvironmentRouterResponseDTOs.CreateEnvironmentResponseDTO {
	response := EnvironmentRouter.HitCreateEnvApi(payload, authToken)
	return response
}

func DeleteEnv(payload []byte, authToken string) EnvironmentRouterResponseDTOs.DeleteEnvResponseDTO {
	response := EnvironmentRouter.HitDeleteEnvApi(payload, authToken)
	return response
}

func CreateDevtronApp(appName string, authToken string) PipelineConfigRouter.CreateAppResponseDto {
	//appName := "app" + strings.ToLower(Base.GetRandomStringOfGivenLength(10))
	createAppRequestDto := PipelineConfigRouter.GetAppRequestDto(appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
	response := PipelineConfigRouter.HitCreateAppApi(byteValueOfCreateApp, appName, 1, 0, authToken)
	return response
}
func DeleteDevtronApp(appId int, appName string, teamId int, TemplateId int, authToken string) PipelineConfigRouter.DeleteResponseDto {
	byteValueOfDeleteApp := PipelineConfigRouter.GetPayLoadForDeleteAppAPI(appId, appName, teamId, TemplateId)
	response := PipelineConfigRouter.HitDeleteAppApi(byteValueOfDeleteApp, appId, authToken)
	return response
}

func CreateHelmApp() {

}
func DeleteHelmApp() {

}

func getExpectedStatusCode(action string, apiName string) int {
	isValid := false
	switch action {
	case "manager":
		isValid = (apiName == GLOBALCONFIGURATIONS) ||
			(apiName == CREATEAPP) || (apiName == APPLISTFETCH) ||
			(apiName == PIPELINECREATE) || (apiName == PIPELINEFETCH) ||
			(apiName == TRIGGERPIPELINE) || (apiName == APPDETAILS) ||
			(apiName == ENVIRONMENTOVERRIDES) || (apiName == DELETE)
	case "admin":
		isValid = (apiName == CREATEAPP) || (apiName == APPLISTFETCH) || (apiName == PIPELINECREATE) || (apiName == PIPELINEFETCH) || (apiName == TRIGGERPIPELINE) || (apiName == APPDETAILS) || (apiName == ENVIRONMENTOVERRIDES)

	case "view":
		isValid = (apiName == APPLISTFETCH) || (apiName == PIPELINEFETCH) || (apiName == APPDETAILS)
	case "trigger":
		isValid = (apiName == APPLISTFETCH) || (apiName == PIPELINECREATE) || (apiName == PIPELINEFETCH) || (apiName == TRIGGERPIPELINE) || (apiName == APPDETAILS)

	}
	if isValid {
		return 200
	}
	return 403

}

type RbacFlowTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *RbacFlowTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
