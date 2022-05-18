package IntegrationTestCases

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestCases struct {
	suite.Suite
	authToken string
}

func (suite *IntegrationTestCases) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
