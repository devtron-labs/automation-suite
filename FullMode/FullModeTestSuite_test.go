package FullMode

import (
	"automation-suite/AppStoreRouter"
	"automation-suite/AttributesRouter"
	"automation-suite/ChartRepositoryRouter"
	"automation-suite/HelmAppRouter"
	"automation-suite/SSOLoginRouter"
	"automation-suite/TeamRouter"
	"automation-suite/UserRouter"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSSOLoginRouterSuite(t *testing.T) {
	suite.Run(t, new(SSOLoginRouter.SSOLoginTestSuite))
}

func TestTeamRouterSuite(t *testing.T) {
	suite.Run(t, new(TeamRouter.TeamTestSuite))
}

func TestUserRouterSuite(t *testing.T) {
	suite.Run(t, new(UserRouter.UserTestSuite))
}

func TestHelmAppRouterSuite(t *testing.T) {
	suite.Run(t, new(HelmAppRouter.HelmAppTestSuite))
}

func TestAppStoreRouterSuite(t *testing.T) {
	suite.Run(t, new(AppStoreRouter.AppStoreTestSuite))
}

func TestChartRepoRouterSuite(t *testing.T) {
	suite.Run(t, new(ChartRepositoryRouter.ChartRepoTestSuite))
}

func TestAttributeRouterSuite(t *testing.T) {
	suite.Run(t, new(AttributesRouter.AttributeRouterTestSuite))
}
