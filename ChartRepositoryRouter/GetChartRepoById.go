package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func (suite *ChartRepoTestSuite) TestClassC3GetChartRepoById() {

	suite.Run("A=1=GetRepoByValidId", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		respGetRepoApi := HitCreateChartRepoApi(string(byteValueOfStruct), suite.authToken)

		respGetRepoListApi := HitGetChartRepoViaId(suite.authToken, strconv.Itoa(respGetRepoApi.Result.Id))
		assert.Equal(suite.T(), RepoName, respGetRepoListApi.Result.Name)
		assert.Equal(suite.T(), AUTH_MODE_ANONYMOUS, respGetRepoListApi.Result.AuthMode)

		createChartRepoRequestDto = createChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, respGetRepoApi.Result.Id, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
		HitDeleteChartRepo(string(byteValueOfStruct), suite.authToken)
	})

	suite.Run("A=2=GetRepoByInvalidId", func() {
		randomId := Base.GetRandomNumberOf9Digit()
		respGetRepoListApi := HitGetChartRepoViaId(suite.authToken, strconv.Itoa(randomId))
		assert.False(suite.T(), respGetRepoListApi.Result.Active)
		assert.False(suite.T(), respGetRepoListApi.Result.Default)
	})
}
