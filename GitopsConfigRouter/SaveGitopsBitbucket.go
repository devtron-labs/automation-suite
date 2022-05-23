package GitopsConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *GitOpsRouterTestSuite) TestCreateGitopsConfigWithValidPayload() {
	gitopsConfig, _ := GetGitopsConfig()
	createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
	byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

	log.Println("Hitting The post gitops config API")
	createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 200, createLinkResponseDto.Code)
	assert.Equal(suite.T(), "Create Repo", createLinkResponseDto.Result.SuccessfulStages[0])

}

func (suite *GitOpsRouterTestSuite) TestCreateGitopsConfigWithInValidProvider() {
	provider := Base.GetRandomStringOfGivenLength(10)
	gitopsConfig, _ := GetGitopsConfig()
	createGitopsConfigRequestDto := GetGitopsConfigRequestDto(provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
	byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

	log.Println("Hitting The post gitops config API")
	createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 0, createLinkResponseDto.Code)
}

func (suite *GitOpsRouterTestSuite) TestCreateGitopsConfigWithInValidToken() {
	token := Base.GetRandomStringOfGivenLength(10)
	gitopsConfig, _ := GetGitopsConfig()

	createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, token, gitopsConfig.GitHubOrgId)
	byteValueOfStruct, _ := json.Marshal(createGitopsConfigRequestDto)

	log.Println("Hitting The post gitops config API")
	createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfStruct, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, token, gitopsConfig.GitHubOrgId, suite.authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(suite.T(), 0, len(createLinkResponseDto.Result.SuccessfulStages))
	assert.Equal(suite.T(), "401 Bad credentials []", createLinkResponseDto.Result.StageErrorMap.ErrorInConnectingWithGITHUB[80:102])

}
