package externalLinkoutRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestLinkOutRouterSuite(t *testing.T) {
	suite.Run(t, new(ExternalLinkOutRouterTestSuite))
}
