package UrlsRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestUrlsTestSuiteSuperAdmin(t *testing.T) {
	suite.Run(t, new(UrlsTestSuite))
}
