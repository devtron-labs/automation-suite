package testUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-resty/resty/v2"
	"github.com/r3labs/sse/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const createSessionApiUrl string = "/orchestrator/api/v1/session"

type Errors struct {
	Code            string `json:"code"`
	InternalMessage string `json:"internalMessage"`
	UserMessage     string `json:"userMessage"`
}
type LogInResult struct {
	Token string `json:"token"`
}

type LogInResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Result LogInResult `json:"result"`
}

type ApiErrorDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
}

type CreateAppResponseDto struct {
	Code   int                 `json:"code"`
	Status string              `json:"status"`
	Result CreateAppRequestDto `json:"result"`
	Errors []Errors            `json:"errors"`
}
type CreateAppRequestDto struct {
	Id         int    `json:"id"`
	AppName    string `json:"appName"`
	TeamId     int    `json:"teamId"`
	TemplateId int    `json:"templateId"`
}

type DeleteResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type BaseClassEnvironmentConfig struct {
	BaseServerUrl          string `json:"BASE_SERVER_URL"`
	LogInUserName          string `json:"LOGIN_USERNAME"`
	LogInUserPwd           string `json:"LOGIN_PASSWORD"`
	SSOClientSecret        string `json:"CLIENT_SECRET"`
	Provider               string `json:"PROVIDER"`
	GitUsername            string `json:"GIT_USERNAME"`
	Host                   string `json:"HOST"`
	GitToken               string `json:"GIT_TOKEN"`
	GitHubOrgId            string `json:"GITHUB_ORG_ID"`
	PluginId               string `json:"PLUGIN_ID"`
	RegistryType           string `json:"REGISTRY_TYPE"`
	RegistryUrl            string `json:"REGISTRY_URL"`
	DockerUsername         string `json:"DOCKER_USERNAME"`
	Password               string `json:"PASSWORD"`
	ClusterBearerToken     string `json:"CLUSTER_BEARER_TOKEN"`
	ClusterServerUrl       string `json:"CLUSTER_SERVER_URL"`
	BearerToken            string `json:"BEARER_TOKEN"`
	GitHubProjectUrl       string `json:"GITHUB_URL_TO_CLONE_PROJECT"`
	DockerRegistry         string `json:"DOCKER_REGISTRY" `
	DockerfilePath         string `json:"DOCKER_FILE_PATH" `
	DockerfileRepository   string `json:"DOCKER_FILE_REPO" `
	DockerfileRelativePath string `json:"DOCKER_FILE_RELATIVE_PATH"`
}

func getRestyClient() *resty.Client {
	baseConfig := ReadBaseEnvConfig()
	fileData := ReadAnyJsonFile(baseConfig.BaseCredentialsFile)
	client := resty.New()
	client.SetBaseURL(fileData.BaseServerUrl)
	return client
}

// MakeApiCall make the api call to the requested url based on http method requested
func MakeApiCall(apiUrl string, method string, body string, queryParams map[string]string, authToken string) (*resty.Response, error) {
	var resp *resty.Response
	var err error
	switch method {
	case "GET":
		if queryParams != nil {
			return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().SetQueryParams(queryParams).Get(apiUrl)
		}
		return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().Get(apiUrl)
	case "POST":
		return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().SetBody(body).Post(apiUrl)
	case "PUT":
		return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().SetBody(body).Put(apiUrl)
	case "DELETE":
		return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().SetBody(body).Delete(apiUrl)
	}
	return resp, err
}

// HandleError Log the error and return boolean value indicating whether error occurred or not
func HandleError(err error, testName string) {
	if nil != err {
		log.Println("Error occurred while invoking api for test:"+testName, "err", err)
	}
}

func GetByteArrayOfGivenJsonFile(filePath string) ([]byte, error) {
	testDataJsonFile, err := os.Open(filePath)
	if nil != err {
		log.Println("Unable to open the file. Error occurred !!", "err", err)
	}
	log.Println("Opened the given json file successfully !!!")
	//defer testDataJsonFile.Close()

	byteValue, err := ioutil.ReadAll(testDataJsonFile)
	return byteValue, err
}

// GetAuthToken support function to return auth token after log in
func GetAuthToken() string {
	envConf := ReadBaseEnvConfig()
	file := ReadAnyJsonFile(envConf.BaseCredentialsFile)
	jsonString := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, file.LogInUserName, file.LogInUserPwd)
	resp, err := MakeApiCall(createSessionApiUrl, http.MethodPost, jsonString, nil, "")
	HandleError(err, "getAuthToken")
	var logInResponse LogInResponse
	json.Unmarshal(resp.Body(), &logInResponse)
	return logInResponse.Result.Token
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charset2 = "abcdefghijklmnopqrstuvwxyz" +
	"0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GetRandomStringOfGivenLength(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandomStringOfGivenLengthOfLowerCaseAndNumber(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset2[seededRand.Intn(len(charset2))]
	}
	return string(b)
}
func GetRandomNumberOf9Digit() int {
	return 100000000 + rand.Intn(999999999-100000000)
}

