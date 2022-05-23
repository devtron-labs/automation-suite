package AppLabelsRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func (suite *AppLabelRouterTestSuite) TestGetAppLabelsWithValidAppId() {
	config, _ := GetEnvironmentConfigForAppLabelsRouter()
	appMetaInfo := HitGetAppMetaInfoByIdApi(config.AppIdForAppLabelRouter, suite.authToken)
	assert.NotNil(suite.T(), appMetaInfo.Result.CreatedOn)
	assert.True(suite.T(), appMetaInfo.Result.Active)
}

func (suite *AppLabelRouterTestSuite) TestGetAppLabelsWithInvalidAppId() {
	randomNumber := strconv.Itoa(testUtils.GetRandomNumberOf9Digit())
	appMetaInfo := HitGetAppMetaInfoByIdApi(randomNumber, suite.authToken)
	assert.Equal(suite.T(), appMetaInfo.Errors[0].UserMessage, "pg: no rows in result set")
}
