package ApiTokenRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestApiTokenRouterSuite(t *testing.T) {
	suite.Run(t, new(ApiTokenRoutersTestSuite))
}
