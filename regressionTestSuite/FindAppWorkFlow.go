package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *regressionTestSuite) TestFetchAllAppWorkflowWithValidAppId() {
	AppId := map[string]string{
		"id": Base.ReadDataByFilenameAndKey("createApp", "app_id"),
	}
	fetchAllAppWorkflowResponseDto := FetchAllAppWorkflow(AppId, suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 200, fetchAllAppWorkflowResponseDto.Code)
}
func (suite *regressionTestSuite) TestFetchAllAppWorkflowWithInvalidAppId() {
	AppId := map[string]string{
		"id": Base.ReadDataByFilenameAndKey("createApp", "app_id"),
	}
	fetchAllAppWorkflowResponseDto := FetchAllAppWorkflow(AppId, suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 404, fetchAllAppWorkflowResponseDto.Code)

}
