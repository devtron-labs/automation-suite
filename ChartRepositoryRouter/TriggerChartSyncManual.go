package ChartRepositoryRouter

import (
	"github.com/stretchr/testify/assert"
)

func (suite ChartRepoTestSuite) TestTriggerChartSyncManualApi() {
	triggerChartSyncApiResponse := HitTriggerChartSyncManualApi(suite.authToken)
	assert.Equal(suite.T(), "ok", triggerChartSyncApiResponse.Result.Status)
}
