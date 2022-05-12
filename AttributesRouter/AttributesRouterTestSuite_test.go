package AttributesRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAttributeRouterSuite(t *testing.T) {
	suite.Run(t, new(AttributeRouterTestSuite))
}
