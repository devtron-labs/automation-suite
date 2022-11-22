package UrlsRouter

import (
	"automation-suite/ApiTokenRouter"
	helmRouter "automation-suite/HelmAppRouter"
	userRouter "automation-suite/UserRouter"
	"automation-suite/UserRouter/RequestDTOs"
	"automation-suite/UserRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type UrlsDto struct {
	Code   int            `json:"code"`
	Status string         `json:"status"`
	Result []UrlsResponse `json:"result"`
}

type DevtronAppConfig struct {
	AppId string `env:"APP_ID" envDefault:"1"`
	EnvId string `env:"ENV_ID" envDefault:"1"`
}

type InstalledAppConfig struct {
	InstalledAppId string `env:"INSTALLED_APP_ID" envDefault:"1"`
	EnvId          string `env:"ENV_ID" envDefault:"1"`
}

type StructUrlsRouter struct {
	urlsResponse UrlsDto
}

type UrlsTestSuite struct {
	suite.Suite
	authToken string
}

type UrlsResponse struct {
	Kind     string   `json:"kind"`
	Name     string   `json:"name"`
	PointsTo string   `json:"pointsTo"`
	Urls     []string `json:"urls"`
}

func (suite *UrlsTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}

func HitGetUrls(queryParams map[string]string, authToken string, url string) UrlsDto {
	resp, err := Base.MakeApiCall(url, http.MethodGet, "", queryParams, authToken)
	Base.HandleError(err, GetUrls)

	structUrlsRouter := StructUrlsRouter{}
	urlsRouter := structUrlsRouter.UnmarshalGivenResponseBody(resp.Body())
	return urlsRouter.urlsResponse
}
func (r StructUrlsRouter) UnmarshalGivenResponseBody(response []byte) StructUrlsRouter {
	json.Unmarshal(response, &r.urlsResponse)
	return r
}
func GetEnvironmentConfigForDevtronApp() (*DevtronAppConfig, error) {
	cfg := &DevtronAppConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}
