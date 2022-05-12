package HelmAppRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

/*Test Data ExternalApp should be deployed in default cluster and group should be apps*/
func (suite *HelmAppTestSuite) TestHibernateWorkloadApiWithValidArgsInPayload() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	errorMessage := resp.Result[0].ErrorMessage
	if errorMessage == "object is already scaled down" {
		assert.False(suite.T(), resp.Result[0].Success)
		//Unhibernating the Workload First
		HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		time.Sleep(2 * time.Second)
		//Again Hibernating and verifying the response
		HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "", resp.Result[0].ErrorMessage)
		assert.True(suite.T(), resp.Result[0].Success)
	}
}

func (suite *HelmAppTestSuite) TestHibernateApiWithInvalidKind() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "InvalidDeployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "no matches for kind \"InvalidDeployment\" in version \"apps/v1\"", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestHibernateApiWithInvalidName() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", "InvalidName", "v1", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "deployments.apps \"InvalidName\" not found", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestHibernateApiWithInvalidGroup() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "Invalid", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "no matches for kind \"Deployment\" in version \"Invalid/v1\"", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestHibernateApiWithInvalidVersion() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "Invalid", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "no matches for kind \"Deployment\" in version \"apps/Invalid\"", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestHibernateApiWithInvalidNamespace() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "Invalid")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "deployments.apps \"envoy-deepak-testing-v1\" not found", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestHibernateApiWithInvalidAppId() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi("InvalidAppId", "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "[{malformed app id InvalidAppId}]", resp.Errors[0].InternalMessage)
}
