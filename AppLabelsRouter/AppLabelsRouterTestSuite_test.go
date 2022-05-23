package AppLabelsRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAppLabelsRouterSuite(t *testing.T) {
	suite.Run(t, new(AppLabelRouterTestSuite))
}
