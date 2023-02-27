package RbacFlows

import (
	"automation-suite/ApiTokenRouter"
	"automation-suite/ApiTokenRouter/ResponseDTOs"
	AppListingRouter "automation-suite/AppListingRouter"
	"automation-suite/PipelineConfigRouter"
	"automation-suite/RbacFlows/RequestDTOs"
	"automation-suite/TeamRouter"
	"time"

	//"automation-suite/RbacFlows/RequestDTOs"
	"automation-suite/UserRouter"
	abcd "automation-suite/UserRouter/RequestDTOs"
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func getStatusCheck(expectedCode, actualCode int) bool {
	if expectedCode == 403 && actualCode == 200 {
		return false
	} else if expectedCode == 200 && (actualCode == 403 || actualCode == 401) {
		return false
	}
	return true
}

func CreateUserPayloadForDynamicToken(responseOfCreateApiToken ResponseDTOs.CreateApiTokenResponseDTO, authToken string, superAdmin bool) (abcd.UserInfo, string, int) {
	updateUserDto, roleGroupId := UserRouter.CreateUserRequestPayload(UserRouter.GroupsAndRoleFilter, authToken)

	resultToken := responseOfCreateApiToken.Result.Token
	updateUserDto.EmailId = responseOfCreateApiToken.Result.UserIdentifier
	updateUserDto.Id = int32(responseOfCreateApiToken.Result.UserId)
	if superAdmin {
		updateUserDto.SuperAdmin = true
	}
	return updateUserDto, resultToken, roleGroupId
}

func (suite *RbacFlowTestSuite) TestRbacFlowsForDevtronApps() {

	var allRoles = []string{"admin", "manager", "view", "trigger"}
	for _, role := range allRoles {
		suite.Run("A=0=AllApisHitsForASpecificRole", func() {
			// Creating Project with Super Admin
			var devtronDeletion RequestDTOs.RbacDevtronDeletion
			saveTeamRequestDto := TeamRouter.GetSaveTeamRequestDto()
			saveTeamRequestDto.Name = UserRouter.PROJECT
			byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)

			responseOfCreateProject := CreateProject(byteValueOfStruct, suite.authToken)
			assert.Equal(suite.T(), 200, responseOfCreateProject.Code)
			assert.Equal(suite.T(), UserRouter.PROJECT, responseOfCreateProject.Result.Name)
			devtronDeletion.ProjectPayload = byteValueOfStruct

			// Creating environment with SuperAdmin
			environments := strings.Split(UserRouter.ENV, ",")
			saveEnvRequestDto := GetSaveEnvRequestDto()
			saveEnvRequestDto.Environment = environments[0]
			byteValueOfStruct, _ = json.Marshal(saveEnvRequestDto)
			responseOfCreateEnvironment := CreateEnv(byteValueOfStruct, suite.authToken)
			assert.Equal(suite.T(), 200, responseOfCreateEnvironment.Code)
			assert.Equal(suite.T(), environments[0], responseOfCreateEnvironment.Result.Environment)
			devtronDeletion.EnvPayLoad = byteValueOfStruct

			//Application With SuperAdmin
			applications := strings.Split(UserRouter.APP, ",")
			responseOfCreateDevtronApp := CreateDevtronApp(applications[0], suite.authToken, responseOfCreateProject.Result.Id)
			assert.Equal(suite.T(), 200, responseOfCreateDevtronApp.Code)
			assert.Equal(suite.T(), applications[0], responseOfCreateDevtronApp.Result.AppName)
			assert.Equal(suite.T(), responseOfCreateProject.Result.Id, responseOfCreateDevtronApp.Result.TeamId)
			devtronDeletion.DevtronPayload = responseOfCreateDevtronApp

			//Creating RoleGroup
			createRoleGroupPayload := UserRouter.CreateRoleGroupPayloadDynamicForDevtronApp(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, UserRouter.ACTION, UserRouter.ACCESS_TYPE)
			createRoleGroupPayload.RoleFilters[0].Action = role
			byteValueOfStruct, _ = json.Marshal(createRoleGroupPayload)
			log.Println("Hitting Create Role Group API")
			createRoleGroupResponseBody := UserRouter.HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)
			assert.Equal(suite.T(), createRoleGroupPayload.Name, createRoleGroupResponseBody.Result.Name)

			devtronDeletion.RoleGroupId = createRoleGroupResponseBody.Result.Id
			log.Println("Verifying the response of Create Role Group API using getRoleGroupById API")
			getRoleGroupByIdResponse := UserRouter.HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
			assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))
			assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Team, createRoleGroupResponseBody.Result.RoleFilters[0].Team)
			assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Action, createRoleGroupResponseBody.Result.RoleFilters[0].Action)

			//Creating API for specific Permission Groups
			var tokenDeletion RequestDTOs.RbacApiTokenDeletion
			createApiTokenRequestDTO := ApiTokenRouter.GetPayLoadForCreateApiToken()
			payloadForCreateApiTokenRequest, _ := json.Marshal(createApiTokenRequestDTO)
			responseOfCreateApiToken := ApiTokenRouter.HitCreateApiTokenApi(string(payloadForCreateApiTokenRequest), suite.authToken)
			assert.Equal(suite.T(), "API-TOKEN:"+createApiTokenRequestDTO.Name, responseOfCreateApiToken.Result.UserIdentifier)
			tokenDeletion.ApiTokenId = responseOfCreateApiToken.Result.UserId

			//createUserDto, roleGroupId := UserRouter.CreateUserRequestPayload(UserRouter.GroupsAndRoleFilter, suite.authToken)
			createUserDto, apiToken, roleGroupId := CreateUserPayloadForDynamicToken(responseOfCreateApiToken, suite.authToken, false)
			//createRoleGroupPayload := UserRouter.CreateRoleGroupPayloadDynamicForDevtronApp(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, role, UserRouter.ACCESS_TYPE)
			createUserDto.RoleFilters = createRoleGroupPayload.RoleFilters

			byteValueOfStruct, _ = json.Marshal(createUserDto)
			log.Println("Hitting the Create User API")
			responseOfCreateUserApi := UserRouter.HitCreateUserApi(byteValueOfStruct, suite.authToken)
			assert.Equal(suite.T(), false, responseOfCreateUserApi.Result[0].SuperAdmin)
			assert.Equal(suite.T(), createUserDto.EmailId, responseOfCreateUserApi.Result[0].EmailId)
			assert.Equal(suite.T(), createUserDto.Groups[0], responseOfCreateUserApi.Result[0].Groups[0])
			assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Action, responseOfCreateUserApi.Result[0].RoleFilters[0].Action)
			assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Team, responseOfCreateUserApi.Result[0].RoleFilters[0].Team)
			tokenDeletion.UserId = responseOfCreateUserApi.Result[0].Id

			log.Println("Hitting the get user by id for verifying the functionality of UpdateUserApi")
			responseOfGetUserById := UserRouter.HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
			assert.Equal(suite.T(), false, responseOfGetUserById.Result.SuperAdmin)
			assert.Equal(suite.T(), responseOfCreateUserApi.Result[0].EmailId, responseOfGetUserById.Result.EmailId)
			assert.Equal(suite.T(), responseOfCreateUserApi.Result[0].Groups, responseOfGetUserById.Result.Groups)
			assert.Equal(suite.T(), responseOfCreateUserApi.Result[0].RoleFilters, responseOfGetUserById.Result.RoleFilters)
			//

			//log.Println("Deleting the Test data Created via Automation")
			//UserRouter.HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
			////UserRouter.HitDeleteRoleGroupByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
			//UserRouter.HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
			log.Println(roleGroupId)
			//})

			//return resultToken, tokenDeletion

			//createRoleGroupResponseBody, createRoleGroupPayload, deleteDevtron := suite.CreateSpecificPermissionGroup(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, role, UserRouter.ACCESS_TYPE)
			//apiToken, deleteToken := suite.CreateTokenForSpecificPermissionGroup(createRoleGroupResponseBody, false, role, createRoleGroupPayload)

			log.Println("Test Case for User ===>", apiToken)
			appName := "app" + strings.ToLower(testUtils.GetRandomStringOfGivenLength(5))
			createAppRequestDto := PipelineConfigRouter.GetAppRequestDto(appName, 1, 0)
			byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
			createAppResponseDto := PipelineConfigRouter.HitCreateAppApi(byteValueOfCreateApp, appName, responseOfCreateProject.Result.Id, 0, apiToken)
			statusCode := getExpectedStatusCode(role, UserRouter.CREATEAPP)

			assert.Equal(suite.T(), true, getStatusCheck(statusCode, createAppResponseDto.Code))
			if createAppResponseDto.Code != 200 && getStatusCheck(statusCode, createAppResponseDto.Code) {

				//createRoleGroupResponseBody := suite.CreateSpecificPermissionGroup(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, UserRouter.ACTION, UserRouter.ACCESS_TYPE)
				//superAdminToken := suite.CreateTokenForSpecificPermissionGroup(createRoleGroupResponseBody, true, role)
				createAppResponseDto = PipelineConfigRouter.HitCreateAppApi(byteValueOfCreateApp, appName, 1, 0, suite.authToken)
			}

			//createRoleGroupResponseBody := suite.CreateSpecificPermissionGroup(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, role, UserRouter.ACCESS_TYPE)
			//apiToken := suite.CreateTokenForSpecificPermissionGroup(createRoleGroupResponseBody, false, role)
			PayloadForApiFetchAppsByEnvironment := AppListingRouter.GetPayloadForApiFetchAppsByEnvironment()
			bytePayloadForTriggerCiPipeline, _ := json.Marshal(PayloadForApiFetchAppsByEnvironment)

			log.Println("Test Case for User ===>", apiToken)
			allAppsByEnvironment := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, apiToken)
			statusCode = getExpectedStatusCode(role, UserRouter.APPLISTFETCH)
			assert.Equal(suite.T(), true, getStatusCheck(statusCode, allAppsByEnvironment.Code))
			//to check	assert.Equal(suite.T(), len(strings.Split(APP,",")), allAppsByEnvironment.Result.AppCount)
			assert.Equal(suite.T(), UserRouter.APP, allAppsByEnvironment.Result.AppContainers[0].AppName)
			//assert.Equal(suite.T(), UserRouter.ENV, allAppsByEnvironment.Result.AppContainers[0].Environments[0].EnvironmentName)
			assert.Equal(suite.T(), UserRouter.PROJECT, allAppsByEnvironment.Result.AppContainers[0].Environments[0].TeamName)

			PipelineConfigRouter.HitDeleteAppApi(byteValueOfCreateApp, createAppResponseDto.Result.Id, suite.authToken)

			//Test case for Getting CI-CDPipeline
			createAppApiResponse, workflowResponse := PipelineConfigRouter.CreateNewAppWithCiCd(suite.authToken, false)
			time.Sleep(2 * time.Second)
			log.Println("=== Here we are getting pipeline material ===")
			appId := strconv.Itoa(createAppApiResponse.Id)

			//createRoleGroupResponseBody := suite.CreateSpecificPermissionGroup(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, role, UserRouter.ACCESS_TYPE)
			//apiToken := suite.CreateTokenForSpecificPermissionGroup(createRoleGroupResponseBody, false, role)
			log.Println("Test Case for User ===>", apiToken)
			cdPipelineStrategiesResponse := PipelineConfigRouter.HitGetCdPipelineStrategies(appId, apiToken)
			statusCode = getExpectedStatusCode(role, UserRouter.PIPELINEFETCH)
			assert.Equal(suite.T(), true, getStatusCheck(statusCode, cdPipelineStrategiesResponse.Code))

			//_, workflowResponse := PipelineConfigRouter.CreateNewAppWithCiCd(suite.authToken)
			////time.Sleep(2 * time.Second)
			log.Println("=== Here we are getting pipeline material ===")
			pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.Result.CiPipelines[0].Id, suite.authToken)
			payloadForTriggerCiPipeline := PipelineConfigRouter.CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, workflowResponse.Result.CiPipelines[0].Id, pipelineMaterial.Result[0].Id, true)
			bytePayloadForTriggerCiPipeline, _ = json.Marshal(payloadForTriggerCiPipeline)

			//createRoleGroupResponseBody := suite.CreateSpecificPermissionGroup(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, role, UserRouter.ACCESS_TYPE)
			//apiToken := suite.CreateTokenForSpecificPermissionGroup(createRoleGroupResponseBody, false, role)
			log.Println("Test Case for User ===>", apiToken)
			triggerCiPipelineResponse := PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, apiToken)
			statusCode = getExpectedStatusCode(role, UserRouter.PIPELINECREATE)
			assert.Equal(suite.T(), true, getStatusCheck(statusCode, triggerCiPipelineResponse.Code))
			if getStatusCheck(statusCode, triggerCiPipelineResponse.Code) && triggerCiPipelineResponse.Code != 200 {
				triggerCiPipelineResponse = PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
			}

			PipelineConfigRouter.DeleteAppWithCiCd(suite.authToken)
			UserRouter.HitDeleteUserApi(strconv.Itoa(tokenDeletion.UserId), suite.authToken)

			//log.Println("=== Here We Deleting the Token After Verification")
			responseOfDeleteApi := ApiTokenRouter.HitDeleteApiToken(strconv.Itoa(tokenDeletion.ApiTokenId), suite.authToken)
			assert.True(suite.T(), responseOfDeleteApi.Result.Success)

			UserRouter.HitDeleteRoleGroupByIdApi(strconv.Itoa(devtronDeletion.RoleGroupId), suite.authToken)
			DeleteDevtronApp(devtronDeletion.DevtronPayload.Result.Id, devtronDeletion.DevtronPayload.Result.AppName, devtronDeletion.DevtronPayload.Result.TeamId, devtronDeletion.DevtronPayload.Result.TemplateId, suite.authToken)
			DeleteEnv(devtronDeletion.EnvPayLoad, suite.authToken)
			DeleteProject(devtronDeletion.ProjectPayload, suite.authToken)

			//UserRouter.HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
			//UserRouter.HitDeleteRoleGroupByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
			//UserRouter.HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
		})

	}
	suite.Run("A=5=HitApiGetAppsListWithSuperAdminUsersAccess", func() {

		//createRoleGroupResponseBody, createRoleGroupPayload, deleteDevtron := suite.CreateSpecificPermissionGroup(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, UserRouter.ACTION, UserRouter.ACCESS_TYPE)
		//superAdminToken, deleteToken := suite.CreateTokenForSpecificPermissionGroup(createRoleGroupResponseBody, true, "", createRoleGroupPayload)
		superAdminToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzdXBlci1hZG1pbi10b2tlbiIsImlzcyI6ImFwaVRva2VuSXNzdWVyIiwiZXhwIjoxNjc5MzE5MjcxfQ.STZ4ptJkaCNral6vHMIA0zL_p8kC3NlcKX69unEfiwg"
		//log.Println(deleteDevtron, deleteToken)
		PayloadForApiFetchAppsByEnvironment := AppListingRouter.GetPayloadForApiFetchAppsByEnvironment()
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(PayloadForApiFetchAppsByEnvironment)
		allAppsViaArgoAdminToken := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, superAdminToken)
		allAppsViaSuperAdminToken := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, suite.authToken)
		assert.Equal(suite.T(), allAppsViaArgoAdminToken.Result.AppCount, allAppsViaSuperAdminToken.Result.AppCount)
	})
}
