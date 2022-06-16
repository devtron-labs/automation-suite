package SSOLoginRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *SSOLoginTestSuite) TestClass3GetSsoLogin() {
	envConf, _ := Base.ReadAnyJsonFile()

	suite.Run("A=1=GetSsoLoginWithCorrectId", func() {
		log.Println("Hitting the Get SSO Details API")
		actualSSODetailsResponse := HitGetSSODetailsApi("1", suite.authToken)

		log.Println("Asserting the API Response...")
		assert.Equal(suite.T(), envConf.BaseServerUrl+"/orchestrator", actualSSODetailsResponse.CreateSSODetailsRequestDto.Url)
		assert.Equal(suite.T(), envConf.SSOClientSecret, actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.ClientSecret)
		assert.Equal(suite.T(), envConf.BaseServerUrl+"/orchestrator/api/dex/callback", actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.RedirectURI)
	})

	suite.Run("A=2=GetSsoLoginWithIncorrectId", func() {
		log.Println("Hitting the Get SSO Details API")
		actualSSODetailsResponse := HitGetSSODetailsApi("99999999", suite.authToken)

		log.Println("Asserting the API Response...")
		assert.Nil(suite.T(), actualSSODetailsResponse.CreateSSODetailsRequestDto)
	})
}
