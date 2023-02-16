package SSOLoginRouter

import (
	"github.com/stretchr/testify/assert"
	"log"
)

func (suite *SSOLoginTestSuite) TestClass2GetList() {
	suite.Run("A=1=TestGetList", func() {
		getListResponse := HitGetListApi(suite.authToken)
		log.Println("Asserting the API Response...")
		assert.Equal(suite.T(), 2, getListResponse.Result[0].Id)
		assert.True(suite.T(), getListResponse.Result[0].Active)
		assert.NotNil(suite.T(), getListResponse.Result[0].Url)
	})
}
