package HelmAppRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

// TestGetApplicationDetail Test Data should be created already via installing envoy helm chart
func (suite *HelmAppTestSuite) TestGetApplicationDetail() {
	suite.Run("A=1=ApplicationDetailWithValidAppId", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		queryParams := map[string]string{"appId": envConf.HAppId}
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		respHibernateApi := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		errorMessage := respHibernateApi.Result[0].ErrorMessage
		if errorMessage == "object is already scaled down" {
			respOfGetApplicationDetailApi := HitGetApplicationDetailApi(queryParams, suite.authToken)
			assert.Equal(suite.T(), "Hibernated", respOfGetApplicationDetailApi.Result.AppDetail.ApplicationStatus)
			assert.Equal(suite.T(), "deployed", respOfGetApplicationDetailApi.Result.AppDetail.ReleaseStatus.Status)
			assert.Equal(suite.T(), "envoy", respOfGetApplicationDetailApi.Result.AppDetail.ChartMetadata.ChartName)
			assert.Equal(suite.T(), 1, respOfGetApplicationDetailApi.Result.AppDetail.EnvironmentDetails.ClusterId)
			// here we are hitting the UnHibernateWorkloadApi and verifying the status of app
			HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
			assert.True(suite.T(), PollForAppStatus(queryParams, suite.authToken))
		}
		//Un-hibernating again for saving cost
		HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	})

	suite.Run("A=2=GetApplicationDetailWithInvalidAppId", func() {
		queryParams := map[string]string{"appId": "InvalidAppId"}
		respOfGetApplicationDetailApi := HitGetApplicationDetailApi(queryParams, suite.authToken)
		assert.Equal(suite.T(), "malformed app id InvalidAppId", respOfGetApplicationDetailApi.Errors[0].UserMessage)
	})
}

func PollForAppStatus(queryParams map[string]string, authToken string) bool {
	count := 0
	for {
		respOfGetApplicationDetailApi := HitGetApplicationDetailApi(queryParams, authToken)
		deploymentStatus := respOfGetApplicationDetailApi.Result.AppDetail.ApplicationStatus
		time.Sleep(1 * time.Second)
		count = count + 1
		if deploymentStatus == "Healthy" || count >= 100 {
			break
		}
	}
	return true
}
