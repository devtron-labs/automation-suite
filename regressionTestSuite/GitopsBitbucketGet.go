package regressionTestSuite

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *regressionTestSuite) TestFetchAllGitopsConfig() {
	log.Println("Hitting GET api for Gitops config")
	fetchAllLinkResponseDto := HitFetchAllGitopsConfigApi()
	noOfTeams := len(fetchAllLinkResponseDto.Result)

	log.Println("Hitting the 'Save Gitops Config' Api for creating a new entry")
	gitopsConfig, _ := GetGitopsConfig()

	createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
	byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

	log.Println("Hitting The post gitops config API")
	createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, suite.authToken)

	log.Println("Hitting the HitFetchAllGitopsConfigApi again for verifying the functionality of it")
	fetchAllLinkResponseDto = HitFetchAllGitopsConfigApi()

	log.Println("Validating the response of FetchAllLink API")

	// as response is not sending id or any parameter we are using if else using return code
	if createLinkResponseDto.Code == 200 {
		assert.Equal(suite.T(), noOfTeams+1, len(fetchAllLinkResponseDto.Result))
		assert.Equal(suite.T(), createGitopsConfigRequestDto.Id, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Id)
		assert.Equal(suite.T(), createGitopsConfigRequestDto.Provider, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Provider)
	}

	log.Println("getting payload for Delete GitOps Config API")
	byteValueOfDeleteApi := GetPayLoadForDeleteGitopsConfigAPI(fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Id, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Provider, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Username, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Host, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Token)
	log.Println("Hitting the Delete link API for Removing the data created via automation")
	HitDeleteLinkApi(byteValueOfDeleteApi, suite.authToken)
}
