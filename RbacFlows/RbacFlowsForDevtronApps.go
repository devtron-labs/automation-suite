package RbacFlows

import (
	"automation-suite/ApiTokenRouter"
	"automation-suite/ApiTokenRouter/ResponseDTOs"
	"automation-suite/AppListingRouter"
	"automation-suite/HelperRouter"
	"automation-suite/PipelineConfigRouter"
	dtos "automation-suite/PipelineConfigRouter/ResponseDTOs"
	"automation-suite/RbacFlows/RequestDTOs"
	"automation-suite/TeamRouter"
	"automation-suite/testdata/testUtils"
	"github.com/tidwall/sjson"
	"time"

	//"automation-suite/RbacFlows/RequestDTOs"
	"automation-suite/UserRouter"
	abcd "automation-suite/UserRouter/RequestDTOs"
	Base "automation-suite/testUtils"
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

	var allRoles = []string{"view"}
	for _, role := range allRoles {
		var (
			createAppApiResponsePtr *PipelineConfigRouter.CreateAppRequestDto
			workflowResponsePtr     *dtos.CreateWorkflowResponseDto
			savePipelineResponsePtr *dtos.SaveCdPipelineResponseDTO
		)
		suite.Run("A=0=AllApisHitsForASpecificRole", func() {
			// Creating Project with Super Admin
			var devtronDeletion RequestDTOs.RbacDevtronDeletion
			saveTeamRequestDto := TeamRouter.GetSaveTeamRequestDto()
			saveTeamRequestDto.Name = UserRouter.PROJECT
			byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)

			responseOfCreateProject := CreateProject(byteValueOfStruct, suite.authToken)
			assert.Equal(suite.T(), 200, responseOfCreateProject.Code)
			assert.Equal(suite.T(), UserRouter.PROJECT, responseOfCreateProject.Result.Name)
			devtronDeletion.ProjectPayload, _ = json.Marshal(responseOfCreateProject.Result)

			// Creating environment with SuperAdmin
			environments := strings.Split(UserRouter.ENV, ",")
			saveEnvRequestDto := GetSaveEnvRequestDto()
			saveEnvRequestDto.Environment = environments[0]
			saveEnvRequestDto.EnvironmentIdentifier = "default_cluster__" + saveEnvRequestDto.Namespace
			byteValueOfStruct, _ = json.Marshal(saveEnvRequestDto)
			responseOfCreateEnvironment := CreateEnv(byteValueOfStruct, suite.authToken)
			assert.Equal(suite.T(), 200, responseOfCreateEnvironment.Code)
			assert.Equal(suite.T(), environments[0], responseOfCreateEnvironment.Result.Environment)

			devtronDeletion.EnvPayLoad, _ = json.Marshal(responseOfCreateEnvironment.Result)

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
			tokenDeletion.ApiTokenId = responseOfCreateApiToken.Result.Id

			createUserDto, apiToken, roleGroupId := CreateUserPayloadForDynamicToken(responseOfCreateApiToken, suite.authToken, false)
			createUserDto.RoleFilters = createRoleGroupPayload.RoleFilters

			// Setting up environment identifier for default Cluster (cluster +"__++namespace)
			createUserDto.RoleFilters[0].Environment = "default_cluster__" + responseOfCreateEnvironment.Result.Namespace
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

			log.Println("Test Case for User ===>", apiToken)

			createAppResponseDto := responseOfCreateDevtronApp
			Envs := []int{}
			Teams := []int{1}
			Namespaces := []string{}
			AppStatuses := []string{}
			requestDTOForApiFetchAppsByEnvironment := AppListingRouter.GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
			bytePayloadForTriggerCiPipeline, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)

			log.Println("Test Case for User ===>", apiToken)
			allAppsByEnvironment := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, apiToken)
			statusCode := getExpectedStatusCode(role, UserRouter.APPLISTFETCH)
			assert.Equal(suite.T(), true, getStatusCheck(statusCode, allAppsByEnvironment.Code))
			//to check	assert.Equal(suite.T(), len(strings.Split(APP,",")), allAppsByEnvironment.Result.AppCount)
			assert.Equal(suite.T(), UserRouter.APP, allAppsByEnvironment.Result.AppContainers[0].AppName)
			//assert.Equal(suite.T(), UserRouter.ENV, allAppsByEnvironment.Result.AppContainers[0].Environments[0].EnvironmentName)
			assert.Equal(suite.T(), UserRouter.PROJECT, allAppsByEnvironment.Result.AppContainers[0].Environments[0].TeamName)

			createAppApiResponse := createAppResponseDto.Result
			workflowResponse := dtos.CreateWorkflowResponseDto{}
			if createAppApiResponsePtr != nil && workflowResponsePtr != nil {
				createAppApiResponse = *createAppApiResponsePtr
				workflowResponse = *workflowResponsePtr
			} else {
				envConf := Base.ReadBaseEnvConfig()
				config := Base.ReadAnyJsonFile(envConf.ClassCredentialsFile)
				var configId int
				log.Println("=== Here we are creating App Material ===")
				createAppMaterialRequestDto := PipelineConfigRouter.GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
				appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
				createAppMaterialResponse := PipelineConfigRouter.HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)
				statusCode := getExpectedStatusCode(role, UserRouter.CREATEAPPMATERIAL)

				assert.Equal(suite.T(), true, getStatusCheck(statusCode, createAppMaterialResponse.Code))
				if createAppMaterialResponse.Code != 200 && getStatusCheck(statusCode, createAppMaterialResponse.Code) {
					createAppMaterialResponse = PipelineConfigRouter.HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)
				}

				log.Println("=== Here we are saving docker build config ===")
				requestPayloadForSaveAppCiPipeline := PipelineConfigRouter.GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, "test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
				byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
				saveAppCiPipelineresponse := PipelineConfigRouter.HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, apiToken)

				assert.Equal(suite.T(), true, getStatusCheck(statusCode, saveAppCiPipelineresponse.Code))
				if saveAppCiPipelineresponse.Code != 200 && getStatusCheck(statusCode, saveAppCiPipelineresponse.Code) {

					saveAppCiPipelineresponse = PipelineConfigRouter.HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)
				}
				log.Println("=== Here we are fetching latestChartReferenceId ===")
				time.Sleep(2 * time.Second)
				getChartReferenceResponse := PipelineConfigRouter.HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), apiToken)
				latestChartRef := getChartReferenceResponse.Result.LatestChartRef

				log.Println("=== Here we are fetching Template using getAppTemplateAPI ===")
				getTemplateResponse := PipelineConfigRouter.HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), apiToken)

				log.Println("=== Here we are fetching DefaultAppOverride from template response ===")
				defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

				log.Println("=== Here we are creating payload for SaveTemplate API ===")
				saveDeploymentTemplate := PipelineConfigRouter.GetRequestPayloadForSaveDeploymentTemplate(createAppApiResponse.Id, latestChartRef, defaultAppOverride)
				byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)
				jsonOfSaveDeploymentTemp := string(byteValueOfSaveDeploymentTemplate)
				jsonWithMicroserviceToleration, _ := sjson.Set(jsonOfSaveDeploymentTemp, "valuesOverride.tolerations.0", map[string]interface{}{"effect": "NoSchedule", "key": "microservice", "operator": "Equal", "value": "true"})
				finalJson, _ := sjson.Set(jsonWithMicroserviceToleration, "valuesOverride.tolerations.1", map[string]interface{}{"effect": "NoSchedule", "key": "kubernetes.azure.com/scalesetpriority", "operator": "Equal", "value": "spot"})
				updatedByteValueOfSaveDeploymentTemplate := []byte(finalJson)

				log.Println("=== Here we are hitting SaveTemplate API ===")
				PipelineConfigRouter.HitSaveDeploymentTemplateApi(updatedByteValueOfSaveDeploymentTemplate, apiToken)

				log.Println("=== Here we are saving Global Configmap ===")
				requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(0, "-config1", createAppApiResponse.Id, "environment", "kubernetes", false, false, false, false)
				byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
				globalConfigMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, apiToken)
				configId = globalConfigMap.Result.Id

				log.Println("=== Here we are saving Global Secret ===")
				requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, "-secret1", createAppApiResponse.Id, "environment", "kubernetes", false, false, true, false)
				byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
				HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, apiToken)

				log.Println("=== Here we are saving workflow with Pre/Post CI ===")
				workflowResponse = PipelineConfigRouter.HitCreateWorkflowApiWithFullPayload(createAppApiResponse.Id, apiToken)
				preStageScript, _ := testUtils.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
				postStageScript, _ := testUtils.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")

				log.Println("=== Here we are saving CD pipeline ===")
				payload := PipelineConfigRouter.GetRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.Result.AppWorkflowId, responseOfCreateEnvironment.Result.Id, workflowResponse.Result.CiPipelines[0].Id, workflowResponse.Result.CiPipelines[0].ParentCiPipeline, "AUTOMATIC", string(preStageScript), string(postStageScript), "AUTOMATIC")
				bytePayload, _ := json.Marshal(payload)
				savePipelineResponse := PipelineConfigRouter.HitSaveCdPipelineApi(bytePayload, apiToken)

				createAppApiResponsePtr = &createAppApiResponse
				workflowResponsePtr = &workflowResponse
				savePipelineResponsePtr = &savePipelineResponse

			}

			//Test case for Getting CI-CDPipeline

			time.Sleep(2 * time.Second)
			log.Println("=== Here we are getting pipeline material ===")
			appId := strconv.Itoa(createAppApiResponse.Id)

			log.Println("Test Case for User ===>", apiToken)
			cdPipelineStrategiesResponse := PipelineConfigRouter.HitGetCdPipelineStrategies(appId, apiToken)
			statusCode = getExpectedStatusCode(role, UserRouter.PIPELINEFETCH)
			assert.Equal(suite.T(), true, getStatusCheck(statusCode, cdPipelineStrategiesResponse.Code))

			time.Sleep(2 * time.Second)
			log.Println("=== Here we are getting pipeline material ===")
			pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.Result.CiPipelines[0].Id, suite.authToken)
			payloadForTriggerCiPipeline := PipelineConfigRouter.CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, workflowResponse.Result.CiPipelines[0].Id, pipelineMaterial.Result[0].Id, true)
			bytePayloadForTriggerCiPipeline, _ = json.Marshal(payloadForTriggerCiPipeline)

			log.Println("Test Case for User ===>", apiToken)
			triggerCiPipelineResponse := PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, apiToken)
			statusCode = getExpectedStatusCode(role, UserRouter.PIPELINECREATE)
			assert.Equal(suite.T(), true, getStatusCheck(statusCode, triggerCiPipelineResponse.Code))
			if getStatusCheck(statusCode, triggerCiPipelineResponse.Code) && triggerCiPipelineResponse.Code != 200 {
				triggerCiPipelineResponse = PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, suite.authToken)
			}

			log.Println("=== Here we are Deleting the CD pipeline ===")
			deletePipelinePayload := PipelineConfigRouter.GetPayloadForDeleteCdPipeline(createAppApiResponsePtr.Id, savePipelineResponsePtr.Result.Pipelines[0].Id)
			deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)

			PipelineConfigRouter.HitForceDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)

			log.Println("=== Here we are Deleting the CI pipeline ===")
			PipelineConfigRouter.DeleteCiPipeline(createAppApiResponsePtr.Id, workflowResponsePtr.Result.CiPipelines[0].Id, suite.authToken)

			log.Println("=== Here we are Deleting CI Workflow ===")
			PipelineConfigRouter.HitDeleteWorkflowApi(createAppApiResponsePtr.Id, workflowResponsePtr.Result.AppWorkflowId, suite.authToken)
			log.Println("=== Here we are Deleting the app after all verifications ===")
			testUtils.DeleteApp(createAppApiResponsePtr.Id, createAppApiResponsePtr.AppName, createAppApiResponsePtr.TeamId, createAppApiResponsePtr.TemplateId, suite.authToken)

			UserRouter.HitDeleteUserApi(strconv.Itoa(tokenDeletion.UserId), suite.authToken)

			//log.Println("=== Here We Deleting the Token After Verification")
			//responseOfDeleteApi := ApiTokenRouter.HitDeleteApiToken(strconv.Itoa(tokenDeletion.ApiTokenId), suite.authToken)
			//assert.True(suite.T(), responseOfDeleteApi.Result.Success)

			UserRouter.HitDeleteRoleGroupByIdApi(strconv.Itoa(devtronDeletion.RoleGroupId), suite.authToken)
			DeleteDevtronApp(devtronDeletion.DevtronPayload.Result.Id, devtronDeletion.DevtronPayload.Result.AppName, devtronDeletion.DevtronPayload.Result.TeamId, devtronDeletion.DevtronPayload.Result.TemplateId, suite.authToken)
			DeleteEnv(devtronDeletion.EnvPayLoad, suite.authToken)
			DeleteProject(devtronDeletion.ProjectPayload, suite.authToken)

			UserRouter.HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
		})

	}
	suite.Run("A=5=HitApiGetAppsListWithSuperAdminUsersAccess", func() {

		//createRoleGroupResponseBody, createRoleGroupPayload, deleteDevtron := suite.CreateSpecificPermissionGroup(UserRouter.ENTITY, UserRouter.PROJECT, UserRouter.ENV, UserRouter.APP, UserRouter.ACTION, UserRouter.ACCESS_TYPE)
		//superAdminToken, deleteToken := suite.CreateTokenForSpecificPermissionGroup(createRoleGroupResponseBody, true, "", createRoleGroupPayload)
		//log.Println(deleteDevtron, deleteToken)
		superAdminToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzdXBlcmFkbWluIiwiaXNzIjoiYXBpVG9rZW5Jc3N1ZXIiLCJleHAiOjE2ODAxMDA1NTJ9.8TB9v0JppMw5YDhZ05H82DQb2sdjRNWwroHtlnmC4DU"

		Envs := []int{}
		Teams := []int{1}
		Namespaces := []string{}
		AppStatuses := []string{}
		requestDTOForApiFetchAppsByEnvironment := AppListingRouter.GetPayloadForApiFetchAppsByEnvironment(Envs, Teams, Namespaces, "", AppStatuses, "ASC", 0, 0, 10)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(requestDTOForApiFetchAppsByEnvironment)
		allAppsViaArgoAdminToken := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, superAdminToken)
		allAppsViaSuperAdminToken := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, suite.authToken)
		assert.Equal(suite.T(), allAppsViaArgoAdminToken.Result.AppCount, allAppsViaSuperAdminToken.Result.AppCount)
	})
}
