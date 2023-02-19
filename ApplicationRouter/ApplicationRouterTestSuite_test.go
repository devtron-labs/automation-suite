package ApplicationRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestApplicationRouterSuite(t *testing.T) {
	applicationsRouterTestSuite := new(ApplicationsRouterTestSuite)
	suite.Run(t, new(ApplicationsRouterTestSuite))
	applicationsRouterTestSuite.AfterSuite()
}
