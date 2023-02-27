package UserRouter

import (
	"automation-suite/UserRouter/RequestDTOs"
	"automation-suite/UserRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type StructUserRouter struct {
	userRole                   RequestDTOs.UserRole
	userInfo                   RequestDTOs.UserInfo
	roleGroup                  RequestDTOs.RoleGroup
	roleFilter                 ResponseDTOs.RoleFilter
	role                       RequestDTOs.Role
	roleData                   RequestDTOs.RoleData
	createUserResponseDto      ResponseDTOs.CreateUserResponseDTO
	deleteUserResponseDto      ResponseDTOs.DeleteUserResponseDTO
	createRoleGroupResponseDto ResponseDTOs.CreateRoleGroupResponseDto
	getUserByIdResponseDto     ResponseDTOs.GetUserByIdResponseDTO
}

func (structUserRouter StructUserRouter) UnmarshalGivenResponseBody(response []byte, apiName string) StructUserRouter {
	switch apiName {
	case "CreateUser":
		json.Unmarshal(response, &structUserRouter.createUserResponseDto)
	case "DeleteUser":
		json.Unmarshal(response, &structUserRouter.deleteUserResponseDto)
	case "CreateRoleGroup":
		json.Unmarshal(response, &structUserRouter.createRoleGroupResponseDto)
	case "GetUserById":
		json.Unmarshal(response, &structUserRouter.getUserByIdResponseDto)
	}
	return structUserRouter
}

func HitGetAllUserApi(authToken string) ResponseDTOs.CreateUserResponseDTO {
	resp, err := Base.MakeApiCall(CreatUserApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, "GetAllUser")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "CreateUser")
	return userRouter.createUserResponseDto
}

func HitGetUserByIdApi(id string, authToken string) ResponseDTOs.GetUserByIdResponseDTO {
	resp, err := Base.MakeApiCall(GetUserByIdApiUrl+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, "GetUserById")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "GetUserById")
	return userRouter.getUserByIdResponseDto
}

func HitCreateUserApi(payload []byte, authToken string) ResponseDTOs.CreateUserResponseDTO {
	resp, err := Base.MakeApiCall(CreateUserApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, "CreateUser")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "CreateUser")
	return userRouter.createUserResponseDto
}

func HitUpdateUserApi(payload []byte, authToken string) ResponseDTOs.GetUserByIdResponseDTO {
	resp, err := Base.MakeApiCall(UpdateUserApiUrl, http.MethodPut, string(payload), nil, authToken)
	Base.HandleError(err, "UpdateUser")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "GetUserById")
	return userRouter.getUserByIdResponseDto
}

func HitDeleteUserApi(id string, authToken string) ResponseDTOs.DeleteUserResponseDTO {
	resp, err := Base.MakeApiCall(DeleteUserApiUrl+id, http.MethodDelete, "", nil, authToken)
	Base.HandleError(err, "DeleteUser")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "DeleteUser")
	return userRouter.deleteUserResponseDto
}

func HitCreateRoleGroupApi(payload []byte, authToken string) ResponseDTOs.CreateRoleGroupResponseDto {
	resp, err := Base.MakeApiCall(CreateRoleGroupApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, "CreateRoleGroup")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "CreateRoleGroup")
	return userRouter.createRoleGroupResponseDto
}

func HitGetRoleGroupByIdApi(id string, authToken string) ResponseDTOs.CreateRoleGroupResponseDto {
	resp, err := Base.MakeApiCall(GetRoleGroupByIdApiUrl+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, "GetRoleGroupById")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "CreateRoleGroup")
	return userRouter.createRoleGroupResponseDto
}

func HitDeleteRoleGroupByIdApi(id string, authToken string) ResponseDTOs.DeleteUserResponseDTO {
	resp, err := Base.MakeApiCall(DeleteRoleGroupByIdApiUrl+id, http.MethodDelete, "", nil, authToken)
	Base.HandleError(err, "DeleteRoleGroupById")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "DeleteUser")
	return userRouter.deleteUserResponseDto
}

func CreateUserRequestPayload(caseName string, authToken string) (RequestDTOs.UserInfo, int) {
	var userInfo RequestDTOs.UserInfo
	var createRoleGroupApiResponse ResponseDTOs.CreateRoleGroupResponseDto
	switch caseName {
	case GroupsAndRoleFilter:
		var listOfRoleFilter []ResponseDTOs.RoleFilter
		var listOfGroups []string

		roleFilter := CreateRoleFilter("", "unassigned", "", "manager", "")
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)

		createRoleGroupPayload := CreateRoleGroupPayload(WithAllFilter)
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
		createRoleGroupApiResponse = HitCreateRoleGroupApi(byteValueOfStruct, authToken)
		GroupName := createRoleGroupApiResponse.Result.Name
		listOfGroups = append(listOfGroups, GroupName)

		userName := Base.GetRandomStringOfGivenLength(10)
		userInfo.EmailId = userName + "@yopmail.com"
		userInfo.SuperAdmin = false
		userInfo.RoleFilters = listOfRoleFilter
		userInfo.Groups = listOfGroups

	case SuperAdmin:
		userName := Base.GetRandomStringOfGivenLength(10)
		userInfo.EmailId = userName + "@yopmail.com"
		userInfo.SuperAdmin = true
		userInfo.RoleFilters = []ResponseDTOs.RoleFilter{}
		userInfo.Groups = []string{}

	case RoleFilterOnly:
		var listOfRoleFilter []ResponseDTOs.RoleFilter
		roleFilter := CreateRoleFilter("", "unassigned", "", "manager", "")
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)
		userName := Base.GetRandomStringOfGivenLength(10)
		userInfo.EmailId = userName + "@yopmail.com"
		userInfo.SuperAdmin = false
		userInfo.RoleFilters = listOfRoleFilter
		userInfo.Groups = []string{}

	case GroupsOnly:
		var listOfGroups []string
		createRoleGroupPayload := CreateRoleGroupPayload(WithAllFilter)
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
		createRoleGroupApiResponse = HitCreateRoleGroupApi(byteValueOfStruct, authToken)
		GroupName := createRoleGroupApiResponse.Result.Name
		listOfGroups = append(listOfGroups, GroupName)

		userName := Base.GetRandomStringOfGivenLength(10)
		userInfo.EmailId = userName + "@yopmail.com"
		userInfo.SuperAdmin = false
		userInfo.RoleFilters = []ResponseDTOs.RoleFilter{}
		userInfo.Groups = listOfGroups
	}
	return userInfo, createRoleGroupApiResponse.Result.Id
}

