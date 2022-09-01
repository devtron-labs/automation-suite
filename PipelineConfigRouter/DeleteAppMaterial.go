package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassDeleteAppMaterial() {
	log.Println("=== Here we are creating an App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appId := createAppApiResponse.Id
	log.Println("=== created App name is ===>", createAppApiResponse.AppName)

	createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, false)
	byteValueOfStruct2, _ := json.Marshal(createAppMaterialRequestDto)
	log.Println("=== Hitting The create material API ===")
	createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct2, appId, 1, false, suite.authToken)

	suite.Run("A=1=DeleteAppMaterialWithInvalidMaterialId", func() {
		invalidMaterialId := Base.GetRandomNumberOf9Digit()
		log.Println("getting payload for Delete material API")
		AppMaterials := createAppMaterialResponseDto.Result.Material[0]
		AppMaterials.Id = invalidMaterialId
		byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(createAppMaterialResponseDto.Result.AppId, AppMaterials)
		log.Println("Hitting the Delete material API for Removing the data created via automation")
		responseOfDeleteMaterialApi := HitDeleteAppMaterialApi(byteValueOfDeleteApp, suite.authToken)
		assert.Equal(suite.T(), responseOfDeleteMaterialApi.Errors[0].UserMessage, "pg: no rows in result set")
		assert.Equal(suite.T(), responseOfDeleteMaterialApi.Code, 404)
	})

	/*suite.Run("A=2=DeleteAppMaterialWithInvalidAppId", func() {
		invalidAppId := Base.GetRandomNumberOf9Digit()
		log.Println("getting payload for Delete material API")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(invalidAppId, createAppMaterialResponseDto.Result.Material[0])
		log.Println("Hitting the Delete material API for Removing the data created via automation")
		responseOfDeleteMaterialApi := HitDeleteAppMaterialApi(byteValueOfDeleteApp, suite.authToken)
		assert.Equal(suite.T(), responseOfDeleteMaterialApi.Errors[0].UserMessage, "pg: no rows in result set")
		assert.Equal(suite.T(), responseOfDeleteMaterialApi.Code, 404)
	})

	suite.Run("A=3=DeleteAppMaterialWithInvalidGitProviderId", func() {
		invalidGitProviderId := Base.GetRandomNumberOf9Digit()
		log.Println("getting payload for Delete material API")
		AppMaterials := createAppMaterialResponseDto.Result.Material[0]
		AppMaterials.GitProviderId = invalidGitProviderId
		byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(appId, AppMaterials)
		log.Println("Hitting the Delete material API for Removing the data created via automation")
		responseOfDeleteMaterialApi := HitDeleteAppMaterialApi(byteValueOfDeleteApp, suite.authToken)
		assert.Equal(suite.T(), responseOfDeleteMaterialApi.Errors[0].UserMessage, "pg: no rows in result set")
		assert.Equal(suite.T(), responseOfDeleteMaterialApi.Code, 404)
	})*/

	suite.Run("A=4=DeleteAppMaterialWithValidPayload", func() {
		log.Println("getting payload for Delete material API")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(createAppMaterialResponseDto.Result.AppId, createAppMaterialResponseDto.Result.Material[0])
		log.Println("Hitting the Delete material API for Removing the data created via automation")
		responseOfDeleteMaterialApi := HitDeleteAppMaterialApi(byteValueOfDeleteApp, suite.authToken)
		assert.Equal(suite.T(), responseOfDeleteMaterialApi.Result, "Git material deleted successfully.")
		log.Println("=== Hitting the Delete material API ===")
		fetchAppGetResponseDto := HitGetApp(createAppMaterialResponseDto.Result.AppId, suite.authToken)
		log.Println("=== Validating the Response of Delete material API ===")
		assert.Nil(suite.T(), fetchAppGetResponseDto.Result.Material)
	})

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId)

	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfDeleteApp, createAppApiResponse.Id, suite.authToken)
}

//todo need to add test cases for other validation once they are handled by Devs
