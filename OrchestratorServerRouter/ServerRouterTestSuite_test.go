package OrchestratorServerRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestOrchestratorRouterSuite(t *testing.T) {
	suite.Run(t, new(ServerRouterTestSuite))
}