func GetEnvironmentConfigForInstalledApp() (*InstalledAppConfig, error) {
	cfg := &InstalledAppConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

func GetTestExpectedUrlsData() []UrlsResponse {
	testFile, err := os.Open("../testdata/UrlsRouter/urlsTestData.json")
	if err != nil {
		fmt.Println(err)
		return []UrlsResponse{}
	}
	defer testFile.Close()
	data, _ := ioutil.ReadAll(testFile)
	res := make([]UrlsResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println(err)
		return []UrlsResponse{}
	}
	return res
}

func testGetUrlsForHelmApp(suite *UrlsTestSuite, token string) {
	envConf, _ := helmRouter.GetEnvironmentConfigForHelmApp()
	queryParams := map[string]string{"appId": envConf.HAppId}
	log.Println("Hitting Get urls API")
	if token == "" {
		token = suite.authToken
	}
	resp := HitGetUrls(queryParams, token, GetUrlsUrlHelm)
	assert.Equal(suite.T(), 200, resp.Code)
}

func testGetUrlsForHelmAppWithIncorrectAppId(suite *UrlsTestSuite, token string) {
	randomHAppId := Base.GetRandomNumberOf9Digit()
	queryParams := map[string]string{"appId": strconv.Itoa(randomHAppId)}
	if token == "" {
		token = suite.authToken
	}
	resp := HitGetUrls(queryParams, token, GetUrlsUrlHelm)
	assert.Equal(suite.T(), 400, resp.Code)
	assert.Equal(suite.T(), "Bad Request", resp.Status)
}

func testGetUrlsForDevtronApp(suite *UrlsTestSuite, token string) {
	envConf, _ := GetEnvironmentConfigForDevtronApp()
	queryParams := map[string]string{"appId": envConf.AppId, "envId": envConf.EnvId}
	log.Println("Hitting Get urls API")
	if token == "" {
		token = suite.authToken
	}
	resp := HitGetUrls(queryParams, token, GetUrlsUrl)
	assert.Equal(suite.T(), 200, resp.Code)
}

func testGetUrlsForDevtronAppWithIncorrectAppId(suite *UrlsTestSuite, token string) {
	randomInstalledAppId := "installedAppId-1"
	randomEnvId := "envid-1"
	queryParams := map[string]string{"installedAppId": randomInstalledAppId, "envId": randomEnvId}
	if token == "" {
		token = suite.authToken
	}
	resp := HitGetUrls(queryParams, token, GetUrlsUrl)
	assert.Equal(suite.T(), 400, resp.Code)
	assert.Equal(suite.T(), "Bad Request", resp.Status)
}

func testGetUrlsForInstalledApp(suite *UrlsTestSuite, token string) {
	envConf, _ := GetEnvironmentConfigForInstalledApp()
	queryParams := map[string]string{"installedAppId": envConf.InstalledAppId, "envId": envConf.EnvId}
	log.Println("Hitting Get urls API")
	if token == "" {
		token = suite.authToken
	}
	resp := HitGetUrls(queryParams, token, GetUrlsUrl)
	assert.Equal(suite.T(), 200, resp.Code)
}

func testGetUrlsForInstalledAppWithIncorrectAppId(suite *UrlsTestSuite, token string) {
	randomAppId := "appId-1"
	randomEnvId := "envid-1"
	queryParams := map[string]string{"appId": randomAppId, "envId": randomEnvId}
	if token == "" {
		token = suite.authToken
	}
	resp := HitGetUrls(queryParams, token, GetUrlsUrl)
	assert.Equal(suite.T(), 400, resp.Code)
	assert.Equal(suite.T(), "Bad Request", resp.Status)
}

func testGetUrlsdata(suite *UrlsTestSuite, token string) {
	expected := GetTestExpectedUrlsData()
	envConf, _ := GetEnvironmentConfigForDevtronApp()
	queryParams := map[string]string{"appId": envConf.AppId, "envId": envConf.EnvId}
	log.Println("Hitting Get urls API")
	if token == "" {
		token = suite.authToken
	}
	resp := HitGetUrls(queryParams, token, GetUrlsUrl)
	assert.Equal(suite.T(), 200, resp.Code)
	assert.Equal(suite.T(), 3, len(resp.Result))
	for j, _ := range resp.Result {
		respData := resp.Result[j]
		assert.Equal(suite.T(), respData.Name, expected[j].Name)
		assert.Equal(suite.T(), respData.Kind, expected[j].Kind)
		for i, url := range respData.Urls {
			assert.Equal(suite.T(), url, expected[j].Urls[i])
		}
	}
}

func createUserRequestPayloadWithRole(role string) RequestDTOs.UserInfo {
	var listOfRoleFilter []ResponseDTOs.RoleFilter
	var userInfo RequestDTOs.UserInfo
	roleFilter := CreateRoleFilterWithDevtronAppsOnly(role)
	listOfRoleFilter = append(listOfRoleFilter, roleFilter)
	userInfo.EmailId = Base.GetRandomStringOfGivenLength(10) + "@yopmail"
	userInfo.SuperAdmin = false
	userInfo.RoleFilters = listOfRoleFilter
	userInfo.Groups = []string{}
	return userInfo
}

func testUrlsdataWithRoleAccess(suite *UrlsTestSuite, role string) {
	//create api-token
	createApiTokenRequestDTO := ApiTokenRouter.GetPayLoadForCreateApiToken()
	payloadForCreateApiTokenRequest, _ := json.Marshal(createApiTokenRequestDTO)
	responseOfCreateApiToken := ApiTokenRouter.HitCreateApiTokenApi(string(payloadForCreateApiTokenRequest), suite.authToken)
	token := responseOfCreateApiToken.Result.Token
	//update user with permissions

	updateUserDto := createUserRequestPayloadWithRole(role)
	updateUserDto.Id = int32(responseOfCreateApiToken.Result.UserId)
	byteValueOfStruct, _ := json.Marshal(updateUserDto)
	responseOfUpdateUserApi := userRouter.HitUpdateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Code, 200)
	//test with token user
	testGetUrlsForDevtronApp(suite, token)
	testGetUrlsdata(suite, token)
	//delete user before deleting token
	responseOfDeleteUserApi := userRouter.HitDeleteUserApi(strconv.Itoa(int(updateUserDto.Id)), suite.authToken)
	assert.Equal(suite.T(), responseOfDeleteUserApi.Code, 200)
	//delete created token
	var tokenId int
	responseOfGetAllApiTokens := ApiTokenRouter.HitGetAllApiTokens(suite.authToken).Result
	for _, result := range responseOfGetAllApiTokens {
		if result.UserId == responseOfCreateApiToken.Result.UserId {
			assert.Equal(suite.T(), responseOfCreateApiToken.Result.Token, result.Token)
			assert.Equal(suite.T(), createApiTokenRequestDTO.ExpireAtInMs, result.ExpireAtInMs)
			assert.Equal(suite.T(), createApiTokenRequestDTO.Name, result.Name)
			tokenId = result.Id
		}
	}
	log.Println("=== Here We Deleting the Token After Verification")
	_ = ApiTokenRouter.HitDeleteApiToken(strconv.Itoa(tokenId), suite.authToken)
}
