package regressionTestSuite

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *regressionTestSuite) TestCreateGitopsConfigWithValidPayload() {
	gitopsConfig, _ := GetGitopsConfig()
	createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
	byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

	log.Println("Hitting The post gitops config API")
	createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 200, createLinkResponseDto.Code)

	log.Println("getting payload for Delete GitOps Config API")
	byteValueOfDeleteApi := GetPayLoadForDeleteGitopsConfigAPI(createGitopsConfigRequestDto.Id, createGitopsConfigRequestDto.Provider, createGitopsConfigRequestDto.Username, createGitopsConfigRequestDto.Host, createGitopsConfigRequestDto.Token)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfDeleteApi, suite.authToken)

}

func (suite *regressionTestSuite) TestCreateGitopsConfigWithInValidProvider() {
	provider := Base.GetRandomStringOfGivenLength(10)
	gitopsConfig, _ := GetGitopsConfig()
	createGitopsConfigRequestDto := GetGitopsConfigRequestDto(provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
	byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

	log.Println("Hitting The post gitops config API")
	createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 502, createLinkResponseDto.Code)
}

func (suite *regressionTestSuite) TestCreateGitopsConfigWithInValidToken() {
	token := Base.GetRandomStringOfGivenLength(10)
	gitopsConfig, _ := GetGitopsConfig()

	createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, token, gitopsConfig.GitHubOrgId)
	byteValueOfStruct, _ := json.Marshal(createGitopsConfigRequestDto)

	log.Println("Hitting The post gitops config API")
	createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfStruct, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, token, gitopsConfig.GitHubOrgId, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 0, len(createLinkResponseDto.Result.SuccessfulStages))
	assert.Equal(suite.T(), false, createLinkResponseDto.Result.DeleteRepoFailed)

}
