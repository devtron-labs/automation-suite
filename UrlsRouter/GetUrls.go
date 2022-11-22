package UrlsRouter

import (
	"automation-suite/ApiTokenRouter"
	helmRouter "automation-suite/HelmAppRouter"
	userRouter "automation-suite/UserRouter"
	"automation-suite/UserRouter/RequestDTOs"
	"automation-suite/UserRouter/ResponseDTOs"
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UrlsTestSuite) TestGetUrlsForHelmApp() {
	testGetUrlsForHelmApp(suite, suite.authToken)
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
func (suite *UrlsTestSuite) TestGetUrlsForHelmAppWithIncorrectAppId() {
	testGetUrlsForHelmAppWithIncorrectAppId(suite, suite.authToken)
}
func testGetUrlsForHelmAppWithIncorrectAppId(suite *UrlsTestSuite, token string) {
	randomHAppId := testUtils.GetRandomNumberOf9Digit()
	queryParams := map[string]string{"appId": strconv.Itoa(randomHAppId)}
	if token == "" {
		token = suite.authToken
	}
	resp := HitGetUrls(queryParams, token, GetUrlsUrlHelm)
	assert.Equal(suite.T(), 400, resp.Code)
	assert.Equal(suite.T(), "Bad Request", resp.Status)
}
func (suite *UrlsTestSuite) TestGetUrlsForDevtronApp() {
	testGetUrlsForDevtronApp(suite, suite.authToken)
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
func (suite *UrlsTestSuite) TestGetUrlsForDevtronAppWithIncorrectAppId() {
	testGetUrlsForDevtronAppWithIncorrectAppId(suite, suite.authToken)
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

func (suite *UrlsTestSuite) TestGetUrlsForInstalledApp() {
	testGetUrlsForInstalledApp(suite, suite.authToken)
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

func (suite *UrlsTestSuite) TestGetUrlsForInstalledAppWithIncorrectAppId() {
	testGetUrlsForInstalledAppWithIncorrectAppId(suite, suite.authToken)
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

func (suite *UrlsTestSuite) TestGetUrlsdata() {
	testGetUrlsdata(suite, suite.authToken)
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

func createUserRequestPayloadViewOnly() RequestDTOs.UserInfo {
	var listOfRoleFilter []ResponseDTOs.RoleFilter
	var userInfo RequestDTOs.UserInfo
	roleFilter := userRouter.CreateRoleFilterWithDevtronAppsOnly()
	listOfRoleFilter = append(listOfRoleFilter, roleFilter)
	userInfo.EmailId = "@yopmail"
	userInfo.SuperAdmin = false
	userInfo.RoleFilters = listOfRoleFilter
	userInfo.Groups = []string{}
	return userInfo
}

func (suite *UrlsTestSuite) TestUrlsdata() {
	//create api-token
	createApiTokenRequestDTO := ApiTokenRouter.GetPayLoadForCreateApiToken()
	payloadForCreateApiTokenRequest, _ := json.Marshal(createApiTokenRequestDTO)
	responseOfCreateApiToken := ApiTokenRouter.HitCreateApiTokenApi(string(payloadForCreateApiTokenRequest), suite.authToken)
	token := responseOfCreateApiToken.Result.Token
	//update user with permissions

	updateUserDto := createUserRequestPayloadViewOnly()
	updateUserDto.Id = int32(responseOfCreateApiToken.Result.UserId)
	byteValueOfStruct, _ := json.Marshal(updateUserDto)
	responseOfUpdateUserApi := userRouter.HitUpdateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Code, 200)
	//test with token user
	testGetUrlsdata(suite, token)
	testGetUrlsForInstalledAppWithIncorrectAppId(suite, token)
	testGetUrlsForInstalledApp(suite, token)
	testGetUrlsForDevtronAppWithIncorrectAppId(suite, token)
	testGetUrlsForDevtronApp(suite, token)
	testGetUrlsForHelmAppWithIncorrectAppId(suite, token)
	testGetUrlsForHelmApp(suite, token)

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
	responseOfDeleteApi := ApiTokenRouter.HitDeleteApiToken(strconv.Itoa(tokenId), suite.authToken)
	assert.True(suite.T(), responseOfDeleteApi.Result.Success)

}
