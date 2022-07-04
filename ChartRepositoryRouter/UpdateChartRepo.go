package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

func (suite *ChartRepoTestSuite) TestClassC2UpdateChartRepo() {

	suite.Run("A=1=UpdateAuthFromAnonymousToAccessToken", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		updateChartRepoResponse := HitUpdateChartRepoApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), AUTH_MODE_ACCESS_TOKEN, updateChartRepoResponse.Result.AuthMode)
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
	})

	suite.Run("A=2=UpdateAuthFromAccessTokenToAnonymous", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		updateChartRepoResponse := HitUpdateChartRepoApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), AUTH_MODE_ANONYMOUS, updateChartRepoResponse.Result.AuthMode)
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
	})

	suite.Run("A=3=UpdateAccessTokenForChartRepo", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken+"updatedUrl", true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		updateChartRepoResponse := HitUpdateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		assert.Equal(suite.T(), AUTH_MODE_ACCESS_TOKEN, updateChartRepoResponse.Result.AuthMode)
		assert.Equal(suite.T(), chartRepoConfig.ChartAccessToken+"updatedUrl", updateChartRepoResponse.Result.AccessToken)
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
	})

	suite.Run("A=4=UpdateActiveFalseFromTrue", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken+"updatedUrl", false)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		updateChartRepoResponse := HitUpdateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		assert.Equal(suite.T(), AUTH_MODE_ACCESS_TOKEN, updateChartRepoResponse.Result.AuthMode)
		assert.Equal(suite.T(), chartRepoConfig.ChartAccessToken+"updatedUrl", updateChartRepoResponse.Result.AccessToken)
		assert.False(suite.T(), updateChartRepoResponse.Result.Active)
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
	})
}

//todo will add test case for name-update once dev will fix the issue
