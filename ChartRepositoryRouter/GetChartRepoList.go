package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func (suite ChartRepoTestSuite) TestGetRepoList() {
	respGetRepoListApi := HitGetChartRepoList(suite.authToken)
	listSize := len(respGetRepoListApi.Result)
	chartRepoConfig, _ := GetChartRepoRouterConfig()
	RepoName := Base.GetRandomStringOfGivenLength(8)
	createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
	byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
	respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)

	respGetRepoListApi = HitGetChartRepoList(suite.authToken)
	assert.Equal(suite.T(), listSize+1, len(respGetRepoListApi.Result))

	createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
	byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
	deleteChartRepoApiResp := HitDeleteChartRepo(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
}
