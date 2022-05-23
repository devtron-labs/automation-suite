package PipelineConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *PipelinesConfigRouterTestSuite) TestClass5CreateMaterial() {
	appId := suite.createAppResponseDto.Result.Id

	suite.Run("A=1=CreateAppMaterialWithValidPayloadAndFetchSubmodulesFalse", func() {
		createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, false)
		byteValueOfStruct2, _ := json.Marshal(createAppMaterialRequestDto)
		log.Println("Hitting The create material API")
		createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct2, appId, 1, false, suite.authToken)

		log.Println("Validating the Response of the Create material API...")
		assert.Equal(suite.T(), appId, createAppMaterialResponseDto.Result.AppId)
		fetchAppGetResponseDto := HitGetMaterial(createAppMaterialResponseDto.Result.AppId, suite.authToken)
		log.Println("Validating the Response of the Get material API...")
		assert.Equal(suite.T(), createAppMaterialResponseDto.Result.AppId, fetchAppGetResponseDto.Result.Id)
		assert.Equal(suite.T(), false, fetchAppGetResponseDto.Result.Material[0].FetchSubmodules)

		log.Println("getting payload for Delete material API")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(createAppMaterialResponseDto.Result.AppId, createAppMaterialResponseDto.Result.Material[0])
		log.Println("Hitting the Delete material API for Removing the data created via automation")
		HitDeleteAppMaterialApi(byteValueOfDeleteApp, suite.authToken)
	})
	suite.Run("A=2=CreateAppMaterialWithValidPayloadAndFetchSubmodulesTrue", func() {
		createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, true)
		byteValueOfStruct, _ := json.Marshal(createAppMaterialRequestDto)
		log.Println("Hitting The create material API")
		createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct, appId, 1, false, suite.authToken)

		log.Println("Validating the Response of the Create material API...")
		assert.Equal(suite.T(), appId, createAppMaterialResponseDto.Result.AppId)
		fetchAppGetResponseDto := HitGetMaterial(createAppMaterialResponseDto.Result.AppId, suite.authToken)
		log.Println("Validating the Response of the Get material API...")
		assert.Equal(suite.T(), createAppMaterialResponseDto.Result.AppId, fetchAppGetResponseDto.Result.Id)
		assert.Equal(suite.T(), true, fetchAppGetResponseDto.Result.Material[0].FetchSubmodules)

		log.Println("getting payload for Delete material API")
		byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(createAppMaterialResponseDto.Result.AppId, createAppMaterialResponseDto.Result.Material[0])
		log.Println("Hitting the Delete material API for Removing the data created via automation")
		HitDeleteAppMaterialApi(byteValueOfDeleteApp, suite.authToken)
	})

	suite.Run("A=3=CreateAppMaterialWithInvalidGitProviderId", func() {
		gitProviderID := Base.GetRandomNumberOf9Digit()
		createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, gitProviderID, false)
		byteValueOfStruct, _ := json.Marshal(createAppMaterialRequestDto)
		log.Println("Hitting The create material API")
		createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct, appId, gitProviderID, false, suite.authToken)

		log.Println("Validating the Response of the Create material API...")
		assert.Equal(suite.T(), "pg: no rows in result set", createAppMaterialResponseDto.Errors[0].UserMessage)
	})

	suite.Run("A=4=CreateAppMaterialWithInvalidCheckoutPath", func() {
		createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, 1, false)
		createAppMaterialRequestDto.Materials[0].CheckoutPath = Base.GetRandomStringOfGivenLength(5)
		byteValueOfStruct2, _ := json.Marshal(createAppMaterialRequestDto)
		log.Println("Hitting The create material API")
		createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct2, appId, 1, false, suite.authToken)
		log.Println("Validating the Response of the Create material API...")
		assert.Equal(suite.T(), "Key: 'CreateMaterialDTO.Material[0].CheckoutPath' Error:Field validation for 'CheckoutPath' failed on the 'checkout-path-component' tag", createAppMaterialResponseDto.Errors[0].UserMessage)
	})

	// add testcase for ./path
}
