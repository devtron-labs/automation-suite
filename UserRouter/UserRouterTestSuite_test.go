package UserRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestUserRouterSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
