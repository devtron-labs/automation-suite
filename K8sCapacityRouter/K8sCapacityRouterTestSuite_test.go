package K8sCapacityRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestK8sCapacityRouterSuite(t *testing.T) {
	suite.Run(t, new(K8sCapacityRoutersTestSuite))
}
