package HelmAppRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

func (suite *HelmAppTestSuite) TestUnHibernateWorkloadApi() {
	suite.Run("A=1=UnHibernateWithValidArgs", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		errorMessage := resp.Result[0].ErrorMessage
		if errorMessage == "object is already scaled up" {
			assert.False(suite.T(), resp.Result[0].Success)
			//Hibernating the Workload First
			HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
			time.Sleep(15 * time.Second)
			//Again UnHibernating and verifying the response
			HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
			time.Sleep(15 * time.Second)
			assert.Equal(suite.T(), "", resp.Result[0].ErrorMessage)
			assert.True(suite.T(), resp.Result[0].Success)
		}
	})

	suite.Run("A=2=UnHibernateWithValidArgs", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "InvalidDeployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "no matches for kind \"InvalidDeployment\" in version \"apps/v1\"", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=3=TestUnHibernateApiWithInvalidName", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", "InvalidName", "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "deployments.apps \"InvalidName\" not found", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=4=TestUnHibernateApiWithInvalidGroup", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "Invalid", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "no matches for kind \"Deployment\" in version \"Invalid/v1\"", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=5=UnHibernateApiWithInvalidVersion", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "Invalid", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "no matches for kind \"Deployment\" in version \"apps/Invalid\"", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=6=UnHibernateApiWithInvalidNamespace", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "Invalid")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "deployments.apps \"envoy-deepak-testing-v1\" not found", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=7=UnHibernateApiWithInvalidAppId", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi("InvalidAppId", "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "[{malformed app id InvalidAppId}]", resp.Errors[0].InternalMessage)
	})
}
