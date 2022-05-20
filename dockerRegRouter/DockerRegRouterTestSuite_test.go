package dockerRegRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestDockerRegRouterSuite(t *testing.T) {
	suite.Run(t, new(DockerRegRouterTestSuite))
}
