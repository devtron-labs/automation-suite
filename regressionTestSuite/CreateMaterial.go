package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *regressionTestSuite) TestCreateAppMaterialWithValidPayload() {
	appName := Base.GetRandomStringOfGivenLength(10)
	createAppRequestDto := GetAppRequestDto(appName, 1, 0)
	byteValueOfCreateApp, _ := json.Marshal(createAppRequestDto)

	log.Println("Hitting The post team API")
	createAppResponseDto := HitCreateAppApi(byteValueOfCreateApp, appName, 1, 0, suite.authToken)

	appId := createAppResponseDto.Result.Id
	gitopsConfig, _ := GetGitopsConfig()

	createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, gitopsConfig.Url, 1, false)
	byteValueOfStruct2, _ := json.Marshal(createAppMaterialRequestDto)
	log.Println("Hitting The post team API")
	createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct2, appId, gitopsConfig.Url, 1, false, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), appId, createAppMaterialResponseDto.Result.AppId)

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(createAppMaterialResponseDto.Result.AppId, createAppMaterialResponseDto.Result.Material[0])
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppMaterialApi(byteValueOfDeleteApp, suite.authToken)
}
