package ChartRepositoryRouter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestChartRepoRouterSuite(t *testing.T) {
	suite.Run(t, new(ChartRepoTestSuite))
}
