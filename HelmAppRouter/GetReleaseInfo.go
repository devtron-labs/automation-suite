package HelmAppRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func (suite *HelmAppTestSuite) TestGetReleaseInfo() {
	suite.Run("A=1=GetReleaseInfoApiWithValidHAppId", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		queryParams := map[string]string{"appId": envConf.HAppId}
		resp := HitGetReleaseInfoApi(queryParams, suite.authToken)

		expectedDefaultValueString, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/HelmAppRouter/DefaultValuesForReleaseInfo.txt")
		assert.Equal(suite.T(), string(expectedDefaultValueString), resp.Result.ReleaseInfo.DefaultValues)

		expectedMergedValueString, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/HelmAppRouter/MergedValuesForReleaseInfo.txt")
		assert.Equal(suite.T(), string(expectedMergedValueString), resp.Result.ReleaseInfo.MergedValues)
		assert.Equal(suite.T(), envConf.HAppId, resp.Result.ReleaseInfo.DeployedAppDetail.AppId)
	})

	suite.Run("A=2=GetReleaseInfoApiWithInvalidHAppId", func() {
		randomNumber := Base.GetRandomNumberOf9Digit()
		queryParams := map[string]string{"appId": strconv.Itoa(randomNumber)}
		resp := HitGetReleaseInfoApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), 400, resp.Code)
	})

}