func CreateApp(authToken string) CreateAppResponseDto {
	appName := strings.ToLower(GetRandomStringOfGivenLength(10))
	createAppRequestDto := GetAppRequestDto("app"+appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)

	response, err := MakeApiCall("/orchestrator/app", http.MethodPost, string(byteValueOfCreateApp), nil, authToken)
	HandleError(err, "CreateAppApi")
	baseConfigRouter := BaseConfigRouter{}
	pipelineConfigRouter := baseConfigRouter.UnmarshalGivenResponseBody(response.Body(), "SaveConfigmapApi")
	return pipelineConfigRouter.createAppResponseDto
}

func GetAppRequestDto(appName string, teamId int, templateId int) CreateAppRequestDto {
	var createAppRequestDto CreateAppRequestDto
	createAppRequestDto.AppName = appName
	createAppRequestDto.TeamId = teamId
	createAppRequestDto.TemplateId = templateId
	return createAppRequestDto
}

func (baseConfigRouter BaseConfigRouter) UnmarshalGivenResponseBody(response []byte, apiName string) BaseConfigRouter {
	switch apiName {
	case "SaveConfigmapApi":
		json.Unmarshal(response, &baseConfigRouter.createAppResponseDto)
	}
	return baseConfigRouter
}

type BaseConfigRouter struct {
	createAppResponseDto CreateAppResponseDto
	deleteResponseDto    DeleteResponseDto
}

func GetPayLoadForDeleteAppAPI(id int, appName string, teamId int, templateId int) []byte {
	var createAppRequestDto CreateAppRequestDto
	createAppRequestDto.Id = id
	createAppRequestDto.AppName = appName
	createAppRequestDto.TeamId = teamId
	createAppRequestDto.TemplateId = templateId
	byteValueOfStruct, _ := json.Marshal(createAppRequestDto)
	return byteValueOfStruct
}

func DeleteApp(appId int, appName string, TeamId int, TemplateId int, authToken string) DeleteResponseDto {
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(appId, appName, TeamId, TemplateId)
	resp, err := MakeApiCall("/orchestrator/app/"+strconv.Itoa(appId), http.MethodDelete, string(byteValueOfDeleteApp), nil, authToken)
	HandleError(err, "DeleteAppApi")
	baseConfigRouter := BaseConfigRouter{}
	apiRouter := baseConfigRouter.UnmarshalGivenResponseBody(resp.Body(), "DeleteAppApi")
	return apiRouter.deleteResponseDto
}

func ReadAnyJsonFile(filename string) BaseClassEnvironmentConfig {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	data := BaseClassEnvironmentConfig{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

type BaseEnvConfigStruct struct {
	BaseCredentialsFile  string `env:"BASE_CREDENTIALS_FILE" envDefault:"/base-test/credentials.json"`
	ClassCredentialsFile string `env:"CLASS_CREDENTIALS_FILE" envDefault:"/class-test/credentials.json"`
}

func ReadBaseEnvConfig() *BaseEnvConfigStruct {
	cfg := &BaseEnvConfigStruct{}
	err := env.Parse(cfg)
	if err != nil {
		return nil
	}
	return cfg
}

func ReadEventStreamsForSpecificApi(apiUrl string, authToken string, ContainerName string, t *testing.T) {
	baseConfig := ReadBaseEnvConfig()
	fileData := ReadAnyJsonFile(baseConfig.BaseCredentialsFile)
	url := fileData.BaseServerUrl + apiUrl
	client := sse.NewClient(url)
	header := make(map[string]string)
	header["token"] = authToken
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
		require.Nil(t, err)
		if i == 0 {
			assert.True(t, strings.Contains(string(msg.Data), ContainerName))
		}
		fmt.Println(i, "=====>", string(msg.Data))
		dt := time.Now()
		if strings.Contains(string(msg.Data), "{\"result\":{\"content\"") || strings.Contains(string(msg.Data), dt.Format("01-02-2006")) {
			assert.True(t, true)
		}
	}
	assert.Nil(t, cErr)
}

func ReadEventStreamsForSpecificApiAndVerifyResult(apiUrl string, authToken string, t *testing.T, indexOfMessage int, message string) {
	baseConfig := ReadBaseEnvConfig()
	fileData := ReadAnyJsonFile(baseConfig.BaseCredentialsFile)
	url := fileData.BaseServerUrl + apiUrl
	client := sse.NewClient(url)
	header := make(map[string]string)
	header["token"] = authToken
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

	for i := 0; i <= indexOfMessage; i++ {
		msg, err := wait(events, time.Second*60)
		require.Nil(t, err)
		fmt.Println(i, "=====>", string(msg.Data))
		if i == indexOfMessage {
			assert.Equal(t, string(msg.Data), message)
		}
	}
	assert.Nil(t, cErr)
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

func CreateUrlForEventStreamsHavingQueryParam(params map[string]string) string {
	var url string = ""
	var finalUrl string = ""
	for key, value := range params {
		url = key + "=" + value
		finalUrl = url + "&" + finalUrl
	}
	finalTrimmedUrl := strings.TrimSpace(finalUrl)
	return strings.TrimRight(finalTrimmedUrl, "&")
}

func ConvertDateStringIntoTimeStamp(timeString string) int64 {
	dateTime, e := time.Parse(time.RFC3339, timeString)
	if e != nil {
		panic("Parse error")
	}
	timestamp := dateTime.Unix()
	fmt.Println("Date to Timestamp : ", timestamp)
	return timestamp
}

func ConvertYamlIntoJson(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = ConvertYamlIntoJson(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = ConvertYamlIntoJson(v)
		}
	}
	return i
}
