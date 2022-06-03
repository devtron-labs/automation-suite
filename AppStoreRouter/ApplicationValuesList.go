package AppStoreRouter

import (
	"automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

//todo need to add more assertions after setup of stage Environment and deploying a chart for permanent test data

func (suite *AppStoreTestSuite) TestGetApplicationValuesList() {

	suite.Run("A=1=FetchAppValuesWithValidAppStoreId", func() {
		envConf, _ := GetEnvironmentConfigForAppStoreRouter()
		resp := HitGetApplicationValuesList(envConf.AppStoreId, suite.authToken)
		log.Println("Asserting the API Response...")
		assert.Equal(suite.T(), 4, len(resp.Result.Values))
		assert.Equal(suite.T(), "DEFAULT", resp.Result.Values[0].Kind)
		assert.Equal(suite.T(), "EXISTING", resp.Result.Values[3].Kind)
	})
	suite.Run("A=2=FetchAppValuesWithInvalidAppStoreId", func() {
		randomNumber := testUtils.GetRandomNumberOf9Digit()
		resp := HitGetApplicationValuesList(strconv.Itoa(randomNumber), suite.authToken)
		log.Println("Asserting the API Response...")
		assert.Nil(suite.T(), resp.Result.Values[0].Values)
		assert.Nil(suite.T(), resp.Result.Values[1].Values)
		assert.Empty(suite.T(), resp.Result.Values[2].Values)
		assert.Empty(suite.T(), resp.Result.Values[3].Values)
	})
}