func CreateRoleGroupPayload(caseName string) RequestDTOs.RoleGroup {
	var roleGroup RequestDTOs.RoleGroup
	var listOfRoleFilter []ResponseDTOs.RoleFilter

	switch caseName {
	case WithHelmAppsOnly:
		roleFilter := CreateRoleFilterWithHelmAppsOnly()
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)
	case WithDevtronAppsOnly:
		roleFilter := CreateRoleFilterWithDevtronAppsOnly()
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)
	case WithChartGroupsOnly:
		roleFilter := CreateRoleFilterWithChartGroupsOnly()
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)
	case WithAllFilter:
		roleFilter := CreateRoleFilterWithHelmAppsOnly()
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)

		roleFilter = CreateRoleFilterWithDevtronAppsOnly()
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)

		roleFilter = CreateRoleFilterWithChartGroupsOnly()
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)
	case WithDevtronAppsOnlyDynamic:
		roleFilter := CreateRoleFilterWithDevtronAppsOnlyDynamic(ENTITY, PROJECT, ENV, APP, ACTION, ACCESS_TYPE)
		listOfRoleFilter = append(listOfRoleFilter, roleFilter)
	}

	roleGroup.Name = Base.GetRandomStringOfGivenLength(10)
	roleGroup.Description = "This is the sample Description for Testing Purpose via Automation only"
	roleGroup.RoleFilters = listOfRoleFilter
	return roleGroup
}
func CreateRoleGroupPayloadDynamicForDevtronApp(entity, team, env, app, action, accessType string) RequestDTOs.RoleGroup {
	var roleGroup RequestDTOs.RoleGroup
	var listOfRoleFilter []ResponseDTOs.RoleFilter
	roleFilter := CreateRoleFilterWithDevtronAppsOnlyDynamic(entity, team, env, app, action, accessType)
	listOfRoleFilter = append(listOfRoleFilter, roleFilter)

	roleGroup.Name = Base.GetRandomStringOfGivenLength(10)
	roleGroup.Description = "This is the sample Description for Testing Purpose via Automation only"
	roleGroup.RoleFilters = listOfRoleFilter
	return roleGroup
}

func CreateRoleFilterWithDevtronAppsOnly() ResponseDTOs.RoleFilter {
	var roleFilter ResponseDTOs.RoleFilter
	roleFilter = CreateRoleFilter("", "devtron-demo", "", "view", "")
	return roleFilter
}
func CreateRoleFilterWithDevtronAppsOnlyDynamic(entity, team, env, app, action, accessType string) ResponseDTOs.RoleFilter {
	var roleFilter ResponseDTOs.RoleFilter
	roleFilter = CreateRoleFilterForDynamicEntityName(entity, team, env, action, accessType, app)
	return roleFilter
}

func CreateRoleFilterWithHelmAppsOnly() ResponseDTOs.RoleFilter {
	var roleFilter ResponseDTOs.RoleFilter
	roleFilter = CreateRoleFilter("", "unassigned", "default_cluster__*", "edit", "helm-app")
	return roleFilter

}

func CreateRoleFilterWithChartGroupsOnly() ResponseDTOs.RoleFilter {
	var roleFilter ResponseDTOs.RoleFilter
	roleFilter = CreateRoleFilter("chart-group", "", "*", "admin", "")
	return roleFilter
}

func CreateRoleFilter(entity string, teamName string, environment string, action string, accessType string) ResponseDTOs.RoleFilter {
	var roleFilter ResponseDTOs.RoleFilter
	roleFilter.Entity = entity
	roleFilter.Team = teamName
	roleFilter.EntityName = ""
	roleFilter.Environment = environment
	roleFilter.Action = action
	roleFilter.AccessType = accessType
	return roleFilter
}
func CreateRoleFilterForDynamicEntityName(entity string, teamName string, environment string, action string, accessType string, entityName string) ResponseDTOs.RoleFilter {
	var roleFilter ResponseDTOs.RoleFilter
	roleFilter.Entity = entity
	roleFilter.Team = teamName
	roleFilter.EntityName = entityName
	roleFilter.Environment = environment
	roleFilter.Action = action
	roleFilter.AccessType = accessType
	return roleFilter
}

type UserTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *UserTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
