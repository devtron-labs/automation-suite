package SSOLoginRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSSOLoginRouterSuite(t *testing.T) {
	suite.Run(t, new(SSOLoginTestSuite))
}
