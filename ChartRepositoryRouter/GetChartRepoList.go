package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

func (suite *ChartRepoTestSuite) TestClassC4GetRepoList() {

	suite.Run("A=1=GetRepoList", func() {
		respGetRepoListApi := HitGetChartRepoList(suite.authToken)
		listSize := len(respGetRepoListApi.Result)
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := CreateChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		respGetRepoListApi = HitGetChartRepoList(suite.authToken)
		assert.Equal(suite.T(), listSize+1, len(respGetRepoListApi.Result))
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		deleteChartRepoApiResp := HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
	})
}
