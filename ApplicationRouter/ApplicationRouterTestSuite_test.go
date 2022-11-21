package ApplicationRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestApplicationRouterSuite(t *testing.T) {
	suite.Run(t, new(ApplicationsRouterTestSuite))
}
