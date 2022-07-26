package CommonRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestCommonRouterSuite(t *testing.T) {
	suite.Run(t, new(BaseCommonRouterTestSuite))
}
