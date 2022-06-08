package FullMode

import (
	"automation-suite/AppLabelsRouter"
	"automation-suite/AppListingRouter"
	"automation-suite/AppStoreRouter"
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

func TestGitOpsRouterSuite(t *testing.T) {
	suite.Run(t, new(GitopsConfigRouter.GitOpsRouterTestSuite))
}

func TestPipelineConfigSuite(t *testing.T) {
	suite.Run(t, new(PipelineConfigRouter.PipelinesConfigRouterTestSuite))
}
