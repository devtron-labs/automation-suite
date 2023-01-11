package PipelineConfigRouter

//todo first I have to finish this as this is incomplete
/*func (suite *PipelinesConfigRouterTestSuite) TestGetAppDeploymentStatusTimeline() {

	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	var configId int
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result

	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	getChartReferenceResponse := HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching DefaultAppOverride using getAppTemplateAPI ===")
	getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), suite.authToken)

	log.Println("=== Here we are fetching DefaultAppOverride using getAppTemplateAPI ===")
	defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

	log.Println("=== Here we are creating payload for SaveTemplate API ===")
	saveDeploymentTemplate := GetRequestPayloadForSaveDeploymentTemplate(createAppApiResponse.Id, latestChartRef, defaultAppOverride)
	byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)

	log.Println("=== Here we are hitting SaveTemplate API ===")
	HitSaveDeploymentTemplateApi(byteValueOfSaveDeploymentTemplate, suite.authToken)

	log.Println("=== Here we are saving Global Configmap ===")

	requestPayloadForConfigMap := HelperRouter.GetRequestPayloadForSecretOrConfig(0, "-config1", createAppApiResponse.Id, "environment", "kubernetes", false, false, false, false)
	byteValueOfSaverConfigMap, _ := json.Marshal(requestPayloadForConfigMap)
	globalConfigMap := HelperRouter.HitSaveGlobalConfigMap(byteValueOfSaverConfigMap, suite.authToken)
	configId = globalConfigMap.Result.Id

	log.Println("=== Here we are saving Global Secret ===")
	requestPayloadForSecret := HelperRouter.GetRequestPayloadForSecretOrConfig(configId, "-secret1", createAppApiResponse.Id, "environment", "kubernetes", false, false, true, false)
	byteValueOfSecret, _ := json.Marshal(requestPayloadForSecret)
	HelperRouter.HitSaveGlobalSecretApi(byteValueOfSecret, suite.authToken)

	log.Println("=== Here we are saving workflow with Pre/Post CI ===")

	workflowResponse := HitCreateWorkflowApiWithFullPayload(createAppApiResponse.Id, suite.authToken).Result

	preStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
	postStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")

	log.Println("=== Here we are saving CD pipeline ===")
	payload := GetRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, Automatic, string(preStageScript), string(postStageScript), Automatic)
	bytePayload, _ := json.Marshal(payload)
	savePipelineResponse := HitSaveCdPipelineApi(bytePayload, suite.authToken)

	//write the test cases here
	time.Sleep(2 * time.Second)
	suite.Run("TestDeploymentInitiation", func() {
		apiResponse := GetAppDeploymentStatusTimeline(createAppApiResponse.Id, 1, suite.authToken)
		assert.NotEqual(suite.T(), nil, apiResponse)
		assert.Equal(suite.T(), 200, apiResponse.Code)
		assert.Equal(suite.T(), 0, len(apiResponse.Error))
		assert.NotEqual(suite.T(), nil, apiResponse.Result)
		isDeploymentStarted := len(apiResponse.Result.Timelines) > 1
		assert.Equal(suite.T(), true, isDeploymentStarted)
		assert.Equal(suite.T(), TIMELINE_STATUS_DEPLOYMENT_INITIATED, apiResponse.Result.Timelines[0])
		assert.Equal(suite.T(), TIMELINE_STATUS_GIT_COMMIT, apiResponse.Result.Timelines[1])
	})
	time.Sleep(2 * time.Second)
	suite.Run("TestGitCommitSuccessAndKubectlApply", func() {
		apiResponse := GetAppDeploymentStatusTimeline(createAppApiResponse.Id, 1, suite.authToken)
		assert.NotEqual(suite.T(), nil, apiResponse)
		assert.Equal(suite.T(), 200, apiResponse.Code)
		assert.Equal(suite.T(), 0, len(apiResponse.Error))
		assert.NotEqual(suite.T(), nil, apiResponse.Result)
		isDeploymentStarted := len(apiResponse.Result.Timelines) > 2
		assert.Equal(suite.T(), true, isDeploymentStarted)
		assert.Equal(suite.T(), TIMELINE_STATUS_DEPLOYMENT_INITIATED, apiResponse.Result.Timelines[0].Status)
		assert.Equal(suite.T(), TIMELINE_STATUS_GIT_COMMIT, apiResponse.Result.Timelines[1].Status)
		kubectlStatus := apiResponse.Result.Timelines[1].Status == TIMELINE_STATUS_KUBECTL_APPLY_STARTED || apiResponse.Result.Timelines[1].Status == TIMELINE_STATUS_KUBECTL_APPLY_SYNCED
		assert.Equal(suite.T(), true, kubectlStatus)
		if apiResponse.Result.Timelines[1].Status == TIMELINE_STATUS_KUBECTL_APPLY_SYNCED {
			isAtleastOneK8sObjectPresent := len(apiResponse.Result.Timelines[1].ResourceDetails) > 0
			assert.Equal(suite.T(), true, isAtleastOneK8sObjectPresent)
			for _, resource := range apiResponse.Result.Timelines[1].ResourceDetails {
				assert.NotNil(suite.T(), resource.Id)
				assert.NotNil(suite.T(), resource.ResourceStatus)
				assert.NotNil(suite.T(), resource.ResourceKind)
				assert.NotNil(suite.T(), resource.ResourceGroup)
				assert.NotNil(suite.T(), resource.ResourceName)
				assert.NotNil(suite.T(), resource.ResourcePhase)
			}
		}
	})

	//write tests with invalid git-ops configuration
	//end of test cases

	//clean created pipelines,materials and app
	log.Println("=== Here we are Deleting the CD pipeline ===")
	deletePipelinePayload := GetPayloadForDeleteCdPipeline(createAppApiResponse.Id, savePipelineResponse.Result.Pipelines[0].Id)
	deletePipelineByteCode, _ := json.Marshal(deletePipelinePayload)
	HitForceDeleteCdPipelineApi(deletePipelineByteCode, suite.authToken)

	log.Println("=== Here we are Deleting the CI pipeline ===")
	DeleteCiPipeline(createAppApiResponse.Id, workflowResponse.CiPipelines[0].Id, suite.authToken)
	log.Println("=== Here we are Deleting CI Workflow ===")
	HitDeleteWorkflowApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, suite.authToken)
	log.Println("=== Here we Deleting the Test data created after verification ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
*/
