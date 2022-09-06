package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassC7CreateWorkflowBranchFixedWithoutBuilds() {
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appId := createAppApiResponse.Id
	log.Println("=== AppName is===>", createAppApiResponse.AppName)
	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := HitCreateAppMaterialApi(appMaterialByteValue, appId, 1, false, suite.authToken).Result

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(appId, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	time.Sleep(2 * time.Second)
	getChartReferenceResponse := HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching Template using getAppTemplateAPI ===")
	getTemplateResponse := HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), suite.authToken)

	log.Println("=== Here we are fetching DefaultAppOverride from template response ===")
	defaultAppOverride := getTemplateResponse.Result.GlobalConfig.DefaultAppOverride

	log.Println("=== Here we are creating payload for SaveTemplate API ===")
	saveDeploymentTemplate := GetRequestPayloadForSaveDeploymentTemplate(createAppApiResponse.Id, latestChartRef, defaultAppOverride)
	byteValueOfSaveDeploymentTemplate, _ := json.Marshal(saveDeploymentTemplate)

	log.Println("=== Here we are hitting SaveTemplate API ===")
	HitSaveDeploymentTemplateApi(byteValueOfSaveDeploymentTemplate, suite.authToken)

	log.Println("Fetching suggested ci pipeline name ")
	fetchSuggestedCiPipelineName := HitGetPipelineSuggestedCiCd("ci", appId, suite.authToken)

	log.Println("Fetching gitMaterialId ")
	fetchAppGetResponseDto := HitGetApp(appId, suite.authToken)

	log.Println("Retrieving request payload for creating workflow from file")
	createWorkflowRequestDto := getRequestPayloadForCreateWorkflow(false, "1", appId, 0)
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].GitMaterialId = fetchAppGetResponseDto.Result.Material[0].Id
	createWorkflowRequestDto.CiPipeline.Name = fetchSuggestedCiPipelineName.Result
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].Source.Value = "main"

	suite.Run("A=1=PatchCiPipelineBranchFixedWithoutBuilds", func() {
		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)
		log.Println("Validating the Create Workflow Api response with with valid payload")
		assert.Equal(suite.T(), createWorkflowRequestDto.AppId, createWorkflowResponseDto.Result.AppId)
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== pre-build check with random task with scriptType SHELL====//////////////

	suite.Run("A=2=PatchCiPipelineWithBranchFixedWithPreBuildScriptTypeShell", func() {
		// Custom part - creating random number of tasks
		numberOfTasks := 2
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "SHELL")
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload)
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
			assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].OutputDirectoryPath[0], "/test/output")
			expectedScript := "#!/bin/sh \nset -eo pipefail \n#set -v  ## uncomment this to debug the script \n\necho \"Here I am Printing value of VarString\"\necho $VarString\n\necho \"Here I am Printing value of VarBool\"\necho $VarBool\n\necho \"Here I am Printing value of VarNumber\"\necho $VarNumber\n\necho \"Here I am Printing value of VarDate\"\necho $VarDate\n\necho \"Here I am Printing value of VarDockerImage\"\necho $VarDockerImage\n\necho \"Here I am exporting Var1\"\nexport Var1=$Var1\n\necho \"Here I am exporting Var1\"\nexport VarBool=$VarBool\n\necho \"Here I am exporting Var1\"\nexport VarNumber=$VarNumber\n\necho \"Here I am exporting Var1\"\nexport VarDate=$VarDate\n\necho \"Here I am exporting Var1\"\nexport VarDockerImage=$VarDockerImage"
			assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.Script, expectedScript)
			assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[0].Name, "VarString")
			assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.OutputVariables[2].Name, "VarNumber")
			assert.Equal(suite.T(), createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[0].Format, "STRING")
		}
		// after creation call /orchestrator/app/ci-pipeline/app-id/wf-id get for delete ci-pipeline
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== pre-build check with random task with scriptType CONTAINER_IMAGE====//////////////

	suite.Run("A=3=PatchCiPipelineWithBranchFixedWithPreBuildScriptTypeContainerImage", func() {
		log.Println("Fetching suggested ci pipeline name ")
		suggestedCiPipelineName := HitGetPipelineSuggestedCiCd("ci", appId, suite.authToken)

		log.Println("Fetching gitMaterialId ")
		appResponseDto := HitGetApp(appId, suite.authToken)

		log.Println("Retrieving request payload for creating workflow from file")
		createWorkflowRequestPayload := getRequestPayloadForCreateWorkflow(false, "1", appId, 0)
		createWorkflowRequestPayload.CiPipeline.CiMaterial[0].GitMaterialId = appResponseDto.Result.Material[0].Id
		createWorkflowRequestPayload.CiPipeline.Name = suggestedCiPipelineName.Result
		createWorkflowRequestPayload.CiPipeline.CiMaterial[0].Source.Value = "main"

		numberOfTasks := 2
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildRequestPayload := getPreBuildStepRequestPayloadDto(i, "CONTAINER_IMAGE")
			createWorkflowRequestPayload.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestPayload.CiPipeline.PreBuildStage.Steps, preBuildRequestPayload)
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestPayload)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponse := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestPayload.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponse.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestPayload.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponse.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
			assert.Equal(suite.T(), createWorkflowResponse.Result.CiPipelines[0].PreBuildStage.Steps[i].OutputDirectoryPath[0], "/test/output")
			assert.Equal(suite.T(), createWorkflowResponse.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ContainerImagePath, "alpine:latest")
			assert.Equal(suite.T(), createWorkflowResponse.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.CommandArgsMap[0].Command, "sh")
			assert.Equal(suite.T(), createWorkflowResponse.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.MountPathMap[0].FilePathOnContainer, "./")
			assert.Equal(suite.T(), createWorkflowResponse.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.PortMap[0].PortOnContainer, 8080)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponse.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponse.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== pre-build check with random task with scriptType SHELL with Input Variables====//////////////

	suite.Run("A=4=PatchCiPipelineWithBranchFixedWithPreBuildScriptTypeShellWithInputs", func() {
		numberOfTasks := 2
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "SHELL")
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload)
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
			for j := 0; j < 3; j++ {
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format)
			}
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== Post-build check with scriptType SHELL====//////////////

	suite.Run("A=5=PatchCiPipelineWithBranchFixedWithPostBuildScriptTypeShell", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(3-1) + 1
		i := 0
		for i = 0; i <= numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "SHELL")
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload)
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i <= numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})
	/////////////////=== post-build check with random task with scriptType CONTAINER_IMAGE====//////////////

	suite.Run("A=6=PatchCiPipelineWithBranchFixedWithPostBuildScriptTypeContainerImage", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(3-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "CONTAINER_IMAGE")
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload)
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== post-build check with random task with scriptType SHELL with Input Variables====//////////////

	suite.Run("A=7=PatchCiPipelineWithBranchFixedWithPostBuildScriptTypeShellWithInputs", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		log.Println("Getting Post-build payload...")
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "SHELL")
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload)
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ScriptType)
			for j := 0; j < 3; j++ {
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format)
			}
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})
	suite.Run("A=8=PatchCiPipelineBranchFixedPreBuildWithVariableConditions", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		log.Println("Getting Post-build payload...")
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "SHELL")
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload)
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
			for j := 0; j < 3; j++ {
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format)
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ConditionDetails[j].ConditionOperator, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ConditionDetails[j].ConditionOperator)
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ConditionDetails[j].ConditionalValue, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ConditionDetails[j].ConditionalValue)
			}
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== post-build check with random task with variable conditions====//////////////

	suite.Run("A=9=PatchCiPipelineFixedPostBuildWithVariableConditions", func() {

		log.Println("Getting Post-build payload...")
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "CONTAINER_IMAGE")
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload)
		}

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ScriptType)
			for j := 0; j < 3; j++ {
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format)
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ConditionDetails[j].ConditionOperator, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ConditionDetails[j].ConditionOperator)
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ConditionDetails[j].ConditionalValue, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ConditionDetails[j].ConditionalValue)
			}
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	suite.Run("A=10=PatchCiPipelinePreBuildOutputDirectory", func() {

		log.Println("Getting Post-build payload...")
		numberOfTasks := 2
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "SHELL")
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload)
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Checking output directory")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].OutputDirectoryPath, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].OutputDirectoryPath)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})
	suite.Run("A=11=PatchCiPipelinePostBuildOutputDirectory", func() {

		log.Println("Getting Post-build payload...")
		numberOfTasks := 2
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "SHELL")
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload)
		}

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")

		// Add assert conditions here
		log.Println("Checking output directory")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].OutputDirectoryPath, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].OutputDirectoryPath)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	suite.Run("A=12=PatchCiPipelineWithFullPayload", func() {
		createWorkflowResponseDto := HitCreateWorkflowApiWithFullPayload(appId, suite.authToken)
		log.Println("Validating pre-build request payload")
		assert.Equal(suite.T(), appId, createWorkflowResponseDto.Result.AppId)
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	suite.Run("A=13=CreateWorkflowWithTwoMaterials", func() {
		createAppMaterialRequestDto = GetAppMaterialRequestDto(appId, 1, false)
		secondBranchValue := "./repo-two"
		createAppMaterialRequestDto.Materials[0].CheckoutPath = secondBranchValue
		byteValueOfStruct2, _ := json.Marshal(createAppMaterialRequestDto)
		log.Println("Hitting The create material API")
		createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct2, appId, 1, false, suite.authToken)
		assert.Equal(suite.T(), 200, createAppMaterialResponseDto.Code)
		log.Println("Fetching suggested ci pipeline name ")
		suggestedCiPipelineName := HitGetPipelineSuggestedCiCd("ci", appId, suite.authToken)

		log.Println("Fetching gitMaterialId ")
		appResponseDto := HitGetApp(appId, suite.authToken)
		log.Println("Retrieving request payload for creating workflow from file")
		createWorkflowRequestPayload := getRequestPayloadForCreateWorkflow(false, "1", appId, 0)
		createWorkflowRequestPayload.CiPipeline.CiMaterial[0].GitMaterialId = appResponseDto.Result.Material[0].Id
		createWorkflowRequestPayload.CiPipeline.Name = suggestedCiPipelineName.Result
		createWorkflowRequestPayload.CiPipeline.CiMaterial[0].Source.Value = "main"

		numberOfTasks := 2
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(i, "SHELL")
			createWorkflowRequestPayload.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestPayload.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload)
		}

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestPayload)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponsePayload := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestPayload.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponsePayload.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestPayload.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponsePayload.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
			assert.Equal(suite.T(), createWorkflowResponsePayload.Result.CiPipelines[0].PreBuildStage.Steps[i].OutputDirectoryPath[0], "/test/output")
			expectedScript := "#!/bin/sh \nset -eo pipefail \n#set -v  ## uncomment this to debug the script \n\necho \"Here I am Printing value of VarString\"\necho $VarString\n\necho \"Here I am Printing value of VarBool\"\necho $VarBool\n\necho \"Here I am Printing value of VarNumber\"\necho $VarNumber\n\necho \"Here I am Printing value of VarDate\"\necho $VarDate\n\necho \"Here I am Printing value of VarDockerImage\"\necho $VarDockerImage\n\necho \"Here I am exporting Var1\"\nexport Var1=$Var1\n\necho \"Here I am exporting Var1\"\nexport VarBool=$VarBool\n\necho \"Here I am exporting Var1\"\nexport VarNumber=$VarNumber\n\necho \"Here I am exporting Var1\"\nexport VarDate=$VarDate\n\necho \"Here I am exporting Var1\"\nexport VarDockerImage=$VarDockerImage"
			assert.Equal(suite.T(), createWorkflowResponsePayload.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.Script, expectedScript)
			assert.Equal(suite.T(), createWorkflowResponsePayload.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[0].Name, "VarString")
			assert.Equal(suite.T(), createWorkflowResponsePayload.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.OutputVariables[2].Name, "VarNumber")
			assert.Equal(suite.T(), createWorkflowResponsePayload.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[0].Format, "STRING")
			assert.Equal(suite.T(), len(createWorkflowResponsePayload.Result.Materials), 2)
			assert.Equal(suite.T(), createWorkflowResponsePayload.Result.Materials[0].MaterialName, "sample-go-app")
		}
		// after creation call /orchestrator/app/ci-pipeline/app-id/wf-id get for delete ci-pipeline
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponsePayload.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponsePayload.Result.AppWorkflowId, suite.authToken)
	})
	log.Println("=== Here we are Deleting app after verification of flow ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
