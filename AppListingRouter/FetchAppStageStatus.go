package AppListingRouter

import (
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *AppsListingRouterTestSuite) TestFetchAllStageStatusWithValidAppId() {
	fetchAllLinkResponseDto := FetchAllStageStatus(Base.ReadDataByFilenameAndKey("", "app_id"), suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 200, fetchAllLinkResponseDto.Code)
	assert.Equal(suite.T(), true, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1])

}
func (suite *AppsListingRouterTestSuite) TestFetchAllStageStatusWithInvalidAppId() {

	fetchAllLinkResponseDto := FetchAllStageStatus(Base.GetRandomStringOfGivenLength(10), suite.authToken)

	log.Println("Validating the response of FetchAllLink API")
	assert.Equal(suite.T(), 404, fetchAllLinkResponseDto.Code)
	assert.Equal(suite.T(), "pg: no rows in result set", fetchAllLinkResponseDto.Errors[0].UserMessage)
}
