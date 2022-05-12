package HelmAppRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func (suite *HelmAppTestSuite) TestUnHitHibernateWorkloadApiWithValidArgsInPayload() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	errorMessage := resp.Result[0].ErrorMessage
	if errorMessage == "object is already scaled up" {
		assert.False(suite.T(), resp.Result[0].Success)
		//Hibernating the Workload First
		HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		//Again UnHibernating and verifying the response
		HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "", resp.Result[0].ErrorMessage)
		assert.True(suite.T(), resp.Result[0].Success)
	}
}

func (suite *HelmAppTestSuite) TestUnHibernateApiWithInvalidKind() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "InvalidDeployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "no matches for kind \"InvalidDeployment\" in version \"apps/v1\"", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestUnHibernateApiWithInvalidName() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", "InvalidName", "v1", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "deployments.apps \"InvalidName\" not found", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestUnHibernateApiWithInvalidGroup() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "Invalid", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "no matches for kind \"Deployment\" in version \"Invalid/v1\"", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestUnHibernateApiWithInvalidVersion() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "Invalid", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "no matches for kind \"Deployment\" in version \"apps/Invalid\"", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestUnHibernateApiWithInvalidNamespace() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "Invalid")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "deployments.apps \"envoy-deepak-testing-v1\" not found", resp.Result[0].ErrorMessage)
	assert.False(suite.T(), resp.Result[0].Success)
}

func (suite *HelmAppTestSuite) TestUnHibernateApiWithInvalidAppId() {
	envConf, _ := GetEnvironmentConfigForHelmApp()
	requestPayloadForHibernateApi := createRequestPayloadForHibernateApi("InvalidAppId", "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
	byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
	resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
	assert.Equal(suite.T(), "[{malformed app id InvalidAppId}]", resp.Errors[0].InternalMessage)
}
