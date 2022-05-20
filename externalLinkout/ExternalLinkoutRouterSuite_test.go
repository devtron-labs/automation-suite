package externalLinkout

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestLinkOutRouterSuite(t *testing.T) {
	suite.Run(t, new(LinkOutRouterTestSuite))
}
