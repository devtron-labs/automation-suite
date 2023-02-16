package ConfigMapRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/suite"
	"log"
)

// ConfigsMapRouterTestSuite =================PipelineConfigSuite Setup =========================
type ConfigsMapRouterTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *ConfigsMapRouterTestSuite) SetupSuite() {
	log.Println("=== Running Before Suite Method ===")
	suite.authToken = Base.GetAuthToken()
}
