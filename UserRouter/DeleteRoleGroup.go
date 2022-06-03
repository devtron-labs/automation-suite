package UserRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTestSuite) TestDeleteRoleGroup() {

	suite.Run("A=1=DeleteRoleGroupHavingAllRoleFilters", func() {
		createRoleGroupPayload := CreateRoleGroupPayload(WithAllFilter)
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)

		log.Println("Hitting Create Role Group API")
		createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)

		log.Println("Hitting the DeleteRoleGroupById API")
		deleteRoleGroupByIdResponse := HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), true, deleteRoleGroupByIdResponse.Result)

		log.Println("Verifying the response of DeleteRoleGroup API using getRoleGroupById")
		getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), "Failed to get by id", getRoleGroupByIdResponse.Errors[0].UserMessage)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", getRoleGroupByIdResponse.Errors[0].InternalMessage)
	})

	suite.Run("A=2=DeleteRoleGroupHavingRoleFilterHelmAppsOnly", func() {
		createRoleGroupPayload := CreateRoleGroupPayload(WithHelmAppsOnly)
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
		createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)

		log.Println("Hitting the DeleteRoleGroupById API")
		deleteRoleGroupByIdResponse := HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), true, deleteRoleGroupByIdResponse.Result)

		log.Println("Verifying the response of DeleteRoleGroup API using getRoleGroupById")
		getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), "Failed to get by id", getRoleGroupByIdResponse.Errors[0].UserMessage)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", getRoleGroupByIdResponse.Errors[0].InternalMessage)
	})

	suite.Run("A=3=DeleteRoleGroupHavingRoleFilterDevtronAppsOnly", func() {
		createRoleGroupPayload := CreateRoleGroupPayload(WithDevtronAppsOnly)
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
		createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)

		log.Println("Hitting the DeleteRoleGroupById API")
		deleteRoleGroupByIdResponse := HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), true, deleteRoleGroupByIdResponse.Result)

		log.Println("Verifying the response of DeleteRoleGroup API using getRoleGroupById")
		getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), "Failed to get by id", getRoleGroupByIdResponse.Errors[0].UserMessage)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", getRoleGroupByIdResponse.Errors[0].InternalMessage)
	})

	suite.Run("A=4=DeleteRoleGroupHavingRoleFilterChartGroupsOnly", func() {
		createRoleGroupPayload := CreateRoleGroupPayload(WithChartGroupsOnly)
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
		createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)

		log.Println("Hitting the DeleteRoleGroupById API")
		deleteRoleGroupByIdResponse := HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), true, deleteRoleGroupByIdResponse.Result)

		log.Println("Verifying the response of DeleteRoleGroup API using getRoleGroupById")
		getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), "Failed to get by id", getRoleGroupByIdResponse.Errors[0].UserMessage)
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", getRoleGroupByIdResponse.Errors[0].InternalMessage)
	})
	
	suite.Run("A=5=DeleteRoleApiWithInvalidId", func() {
		log.Println("Hitting the getRoleGroupById with invalid argument")
		randomId := testUtils.GetRandomNumberOf9Digit()
		getRoleGroupByIdResponse := HitDeleteRoleGroupByIdApi(strconv.Itoa(randomId), suite.authToken)

		log.Println("verifying the response of getRoleGroupById API")
		assert.Equal(suite.T(), "[{pg: no rows in result set}]", getRoleGroupByIdResponse.Errors[0].InternalMessage)
		assert.Equal(suite.T(), 404, getRoleGroupByIdResponse.Code)
	})
}
