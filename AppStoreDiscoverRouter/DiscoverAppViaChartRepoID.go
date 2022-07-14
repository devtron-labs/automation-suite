package AppStoreDiscoverRouter

import (
	"automation-suite/ChartRepositoryRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestDiscoverHelmAppsViaChartRepoId() {
	log.Println("=== Here we are Adding a chart repo ===")
	chartRepoConfig, _ := ChartRepositoryRouter.GetChartRepoRouterConfig()
	RepoName := Base.GetRandomStringOfGivenLength(8)
	createChartRepoRequestDto := ChartRepositoryRouter.CreateChartRepoRequestPayload(AuthModeAnonymous, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
	byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
	respGetRepoApi := ChartRepositoryRouter.HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
	chartRepoId := respGetRepoApi.Result.Id

	suite.Run("A=1=DiscoverWithCorrectRepoId", func() {
		queryParams := map[string]string{"chartRepoId": strconv.Itoa(chartRepoId)}
		PollForGettingHelmAppData(queryParams, suite.authToken)
		ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), chartRepoId, ActiveDiscoveredApps.Result[0].ChartRepoId)
		assert.False(suite.T(), ActiveDiscoveredApps.Result[0].Deprecated)
	})

	suite.Run("A=2=DiscoverWithInCorrectRepoId", func() {
		randomRepoId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		queryParams := map[string]string{"chartRepoId": randomRepoId}
		time.Sleep(10 * time.Second)
		ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
		assert.Nil(suite.T(), ActiveDiscoveredApps.Result)
	})

	log.Println("=== Here we are Deleting chart repo after verifications ===")
	createChartRepoRequestDto.Id = respGetRepoApi.Result.Id
	byteValueOfStruct, _ = json.Marshal(createChartRepoRequestDto)
	deleteChartRepoApiResp := ChartRepositoryRouter.HitDeleteChartRepo(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), "Chart repo deleted successfully.", deleteChartRepoApiResp.Result)
}
