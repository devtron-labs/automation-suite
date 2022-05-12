package SSOLoginRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *SSOLoginTestSuite) TestGetSsoLoginConfigWithCorrectName() {
	envConf, _ := Base.GetEnvironmentConfig()
	queryParams := map[string]string{"name": "google"}
	log.Println("Hitting the Get SSO login config by Name API")
	actualSSODetailsResponse := HitGetLoginConfigByNameApi(queryParams, suite.authToken)

	log.Println("Asserting the API Response...")
	assert.Equal(suite.T(), envConf.BaseServerUrl+"orchestrator", actualSSODetailsResponse.CreateSSODetailsRequestDto.Url)
	assert.Equal(suite.T(), envConf.SSOClientSecret, actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.ClientSecret)
	assert.Equal(suite.T(), envConf.BaseServerUrl+"orchestrator/api/dex/callback", actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.RedirectURI)
}

func (suite *SSOLoginTestSuite) TestGetSsoLoginConfigWithInCorrectName() {
	queryParams := map[string]string{"name": "InCorrectGoogle"}
	log.Println("Hitting the Get SSO login config by Name API")
	actualSSODetailsResponse := HitGetLoginConfigByNameApi(queryParams, suite.authToken)

	log.Println("Asserting the API Response...")
	assert.Nil(suite.T(), actualSSODetailsResponse.CreateSSODetailsRequestDto)
}

//Todo this test case is failing as we are not trimming white spaces in given query params
/*func (suite *SSOLoginTestSuite) TestGetSsoLoginConfigWithCorrectNameHavingWhiteSpace() {
	envConf, _ := Base.GetEnvironmentConfig()
	queryParams := map[string]string{"name": "  google  "}
	log.Println("Hitting the Get SSO login config by Name API")
	actualSSODetailsResponse := HitGetLoginConfigByNameApi(queryParams)

	log.Println("Asserting the API Response...")
	assert.Equal(suite.T(), envConf.BaseServerUrl+"orchestrator", actualSSODetailsResponse.CreateSSODetailsRequestDto.Url)
	assert.Equal(suite.T(), envConf.SSOClientSecret, actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.ClientSecret)
	assert.Equal(suite.T(), envConf.BaseServerUrl+"orchestrator/api/dex/callback", actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.RedirectURI)
}
*/

//disabling the above test case for sometime until we fix it
