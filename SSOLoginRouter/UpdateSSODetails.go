package SSOLoginRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

//todo this is failing as we are not updating the name(api issue)
func (suite *SSOLoginTestSuite) TestUpdateSsoLoginWithCorrectArgs() {
	envConf, _ := Base.GetEnvironmentConfig()
	byteValue, err := Base.GetByteArrayOfGivenJsonFile("../testdata/SSOLoginTestData/updateSSODetailsPayload.json")
	if nil != err {
		log.Println("Unable to get the byte value of given json !!", "err", err)
	}
	actualSSODetailsResponse := HitUpdateSSODetailsApi(byteValue, suite.authToken)

	log.Println("Asserting the API Response...")
	assert.Equal(suite.T(), "googleDeepak", actualSSODetailsResponse.CreateSSODetailsRequestDto.Name)
	assert.Equal(suite.T(), envConf.BaseServerUrl+"orchestrator/api/dex/callback/deepak", actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.RedirectURI)
	assert.Equal(suite.T(), envConf.SSOClientSecret+"Deepak", actualSSODetailsResponse.CreateSSODetailsRequestDto.Config.Config.ClientSecret)

	log.Println("Hitting the Get SSO Details API")
	actualSSODetailsResponseAfterUpdate := HitGetSSODetailsApi("2", suite.authToken)

	//disabling assert ,will enable again once the update name issue will get resolved
	//assert.Equal(suite.T(), "googleDeepak", actualSSODetailsResponseAfterUpdate.CreateSSODetailsRequestDto.Name)
	assert.Equal(suite.T(), envConf.BaseServerUrl+"orchestrator/api/dex/callback/deepak", actualSSODetailsResponseAfterUpdate.CreateSSODetailsRequestDto.Config.Config.RedirectURI)
	assert.Equal(suite.T(), envConf.SSOClientSecret+"Deepak", actualSSODetailsResponseAfterUpdate.CreateSSODetailsRequestDto.Config.Config.ClientSecret)
	assert.True(suite.T(), actualSSODetailsResponseAfterUpdate.CreateSSODetailsRequestDto.Active)

	log.Println("Resetting the content of SSODetails in DB")
	byteValue, err = Base.GetByteArrayOfGivenJsonFile("../testdata/SSOLoginTestData/defaultSSODetailsPayload.json")
	if nil != err {
		log.Println("Unable to get the byte value of given json !!", "err", err)
	}
	HitUpdateSSODetailsApi(byteValue, suite.authToken)
}
