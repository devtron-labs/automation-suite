package AppStoreDiscoverRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAppStoreDiscoverRouterSuite(t *testing.T) {
	appStoreDiscoverTestSuite := new(AppStoreDiscoverTestSuite)
	suite.Run(t, appStoreDiscoverTestSuite)
	appStoreDiscoverTestSuite.AfterSuite()
}
