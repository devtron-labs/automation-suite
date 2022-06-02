package ConfigMapRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func bTestConfigMapSuite(t *testing.T) {
	suite.Run(t, new(ConfigsMapRouterTestSuite))
}
