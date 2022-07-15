package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
	"time"
)

func (suite *ChartRepoTestSuite) TestClassC3GetChartRepoById() {

	suite.Run("A=1=GetRepoByValidId", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := CreateChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
		time.Sleep(2 * time.Second)
		respGetRepoListApi := HitGetChartRepoViaId(suite.authToken, strconv.Itoa(respGetRepoApi.Result.Id))
		assert.Equal(suite.T(), RepoName, respGetRepoListApi.Result.Name)
		assert.Equal(suite.T(), AUTH_MODE_ANONYMOUS, respGetRepoListApi.Result.AuthMode)
		createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		deleteChartRepoApiResp := HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
	})

	suite.Run("A=2=GetRepoByInvalidId", func() {
		randomId := Base.GetRandomNumberOf9Digit()
		respGetRepoListApi := HitGetChartRepoViaId(suite.authToken, strconv.Itoa(randomId))
		assert.False(suite.T(), respGetRepoListApi.Result.Active)
		assert.False(suite.T(), respGetRepoListApi.Result.Default)
	})
}
