package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClassUpdateAndDeleteAppMaterial() {
	log.Println("=== Here we are creating an App ===")
	createAppApiResponse := Base.CreateApp(suite.authToken).Result
	appId := createAppApiResponse.Id
	log.Println("=== created App name is ===>", createAppApiResponse.AppName)
	createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, false)
	byteValueOfStruct, _ := json.Marshal(createAppMaterialRequestDto)
	log.Println("=== Here we are creating app material ===")
	appMaterial := HitCreateAppMaterialApi(byteValueOfStruct, appId, 1, false, suite.authToken).Result.Material[0]

	suite.Run("A=1=UpdateCheckoutPath", func() {
		requestDTOForUpdateAppMaterial := GetPayloadForUpdateAppMaterial(appId, appMaterial.Url, appMaterial.Id, appMaterial.GitProviderId, "./test", false)
		byteValueOfRequestPayload, _ := json.Marshal(requestDTOForUpdateAppMaterial)
		responseOfUpdateAppMaterialApi := HitUpdateAppMaterialApi(byteValueOfRequestPayload, suite.authToken)
		assert.Equal(suite.T(), responseOfUpdateAppMaterialApi.Result.Material.CheckoutPath, "./test")
		log.Println("Validating the Response of the update material API...")
		fetchAppGetResponseDto := HitGetMaterial(createAppApiResponse.Id, suite.authToken)
		len := len(fetchAppGetResponseDto.Result.Material)
		assert.Equal(suite.T(), fetchAppGetResponseDto.Result.Material[len-1].CheckoutPath, "./test")
	})

	suite.Run("A=2=UpdateUrl", func() {
		updatedRepoUrl := "https://github.com/devtron-labs/getting-started-nodejs.git"
		requestDTOForUpdateAppMaterial := GetPayloadForUpdateAppMaterial(appId, updatedRepoUrl, appMaterial.Id, appMaterial.GitProviderId, "./test", false)
		byteValueOfRequestPayload, _ := json.Marshal(requestDTOForUpdateAppMaterial)
		responseOfUpdateAppMaterialApi := HitUpdateAppMaterialApi(byteValueOfRequestPayload, suite.authToken)
		assert.Equal(suite.T(), responseOfUpdateAppMaterialApi.Result.Material.Url, updatedRepoUrl)
		log.Println("Validating the Response of the update material API...")
		fetchAppGetResponseDto := HitGetMaterial(createAppApiResponse.Id, suite.authToken)
		len := len(fetchAppGetResponseDto.Result.Material)
		assert.Equal(suite.T(), fetchAppGetResponseDto.Result.Material[len-1].Url, updatedRepoUrl)
	})

	suite.Run("A=3=UpdateUrlForInvalidAppId", func() {
		invalidAppId := Base.GetRandomNumberOf9Digit()
		updatedRepoUrl := "https://github.com/devtron-labs/getting-started-nodejs.git"
		requestDTOForUpdateAppMaterial := GetPayloadForUpdateAppMaterial(invalidAppId, updatedRepoUrl, appMaterial.Id, appMaterial.GitProviderId, "./test", false)
		byteValueOfRequestPayload, _ := json.Marshal(requestDTOForUpdateAppMaterial)
		responseOfUpdateAppMaterialApi := HitUpdateAppMaterialApi(byteValueOfRequestPayload, suite.authToken)
		assert.Equal(suite.T(), responseOfUpdateAppMaterialApi.Errors[0].UserMessage, "material to be updated does not exist")
	})

	suite.Run("A=4=UpdateUrlForInvalidAppMaterialId", func() {
		invalidAppMaterialId := Base.GetRandomNumberOf9Digit()
		updatedRepoUrl := "https://github.com/devtron-labs/getting-started-nodejs.git"
		requestDTOForUpdateAppMaterial := GetPayloadForUpdateAppMaterial(appId, updatedRepoUrl, invalidAppMaterialId, appMaterial.GitProviderId, "./test", false)
		byteValueOfRequestPayload, _ := json.Marshal(requestDTOForUpdateAppMaterial)
		responseOfUpdateAppMaterialApi := HitUpdateAppMaterialApi(byteValueOfRequestPayload, suite.authToken)
		assert.Equal(suite.T(), responseOfUpdateAppMaterialApi.Errors[0].UserMessage, "material to be updated does not exist")
	})

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppAPI(createAppApiResponse.Id, createAppApiResponse.AppName, createAppApiResponse.TeamId, createAppApiResponse.TemplateId)

	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppApi(byteValueOfDeleteApp, createAppApiResponse.Id, suite.authToken)
}
