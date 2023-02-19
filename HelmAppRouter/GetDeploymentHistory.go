package HelmAppRouter

func (suite *HelmAppTestSuite) TestGetDeploymentHistory() {
	/*suite.Run("A=1=GetDeploymentHistoryWithCorrectAppId", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		queryParams := map[string]string{"appId": envConf.HAppId}
		log.Println("Hitting Get Deployment History API before creating any new deployment")
		resp := HitGetDeploymentHistoryById(queryParams, suite.authToken)
		totalNoDeploymentHistory := len(resp.Result.DeploymentHistory)

		rollbackApiRequestDto := GetRollbackAppApiRequestDto(envConf.HAppId, resp.Result.DeploymentHistory[totalNoDeploymentHistory-1].Version)
		payloadForRollbackApi, _ := json.Marshal(rollbackApiRequestDto)
		rollbackApiRespDto := HitRollbackApplicationApi(string(payloadForRollbackApi), suite.authToken)
		assert.Equal(suite.T(), 200, rollbackApiRespDto.Code)
		log.Println("Hitting Get Deployment History API after creating any new deployment")
		resp = HitGetDeploymentHistoryById(queryParams, suite.authToken)
		log.Println("Verifying the response of the GetDeploymentHistoryApi")
		assert.Equal(suite.T(), totalNoDeploymentHistory+1, len(resp.Result.DeploymentHistory))
	})
	suite.Run("A=2=GetDeploymentHistoryWithIncorrectAppId", func() {
		randomHAppId := testUtils.GetRandomNumberOf9Digit()
		queryParams := map[string]string{"appId": strconv.Itoa(randomHAppId)}
		resp := HitGetDeploymentHistoryById(queryParams, suite.authToken)
		assert.Equal(suite.T(), 400, resp.Code)
		assert.Equal(suite.T(), "Bad Request", resp.Status)
	})*/
}
