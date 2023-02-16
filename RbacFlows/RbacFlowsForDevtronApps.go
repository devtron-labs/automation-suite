package RbacFlows

import (
	AppListingRouter "automation-suite/AppListingRouter"
	"automation-suite/PipelineConfigRouter"
	"automation-suite/RbacFlows/RequestDTOs"
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *RbacFlowTestSuite) TestRbacFlowsForDevtronApps() {
	token_spec_proj_env_app_view_devtron := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzcGVjLXByb2pfZW52X2FwcC12aWV3LWRldnRyb24iLCJpc3MiOiJhcGlUb2tlbklzc3VlciJ9.4s3KnwlSZELxxWoLxnniUOtaAmR5fD7KpMJGAw6g3hE"
	specific_proj_env_app_manager := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzcGVjaWZpY19wcm9qX2Vudl9hcHBfbWFuYWdlciIsImlzcyI6ImFwaVRva2VuSXNzdWVyIiwiZXhwIjoxNjc5Mjk0MDQ3fQ.CWzv0z50YtG0troArk3MQmrmJQocFmB97VUsXwHlMew"
	specific_proj_env_app_build_deploy_devtron := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzcGVjaWZpY19wcm9qX2Vudl9hcHBfYnVpbGRfZGVwbG95X2RldnRyb24iLCJpc3MiOiJhcGlUb2tlbklzc3VlciJ9.PoHuMqeUAlGHczJWrBcO4X6u5v89lnyj-7V01VA2WQA"
	specific_proj_env_app_admin_devtron_only := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzcGVjaWZpY19wcm9qX2Vudl9hcHBfYWRtaW5fZGV2dHJvbl9vbmx5IiwiaXNzIjoiYXBpVG9rZW5Jc3N1ZXIifQ.zDpcAZS2GamHCWwZFeB4fXTU6JxuirbUxPeSqz9qlgU"

	specific_proj_all_env_spec_app_devtron_manager := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzcGVjaWZpY19wcm9qX2FsbF9lbnZfc3BlY19hcHBfZGV2dHJvbl9tYW5hZ2VyIiwiaXNzIjoiYXBpVG9rZW5Jc3N1ZXIiLCJleHAiOjE2NzkzMjIxMzR9.nUAoSD8ztE6rRo1yqR7ymktN7KSFLqSo5R_A330zCuA"
	specific_proj_all_env_spec_app_devtron_admin := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzcGVjaWZpY19wcm9qX2FsbF9lbnZfc3BlY19hcHBfZGV2dHJvbl9hZG1pbiIsImlzcyI6ImFwaVRva2VuSXNzdWVyIiwiZXhwIjoxNjc5MzIyMjMzfQ.Jkdbi1-ckWM42Vfzcq4bSKzvF2WsWLZzaONrmn5YeYY"
	specific_proj_all_env_spec_app_devtron_build_deploy := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzcGVjaWZpY19wcm9qX2FsbF9lbnZfc3BlY19hcHBfZGV2dHJvbl9idWlsZF9kcGx5IiwiaXNzIjoiYXBpVG9rZW5Jc3N1ZXIiLCJleHAiOjE2NzkzMjIzMDR9.zj9QYC4QSpX2NMjT988pr5jEmQjvsiikVV8Z08EfgJc"
	specific_proj_all_env_spec_app_devtron_view_only := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzcGVjaWZpY19wcm9qX2FsbF9lbnZfc3BlY19hcHBfZGV2dHJvbl92aWV3X29ubHkiLCJpc3MiOiJhcGlUb2tlbklzc3VlciIsImV4cCI6MTY3OTMyMjM1Nn0.wVVrZ721zV7QcPo1b0JVasF8LV_pCKIdWcUUQaAlmhU"

	specific_proj_all_env_all_app_devtron_view_only := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpwZWNpZmljX3Byb2pfYWxsX2Vudl9hbGxfYXBwX2RldnRyb25fdmlld19vbmx5IiwiaXNzIjoiYXBpVG9rZW5Jc3N1ZXIifQ.EDBPHJTOGlokdXAOsyExRROHbvA9ze1Kv4eNgIvTbBY"
	specific_proj_all_env_all_app_devtron_build_deploy := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpwZWNpZmljX3Byb2pfYWxsX2Vudl9hbGxfYXBwX2RldnRyb25fYnVpbGRfZGVwbG95IiwiaXNzIjoiYXBpVG9rZW5Jc3N1ZXIifQ.Cxhikq8qaHx9NhcPFoILYfa_VB-oKT1A1gZEn728FfM"
	specific_proj_all_env_all_app_devtron_admin := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpwZWNpZmljX3Byb2pfYWxsX2Vudl9hbGxfYXBwX2RldnRyb25fYWRtaW4iLCJpc3MiOiJhcGlUb2tlbklzc3VlciJ9.qgS4OfcUHXRUpJwC_gpzOLAL1TwMplsR-ggZw_76Seo"
	specific_proj_all_env_all_app_devtron_manager := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpwZWNpZmljX3Byb2pfYWxsX2Vudl9hbGxfYXBwX2RldnRyb25fbWFuYWdlciIsImlzcyI6ImFwaVRva2VuSXNzdWVyIn0.VD5qeLYYKb1C4_xHZxSTeLpnBsf3Sw26IZjNtu6KqOA"

	var usersWithAccessOfSpecificProjEnvAndAppForGettingAppList = []RequestDTOs.RbacDevtronUserApiDTO{
		{"specific_proj_env_app-view-devtron", token_spec_proj_env_app_view_devtron, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_manager", specific_proj_env_app_manager, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_build_deploy_devtron", specific_proj_env_app_build_deploy_devtron, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_admin_devtron_only", specific_proj_env_app_admin_devtron_only, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_manager", specific_proj_all_env_spec_app_devtron_manager, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_admin", specific_proj_all_env_spec_app_devtron_admin, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_build_deploy", specific_proj_all_env_spec_app_devtron_build_deploy, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_view_only", specific_proj_all_env_spec_app_devtron_view_only, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_view_only", specific_proj_all_env_all_app_devtron_view_only, 200, 2, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_build_deploy", specific_proj_all_env_all_app_devtron_build_deploy, 200, 2, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_admin", specific_proj_all_env_all_app_devtron_admin, 200, 2, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_manager", specific_proj_all_env_all_app_devtron_manager, 200, 2, "test-app-1", "test-project-1", "test-env-1"},
	}
	var usersWithAccessOfSpecificProjAllEnvAndAllAppForGettingAppList = []RequestDTOs.RbacDevtronUserApiDTO{
		{"specific_proj_all_env_all_app_devtron_view_only", specific_proj_all_env_all_app_devtron_view_only, 200, 2, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_build_deploy", specific_proj_all_env_all_app_devtron_build_deploy, 200, 2, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_admin", specific_proj_all_env_all_app_devtron_admin, 200, 2, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_manager", specific_proj_all_env_all_app_devtron_manager, 200, 2, "test-app-1", "test-project-1", "test-env-1"},
	}

	var UnauthorisedUsersToGetCdPipelineStrategies = []RequestDTOs.RbacDevtronUserApiDTO{
		{"spec-proj_env_app-view-devtron", token_spec_proj_env_app_view_devtron, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_manager", specific_proj_env_app_manager, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_build_deploy_devtron", specific_proj_env_app_build_deploy_devtron, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_admin_devtron_only", specific_proj_env_app_admin_devtron_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_manager", specific_proj_all_env_spec_app_devtron_manager, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_admin", specific_proj_all_env_spec_app_devtron_admin, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_build_deploy", specific_proj_all_env_spec_app_devtron_build_deploy, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_view_only", specific_proj_all_env_spec_app_devtron_view_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_view_only", specific_proj_all_env_all_app_devtron_view_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_build_deploy", specific_proj_all_env_all_app_devtron_build_deploy, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_admin", specific_proj_all_env_all_app_devtron_admin, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_manager", specific_proj_all_env_all_app_devtron_manager, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
	}
	var UnauthorisedUsersToTriggerCiPipeline = []RequestDTOs.RbacDevtronUserApiDTO{
		{"spec-proj_env_app-view-devtron", token_spec_proj_env_app_view_devtron, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_manager", specific_proj_env_app_manager, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_build_deploy_devtron", specific_proj_env_app_build_deploy_devtron, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_admin_devtron_only", specific_proj_env_app_admin_devtron_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_manager", specific_proj_all_env_spec_app_devtron_manager, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_admin", specific_proj_all_env_spec_app_devtron_admin, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_build_deploy", specific_proj_all_env_spec_app_devtron_build_deploy, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_view_only", specific_proj_all_env_spec_app_devtron_view_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_view_only", specific_proj_all_env_all_app_devtron_view_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_build_deploy", specific_proj_all_env_all_app_devtron_build_deploy, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_admin", specific_proj_all_env_all_app_devtron_admin, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_manager", specific_proj_all_env_all_app_devtron_manager, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
	}
	var UnauthorisedUsersToCreateApp = []RequestDTOs.RbacDevtronUserApiDTO{
		{"spec-proj_env_app-view-devtron", token_spec_proj_env_app_view_devtron, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_manager", specific_proj_env_app_manager, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_build_deploy_devtron", specific_proj_env_app_build_deploy_devtron, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_admin_devtron_only", specific_proj_env_app_admin_devtron_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_manager", specific_proj_all_env_spec_app_devtron_manager, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_admin", specific_proj_all_env_spec_app_devtron_admin, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_build_deploy", specific_proj_all_env_spec_app_devtron_build_deploy, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_spec_app_devtron_view_only", specific_proj_all_env_spec_app_devtron_view_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_view_only", specific_proj_all_env_all_app_devtron_view_only, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_all_env_all_app_devtron_build_deploy", specific_proj_all_env_all_app_devtron_build_deploy, 403, 1, "test-app-1", "test-project-1", "test-env-1"},
	}

	var AuthorisedUsersToGetCdPipelineStrategies = []RequestDTOs.RbacDevtronUserApiDTO{
		{"specific_proj_env_app_manager", specific_proj_env_app_manager, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_build_deploy_devtron", specific_proj_env_app_build_deploy_devtron, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_admin_devtron_only", specific_proj_env_app_admin_devtron_only, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
	}

	var AuthorisedUsersToTriggerCiPipelines = []RequestDTOs.RbacDevtronUserApiDTO{
		{"specific_proj_env_app_manager", specific_proj_env_app_manager, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_build_deploy_devtron", specific_proj_env_app_build_deploy_devtron, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
		{"specific_proj_env_app_admin_devtron_only", specific_proj_env_app_admin_devtron_only, 200, 1, "test-app-1", "test-project-1", "test-env-1"},
	}

	suite.Run("A=1=HitApiWithUnauthorisedUsersToGetCdPipelineStrategies", func() {
		createAppApiResponse, _ := PipelineConfigRouter.CreateNewAppWithCiCd(suite.authToken)
		time.Sleep(2 * time.Second)
		log.Println("=== Here we are getting pipeline material ===")
		appId := strconv.Itoa(createAppApiResponse.Id)
		for _, user := range UnauthorisedUsersToGetCdPipelineStrategies {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			cdPipelineStrategiesResponse := PipelineConfigRouter.HitGetCdPipelineStrategies(appId, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, cdPipelineStrategiesResponse.Code)
		}
	})

	suite.Run("A=2=HitApiWithUnauthorisedUsersToTriggerCiPipeline", func() {
		_, workflowResponse := PipelineConfigRouter.CreateNewAppWithCiCd(suite.authToken)
		time.Sleep(2 * time.Second)
		log.Println("=== Here we are getting pipeline material ===")
		pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.Result.CiPipelines[0].Id, suite.authToken)
		payloadForTriggerCiPipeline := PipelineConfigRouter.CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, workflowResponse.Result.CiPipelines[0].Id, pipelineMaterial.Result[0].Id, true)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		for _, user := range UnauthorisedUsersToTriggerCiPipeline {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			triggerCiPipelineResponse := PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, triggerCiPipelineResponse.Code)
		}
		PipelineConfigRouter.DeleteAppWithCiCd(suite.authToken)
	})

	suite.Run("A=3=HitApiGetAppsListWithUsersAccessOfSpecificProjEnvAndApp", func() {
		PayloadForApiFetchAppsByEnvironment := AppListingRouter.GetPayloadForApiFetchAppsByEnvironment()
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(PayloadForApiFetchAppsByEnvironment)
		for _, user := range usersWithAccessOfSpecificProjEnvAndAppForGettingAppList {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			allAppsByEnvironment := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, allAppsByEnvironment.Code)
			assert.Equal(suite.T(), user.ExpectedAppCount, allAppsByEnvironment.Result.AppCount)
			assert.Equal(suite.T(), user.ExpectedAppName, allAppsByEnvironment.Result.AppContainers[0].AppName)
			assert.Equal(suite.T(), user.ExpectedEnvironmentName, allAppsByEnvironment.Result.AppContainers[0].Environments[0].EnvironmentName)
			assert.Equal(suite.T(), user.ExpectedTeamName, allAppsByEnvironment.Result.AppContainers[0].Environments[0].TeamName)
		}
	})

	suite.Run("A=4=HitApiGetAppsListWithUsersAccessOfSpecificProjAllEnvAndApp", func() {
		PayloadForApiFetchAppsByEnvironment := AppListingRouter.GetPayloadForApiFetchAppsByEnvironment()
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(PayloadForApiFetchAppsByEnvironment)
		for _, user := range usersWithAccessOfSpecificProjAllEnvAndAllAppForGettingAppList {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			allAppsByEnvironment := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, allAppsByEnvironment.Code)
			assert.Equal(suite.T(), user.ExpectedAppCount, allAppsByEnvironment.Result.AppCount)
			assert.Equal(suite.T(), user.ExpectedAppName, allAppsByEnvironment.Result.AppContainers[0].AppName)
			assert.Equal(suite.T(), user.ExpectedEnvironmentName, allAppsByEnvironment.Result.AppContainers[0].Environments[0].EnvironmentName)
			assert.Equal(suite.T(), user.ExpectedTeamName, allAppsByEnvironment.Result.AppContainers[0].Environments[0].TeamName)
		}
	})

	suite.Run("A=5=HitApiWithUnauthorisedUsersToCreateApp", func() {
		for _, user := range UnauthorisedUsersToCreateApp {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			appName := "app" + strings.ToLower(testUtils.GetRandomStringOfGivenLength(5))
			createAppRequestDto := PipelineConfigRouter.GetAppRequestDto(appName, 1, 0)
			byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
			createAppResponseDto := PipelineConfigRouter.HitCreateAppApi(byteValueOfCreateApp, appName, 1, 0, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, createAppResponseDto.Code)
		}
	})

	//todo PipelineId is hardcoded for RBAC story because of time crunch, need to make this dynamic
	suite.Run("A=6=HitApiWithAuthorisedUsersToTriggerCiPipeline", func() {
		CiPipelineId := 7
		pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(CiPipelineId, suite.authToken)
		payloadForTriggerCiPipeline := PipelineConfigRouter.CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, CiPipelineId, pipelineMaterial.Result[0].Id, true)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		for _, user := range AuthorisedUsersToGetCdPipelineStrategies {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			triggerCiPipelineResponse := PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, triggerCiPipelineResponse.Code)
		}
	})

	//todo PipelineId is hardcoded for RBAC story because of time crunch, need to make this dynamic
	suite.Run("A=7=HitApiWithAuthorisedUsersToTriggerCiPipelines", func() {
		CiPipelineId := 7
		pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(CiPipelineId, suite.authToken)
		payloadForTriggerCiPipeline := PipelineConfigRouter.CreatePayloadForTriggerCiPipeline(pipelineMaterial.Result[0].History[0].Commit, CiPipelineId, pipelineMaterial.Result[0].Id, true)
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(payloadForTriggerCiPipeline)
		for _, user := range AuthorisedUsersToTriggerCiPipelines {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			triggerCiPipelineResponse := PipelineConfigRouter.HitTriggerCiPipelineApi(bytePayloadForTriggerCiPipeline, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, triggerCiPipelineResponse.Code)
		}
	})

	//todo hard coded the AppId in line135 because of time crunch during testing of RBAC feature
	suite.Run("A=8=HitApiWithAuthorisedUsersToGetCdPipelineStrategies", func() {
		time.Sleep(2 * time.Second)
		appId := strconv.Itoa(75)
		for _, user := range AuthorisedUsersToGetCdPipelineStrategies {
			log.Println("Test Case for User ===>", user.ApiTokenName)
			cdPipelineStrategiesResponse := PipelineConfigRouter.HitGetCdPipelineStrategies(appId, user.ApiToken)
			assert.Equal(suite.T(), user.ExpectedResponseCode, cdPipelineStrategiesResponse.Code)
		}
	})

	//todo hard coding of SuperAdmin Access
	suite.Run("A=9=HitApiGetAppsListWithSuperAdminUsersAccess", func() {
		superAdminToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpzdXBlci1hZG1pbi10b2tlbiIsImlzcyI6ImFwaVRva2VuSXNzdWVyIiwiZXhwIjoxNjc5MzE5MjcxfQ.STZ4ptJkaCNral6vHMIA0zL_p8kC3NlcKX69unEfiwg"
		PayloadForApiFetchAppsByEnvironment := AppListingRouter.GetPayloadForApiFetchAppsByEnvironment()
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(PayloadForApiFetchAppsByEnvironment)
		allAppsViaArgoAdminToken := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, superAdminToken)
		allAppsViaSuperAdminToken := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, suite.authToken)
		assert.Equal(suite.T(), allAppsViaArgoAdminToken.Result.AppCount, allAppsViaSuperAdminToken.Result.AppCount)
	})

	//todo hard coding of no-access token for testing as of now
	suite.Run("A=10=HitApiGetAppsListWithSuperAdminUsersAccess", func() {
		noAccessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpuby1hY2Nlc3MiLCJpc3MiOiJhcGlUb2tlbklzc3VlciJ9.czSJfHMgIYzeXj6oTEGKm0jDqk2wUtWOwBuOGerrK28"
		PayloadForApiFetchAppsByEnvironment := AppListingRouter.GetPayloadForApiFetchAppsByEnvironment()
		bytePayloadForTriggerCiPipeline, _ := json.Marshal(PayloadForApiFetchAppsByEnvironment)
		allAppsViaNoAccessToken := AppListingRouter.HitApiFetchAppsByEnvironment(bytePayloadForTriggerCiPipeline, noAccessToken)
		assert.Equal(suite.T(), 200, allAppsViaNoAccessToken.Code)
		assert.Equal(suite.T(), 0, allAppsViaNoAccessToken.Result.AppCount)
	})

	//todo need to take care of teamId dynamically, as of now hard coded the ID because of time crunch in RBAC Feature testing
	suite.Run("A=11=HitApiWithAdminUsersToCreateAppInSpecificProject", func() {
		token_specific_proj_all_env_all_app_devtron_admin := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpwZWNpZmljX3Byb2pfYWxsX2Vudl9hbGxfYXBwX2RldnRyb25fYWRtaW4iLCJpc3MiOiJhcGlUb2tlbklzc3VlciJ9.qgS4OfcUHXRUpJwC_gpzOLAL1TwMplsR-ggZw_76Seo"
		appName := strings.ToLower(testUtils.GetRandomStringOfGivenLength(8))
		createAppRequestDto := PipelineConfigRouter.GetAppRequestDto(appName, 4, 0)
		byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
		createAppResponseDto := PipelineConfigRouter.HitCreateAppApi(byteValueOfCreateApp, appName, 4, 0, token_specific_proj_all_env_all_app_devtron_admin)
		assert.Equal(suite.T(), 200, createAppResponseDto.Code)
		testUtils.DeleteApp(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId, token_specific_proj_all_env_all_app_devtron_admin)
	})
	//todo need to take care of teamId dynamically, as of now hard coded the ID because of time crunch in RBAC Feature testing
	suite.Run("A=11=HitApiWithManagerUsersToCreateAppInSpecificProject", func() {
		Token_Specific_proj_all_env_all_app_devtron_manager := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkFQSS1UT0tFTjpwZWNpZmljX3Byb2pfYWxsX2Vudl9hbGxfYXBwX2RldnRyb25fbWFuYWdlciIsImlzcyI6ImFwaVRva2VuSXNzdWVyIn0.VD5qeLYYKb1C4_xHZxSTeLpnBsf3Sw26IZjNtu6KqOA"
		appName := "app" + strings.ToLower(testUtils.GetRandomStringOfGivenLength(5))
		createAppRequestDto := PipelineConfigRouter.GetAppRequestDto(appName, 4, 0)
		byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)
		createAppResponseDto := PipelineConfigRouter.HitCreateAppApi(byteValueOfCreateApp, appName, 4, 0, Token_Specific_proj_all_env_all_app_devtron_manager)
		assert.Equal(suite.T(), 200, createAppResponseDto.Code)
		testUtils.DeleteApp(createAppResponseDto.Result.Id, createAppResponseDto.Result.AppName, createAppResponseDto.Result.TeamId, createAppResponseDto.Result.TemplateId, Token_Specific_proj_all_env_all_app_devtron_manager)
	})
}
