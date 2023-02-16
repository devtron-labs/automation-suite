package RbacFlows

import (
	"automation-suite/AppStoreDiscoverRouter"
	"automation-suite/RbacFlows/RequestDTOs"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *RbacFlowTestSuite) TestRbacFlowsForHelmApps() {
	spec_proj_env_app_view_helm := ""
	spec_proj_env_app_view_edit_helm := ""
	spec_proj_env_app_admin_helm := ""
	spec_proj_all_exs_env_spec_app_view_helm := ""
	spec_proj_all_exs_env_spec_app_view_edit_helm := ""
	spec_proj_all_exs_env_spec_app_admin_helm := ""
	spec_proj_all_exs_env_all_app_admin_helm := ""
	spec_proj_all_exs_env_all_app_view_edit_helm := ""
	spec_proj_all_exs_env_all_app_view_helm := ""
	var user_spec_proj_env_app_helm_forApiGetAllInstalledApps = []RequestDTOs.RbacHelmUserApiDTO{
		{"spec_proj_env_app_view_helm", spec_proj_env_app_view_helm, 200, 1, "test-app-1", 4, 5},
		{"spec_proj_env_app_view_edit_helm", spec_proj_env_app_view_edit_helm, 200, 1, "test-app-1", 4, 5},
		{"spec_proj_env_app_admin_helm", spec_proj_env_app_admin_helm, 200, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_spec_app_view_helm", spec_proj_all_exs_env_spec_app_view_helm, 200, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_spec_app_view_edit_helm", spec_proj_all_exs_env_spec_app_view_edit_helm, 200, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_spec_app_admin_helm", spec_proj_all_exs_env_spec_app_admin_helm, 200, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_all_app_admin_helm", spec_proj_all_exs_env_all_app_admin_helm, 200, 2, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_all_app_view_edit_helm", spec_proj_all_exs_env_all_app_view_edit_helm, 200, 2, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_all_app_view_helm", spec_proj_all_exs_env_all_app_view_helm, 200, 2, "test-app-1", 4, 5},
	}

	var UnauthorisedUserForApiDeleteInstalledApp = []RequestDTOs.RbacHelmUserApiDTO{
		{"spec_proj_env_app_view_helm", spec_proj_env_app_view_helm, 403, 1, "test-app-1", 4, 5},
		{"spec_proj_env_app_view_edit_helm", spec_proj_env_app_view_edit_helm, 403, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_spec_app_view_helm", spec_proj_all_exs_env_spec_app_view_helm, 403, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_spec_app_view_edit_helm", spec_proj_all_exs_env_spec_app_view_edit_helm, 403, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_all_app_view_edit_helm", spec_proj_all_exs_env_all_app_view_edit_helm, 403, 2, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_all_app_view_helm", spec_proj_all_exs_env_all_app_view_helm, 403, 2, "test-app-1", 4, 5},
	}

	var AdminUserNotAuthorisedForAppInDifferentProject = []RequestDTOs.RbacHelmUserApiDTO{
		{"spec_proj_env_app_admin_helm", spec_proj_env_app_admin_helm, 403, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_spec_app_admin_helm", spec_proj_all_exs_env_spec_app_admin_helm, 403, 1, "test-app-1", 4, 5},
		{"spec_proj_all_exs_env_all_app_admin_helm", spec_proj_all_exs_env_all_app_admin_helm, 403, 2, "test-app-1", 4, 5},
	}

	suite.Run("A=1=HitApiGetAllInstalledAppsForUserHavingPermissionOfSpecificProjEnvAndApp", func() {
		for _, user := range user_spec_proj_env_app_helm_forApiGetAllInstalledApps {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			AllInstalledHelmApps := AppStoreDiscoverRouter.HitApiGetAllInstalledApps(user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, AllInstalledHelmApps.Code)
			assert.Equal(suite.T(), user.ExpectedAppCount, len(AllInstalledHelmApps.Result.HelmApps))
			assert.Equal(suite.T(), user.ExpectedAppCount, len(AllInstalledHelmApps.Result.HelmApps))
			assert.Equal(suite.T(), user.ExpectedAppCount, len(AllInstalledHelmApps.Result.HelmApps))
			assert.Equal(suite.T(), user.ExpectedEnvironmentId, AllInstalledHelmApps.Result.HelmApps[0].EnvironmentDetail.EnvironmentId)
			assert.Equal(suite.T(), user.ExpectedTeamId, AllInstalledHelmApps.Result.HelmApps[0].ProjectId)
		}
	})

	suite.Run("A=2=HitApiGetAllInstalledAppsForUserHavingNoPermissionOfSpecificProjEnvAndApp", func() {
		UserHavingNoPermissionOfSpecificProjEnvAndApp := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpuby1hY2Nlc3MiLCJpc3MiOiJhcGlUb2tlbklzc3VlciJ9.czSJfHMgIYzeXj6oTEGKm0jDqk2wUtWOwBuOGerrK28"
		AllInstalledHelmApps := AppStoreDiscoverRouter.HitApiGetAllInstalledApps(UserHavingNoPermissionOfSpecificProjEnvAndApp)
		assert.Equal(suite.T(), 200, AllInstalledHelmApps.Code)
		assert.Equal(suite.T(), 0, len(AllInstalledHelmApps.Result.HelmApps))
	})

	suite.Run("A=3=HitApiGetAllInstalledAppsForSuperAdminSubuser", func() {
		SuperAdminSubuser := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzdXBlci1hZG1pbi10b2tlbiIsImlzcyI6ImFwaVRva2VuSXNzdWVyIiwiZXhwIjoxNjc5MzE5MjcxfQ.STZ4ptJkaCNral6vHMIA0zL_p8kC3NlcKX69unEfiwg"
		AllInstalledHelmAppsWithSuperAdminSubuser := AppStoreDiscoverRouter.HitApiGetAllInstalledApps(SuperAdminSubuser)
		AllInstalledHelmsAppsWithSuperAdminUser := AppStoreDiscoverRouter.HitApiGetAllInstalledApps(suite.authToken)
		assert.Equal(suite.T(), len(AllInstalledHelmAppsWithSuperAdminSubuser.Result.HelmApps), len(AllInstalledHelmsAppsWithSuperAdminUser.Result.HelmApps))
	})

	//todo hardcoded appId as of now because of time crunch in RBAC story
	suite.Run("A=4=HitApiDeleteInstalledAppViaUnauthorisedUser", func() {
		appId := strconv.Itoa(58)
		for _, user := range UnauthorisedUserForApiDeleteInstalledApp {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			ResponseOfDeleteInstalledAppApi := AppStoreDiscoverRouter.HitDeleteInstalledAppApi(appId, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, ResponseOfDeleteInstalledAppApi.Code)
		}
	})

	//todo hardcoded appId as of now because of time crunch in RBAC story
	suite.Run("A=4=HitApiDeleteInstalledAppViaUnauthorisedUser", func() {
		appId := strconv.Itoa(60)
		for _, user := range AdminUserNotAuthorisedForAppInDifferentProject {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			ResponseOfDeleteInstalledAppApi := AppStoreDiscoverRouter.HitDeleteInstalledAppApi(appId, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, ResponseOfDeleteInstalledAppApi.Code)
		}
	})
}
