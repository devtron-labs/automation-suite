package AppStoreDeploymentRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAppStoreDeploymentRouterSuite(t *testing.T) {
	suite.Run(t, new(AppStoreDeploymentTestSuite))
}
