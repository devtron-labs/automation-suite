package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *regressionTestSuite) TestFetchOtherEnvWithValidAppId() {

	AppId := map[string]string{
		"id": Base.ReadDataByFilenameAndKey("createApp", "app_id"),
	}

	fetchOtherEnvResponseDto := FetchOtherEnv(AppId, suite.authToken)
	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 200, fetchOtherEnvResponseDto.Code)
}
func (suite *regressionTestSuite) TestFetchOtherEnvWithInvalidAppId() {
	AppId := map[string]string{
		"id": Base.ReadDataByFilenameAndKey("createApp", "app_id"),
	}
	fetchOtherEnvResponseDto := FetchOtherEnv(AppId, suite.authToken)
	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 404, fetchOtherEnvResponseDto.Code)

}
