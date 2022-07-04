package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

func (suite *ChartRepoTestSuite) TestClassC7DeleteChartRepo() {

	suite.Run("A=1=DeleteRepoHavingAnonymousAuthMode", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		deleteChartRepoApiResp := HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
	})

	suite.Run("A=2=DeleteRepoHavingAuthModeAccessToken", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		deleteChartRepoApiResp := HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
	})

	suite.Run("A=3=DeleteRepoHavingInvalidId", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 0, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ACCESS_TOKEN, 123456789, RepoName, chartRepoConfig.ChartRepoUrl, chartRepoConfig.ChartAccessToken, true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		deleteChartRepoApiResp := HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "pg: no rows in result set", deleteChartRepoApiResp.Errors[0].UserMessage)
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		deleteChartRepoApiResp = HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
	})
}

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
