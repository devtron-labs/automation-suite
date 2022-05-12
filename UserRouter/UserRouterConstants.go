package UserRouter

const (
	SuperAdmin                string = "SuperAdmin"
	GroupsAndRoleFilter       string = "GroupsAndRoleFilter"
	RoleFilterOnly            string = "RoleFilterOnly"
	GroupsOnly                string = "GroupsOnly"
	WithHelmAppsOnly          string = "WithHelmAppsOnly"
	WithDevtronAppsOnly       string = "WithDevtronAppsOnly"
	WithChartGroupsOnly       string = "WithChartGroupsOnly"
	WithAllFilter             string = "WithAllFilter"
	CreatUserApiUrl           string = "/orchestrator/user"
	GetUserByIdApiUrl         string = "/orchestrator/user/"
	CreateUserApiUrl          string = "/orchestrator/user"
	UpdateUserApiUrl          string = "/orchestrator/user"
	DeleteUserApiUrl          string = "/orchestrator/user/"
	CreateRoleGroupApiUrl     string = "/orchestrator/user/role/group"
	GetRoleGroupByIdApiUrl    string = "/orchestrator/user/role/group/"
	DeleteRoleGroupByIdApiUrl string = "/orchestrator/user/role/group/"
)
