package IntegrationTestCases

import (
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"strings"
	"testing"
)

func TestFoo(t *testing.T) {
	authToken := ""
	noOfteams := 0
	teamName := ""
	t.Run("TestSSOLoginRouterSuite", func(t *testing.T) {
		suite.Run(t, new(CiCdAutomationFlow))
	})
	t.Run("TestFetchAllGitopsConfig", func(t *testing.T) {
		log.Println("Hitting GET api for Gitops config")
		fetchAllLinkResponseDto := HitFetchAllGitopsConfigApi(authToken)
		noOfGitopsConfig := len(fetchAllLinkResponseDto.Result)

		log.Println("Hitting the 'Save Gitops Config' Api for creating a new entry")
		gitopsConfig, _ := GetGitopsConfig()

		createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
		byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

		log.Println("Hitting The post gitops config API")
		createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, authToken)

		log.Println("Hitting the HitFetchAllGitopsConfigApi again for verifying the functionality of it")
		fetchAllLinkResponseDto = HitFetchAllGitopsConfigApi(authToken)

		log.Println("Validating the response of FetchAllLink API")

		// as response is not sending id or any parameter we are using if else using return code
		if createLinkResponseDto.Code == 200 {
			assert.Equal(t, noOfGitopsConfig+1, len(fetchAllLinkResponseDto.Result))
			assert.Equal(t, createGitopsConfigRequestDto.Provider, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Provider)
		}
	})
	t.Run("TestCreateGitopsConfigWithValidPayload", func(t *testing.T) {
		gitopsConfig, _ := GetGitopsConfig()
		createGitopsConfigRequestDto := GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
		byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

		log.Println("Hitting The post gitops config API")
		createLinkResponseDto := HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, authToken)

		log.Println("Validating the Response of the Create Gitops Config API...")
		assert.Equal(t, 200, createLinkResponseDto.Code)
		assert.Equal(t, "Create Repo", createLinkResponseDto.Result.SuccessfulStages[0])
	})
	t.Run("TestSaveTeamWithValidPayload", func(t *testing.T) {
		name := strings.ToLower(Base.GetRandomStringOfGivenLength(10))
		saveTeamRequestDto := GetSaveTeamRequestDto()
		byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)

		log.Println("Hitting The Save Team API")
		saveTeamResponseDto := HitCreateTeamApi(byteValueOfStruct, name, true, authToken)
		//for team autocomplete
		fetchAllTeamResponseDto := HitFetchAllTeamApi(authToken)
		noOfteams = len(fetchAllTeamResponseDto.Result)
		teamName = saveTeamResponseDto.Result.Name
		log.Println("Validating the Response of the Save API...")
		assert.Equal(t, saveTeamRequestDto.Name, saveTeamResponseDto.Result.Name)
		assert.NotNil(t, saveTeamResponseDto.Result.Id)
	})
	t.Run("TestSaveDockerRegistryWithValidPayload", func(t *testing.T) {
		saveDockerRegistryRequestDto := GetDockerRegistryRequestDto(false, "", "", "", "", false, "", "")
		byteValueOfSaveDockerRegistry, _ := json.Marshal(saveDockerRegistryRequestDto)

		log.Println("Hitting The post Docker registry API")
		saveDockerRegistryResponseDto := HitSaveDockerRegistryApi(false, byteValueOfSaveDockerRegistry, "", "", "", "", "", "", false, authToken)

		log.Println("Validating the Response of the save docker registry API...")
		assert.Equal(t, saveDockerRegistryRequestDto.Id, saveDockerRegistryResponseDto.Result.Id)

	})
	t.Run("TestFetchForAutocompleteApiWithValidPayload", func(t *testing.T) {

		log.Println("Hitting the FetchForAutocomplete API again for verifying the functionality of it")
		fetchAllTeamResponseDto := HitFetchAllTeamApi(authToken)
		log.Println("Validating the response of FetchForAutocomplete API")
		assert.Equal(t, noOfteams+1, len(fetchAllTeamResponseDto.Result))
		assert.Equal(t, teamName, fetchAllTeamResponseDto.Result[len(fetchAllTeamResponseDto.Result)-1].Name)

	})
	t.Run("B=1", func(t *testing.T) {

	})
	// <tear-down code>
}
