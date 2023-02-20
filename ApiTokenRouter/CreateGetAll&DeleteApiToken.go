package ApiTokenRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
)

func (suite *ApiTokenRoutersTestSuite) TestCreateGetAllAndDeleteApiToken() {

	suite.Run("A=1=CreateApiTokenWithValidArgs", func() {
		var tokenId int
		createApiTokenRequestDTO := GetPayLoadForCreateApiToken()
		payloadForCreateApiTokenRequest, _ := json.Marshal(createApiTokenRequestDTO)
		responseOfCreateApiToken := HitCreateApiTokenApi(string(payloadForCreateApiTokenRequest), suite.authToken)
		assert.Equal(suite.T(), "API-TOKEN:"+createApiTokenRequestDTO.Name, responseOfCreateApiToken.Result.UserIdentifier)
		responseOfGetAllApiTokens := HitGetAllApiTokens(suite.authToken).Result
		for _, result := range responseOfGetAllApiTokens {
			if result.UserId == responseOfCreateApiToken.Result.UserId {
				assert.Equal(suite.T(), responseOfCreateApiToken.Result.Token, result.Token)
				assert.Equal(suite.T(), createApiTokenRequestDTO.ExpireAtInMs, result.ExpireAtInMs)
				assert.Equal(suite.T(), createApiTokenRequestDTO.Name, result.Name)
				tokenId = result.Id
			}
		}
		log.Println("=== Here We Deleting the Token After Verification")
		responseOfDeleteApi := HitDeleteApiToken(strconv.Itoa(tokenId), suite.authToken)
		assert.True(suite.T(), responseOfDeleteApi.Result.Success)
	})

	suite.Run("A=2=CreateApiTokenWithExistingName", func() {
		var tokenId int
		createApiTokenRequestDTO := GetPayLoadForCreateApiToken()

		payloadForCreateApiTokenRequest, _ := json.Marshal(createApiTokenRequestDTO)
		responseOfCreateApiToken := HitCreateApiTokenApi(string(payloadForCreateApiTokenRequest), suite.authToken)
		log.Println("=== Here we are getting All the API token for finding the TokenID ===")
		responseOfGetAllApiTokens := HitGetAllApiTokens(suite.authToken).Result
		for _, result := range responseOfGetAllApiTokens {
			if result.UserId == responseOfCreateApiToken.Result.UserId {
				assert.Equal(suite.T(), responseOfCreateApiToken.Result.Token, result.Token)
				tokenId = result.Id
			}
		}
		log.Println("=== Here we are hitting the CreateToken Api with existing name ===")
		responseOfCreateApiToken = HitCreateApiTokenApi(string(payloadForCreateApiTokenRequest), suite.authToken)
		assert.True(suite.T(), strings.Contains(responseOfCreateApiToken.Errors[0].UserMessage, "is already used. please use another name"))
		log.Println("=== Here We Deleting the Token After Verification ===")
		responseOfDeleteApi := HitDeleteApiToken(strconv.Itoa(tokenId), suite.authToken)
		assert.True(suite.T(), responseOfDeleteApi.Result.Success)
	})
}
