package AppStoreDiscoverRouter

import (
	"automation-suite/ChartRepositoryRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"time"
)

func (suite *AppStoreDiscoverTestSuite) TestDiscoverHelmAppsViaAppstoreName() {
	log.Println("=== Here we are Adding a chart repo ===")
	chartRepoConfig, _ := ChartRepositoryRouter.GetChartRepoRouterConfig()
	RepoName := Base.GetRandomStringOfGivenLength(8)
	createChartRepoRequestDto := ChartRepositoryRouter.CreateChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
	byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
	respGetRepoApi := ChartRepositoryRouter.HitCreateChartRepoApi(byteValueOfStruct, suite.authToken)
	chartRepoName := respGetRepoApi.Result.Name

	suite.Run("A=1=DiscoverWithCorrectRepoName", func() {
		queryParams := map[string]string{"appStoreName": chartRepoName}
		PollForGettingHelmAppData(queryParams, suite.authToken)
		ActiveDiscoveredApps := HitDiscoverAppApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), chartRepoName, ActiveDiscoveredApps.Result[0].Name)
		assert.False(suite.T(), ActiveDiscoveredApps.Result[0].Deprecated)
	})

	suite.Run("A=2=DiscoverWithInCorrectRepoName", func() {
		randomRepoName := Base.GetRandomStringOfGivenLength(8)
		queryParams := map[string]string{"chartRepoId": randomRepoName}
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
