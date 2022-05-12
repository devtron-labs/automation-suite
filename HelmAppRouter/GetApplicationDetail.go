package HelmAppRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

//Test Data should be created already via installing envoy helm chart
func (suite *HelmAppTestSuite) TestGetApplicationDetailWithValidAppId() {
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

		HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		respOfGetApplicationDetailApi = HitGetApplicationDetailApi(queryParams, suite.authToken)
		time.Sleep(2 * time.Second)
		assert.Equal(suite.T(), "Healthy", respOfGetApplicationDetailApi.Result.AppDetail.ApplicationStatus)
	}
}

func (suite *HelmAppTestSuite) TestGetApplicationDetailWithInvalidAppId() {
	queryParams := map[string]string{"appId": "InvalidAppId"}
	respOfGetApplicationDetailApi := HitGetApplicationDetailApi(queryParams, suite.authToken)
	assert.Equal(suite.T(), "malformed app id InvalidAppId", respOfGetApplicationDetailApi.Errors[0].UserMessage)
}
