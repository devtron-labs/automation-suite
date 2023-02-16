package GitopsConfigRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestGitOpsRouterSuite(t *testing.T) {
	suite.Run(t, new(GitOpsRouterTestSuite))
}
