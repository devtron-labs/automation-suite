package HelmAppRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"time"
)

/*Test Data ExternalApp should be deployed in default cluster and group should be apps*/

func (suite *HelmAppTestSuite) TestHibernateWorkloadApi() {

	suite.Run("A=1=HibernateWithValidArgs", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		errorMessage := resp.Result[0].ErrorMessage
		if errorMessage == "object is already scaled down" {
			assert.False(suite.T(), resp.Result[0].Success)
			//Un-hibernating the Workload First
			HitUnHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
			time.Sleep(15 * time.Second)
			//Again Hibernating and verifying the response
			HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
			assert.Equal(suite.T(), "object is already scaled down", resp.Result[0].ErrorMessage)
		} else {
			assert.Equal(suite.T(), "", resp.Result[0].ErrorMessage)
		}
	})

	suite.Run("A=2=HibernateApiWithInvalidKind", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "InvalidDeployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "no matches for kind \"InvalidDeployment\" in version \"apps/v1\"", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=3=HibernateApiWithInvalidName", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", "InvalidName", "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "deployments.apps \"InvalidName\" not found", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=4=HibernateApiWithInvalidGroup", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "Invalid", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "no matches for kind \"Deployment\" in version \"Invalid/v1\"", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=5=HibernateApiWithInvalidVersion", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "Invalid", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "no matches for kind \"Deployment\" in version \"apps/Invalid\"", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=6=HibernateApiWithInvalidNamespace", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi(envConf.HAppId, "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "Invalid")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "deployments.apps \"envoy-deepak-testing-v1\" not found", resp.Result[0].ErrorMessage)
		assert.False(suite.T(), resp.Result[0].Success)
	})

	suite.Run("A=7=HibernateApiWithInvalidAppId", func() {
		envConf, _ := GetEnvironmentConfigForHelmApp()
		requestPayloadForHibernateApi := createRequestPayloadForHibernateApi("InvalidAppId", "Deployment", envConf.ResourceNameToHibernate, "v1", "apps", "default")
		byteValueOfStruct, _ := json.Marshal(requestPayloadForHibernateApi)
		resp := HitHibernateWorkloadApi(string(byteValueOfStruct), suite.authToken)
		assert.Equal(suite.T(), "[{malformed app id InvalidAppId}]", resp.Errors[0].InternalMessage)
	})
}
