package AppListingRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAppListingRouterSuite(t *testing.T) {
	suite.Run(t, new(AppListingRouterTestSuite))
}
