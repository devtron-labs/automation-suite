package PipelineConfigRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPipelineConfigSuite(t *testing.T) {
	suite.Run(t, new(PipelinesConfigRouterTestSuite))
}
