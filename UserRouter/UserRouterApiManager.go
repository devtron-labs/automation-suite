package UserRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"net/http"
)

type UserRole struct {
	Id      int32  `json:"id" validate:"number"`
	EmailId string `json:"email_id" validate:"email"`
	Role    string `json:"role"`
}

type UserInfo struct {
	Id          int32        `json:"id" validate:"number"`
	EmailId     string       `json:"email_id" validate:"required"`
	Roles       []string     `json:"roles,omitempty"`
	AccessToken string       `json:"access_token,omitempty"`
	Exist       bool         `json:"-"`
	UserId      int32        `json:"-"` // created or modified user id
	RoleFilters []RoleFilter `json:"roleFilters"`
	Status      string       `json:"status,omitempty"`
	Groups      []string     `json:"groups"`
	SuperAdmin  bool         `json:"superAdmin,notnull"`
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

type CreateUserResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	Result []struct {
		Id          int          `json:"id"`
		EmailId     string       `json:"email_id"`
		RoleFilters []RoleFilter `json:"roleFilters"`
		Groups      []string     `json:"groups"`
		SuperAdmin  bool         `json:"superAdmin"`
	} `json:"result"`
}

type DeleteUserResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result bool   `json:"result"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
}

type CreateRoleGroupResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	Result struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		RoleFilters []struct {
			Entity      string `json:"entity"`
			Team        string `json:"team"`
			EntityName  string `json:"entityName"`
			Environment string `json:"environment"`
			Action      string `json:"action"`
			AccessType  string `json:"accessType"`
		} `json:"roleFilters"`
	} `json:"result"`
}

type GetUserByIdResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
	Result struct {
		Id          int          `json:"id"`
		EmailId     string       `json:"email_id"`
		RoleFilters []RoleFilter `json:"roleFilters"`
		Groups      []string     `json:"groups"`
		SuperAdmin  bool         `json:"superAdmin"`
	} `json:"result"`
}

type StructUserRouter struct {
	userRole                   UserRole
	userInfo                   UserInfo
	roleGroup                  RoleGroup
	roleFilter                 RoleFilter
	role                       Role
	roleData                   RoleData
	createUserResponseDto      CreateUserResponseDto
	deleteUserResponseDto      DeleteUserResponseDto
	createRoleGroupResponseDto CreateRoleGroupResponseDto
	getUserByIdResponseDto     GetUserByIdResponseDto
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

func HitGetAllUserApi(authToken string) CreateUserResponseDto {
	resp, err := Base.MakeApiCall(CreatUserApiUrl, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, "GetAllUser")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "CreateUser")
	return userRouter.createUserResponseDto
}

func HitGetUserByIdApi(id string, authToken string) GetUserByIdResponseDto {
	resp, err := Base.MakeApiCall(GetUserByIdApiUrl+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, "GetUserById")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "GetUserById")
	return userRouter.getUserByIdResponseDto
}

func HitCreateUserApi(payload []byte, authToken string) CreateUserResponseDto {
	resp, err := Base.MakeApiCall(CreateUserApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, "CreateUser")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "CreateUser")
	return userRouter.createUserResponseDto
}

func HitUpdateUserApi(payload []byte, authToken string) GetUserByIdResponseDto {
	resp, err := Base.MakeApiCall(UpdateUserApiUrl, http.MethodPut, string(payload), nil, authToken)
	Base.HandleError(err, "UpdateUser")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "GetUserById")
	return userRouter.getUserByIdResponseDto
}

func HitDeleteUserApi(id string, authToken string) DeleteUserResponseDto {
	resp, err := Base.MakeApiCall(DeleteUserApiUrl+id, http.MethodDelete, "", nil, authToken)
	Base.HandleError(err, "DeleteUser")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "DeleteUser")
	return userRouter.deleteUserResponseDto
}

func HitCreateRoleGroupApi(payload []byte, authToken string) CreateRoleGroupResponseDto {
	resp, err := Base.MakeApiCall(CreateRoleGroupApiUrl, http.MethodPost, string(payload), nil, authToken)
	Base.HandleError(err, "CreateRoleGroup")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "CreateRoleGroup")
	return userRouter.createRoleGroupResponseDto
}

func HitGetRoleGroupByIdApi(id string, authToken string) CreateRoleGroupResponseDto {
	resp, err := Base.MakeApiCall(GetRoleGroupByIdApiUrl+id, http.MethodGet, "", nil, authToken)
	Base.HandleError(err, "GetRoleGroupById")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "CreateRoleGroup")
	return userRouter.createRoleGroupResponseDto
}

func HitDeleteRoleGroupByIdApi(id string, authToken string) DeleteUserResponseDto {
	resp, err := Base.MakeApiCall(DeleteRoleGroupByIdApiUrl+id, http.MethodDelete, "", nil, authToken)
	Base.HandleError(err, "DeleteRoleGroupById")

	structUserRouter := StructUserRouter{}
	userRouter := structUserRouter.UnmarshalGivenResponseBody(resp.Body(), "DeleteUser")
	return userRouter.deleteUserResponseDto
}

func CreateUserRequestPayload(caseName string, authToken string) (UserInfo, int) {
	var userInfo UserInfo
	var createRoleGroupApiResponse CreateRoleGroupResponseDto
	switch caseName {
	case GroupsAndRoleFilter:
		var listOfRoleFilter []RoleFilter
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
		userInfo.RoleFilters = []RoleFilter{}
		userInfo.Groups = []string{}

	case RoleFilterOnly:
		var listOfRoleFilter []RoleFilter
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
		userInfo.RoleFilters = []RoleFilter{}
		userInfo.Groups = listOfGroups
	}
	return userInfo, createRoleGroupApiResponse.Result.Id
}

func CreateRoleGroupPayload(caseName string) RoleGroup {
	var roleGroup RoleGroup
	var listOfRoleFilter []RoleFilter

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
	}

	roleGroup.Name = Base.GetRandomStringOfGivenLength(10)
	roleGroup.Description = "This is the sample Description for Testing Purpose via Automation only"
	roleGroup.RoleFilters = listOfRoleFilter
	return roleGroup
}

func CreateRoleFilterWithDevtronAppsOnly() RoleFilter {
	var roleFilter RoleFilter
	roleFilter = CreateRoleFilter("", "devtron-demo", "", "view", "")
	return roleFilter
}

func CreateRoleFilterWithHelmAppsOnly() RoleFilter {
	var roleFilter RoleFilter
	roleFilter = CreateRoleFilter("", "unassigned", "default_cluster__*", "edit", "helm-app")
	return roleFilter

}

func CreateRoleFilterWithChartGroupsOnly() RoleFilter {
	var roleFilter RoleFilter
	roleFilter = CreateRoleFilter("chart-group", "", "*", "admin", "")
	return roleFilter
}

func CreateRoleFilter(entity string, teamName string, environment string, action string, accessType string) RoleFilter {
	var roleFilter RoleFilter
	roleFilter.Entity = entity
	roleFilter.Team = teamName
	roleFilter.EntityName = ""
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
