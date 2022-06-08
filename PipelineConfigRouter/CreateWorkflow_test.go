package PipelineConfigRouter

import (
	"automation-suite/testUtils"
	Base "automation-suite/testUtils"
	"encoding/json"
	"log"
	"math/rand"
	"strings"

	"github.com/stretchr/testify/assert"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassC7CreateWorkflowBranchFixedWithoutBuilds() {
	appId := suite.createAppResponseDto.Result.Id
	config, _ := GetEnvironmentConfigPipelineConfigRouter()
	createAppApiResponse := suite.createAppResponseDto.Result
	createAppMaterialResponse := suite.createAppMaterialResponseDto.Result

	requestPayloadForSaveAppCiPipeline := GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	log.Println("=== Hitting the SaveAppCiPipeline API ====")
	HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("Fetching suggested ci pipeline name ")
	fetchSuggestedCiPipelineName := HitGetPipelineSuggestedCiCd("ci", appId, suite.authToken)
	log.Println("Fetching gitMaterialId ")
	fetchAppGetResponseDto := HitGetMaterial(appId, suite.authToken)

	log.Println("Retrieving request payload from file")
	createWorkflowRequestDto := getRequestPayloadForCreateWorkflow(false, "1", appId, 0)

	createWorkflowRequestDto.CiPipeline.CiMaterial[0].GitMaterialId = fetchAppGetResponseDto.Result.Material[0].Id
	createWorkflowRequestDto.CiPipeline.Name = fetchSuggestedCiPipelineName.Result
	createWorkflowRequestDto.CiPipeline.CiMaterial[0].Source.Value = strings.ToLower(testUtils.GetRandomStringOfGivenLength(10))

	suite.Run("A=1=CreateWorkflowBranchFixedWithoutBuilds", func() {

		byteValueOfCreateWorkflow, _ := json.Marshal(createWorkflowRequestDto)
		log.Println("Hitting the Create Workflow Api with valid payload")
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)
		log.Println("Validating the Create Workflow Api response with with valid payload")
		assert.Equal(suite.T(), createWorkflowRequestDto.AppId, createWorkflowResponseDto.Result.AppId)
		log.Println("Hitting get workflow details api")

		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)

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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i <= numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		// after creation call /orchestrator/app/ci-pipeline/app-id/wf-id get for delete ci-pipeline
		log.Println("Hitting get workflow details api")

		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)

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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.ScriptType)
			for j := 0; j < 3; j++ {
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format)
			}
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format)
		}
		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i <= numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}
		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ScriptType)
		}

		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].Name, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].Name)
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.ScriptType, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.ScriptType)
			for j := 0; j < 3; j++ {
				assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.InputVariables[j].Format)
			}
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].InlineStepDetail.InputVariables[4].Format)
		}

		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

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

		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
	})

	/////////////////=== post-build check with random task with variable conditions====//////////////

	suite.Run("A=11=CreateWorkflowBranchFixedPostBuildWithVariableConditions", func() {

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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

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
		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
	})

	suite.Run("A=12=CreateWorkflowPreBuildOutoutDirectory", func() {

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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")

		// Add assert conditions here
		log.Println("Checking output directory")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PreBuildStage.Steps[i].OutputDirectoryPath, createWorkflowResponseDto.Result.CiPipelines[0].PreBuildStage.Steps[i].OutputDirectoryPath)
		}
		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
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
		createWorkflowResponseDto := HitCreateWorkflowApi(byteValueOfCreateWorkflow, suite.authToken)

		log.Println("Validating pre-build request payload")

		// Add assert conditions here
		log.Println("Checking output directory")
		for i = 0; i < numberOfTasks; i++ {
			assert.Equal(suite.T(), createWorkflowRequestDto.CiPipeline.PostBuildStage.Steps[i].OutputDirectoryPath, createWorkflowResponseDto.Result.CiPipelines[0].PostBuildStage.Steps[i].OutputDirectoryPath)
		}

		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
	})

	suite.Run("A=14=CreateWorkflowWithFullPayload", func() {
		createWorkflowResponseDto := HitCreateWorkflowApiWithFullPayload(appId, suite.authToken)

		log.Println("Validating pre-build request payload")
		assert.Equal(suite.T(), appId, createWorkflowResponseDto.Result.AppId)

		wfId := createWorkflowResponseDto.Result.AppWorkflowId

		DeleteWorkflow(appId, wfId, suite.authToken)
	})
	Base.DeleteApp(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId, suite.authToken)

}

/*
func TestCreateMaterialWithMultipleMaterial(t *testing.T) {
	appId := 1205
	authToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImQwNzg1NmMyMjA5YTlmYzk2ZTIzNDBhYzZlYmZhODMxYjUwZGFjZGIifQ.eyJpc3MiOiJodHRwczovL3N0YWdpbmcuZGV2dHJvbi5pbmZvL29yY2hlc3RyYXRvci9hcGkvZGV4Iiwic3ViIjoiQ2hVeE1ESTJPVEk1TmpBd056RXhNekU0TVRFMU5qY1NCbWR2YjJkc1pRIiwiYXVkIjoiYXJnby1jZCIsImV4cCI6MTY1NDc0NjcyMiwiaWF0IjoxNjU0NjYwMzIyLCJhdF9oYXNoIjoiU3h2TlEwY0ZZZUR5YklkQWtuNTE1USIsImVtYWlsIjoibmlrZXNoQGRldnRyb24uYWkiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6Ik5pa2VzaCBSYXRob2QifQ.O4klGQkFzO8tB8ZtChFMi_EBq7GQyzqXuWAq1QeWk8u3EeFd8euQzF1TDExkDPrvKF3W_5-5sMH-hTiN0s7J1rje2GsZZPEY148CL0MSdkRgiqG-zV1ZFXbLcemxu9cnc22XQNvsC7NtKGbUuJaNEuyVlVIC-9HJoiLuexR12aGJEs3CWb2-AZkqb5Y8rzd097kuautuKuRPHTSqDzLv4B-vnpbLMLxz6cYhgALnu2XK6ansLOMS8jdNGFsIfPwml90D6D41u7G3wyzxVdL8nlrt_rCskJOJeF9UsYKJM0lq5J5Df-UBwmzEQAZI1hPnY0gBadBr48LrcBBQBZDwBA"
	createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 2, false)
	createAppMaterialRequestDto.Materials[0].CheckoutPath = "./" + Base.GetRandomStringOfGivenLength(3)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	HitCreateAppMaterialApi(appMaterialByteValue, appId, 1, false, authToken)

	createWorkflowResponseDto := HitCreateWorkflowApiWithFullPayload(appId, authToken)

	log.Println("Validating pre-build request payload")
	assert.Equal(t, appId, createWorkflowResponseDto.Result.AppId)

	wfId := createWorkflowResponseDto.Result.AppWorkflowId

	DeleteWorkflow(appId, wfId, authToken)
}
*/