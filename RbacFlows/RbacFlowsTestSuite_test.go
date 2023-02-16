package RbacFlows

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestRbacFlowsSuite(t *testing.T) {
	suite.Run(t, new(RbacFlowTestSuite))
}
