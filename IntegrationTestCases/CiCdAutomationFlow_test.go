package IntegrationTestCases

import (
	"automation-suite/DockerRegRouter"
	"automation-suite/GitopsConfigRouter"
	"automation-suite/TeamRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"strconv"
	"testing"
)

func TestFoo(t *testing.T) {
	authToken := ""
	t.Run("TestSSOLoginRouterSuite", func(t *testing.T) {
		suite.Run(t, new(IntegrationTestCases))
	})
	t.Run("TestFetchAllGitopsConfig", func(t *testing.T) {
		log.Println("Hitting GET api for Gitops config")
		fetchAllLinkResponseDto := GitopsConfigRouter.HitFetchAllGitopsConfigApi(authToken)
		noOfGitopsConfig := len(fetchAllLinkResponseDto.Result)

		log.Println("Hitting the 'Save Gitops Config' Api for creating a new entry")
		gitopsConfig, _ := GitopsConfigRouter.GetGitopsConfig()

		createGitopsConfigRequestDto := GitopsConfigRouter.GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
		byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

		log.Println("Hitting The post gitops config API")
		createLinkResponseDto := GitopsConfigRouter.HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, authToken)

		log.Println("Hitting the HitFetchAllGitopsConfigApi again for verifying the functionality of it")
		fetchAllLinkResponseDto = GitopsConfigRouter.HitFetchAllGitopsConfigApi(authToken)

		log.Println("Validating the response of FetchAllLink API")

		// as response is not sending id or any parameter we are using if else using return code
		if createLinkResponseDto.Code == 200 {
			assert.Equal(t, noOfGitopsConfig+1, len(fetchAllLinkResponseDto.Result))
			assert.Equal(t, createGitopsConfigRequestDto.Provider, fetchAllLinkResponseDto.Result[len(fetchAllLinkResponseDto.Result)-1].Provider)
		}
	})
	t.Run("TestCreateGitopsConfigWithValidPayload", func(t *testing.T) {
		gitopsConfig, _ := GitopsConfigRouter.GetGitopsConfig()
		createGitopsConfigRequestDto := GitopsConfigRouter.GetGitopsConfigRequestDto(gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId)
		byteValueOfCreateGitopsConfig, _ := json.Marshal(createGitopsConfigRequestDto)

		log.Println("Hitting The post gitops config API")
		createLinkResponseDto := GitopsConfigRouter.HitCreateGitopsConfigApi(byteValueOfCreateGitopsConfig, gitopsConfig.Provider, gitopsConfig.Username, gitopsConfig.Host, gitopsConfig.Token, gitopsConfig.GitHubOrgId, authToken)

		log.Println("Validating the Response of the Create Gitops Config API...")
		assert.Equal(t, "Create Repo", createLinkResponseDto.Result.SuccessfulStages[0])
	})
	t.Run("TestSaveTeamWithValidPayload", func(t *testing.T) {
		saveTeamRequestDto := TeamRouter.GetSaveTeamRequestDto()
		byteValueOfStruct, _ := json.Marshal(saveTeamRequestDto)

		log.Println("Hitting The Save Team API")
		saveTeamResponseDto := TeamRouter.HitSaveTeamApi(byteValueOfStruct, authToken)

		//for team autocomplete
		fetchAllTeamResponseDto := TeamRouter.HitFetchAllTeamApi(authToken)
		Base.CreateFileAndEnterData("team", "noOfTeams", strconv.Itoa(len(fetchAllTeamResponseDto.Result)))
		Base.CreateFileAndEnterData("team", "teamName", saveTeamResponseDto.Result.Name)

		log.Println("Validating the Response of the Save API...")
		assert.Equal(t, saveTeamRequestDto.Name, saveTeamResponseDto.Result.Name)
		assert.NotNil(t, saveTeamResponseDto.Result.Id)
	})
	t.Run("TestSaveDockerRegistryWithValidPayload", func(t *testing.T) {
		saveDockerRegistryRequestDto := DockerRegRouter.GetDockerRegistryRequestDto(false, "", "", "", "", false, "", "")
		byteValueOfSaveDockerRegistry, _ := json.Marshal(saveDockerRegistryRequestDto)

		log.Println("Hitting The post Docker registry API")
		saveDockerRegistryResponseDto := DockerRegRouter.HitSaveDockerRegistryApi(false, byteValueOfSaveDockerRegistry, "", "", "", "", "", "", false, authToken)

		log.Println("Validating the Response of the save docker registry API...")
		assert.Equal(t, saveDockerRegistryRequestDto.Id, saveDockerRegistryResponseDto.Result.Id)

	})
	t.Run("TestFetchForAutocompleteApiWithValidPayload", func(t *testing.T) {

		log.Println("Hitting the FetchForAutocomplete API again for verifying the functionality of it")
		fetchAllTeamResponseDto := TeamRouter.HitFetchAllTeamApi(authToken)
		log.Println("Validating the response of FetchForAutocomplete API")
		noOfTeams, _ := strconv.Atoi(Base.ReadDataByFilenameAndKey("team", "noOfTeams"))
		assert.Equal(t, noOfTeams+1, len(fetchAllTeamResponseDto.Result))
		assert.Equal(t, Base.ReadDataByFilenameAndKey("team", "teamName"), fetchAllTeamResponseDto.Result[len(fetchAllTeamResponseDto.Result)-1].Name)
	})
	t.Run("B=1", func(t *testing.T) {

	})
	// <tear-down code>
}
