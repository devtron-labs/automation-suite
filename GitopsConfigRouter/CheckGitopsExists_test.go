package GitopsConfigRouter

import (
	"encoding/json"
	"log"

	"github.com/stretchr/testify/assert"
)

func (suite *GitOpsRouterTestSuite) TestClassA1CheckGitopsExists() {

	suite.Run("A=1=CheckGitopsExistsIsTrue", func() {
		checkGitopsExistsResponseDto := HitGitopsConfigured(suite.authToken)

		log.Println("Validating the response of check gitops Exists API")
		assert.Equal(suite.T(), 200, checkGitopsExistsResponseDto.Code)
		assert.Equal(suite.T(), true, checkGitopsExistsResponseDto.Result.Exists)

	})
	suite.Run("A=2=CheckGitopsExistsIsFalse", func() {
		log.Println("Fetching all gitops configs")
		payload := UpdateGitops(suite.authToken)
		log.Println("Hitting HitGitopsConfigured api ")
		checkGitopsExistsResponseDto := HitGitopsConfigured(suite.authToken)
		log.Println("Validating the response of check gitops Exists API")
		assert.Equal(suite.T(), false, checkGitopsExistsResponseDto.Result.Exists)
		log.Println("Updating gitops to True")
		byteValueOfCreateGitopsConfig, _ := json.Marshal(payload)
		HitUpdateGitopsConfigApi(byteValueOfCreateGitopsConfig, suite.authToken)
		checkGitopsExistsAgainResponseDto := HitGitopsConfigured(suite.authToken)
		log.Println("Rechecking the response of check gitops Exists API")
		assert.Equal(suite.T(), true, checkGitopsExistsAgainResponseDto.Result.Exists)

	})
}
