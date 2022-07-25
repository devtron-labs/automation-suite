package AppStoreDeploymentRouter

import (
	"automation-suite/AppStoreDeploymentRouter/RequestDTOs"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"time"
)

func (suite *AppStoreDeploymentTestSuite) TestGetInstalledAppVersion() {
	log.Println("=== Here We are installing Helm chart from chart-store ===")
	expectedPayload, _ := Base.GetByteArrayOfGivenJsonFile("../testdata/AppStoreRouter/InstallAppRequestPayload.json")
	log.Println("Hitting the InstallAppApi with valid payload")
	installAppRequestDTO := RequestDTOs.InstallAppRequestDTO{}
	json.Unmarshal(expectedPayload, &installAppRequestDTO)
	installAppRequestDTO.AppName = "deepak-helm-airflow" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
	requestPayload, _ := json.Marshal(installAppRequestDTO)
	responseAfterInstallingApp := HitInstallAppApi(string(requestPayload), suite.authToken)
	time.Sleep(2 * time.Second)
	installedAppVersionId := responseAfterInstallingApp.Result.InstalledAppVersionId

	suite.Run("A=1=GetDetailsWithCorrectAppId", func() {
		installedAppVersion := HitGetInstalledAppVersionApi(strconv.Itoa(installedAppVersionId), suite.authToken)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.AppName, installedAppVersion.Result.AppName)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.AppStoreVersion, installedAppVersion.Result.AppStoreVersion)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.GitOpsRepoName, installedAppVersion.Result.GitOpsRepoName)
		assert.Equal(suite.T(), responseAfterInstallingApp.Result.ValuesOverrideYaml, installedAppVersion.Result.ValuesOverrideYaml)
	})

	suite.Run("A=2=GetDetailsWithIncorrectAppId", func() {
		randomAppId := Base.GetRandomNumberOf9Digit()
		installedAppVersion := HitGetInstalledAppVersionApi(strconv.Itoa(randomAppId), suite.authToken)
		assert.Equal(suite.T(), 404, installedAppVersion.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", installedAppVersion.Error[0].UserMessage)
	})

	log.Println("Removing the data created via API")
	respOfDeleteInstallAppApi := HitDeleteInstalledAppApi(strconv.Itoa(responseAfterInstallingApp.Result.InstalledAppId), suite.authToken)
	assert.Equal(suite.T(), responseAfterInstallingApp.Result.InstalledAppId, respOfDeleteInstallAppApi.Result.InstalledAppId)
}
