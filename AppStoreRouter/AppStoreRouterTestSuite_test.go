package AppStoreRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAppStoreRouterSuite(t *testing.T) {
	suite.Run(t, new(AppStoreTestSuite))
}
