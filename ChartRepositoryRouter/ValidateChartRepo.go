package ChartRepositoryRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func (suite *ChartRepoTestSuite) TestClassC5ValidateChartRepo() {

	suite.Run("A=1=ValidateChartRepoWithValidPayload", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := CreateChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		validateChartRepoApiResponse := HitValidateChartRepo(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "Configurations are validated successfully", validateChartRepoApiResponse.Result.CustomErrMsg)
	})
	suite.Run("A=2=ValidateChartRepoWithInvalidChartRepoUrl", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := CreateChartRepoRequestPayload(AUTH_MODE_ANONYMOUS, 0, RepoName, chartRepoConfig.ChartRepoUrl+"InvalidUrl", "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		validateChartRepoApiResponse := HitValidateChartRepo(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "Could not find an index.yaml file in the repo directory. Please try another chart repo.", validateChartRepoApiResponse.Result.CustomErrMsg)
	})
	suite.Run("A=3=ValidateChartRepoWithEmptyValueOfAuthMode", func() {
		chartRepoConfig, _ := GetChartRepoRouterConfig()
		RepoName := Base.GetRandomStringOfGivenLength(8)
		createChartRepoRequestDto := CreateChartRepoRequestPayload("", 0, RepoName, chartRepoConfig.ChartRepoUrl, "", true)
		byteValueOfStruct, _ := json.Marshal(createChartRepoRequestDto)
		validateChartRepoApiResponse := HitValidateChartRepo(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "Key: 'ChartRepoDto.AuthMode' Error:Field validation for 'AuthMode' failed on the 'required' tag", validateChartRepoApiResponse.Errors[0].InternalMessage)
	})
}

//todo value of "accessToken/authMode" key doesn't matter as it is working for any random value
