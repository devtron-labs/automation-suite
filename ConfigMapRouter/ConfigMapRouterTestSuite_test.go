package ConfigMapRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestConfigMapSuite(t *testing.T) {
	suite.Run(t, new(ConfigsMapRouterTestSuite))
}
