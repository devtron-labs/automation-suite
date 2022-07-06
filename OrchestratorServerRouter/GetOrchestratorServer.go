package OrchestratorServerRouter

import (
	"github.com/stretchr/testify/assert"
)

func (suite *ServerRouterTestSuite) TestGetServerStatus() {

	suite.Run("A=1=GetServerStatus", func() {
		serverStatus := HitGetOrchestratorServerApi(suite.authToken)
		assert.NotNil(suite.T(), serverStatus.Result.CurrentVersion)
		assert.NotNil(suite.T(), serverStatus.Result.Status)
		assert.Equal(suite.T(), serverStatus.Result.ReleaseName, "devtron")
	})
}
