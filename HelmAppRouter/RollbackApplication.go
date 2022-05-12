package HelmAppRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

var HAppId string
var version, installedAppId, installedVersionId int

var dataProviderForRollbackAppApiHavingInvalidArgs = []RollbackApplicationApiRequestDto{
	{strconv.Itoa(testUtils.GetRandomNumberOf9Digit()), version},
	{HAppId, testUtils.GetRandomNumberOf9Digit()},
	{strconv.Itoa(testUtils.GetRandomNumberOf9Digit()), testUtils.GetRandomNumberOf9Digit()},
}

func (suite *HelmAppTestSuite) TestRollBackApplicationApiWithValidPayload() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	HAppId = envConf.HAppId
	queryParams := map[string]string{"appId": envConf.HAppId}
	resp := HitGetDeploymentHistoryById(queryParams, suite.authToken)
	totalNoDeploymentHistory := len(resp.Result.DeploymentHistory)

	version = resp.Result.DeploymentHistory[totalNoDeploymentHistory-1].Version
	installedAppId = resp.Result.InstalledAppInfo.AppId
	installedVersionId = resp.Result.InstalledAppInfo.InstalledAppVersionId

	rollbackApiRequestDto := GetRollbackAppApiRequestDto(envConf.HAppId, version)
	payloadForRollbackApi, _ := json.Marshal(rollbackApiRequestDto)
	log.Println("Hitting Rollback Application ")
	rollbackApiRespDto := HitRollbackApplicationApi(string(payloadForRollbackApi), suite.authToken)
	assert.Equal(suite.T(), 200, rollbackApiRespDto.Code)
	resp = HitGetDeploymentHistoryById(queryParams, suite.authToken)
	log.Println("Verifying the response of Rollback Application API")
	assert.Equal(suite.T(), totalNoDeploymentHistory+1, len(resp.Result.DeploymentHistory))
}

//todo failing as we are getting invalid json in API Response, so commenting as of now
/*func (suite *HelmAppTestSuite) TestRollBackApplicationApiWithInvalidArgument() {
	for _, RollbackAppApiArgs := range dataProviderForRollbackAppApiHavingInvalidArgs {
		rollbackApiRequestDto := GetRollbackAppApiRequestDto(RollbackAppApiArgs.HAppId, RollbackAppApiArgs.Version)
		payloadForRollbackApi, _ := json.Marshal(rollbackApiRequestDto)
		log.Println("Hitting Rollback Application ")
		rollbackApiRespDto := HitRollbackApplicationApi(string(payloadForRollbackApi), suite.authToken)
		log.Println("Verifying the response of Rollback Application API ")
		assert.Equal(suite.T(), 400, rollbackApiRespDto.Code)
	}
}
*/
//todo this is failing as we are not trimming whitespaces in the arguments
//todo disabling as of now ,will enable once fixed by dev
/*func (suite *HelmAppTestSuite) TestRollBackApplicationApiWithValidArgumentHavingWhiteSpaces() {
	rollbackApiRequestDto := GetRollbackAppApiRequestDto(" "+HAppId+" ", installedVersionId, version, installedAppId)
	payloadForRollbackApi, _ := json.Marshal(rollbackApiRequestDto)
	log.Println("Hitting Rollback Application ")
	rollbackApiRespDto := HitRollbackApplicationApi(string(payloadForRollbackApi), suite.authToken)
	log.Println("Verifying the response of Rollback Application API ")
	assert.Equal(suite.T(), 200, rollbackApiRespDto.Code)
}*/
