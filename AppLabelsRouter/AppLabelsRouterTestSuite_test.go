package AppLabelsRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAttributeRouterSuite(t *testing.T) {
	suite.Run(t, new(AppLabelRouterTestSuite))
}
