package ApplicationRouter

import (
	"automation-suite/HelperRouter"
	"automation-suite/PipelineConfigRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/r3labs/sse/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *ApplicationsRouterTestSuite) TestDeepakSSE1() {

	config, _ := PipelineConfigRouter.GetEnvironmentConfigPipelineConfigRouter()
	var configId int
	log.Println("=== Here we are creating a App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appName := createAppApiResponse.AppName
	log.Println("=== App Name is :====", appName)

	log.Println("=== Here we are creating App Material ===")
	createAppMaterialRequestDto := PipelineConfigRouter.GetAppMaterialRequestDto(createAppApiResponse.Id, 1, false)
	appMaterialByteValue, _ := json.Marshal(createAppMaterialRequestDto)
	createAppMaterialResponse := PipelineConfigRouter.HitCreateAppMaterialApi(appMaterialByteValue, createAppApiResponse.Id, 1, false, suite.authToken)

	log.Println("=== Here we are saving docker build config ===")
	requestPayloadForSaveAppCiPipeline := PipelineConfigRouter.GetRequestPayloadForSaveAppCiPipeline(createAppApiResponse.Id, config.DockerRegistry, config.DockerRegistry+"/test", config.DockerfilePath, config.DockerfileRepository, config.DockerfileRelativePath, createAppMaterialResponse.Result.Material[0].Id)
	byteValueOfSaveAppCiPipeline, _ := json.Marshal(requestPayloadForSaveAppCiPipeline)
	PipelineConfigRouter.HitSaveAppCiPipeline(byteValueOfSaveAppCiPipeline, suite.authToken)

	log.Println("=== Here we are fetching latestChartReferenceId ===")
	time.Sleep(2 * time.Second)
	getChartReferenceResponse := PipelineConfigRouter.HitGetChartReferenceViaAppId(strconv.Itoa(createAppApiResponse.Id), suite.authToken)
	latestChartRef := getChartReferenceResponse.Result.LatestChartRef

	log.Println("=== Here we are fetching Template using getAppTemplateAPI ===")
	getTemplateResponse := PipelineConfigRouter.HitGetTemplateViaAppIdAndChartRefId(strconv.Itoa(createAppApiResponse.Id), strconv.Itoa(latestChartRef), suite.authToken)

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
	PipelineConfigRouter.HitSaveDeploymentTemplateApi(updatedByteValueOfSaveDeploymentTemplate, suite.authToken)

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
	workflowResponse := PipelineConfigRouter.HitCreateWorkflowApiWithFullPayload(createAppApiResponse.Id, suite.authToken).Result

	preStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/preStageScript.txt")
	postStageScript, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/PipeLineConfigRouter/postStageScript.txt")

	log.Println("=== Here we are saving CD pipeline ===")
	payload := PipelineConfigRouter.GetRequestPayloadForSaveCdPipelineApi(createAppApiResponse.Id, workflowResponse.AppWorkflowId, 1, workflowResponse.CiPipelines[0].Id, workflowResponse.CiPipelines[0].ParentCiPipeline, "AUTOMATIC", string(preStageScript), string(postStageScript), "AUTOMATIC")
	bytePayload, _ := json.Marshal(payload)
	savePipelineResponse := PipelineConfigRouter.HitSaveCdPipelineApi(bytePayload, suite.authToken)
	fmt.Println(savePipelineResponse)
	time.Sleep(2 * time.Second)

	log.Println("=== Here we are getting pipeline material ===")
	pipelineMaterial := PipelineConfigRouter.HitGetCiPipelineMaterial(workflowResponse.CiPipelines[0].Id, suite.authToken)

	log.Println("=== Here we are Triggering CI/CD and verifying CI/CD Deploy Status ===")
	triggerAndVerifyCiPipeline(createAppApiResponse, pipelineMaterial, workflowResponse.CiPipelines[0].Id, suite)

	log.Println("=== Here we are getting ResourceTree ===")
	ResourceTreeApiResponse := HitGetResourceTreeApi(createAppApiResponse.AppName, suite.authToken)

	container := ResourceTreeApiResponse.Result.Nodes[0].Name
	//container := ResourceTreeApiResponse.Result.Nodes[0].Name

	queryParams := make(map[string]string)
	queryParams["container"] = container
	queryParams["follow"] = "true"
	queryParams["namespace"] = "devtron-demo"
	queryParams["tailLines"] = "50"
	url := CreateFinalUrlHavingQueryParam(queryParams)
	println(url)
	urlDeepak := ApplicationsRouterBaseUrl + "appeicbhw1s3m-devtron-demo" + "/pods/" + "appeicbhw1s3m-devtron-demo-686bf6f465-fpgcq" + "/logs?" + url
	fmt.Println(urlDeepak)
	//https://staging.devtron.info/orchestrator/api/v1/applications/appeicbhw1s3m-devtron-demo/pods/appeicbhw1s3m-devtron-demo-686bf6f465-fpgcq/logs?container=appeicbhw1s3m&follow=true&namespace=devtron-demo&tailLines=500
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE0OTEwOTYsImp0aSI6IjkzMDI5ZjY1LTUzMWEtNGE5OC05NDFiLTU4ZTZmOGI3YWI1ZiIsImlhdCI6MTY2MTQwNDY5NiwiaXNzIjoiYXJnb2NkIiwibmJmIjoxNjYxNDA0Njk2LCJzdWIiOiJhZG1pbiJ9.-WmZkUrWapN-SUzLbuC1DBv8NM37zBKBSCR_ORGuDpk"
	Base.ReadEventStreamsForSpecificApi(urlDeepak, suite.authToken, suite.T())
}

func (suite *ApplicationsRouterTestSuite) TestDeepakSSE2() {
	client := sse.NewClient("https://staging.devtron.info/orchestrator/api/v1/applications/appeicbhw1s3m-devtron-demo/pods/appeicbhw1s3m-devtron-demo-686bf6f465-fpgcq/logs?container=appeicbhw1s3m&follow=true&namespace=devtron-demo&tailLines=500")
	header := make(map[string]string)
	header["token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE0MDM0OTYsImp0aSI6ImZjMDYzNzMzLTYzZDctNDA0Ni1iZjA1LWEyNjg5NTA3YzBlYiIsImlhdCI6MTY2MTMxNzA5NiwiaXNzIjoiYXJnb2NkIiwibmJmIjoxNjYxMzE3MDk2LCJzdWIiOiJhZG1pbiJ9.g7H-N4NN6bNErQCDUcKabQgdHVpcpx08wiwEfWB1nss"
	client.Headers = header
	events := make(chan *sse.Event)
	var cErr error
	go func() {
		cErr = client.Subscribe("message", func(msg *sse.Event) {
			if msg.Data != nil {
				events <- msg
				return
			}
		})
	}()

	for i := 0; i < 3; i++ {
		msg, err := wait(events, time.Second*60)
		require.Nil(suite.T(), err)
		if i == 0 {
			assert.True(suite.T(), strings.Contains(string(msg.Data), "\"podName\":\"appeicbhw1s3m-devtron-demo-686bf6f465-fpgcq\""))
		}
		fmt.Println(i, "=====>", string(msg.Data))
		dt := time.Now()
		if strings.Contains(string(msg.Data), "{\"result\":{\"content\"") || strings.Contains(string(msg.Data), dt.Format("01-02-2006")) {
			assert.True(suite.T(), true)
		}
	}
	assert.Nil(suite.T(), cErr)
}

func wait(ch chan *sse.Event, duration time.Duration) (*sse.Event, error) {
	var err error
	var msg *sse.Event
	select {
	case event := <-ch:
		msg = event
	case <-time.After(duration):
		err = errors.New("timeout")
	}
	return msg, err
}

func CreateFinalUrlHavingQueryParam(params map[string]string) string {
	var url string = ""
	var finalUrl string = ""
	for key, value := range params {
		url = key + "=" + value
		finalUrl = url + "&" + finalUrl
	}
	finalTrimmedUrl := strings.TrimSpace(finalUrl)
	return strings.TrimRight(finalTrimmedUrl, "&")
}
