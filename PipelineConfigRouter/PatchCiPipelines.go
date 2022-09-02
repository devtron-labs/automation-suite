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

	suite.Run("A=1=CreateWorkflowBranchFixedWithoutBuilds", func() {
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

	suite.Run("A=2=CreateWorkflowWithBranchFixedWithPreBuildScriptTypeShell", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(3-1) + 1
		i := 0
		for i = 0; i <= numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(1)
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload[0])
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i <= numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		// after creation call /orchestrator/app/ci-pipeline/app-id/wf-id get for delete ci-pipeline
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== pre-build check with random task with scriptType CONTAINER_IMAGE====//////////////

	suite.Run("A=3=CreateWorkflowWithBranchFixedWithPreBuildScriptTypeContainerImage", func() {

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(3-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(2)
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload[0])
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== pre-build check with random task with scriptType either SHELL or CONTAINER_IMAGE====//////////////
	suite.Run("A=4=CreateWorkflowWithBranchFixedWithPreBuildScriptTypeEitherShellOrContainerImage", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(rand.Intn(2-1) + 1)
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload[0])
		}
		// Custom part end here

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitPatchCiPipelinesApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	/////////////////=== pre-build check with random task with scriptType SHELL with Input Variables====//////////////

	suite.Run("A=5=CreateWorkflowWithBranchFixedWithPreBuildScriptTypeShellWithInputs", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(rand.Intn(2-1) + 1)
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload[0])
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

	suite.Run("A=6=CreateWorkflowWithBranchFixedWithPostBuildScriptTypeShell", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(3-1) + 1
		i := 0
		for i = 0; i <= numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(1)
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload[0])
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

	suite.Run("A=7=CreateWorkflowWithBranchFixedWithPostBuildScriptTypeContainerImage", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(3-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(2)
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload[0])
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

	/////////////////=== post-build check with random task with scriptType either SHELL or CONTAINER_IMAGE====//////////////

	suite.Run("A=8=CreateWorkflowWithBranchFixedWithPostBuildScriptTypeEitherShellOrContainerImage", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		log.Println("Creating Payload with Random Script type")
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(rand.Intn(2-1) + 1)
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload[0])
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

	suite.Run("A=9=CreateWorkflowWithBranchFixedWithPostBuildScriptTypeShellWithInputs", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		log.Println("Getting Post-build paload...")
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(rand.Intn(2-1) + 1)
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload[0])
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
	/////////////////=== pre-build check with random task with variable conditions====//////////////
	suite.Run("A=10=CreateWorkflowBranchFixedPreBuildWithVariableConditions", func() {

		// Pre-requirements end here

		// Custom part - creating random number of tasks
		log.Println("Getting Post-build paload...")
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(rand.Intn(2-1) + 1)
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload[0])
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

	suite.Run("A=11=CreateWorkflowBranchFixedPostBuildWithVariableConditions", func() {

		log.Println("Getting Post-build paload...")
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(rand.Intn(2-1) + 1)
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload[0])
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

	suite.Run("A=12=CreateWorkflowPreBuildOutoutDirectory", func() {

		log.Println("Getting Post-build paload...")
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			preBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(rand.Intn(2-1) + 1)
			createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps, preBuildStepRequestPayload[0])
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
	suite.Run("A=13=CreateWorkflowPostBuildOutoutDirectory", func() {

		log.Println("Getting Post-build paload...")
		numberOfTasks := rand.Intn(4-1) + 1
		i := 0
		for i = 0; i < numberOfTasks; i++ {
			postBuildStepRequestPayload := getPreBuildStepRequestPayloadDto(rand.Intn(2-1) + 1)
			createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps = append(createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps, postBuildStepRequestPayload[0])
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

	suite.Run("A=14=CreateWorkflowWithFullPayload", func() {
		createWorkflowResponseDto := HitCreateWorkflowApiWithFullPayload(appId, suite.authToken)
		log.Println("Validating pre-build request payload")
		assert.Equal(suite.T(), appId, createWorkflowResponseDto.Result.AppId)
		log.Println("=== Here we are Deleting the CI pipeline ===")
		DeleteCiPipeline(appId, createWorkflowResponseDto.Result.CiPipelines[0].Id, suite.authToken)
		log.Println("=== Here we are Deleting CI Workflow ===")
		HitDeleteWorkflowApi(appId, createWorkflowResponseDto.Result.AppWorkflowId, suite.authToken)
	})

	//todo disabling this test case as we are not deleting test data for this test case ,will handle the deletion of data and enable this

	/*suite.Run("A=15=CreateWorkflowWithTwoMaterials", func() {
		createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, false)

		secondBranchValue := "./repo-two"
		createAppMaterialRequestDto.Materials[0].CheckoutPath = secondBranchValue
		byteValueOfStruct2, _ := json.Marshal(createAppMaterialRequestDto)
		log.Println("Hitting The create material API")
		createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct2, appId, 1, false, suite.authToken)

		assert.Equal(suite.T(), 200, createAppMaterialResponseDto.Code)
		createWorkflowResponseDto := HitCreateWorkflowApiWithFullPayload(appId, suite.authToken)

		log.Println("Validating material request payload")
		noOfMaterials := len(createWorkflowResponseDto.Result.CiPipelines[0].CiMaterial)
		assert.Equal(suite.T(), createAppMaterialResponseDto.Result.Material[0].Id, createWorkflowResponseDto.Result.CiPipelines[0].CiMaterial[noOfMaterials-1].GitMaterialId)
		assert.Equal(suite.T(), appId, createWorkflowResponseDto.Result.AppId)

		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteCiPipeline(appId, wfId, suite.authToken)

		log.Println("getting payload for Delete material API")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(createAppMaterialResponseDto.Result.AppId, createAppMaterialResponseDto.Result.Material[0])
		log.Println("Hitting the Delete material API for Removing the data created via automation")
		HitDeleteAppMaterialApi(byteValueOfDeleteApp, suite.authToken)

	})*/
	log.Println("=== Here we are Deleting app after verification of flow ===")
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)
}
