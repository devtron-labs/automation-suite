package SSOLoginRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *SSOLoginTestSuite) TestGetSsoLoginWithCorrectId() {
	envConf, _ := Base.GetEnvironmentConfig()
	log.Println("Hitting the Get SSO Details API")
	actualSSODetailsResponse := HitGetSSODetailsApi("1", suite.authToken)

	log.Println("Asserting the API Response...")
	assert.Equal(suite.T(), envConf.BaseServerUrl+"orchestrator", actualSSODetailsResponse.CreateSSODetailsRequestDto.Url)
	assert.Equal(suite.T(), envConf.SSOClientSecret, actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.ClientSecret)
	assert.Equal(suite.T(), envConf.BaseServerUrl+"orchestrator/api/dex/callback", actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.RedirectURI)
}

func (suite *SSOLoginTestSuite) TestGetSsoLoginWithInCorrectId() {
	log.Println("Hitting the Get SSO Details API")
	actualSSODetailsResponse := HitGetSSODetailsApi("99999999", suite.authToken)

	log.Println("Asserting the API Response...")
	assert.Nil(suite.T(), actualSSODetailsResponse.CreateSSODetailsRequestDto)
}
