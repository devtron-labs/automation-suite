package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

//todo currently user is able to add repo with invalid access token as well so commenting this test case
/*func TestCreateChartRepoWithInvalidAuthModeAccessToken(t *testing.T) {
	chartRepoConfig, _ := GetChartRepoRouterConfig()
	RepoName := Base.GetRandomStringOfGivenLength(8)
	createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, "invalidAccessToken")
	byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
	authToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImE0YTg1YWIzYjExZGM2ZDMwYTQ2Mzk4MjdlNGZiNjVlM2UzZWJjZjAifQ.eyJpc3MiOiJodHRwczovL3N0YWdpbmcuZGV2dHJvbi5pbmZvL29yY2hlc3RyYXRvci9hcGkvZGV4Iiwic3ViIjoiQ2hVeE1EY3hNRGszTmpBNE5Ea3lOamswTlRjNU16a1NCbWR2YjJkc1pRIiwiYXVkIjoiYXJnby1jZCIsImV4cCI6MTY1MDk1OTEzMywiaWF0IjoxNjUwODcyNzMzLCJhdF9oYXNoIjoibkFycjFYVDlLeEY0cmkyR0dLaldsUSIsImVtYWlsIjoiZGVlcGFrQGRldnRyb24uYWkiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6ImRlZXBhayBwYW53YXIifQ.NrA8Nho8zTyX5kdXKjulIbarHReKTARcVUr5MnTekOdRE6lGb7x-b54Xatk-3151OK98jqZ7YVH9cujjcV_4IhF5lENjBqBP5TcuKwtWsBeLTVteqBOc5kl4I0vZSTjAcGJQN_y--yLQLv0DtujOXvWgzNaLRGUnvx3YtMRWXnHos0Du2062gravGKk_Rgru9YYhWQSung4zxw0awdKx6qkKy0CvoB8QUmQvVbpeVhnm2DTshcj9_rktHtpv6ebqrwkOTTTIpE9eMqmLNilqSkmmQMdr7wFt1y5p7Yat-nReIma8Lel48IAMz8yADm1LioJgDdhlXSC8jjdDZN6hGg"
	respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), authToken)
	assert.Equal(t, AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.AuthMode)
	assert.Equal(t, RepoName, respGetRepoApi.Result.Name)

	createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, "invalidAccessToken")
	byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
	deleteChartRepoApiResp := HitDeleteChartRepo(string(byteValueOfStruct), authToken)
	assert.Equal(t, "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
}
*/

//todo need to add test cases for ssh and username password auth type once issue resolved from backend

func (suite *ChartRepoTestSuite) TestClassC1CreateChartRepo() {

	suite.Run("A=1=CreateRepoWithValidArgsOnly", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), AUTH_MODE_ANONYMOUS, respGetRepoApi.Result.AuthMode)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		deleteChartRepoApiResp := HitDeleteChartRepo(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
	})

	suite.Run("A=2=CreateRepoWithInvalidUrl", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl+"invalid", "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "Could not find an index.yaml file in the repo directory. Please try another chart repo.", respGetRepoApi.Result.CustomErrMsg)

	})

	suite.Run("A=3=CreateRepoWithValidAuthModeAccessToken", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.AuthMode)
		assert.Equal(suite.T(), RepoName, respGetRepoApi.Result.Name)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		deleteChartRepoApiResp := HitDeleteChartRepo(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)

	})

	suite.Run("A=4=CreateRepoWithInValidChartRepoUrl", func() {
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, "https://invalid-chart-repo-url.com", "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "Could not validate the repo. Please try again.", respGetRepoApi.Result.CustomErrMsg)
	})
}
