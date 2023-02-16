package UserTerminalAccessRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestUserTerminalAccessRouterTestSuite(t *testing.T) {
	suite.Run(t, new(UserTerminalAccessRoutersTestSuite))
}
