package AppListingRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *AppListingRouter) TestFetchAllStageStatusWithValidAppId() {
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
	fetchAllLinkResponseDto := FetchAllStageStatus(AppId, suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 200, fetchAllLinkResponseDto.Code)

}
func (suite *AppListingRouter) TestFetchAllStageStatusWithInvalidAppId() {
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
	fetchAllLinkResponseDto := FetchAllStageStatus(AppId, suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 404, fetchAllLinkResponseDto.Code)

}
