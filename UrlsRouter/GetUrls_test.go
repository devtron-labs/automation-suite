package UrlsRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestUrlsTestSuite(t *testing.T) {
	suite.Run(t, new(UrlsTestSuite))
}
