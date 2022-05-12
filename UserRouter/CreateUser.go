package UserRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTestSuite) TestCreateUserAsSuperAdmin() {
	createUserDto, _ := CreateUserRequestPayload(SuperAdmin, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	log.Println("Hitting the Create User API")
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), true, responseOfCreateUserApi.Result[0].SuperAdmin)
	assert.Equal(suite.T(), createUserDto.EmailId, responseOfCreateUserApi.Result[0].EmailId)
	assert.Empty(suite.T(), responseOfCreateUserApi.Result[0].Groups)
	assert.Empty(suite.T(), responseOfCreateUserApi.Result[0].RoleFilters)
	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
}

func (suite *UserTestSuite) TestCreateUserWithValidGroupsAndRoleFilters() {
	createUserDto, roleGroupId := CreateUserRequestPayload(GroupsAndRoleFilter, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	log.Println("Hitting the Create User API")
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), false, responseOfCreateUserApi.Result[0].SuperAdmin)
	assert.Equal(suite.T(), createUserDto.EmailId, responseOfCreateUserApi.Result[0].EmailId)
	assert.Equal(suite.T(), createUserDto.Groups[0], responseOfCreateUserApi.Result[0].Groups[0])
	assert.Equal(suite.T(), createUserDto.RoleFilters[0].Action, responseOfCreateUserApi.Result[0].RoleFilters[0].Action)
	assert.Equal(suite.T(), createUserDto.RoleFilters[0].Team, responseOfCreateUserApi.Result[0].RoleFilters[0].Team)

	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
}

func (suite *UserTestSuite) TestCreateUserWithValidGroupsOnly() {
	createUserDto, roleGroupId := CreateUserRequestPayload(GroupsOnly, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	log.Println("Hitting the Create User API")
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), false, responseOfCreateUserApi.Result[0].SuperAdmin)
	assert.Equal(suite.T(), createUserDto.EmailId, responseOfCreateUserApi.Result[0].EmailId)
	assert.Equal(suite.T(), createUserDto.Groups[0], responseOfCreateUserApi.Result[0].Groups[0])

	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
}

func (suite *UserTestSuite) TestCreateUserWithValidFiltersOnly() {
	createUserDto, roleGroupId := CreateUserRequestPayload(RoleFilterOnly, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	log.Println("Hitting the Create User API")
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), false, responseOfCreateUserApi.Result[0].SuperAdmin)
	assert.Equal(suite.T(), createUserDto.EmailId, responseOfCreateUserApi.Result[0].EmailId)

	assert.Equal(suite.T(), createUserDto.RoleFilters[0].Action, responseOfCreateUserApi.Result[0].RoleFilters[0].Action)
	assert.Equal(suite.T(), createUserDto.RoleFilters[0].Team, responseOfCreateUserApi.Result[0].RoleFilters[0].Team)

	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
}
