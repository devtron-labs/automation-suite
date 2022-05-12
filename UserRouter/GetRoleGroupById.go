package UserRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

//todo need to add more assert once order of RoleFilters data in response will fix (currently coming randomly)
func (suite *UserTestSuite) TestGetRoleGroupHavingAllFiltersWithValidId() {
	createRoleGroupPayload := CreateRoleGroupPayload(WithAllFilter)
	byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
	createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the getRoleGroupById API")
	getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)

	log.Println("verifying the response of getRoleGroupById API")
	assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))
	log.Println("Deleting the Test data Created via Automation")
	HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
}

func (suite *UserTestSuite) TestGetRoleGroupHavingHelmAppFilterWithValidId() {
	createRoleGroupPayload := CreateRoleGroupPayload(WithHelmAppsOnly)
	byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
	createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the getRoleGroupById API")
	getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)

	log.Println("verifying the response of getRoleGroupById API")
	assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))
	assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Action, getRoleGroupByIdResponse.Result.RoleFilters[0].Action)
	assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Team, getRoleGroupByIdResponse.Result.RoleFilters[0].Team)
	assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Environment, getRoleGroupByIdResponse.Result.RoleFilters[0].Environment)
	assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].AccessType, getRoleGroupByIdResponse.Result.RoleFilters[0].AccessType)
	log.Println("Deleting the Test data Created via Automation")
	HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
}

func (suite *UserTestSuite) TestGetRoleGroupHavingDevtronAppFilterWithValidId() {
	createRoleGroupPayload := CreateRoleGroupPayload(WithDevtronAppsOnly)
	byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
	createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the getRoleGroupById API")
	getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)

	log.Println("verifying the response of getRoleGroupById API")
	assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))
	assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Action, getRoleGroupByIdResponse.Result.RoleFilters[0].Action)
	assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Team, getRoleGroupByIdResponse.Result.RoleFilters[0].Team)
	log.Println("Deleting the Test data Created via Automation")
	HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
}

func (suite *UserTestSuite) TestGetRoleGroupHavingChartGroupFilterWithValidId() {
	createRoleGroupPayload := CreateRoleGroupPayload(WithChartGroupsOnly)
	byteValueOfStruct, _ := json.Marshal(createRoleGroupPayload)
	createRoleGroupResponseBody := HitCreateRoleGroupApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the getRoleGroupById API")
	getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)

	log.Println("verifying the response of getRoleGroupById API")
	assert.Equal(suite.T(), len(getRoleGroupByIdResponse.Result.RoleFilters), len(createRoleGroupResponseBody.Result.RoleFilters))
	assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Action, getRoleGroupByIdResponse.Result.RoleFilters[0].Action)
	assert.Equal(suite.T(), createRoleGroupResponseBody.Result.RoleFilters[0].Entity, getRoleGroupByIdResponse.Result.RoleFilters[0].Entity)
	log.Println("Deleting the Test data Created via Automation")
	HitDeleteRoleGroupByIdApi(strconv.Itoa(createRoleGroupResponseBody.Result.Id), suite.authToken)
}

func (suite *UserTestSuite) TestGetRoleGroupWithInvalidId() {
	log.Println("Hitting the getRoleGroupById with invalid argument")
	randomId := testUtils.GetRandomNumberOf9Digit()
	getRoleGroupByIdResponse := HitGetRoleGroupByIdApi(strconv.Itoa(randomId), suite.authToken)

	log.Println("verifying the response of getRoleGroupById API")
	assert.Equal(suite.T(), "[{pg: no rows in result set}]", getRoleGroupByIdResponse.Errors[0].InternalMessage)
	assert.Equal(suite.T(), "Failed to get by id", getRoleGroupByIdResponse.Errors[0].UserMessage)
}
