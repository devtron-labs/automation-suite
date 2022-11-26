package FullMode

import (
	"automation-suite/ApiTokenRouter"
	"automation-suite/AppLabelsRouter"
	"automation-suite/AppListingRouter"
	"automation-suite/AppStoreDiscoverRouter"
	"automation-suite/ApplicationRouter"
	"automation-suite/AttributesRouter"
	"automation-suite/ChartRepositoryRouter"
	"automation-suite/GitopsConfigRouter"
	"automation-suite/HelmAppRouter"
	"automation-suite/PipelineConfigRouter"
	"automation-suite/SSOLoginRouter"
	"automation-suite/TeamRouter"
	"automation-suite/UserRouter"
	"automation-suite/dockerRegRouter"
	"automation-suite/externalLinkoutRouter"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSSOLoginRouterSuite(t *testing.T) {
	suite.Run(t, new(SSOLoginRouter.SSOLoginTestSuite))
}

func TestAttributeRouterSuite(t *testing.T) {
	suite.Run(t, new(AttributesRouter.AttributeRouterTestSuite))
}

func TestGitOpsRouterSuite(t *testing.T) {
	suite.Run(t, new(GitopsConfigRouter.GitOpsRouterTestSuite))
}

func TestTeamRouterSuite(t *testing.T) {
	suite.Run(t, new(TeamRouter.TeamTestSuite))
}

func TestUserRouterSuite(t *testing.T) {
	suite.Run(t, new(UserRouter.UserTestSuite))
}

func TestChartRepoRouterSuite(t *testing.T) {
	suite.Run(t, new(ChartRepositoryRouter.ChartRepoTestSuite))
}

func TestAppLabelsRouterSuite(t *testing.T) {
	suite.Run(t, new(AppLabelsRouter.AppLabelRouterTestSuite))
}

func TestAppListingRouterSuite(t *testing.T) {
	suite.Run(t, new(AppListingRouter.AppsListingRouterTestSuite))
}

func TestDockerRegRouterSuite(t *testing.T) {
	suite.Run(t, new(dockerRegRouter.DockersRegRouterTestSuite))
}

func TestLinkOutRouterSuite(t *testing.T) {
	suite.Run(t, new(externalLinkoutRouter.LinkOutRouterTestSuite))
}

func TestHelmAppRouterSuite(t *testing.T) {
	suite.Run(t, new(HelmAppRouter.HelmAppTestSuite))
}

func TestAppStoreDiscoverRouterSuite(t *testing.T) {
	suite.Run(t, new(AppStoreDiscoverRouter.AppStoreDiscoverTestSuite))
}

func TestPipelineConfigSuite(t *testing.T) {
	suite.Run(t, new(PipelineConfigRouter.PipelinesConfigRouterTestSuite))
}

func TestApplicationRouterSuite(t *testing.T) {
	suite.Run(t, new(ApplicationRouter.ApplicationsRouterTestSuite))
}

func TestApiTokenRouterSuite(t *testing.T) {
	suite.Run(t, new(ApiTokenRouter.ApiTokenRoutersTestSuite))
}
