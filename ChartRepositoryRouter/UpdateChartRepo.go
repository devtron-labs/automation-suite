package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func (suite *ChartRepoTestSuite) TestUpdateChartRepo() {

	suite.Run("A=1=UpdateAuthFromAnonymousToAccessToken", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		updateChartRepoResponse := HitUpdateChartRepoApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), AUTH_MODE_ACCESS_TOKEN, updateChartRepoResponse.Result.AuthMode)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		HitDeleteChartRepo(string(byteValueOfStruct), suite.authToken)
	})

	suite.Run("A=2=UpdateAuthFromAccessTokenToAnonymous", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		updateChartRepoResponse := HitUpdateChartRepoApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), AUTH_MODE_ANONYMOUS, updateChartRepoResponse.Result.AuthMode)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		HitDeleteChartRepo(string(byteValueOfStruct), suite.authToken)
	})

	suite.Run("A=3=UpdateAccessTokenForChartRepo", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken+"updatedUrl", true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		updateChartRepoResponse := HitUpdateChartRepoApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), AUTH_MODE_ACCESS_TOKEN, updateChartRepoResponse.Result.AuthMode)
		assert.Equal(suite.T(), chartRepoConfig.ChartAccessToken+"updatedUrl", updateChartRepoResponse.Result.AccessToken)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken+"updatedUrl", true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		HitDeleteChartRepo(string(byteValueOfStruct), suite.authToken)
	})

	suite.Run("A=4=UpdateActiveFalseFromTrue", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken+"updatedUrl", false)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		updateChartRepoResponse := HitUpdateChartRepoApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), AUTH_MODE_ACCESS_TOKEN, updateChartRepoResponse.Result.AuthMode)
		assert.Equal(suite.T(), chartRepoConfig.ChartAccessToken+"updatedUrl", updateChartRepoResponse.Result.AccessToken)
		assert.False(suite.T(), updateChartRepoResponse.Result.Active)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken+"updatedUrl", false)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		HitDeleteChartRepo(string(byteValueOfStruct), suite.authToken)
	})
}

//todo will add test case for name-update once dev will fix the issue
