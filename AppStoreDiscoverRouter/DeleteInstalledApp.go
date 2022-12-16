package AppStoreDiscoverRouter

//func (suite *AppStoreDiscoverTestSuite) TestDeleteInstalledApp() {
//	log.Println("=== Here We are installing Helm chart from chart-store ===")
//	expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
//	log.Println("Hitting the InstallAppApi with valid payload")
//	installAppRequestDTO := RequestDTOs.InstallAppRequestDTO{}
//	json.Unmarshal(expectedPayload, &installAppRequestDTO)
//	installAppRequestDTO.AppName = "deepak-helm-apache" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
//	requestPayload, _ := json.Marshal(installAppRequestDTO)
//	responseAfterInstallingApp := HitInstallAppApi(string(requestPayload), suite.authToken)
//	time.Sleep(2 * time.Second)
//	installedAppVersionId := responseAfterInstallingApp.Result.InstalledAppVersionId
//
//	suite.Run("A=1=DeleteWithCorrectAppId", func() {
//		log.Println("=== Here We are getting installed App versionId ===")
//		installedAppVersion := HitGetInstalledAppVersionApi(strconv.Itoa(installedAppVersionId), suite.authToken)
//		assert.Equal(suite.T(), responseAfterInstallingApp.Result.AppName, installedAppVersion.Result.AppName)
//		log.Println("=== Here We are Deleting installed App ===")
//		respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
//		assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
//		log.Println("=== Here We are getting installed App versionId again for verifying the response of Delete API ===")
//		installedAppVersion = HitGetInstalledAppVersionApi(strconv.Itoa(installedAppVersionId), suite.authToken)
//		assert.Equal(suite.T(), 404, installedAppVersion.Code)
//		assert.Equal(suite.T(), "pg: no rows in result set", installedAppVersion.Error[0].UserMessage)
//	})
//
//	suite.Run("A=2=DeleteWithIncorrectAppId", func() {
//		randomAppId := Base.GetRandomNumberOf9Digit()
//		respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(randomAppId), suite.authToken)
//		assert.Equal(suite.T(), 404, respOfDeleteInstallAppApi.Code)
//		assert.Equal(suite.T(), "pg: no rows in result set", respOfDeleteInstallAppApi.Errors[0].UserMessage)
//	})
//}
