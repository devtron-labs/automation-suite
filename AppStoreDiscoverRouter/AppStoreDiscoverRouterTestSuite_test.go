package AppStoreDiscoverRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAppStoreDiscoverRouterSuite(t *testing.T) {
	suite.Run(t, new(AppStoreDiscoverTestSuite))
}
