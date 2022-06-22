package GitopsConfigRouter

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"log"
	"strings"

	"github.com/stretchr/testify/assert"
)

func (suite *GitOpsRouterTestSuite) TestClassA3SaveGitopsConfig() {

	suite.Run("A=1=CreateGitopsConfigWithValidPayload", func() {
		//gitopsConfig, _ := GetGitopsConfig()
		envConf := Base.ReadBaseEnvConfig()
		gitopsConfig := Base.ReadAnyJsonFile(envConf.BaseCredentialsFile)
		createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.GitUsername, gitopsConfig.Host, gitopsConfig.GitToken, gitopsConfig.GitHubOrgId)
		byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

		log.Println("Hitting The post gitops config API")
		createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, gitopsConfig.Provider, gitopsConfig.GitUsername, gitopsConfig.Host, gitopsConfig.GitToken, gitopsConfig.GitHubOrgId, suite.authToken)

		log.Println("Validating the Response of the Create Gitops Config API...")
		assert.Equal(suite.T(), 200, createLinkResponseDto.Code)
		assert.Equal(suite.T(), "Create Repo", createLinkResponseDto.Result.SuccessfulStages[0])
	})

	suite.Run("A=2=CreateGitopsConfigWithInValidProvider", func() {
		provider := Base.GetRandomStringOfGivenLength(10)
		//gitopsConfig, _ := GetGitopsConfig()
		envConf := Base.ReadBaseEnvConfig()
		gitopsConfig := Base.ReadAnyJsonFile(envConf.BaseCredentialsFile)
		createGitopsConfigRequestDto := GetGitopsConfigRequestDto(provider, gitopsConfig.GitUsername, gitopsConfig.Host, gitopsConfig.GitToken, gitopsConfig.GitHubOrgId)
		byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

		log.Println("Hitting The post gitops config API")
		createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, provider, gitopsConfig.GitUsername, gitopsConfig.Host, gitopsConfig.GitToken, gitopsConfig.GitHubOrgId, suite.authToken)

		log.Println("Validating the Response of the Create Gitops Config API...")
		assert.Equal(suite.T(), 0, createLinkResponseDto.Code)
	})
	suite.Run("A=3=CreateGitopsConfigWithValidPayload", func() {
		token := Base.GetRandomStringOfGivenLength(10)
		//gitopsConfig, _ := GetGitopsConfig()
		envConf := Base.ReadBaseEnvConfig()
		gitopsConfig := Base.ReadAnyJsonFile(envConf.BaseCredentialsFile)
		createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.GitUsername, gitopsConfig.Host, token, gitopsConfig.GitHubOrgId)
		byteValueOfStruct, _ := json.Marshal(createGitopsConfigRequestDto)

		log.Println("Hitting The post gitops config API")
		createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfStruct, gitopsConfig.Provider, gitopsConfig.GitUsername, gitopsConfig.Host, token, gitopsConfig.GitHubOrgId, suite.authToken)

		log.Println("Validating the Response of the Create Gitops Config API...")
		assert.Equal(suite.T(), 0, len(createLinkResponseDto.Result.SuccessfulStages))
		assert.True(suite.T(), strings.Contains(createLinkResponseDto.Result.StageErrorMap.ErrorInConnectingWithGITHUB, "401 Bad credentials []"))
	})
}
