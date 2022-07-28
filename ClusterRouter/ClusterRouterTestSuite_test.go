package ClusterRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestClusterRouterSuite(t *testing.T) {
	suite.Run(t, new(ClustersRouterTestSuite))
}
