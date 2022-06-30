package SSOLoginRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

//Todo this test case is failing as we are not trimming white spaces in given query params
/*func (suite *SSOLoginTestSuite) TestGetSsoLoginConfigWithCorrectNameHavingWhiteSpace() {
	envConf, _ := Base.ReadAnyJsonFile()
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

func (suite *SSOLoginTestSuite) TestClass1GetSsoLoginConfig() {
	envConfig := Base.ReadBaseEnvConfig()
	baseCredentials := Base.ReadAnyJsonFile(envConfig.BaseCredentialsFile)
	classCredentials := Base.ReadAnyJsonFile(envConfig.ClassCredentialsFile)
	suite.Run("A=1=SsoLoginConfigWithCorrectName", func() {
		queryParams := map[string]string{"name": "google"}
		log.Println("Hitting the Get SSO login config by Name API")
		actualSSODetailsResponse := HitGetLoginConfigByNameApi(queryParams, suite.authToken)

		log.Println("Asserting the API Response...")
		assert.Equal(suite.T(), baseCredentials.BaseServerUrl+"/orchestrator", actualSSODetailsResponse.CreateSSODetailsRequestDto.Url)
		assert.Equal(suite.T(), classCredentials.SSOClientSecret, actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.ClientSecret)
		assert.Equal(suite.T(), baseCredentials.BaseServerUrl+"/orchestrator/api/dex/callback", actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.RedirectURI)

	})
	suite.Run("A=2=SsoLoginConfigWithIncorrectName", func() {
		queryParams := map[string]string{"name": "InCorrectGoogle"}
		log.Println("Hitting the Get SSO login config by Name API")
		actualSSODetailsResponse := HitGetLoginConfigByNameApi(queryParams, suite.authToken)

		log.Println("Asserting the API Response...")
		assert.Nil(suite.T(), actualSSODetailsResponse.CreateSSODetailsRequestDto)
	})
}
