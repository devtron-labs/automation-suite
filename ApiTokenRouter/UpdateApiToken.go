package ApiTokenRouter

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"time"
)

func (suite *ApiTokenRoutersTestSuite) TestUpdateApiToken() {

	suite.Run("A=1=CreateApiTokenWithValidArgs", func() {
		timeStampBeforeUpdating := time.Now().Unix()
		time.Sleep(5 * time.Second)
		var tokenId int
		createApiTokenRequestDTO := GetPayLoadForCreateApiToken()
		payloadForCreateApiTokenRequest, _ := json.Marshal(createApiTokenRequestDTO)
		responseOfCreateApiToken := HitCreateApiTokenApi(string(payloadForCreateApiTokenRequest), suite.authToken)
		assert.Equal(suite.T(), "API-TOKEN:"+createApiTokenRequestDTO.Name, responseOfCreateApiToken.Result.UserIdentifier)
		responseOfGetAllApiTokens := HitGetAllApiTokens(suite.authToken).Result
		for _, result := range responseOfGetAllApiTokens {
			if result.UserId == responseOfCreateApiToken.Result.UserId {
				assert.Equal(suite.T(), responseOfCreateApiToken.Result.Token, result.Token)
				tokenId = result.Id
			}
		}
		log.Println("=== Here We updating & verifying the Token After creation ===")
		HitUpdateApiToken(strconv.Itoa(tokenId), suite.authToken)
		responseOfGetAllApiTokens = HitGetAllApiTokens(suite.authToken).Result
		var DateStringForUpdateToken string
		for _, result := range responseOfGetAllApiTokens {
			if result.UserId == responseOfCreateApiToken.Result.UserId {

				assert.Equal(suite.T(), responseOfCreateApiToken.Result.Token, result.Token)
				DateStringForUpdateToken = result.UpdatedAt
			}
		}

		timeStampWhenTokenUpdated := ConvertDateStringIntoTimeStamp(DateStringForUpdateToken)
		fmt.Println("Here I am printing timeStampWhenTokenUpdated===>", timeStampWhenTokenUpdated)
		fmt.Println("Here I am printing timeStampBeforeUpdating===>", timeStampBeforeUpdating)
		assert.True(suite.T(), timeStampBeforeUpdating < timeStampWhenTokenUpdated)
		log.Println("=== Here We Deleting the Token After Verification ===")
		responseOfDeleteApi := HitDeleteApiToken(strconv.Itoa(tokenId), suite.authToken)
		assert.True(suite.T(), responseOfDeleteApi.Result.Success)
	})
}

func ConvertDateStringIntoTimeStamp(timeString string) int64 {
	layout := "2006-01-02 15:04:05 -0700 MST"
	t, _ := time.Parse(layout, timeString)
	return t.Unix()
}
