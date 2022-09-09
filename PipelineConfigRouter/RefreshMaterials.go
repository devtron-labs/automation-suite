package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassRefreshMaterials() {
	log.Println("=== Here we are creating an App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appId := createAppApiResponse.Id
	log.Println("=== created App name is ===>", createAppApiResponse.AppName)
	createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, false)
	byteValueOfStruct, _ := json.Marshal(createAppMaterialRequestDto)
	log.Println("=== Here we are creating app material ===")
	appMaterial := HitCreateAppMaterialApi(byteValueOfStruct, appId, 1, false, suite.authToken).Result.Material[0]

	suite.Run("A=1=RefreshAppMaterialsWithCorrectGitMaterialId", func() {
		currentTimestamp := time.Now().Unix()
		time.Sleep(5 * time.Second)
		refreshMaterialResponse := HitRefreshMaterialsApi(strconv.Itoa(appMaterial.Id), suite.authToken)
		assert.Equal(suite.T(), "successfully refreshed material", refreshMaterialResponse.Result.Message)
		ts := Base.ConvertDateStringIntoTimeStamp(refreshMaterialResponse.Result.LastFetchTime)
		fmt.Println("Here I am printing currentTimestamp===>", currentTimestamp)
		fmt.Println("Here I am printing TimeStampDuringRefresh===>", ts)
		assert.True(suite.T(), currentTimestamp < ts)
	})

	suite.Run("A=2=RefreshAppMaterialsWithIncorrectGitMaterialId", func() {
		randomGitMaterialId := strconv.Itoa(Base.GetRandomNumberOf9Digit())
		refreshMaterialResponse := HitRefreshMaterialsApi(randomGitMaterialId, suite.authToken)
		assert.Equal(suite.T(), 404, refreshMaterialResponse.Code)
		assert.Equal(suite.T(), "pg: no rows in result set", refreshMaterialResponse.Errors[0].UserMessage)
	})

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId)

	log.Println("Hitting the Delete App API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfDeleteApp, createAppApiResponse.Id, suite.authToken)
}
