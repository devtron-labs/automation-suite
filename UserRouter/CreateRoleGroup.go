package UserRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTestSuite) TestCreateRoleGroup() {

	suite.Run("A=1=CreateRoleGroupForDevtronAppsOnly", func() {
		createRoleGroupPayload := CreateRoleGroupPayload("WithDevtronAppsOnly")
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)

		log.Println("Hitting Create Role Group API")
		createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), createRoleGroupPayload.Name, createRoleGroupResponseBody.Result.Name)
		assert.Equal(suite.T(), createRoleGroupPayload.Description, createRoleGroupResponseBody.Result.Description)

		log.Println("Verifying the response of Create Role Group API using getRoleGroupById API")
		getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))
		assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Team, createRoleGroupResponseBody.Result.RoleFilters[0].Team)
		assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Action, createRoleGroupResponseBody.Result.RoleFilters[0].Action)

		log.Println("Deleting the Test Data created via Automation")
		HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
	})

	suite.Run("A=2=CreateRoleGroupForHelmAppsOnly", func() {
		createRoleGroupPayload := CreateRoleGroupPayload("WithHelmAppsOnly")
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)

		log.Println("Hitting Create Role Group API")
		createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), createRoleGroupPayload.Name, createRoleGroupResponseBody.Result.Name)
		assert.Equal(suite.T(), createRoleGroupPayload.Description, createRoleGroupResponseBody.Result.Description)

		log.Println("Verifying the response of Create Role Group API using getRoleGroupById API")
		getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))
		assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Team, createRoleGroupResponseBody.Result.RoleFilters[0].Team)
		assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Action, createRoleGroupResponseBody.Result.RoleFilters[0].Action)
		assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Environment, createRoleGroupResponseBody.Result.RoleFilters[0].Environment)
		assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].AccessType, createRoleGroupResponseBody.Result.RoleFilters[0].AccessType)

		log.Println("Deleting the Test Data created via Automation")
		HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
	})

	suite.Run("A=3=CreateRoleGroupForChartGroupsOnly", func() {
		createRoleGroupPayload := CreateRoleGroupPayload("WithChartGroupsOnly")
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)

		log.Println("Hitting Create Role Group API")
		createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), createRoleGroupPayload.Name, createRoleGroupResponseBody.Result.Name)
		assert.Equal(suite.T(), createRoleGroupPayload.Description, createRoleGroupResponseBody.Result.Description)

		log.Println("Verifying the response of Create Role Group API using getRoleGroupById API")
		getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))
		assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Action, createRoleGroupResponseBody.Result.RoleFilters[0].Action)
		assert.Equal(suite.T(), getRoleGroupByIdResponse.Result.RoleFilters[0].Entity, createRoleGroupResponseBody.Result.RoleFilters[0].Entity)
		log.Println("Deleting the Test Data created via Automation")
		HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
	})

	suite.Run("A=1=CreateRoleGroupForAllFilters", func() {
		createRoleGroupPayload := CreateRoleGroupPayload("WithAllFilter")
		byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)

		log.Println("Hitting Create Role Group API")
		createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)
		assert.Equal(suite.T(), createRoleGroupPayload.Name, createRoleGroupResponseBody.Result.Name)
		assert.Equal(suite.T(), createRoleGroupPayload.Description, createRoleGroupResponseBody.Result.Description)

		log.Println("Verifying the response of Create Role Group API using getRoleGroupById API")
		getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
		assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))

		log.Println("Deleting the Test Data created via Automation")
		HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
	})
}
