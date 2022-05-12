package HelmAppRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestHelmAppRouterSuite(t *testing.T) {
	suite.Run(t, new(HelmAppTestSuite))
}
