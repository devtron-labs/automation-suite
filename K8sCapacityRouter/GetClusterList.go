package K8sCapacityRouter

import "github.com/stretchr/testify/assert"

func (suite *K8sCapacityRoutersTestSuite) TestGetClusterList() {

	suite.Run("A=1=GetClusterList", func() {
		ClusterList := HitGetClusterListApi(suite.authToken)
		var isDefaultClusterPresent bool
		for _, cluster := range ClusterList.Result {
			if cluster.Name == "default_cluster" {
				isDefaultClusterPresent = true
				break
			}
		}
		assert.True(suite.T(), isDefaultClusterPresent)
	})
}
