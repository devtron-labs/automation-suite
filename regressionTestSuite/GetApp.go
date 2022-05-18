package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *regressionTestSuite) TestFetchAppGetWithValidAppId() {

	type appId struct {
		appId string `json:"app_id"`
	}
	temp := json.RawMessage(Base.ReadFile())
	var app appId
	json.Unmarshal(temp, &app)

	id, _ := strconv.Atoi(app.appId)

	AppId := map[string]string{
		"id": strconv.Itoa(id),
	}
	fetchAppGetResponseDto := FetchAppGet(AppId, suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 200, fetchAppGetResponseDto.Code)
}
func (suite *regressionTestSuite) TestFetchAppGetWithInvalidAppId() {
	type appId struct {
		appId string `json:"app_id"`
	}
	temp := json.RawMessage(Base.ReadFile())
	var app appId
	json.Unmarshal(temp, &app)

	id, _ := strconv.Atoi(app.appId)

	AppId := map[string]string{
		"id": strconv.Itoa(id),
	}
	fetchAppGetResponseDto := FetchAppGet(AppId, suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 404, fetchAppGetResponseDto.Code)

}
