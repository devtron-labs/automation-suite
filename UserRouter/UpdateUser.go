package UserRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTestSuite) TestUpdateUserWithGroupsAndRoleFilters() {
	createUserDto, _ := CreateUserRequestPayload(SuperAdmin, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	log.Println("Hitting the Create User API")
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Getting the payload for updating the user")
	updateUserDto, roleGroupId := CreateUserRequestPayload("GroupsAndRoleFilter", suite.authToken)
	updateUserDto.Id = int32(responseOfCreateUserApi.Result[0].Id)
	updateUserDto.EmailId = responseOfCreateUserApi.Result[0].EmailId
	byteValueOfStruct, _ = json.Marshal(updateUserDto)

	log.Println("Hitting the Update User API")
	responseOfUpdateUserApi := HitUpdateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), false, responseOfUpdateUserApi.Result.SuperAdmin)
	assert.Equal(suite.T(), updateUserDto.EmailId, responseOfUpdateUserApi.Result.EmailId)
	assert.Equal(suite.T(), updateUserDto.Groups[0], responseOfUpdateUserApi.Result.Groups[0])
	assert.Equal(suite.T(), updateUserDto.RoleFilters[0].Action, responseOfUpdateUserApi.Result.RoleFilters[0].Action)
	assert.Equal(suite.T(), updateUserDto.RoleFilters[0].Team, responseOfUpdateUserApi.Result.RoleFilters[0].Team)

	log.Println("Hitting the get user by id for verifying the functionality of UpdateUserApi")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), false, responseOfGetUserById.Result.SuperAdmin)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Result.EmailId, responseOfGetUserById.Result.EmailId)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Result.Groups, responseOfGetUserById.Result.Groups)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Result.RoleFilters, responseOfGetUserById.Result.RoleFilters)

	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
}

func (suite *UserTestSuite) TestUpdateUserWithGroupsOnly() {
	createUserDto, _ := CreateUserRequestPayload(SuperAdmin, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	log.Println("Hitting the Create User API")
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Getting the payload for updating the user")
	updateUserDto, roleGroupId := CreateUserRequestPayload("GroupsOnly", suite.authToken)
	updateUserDto.Id = int32(responseOfCreateUserApi.Result[0].Id)
	updateUserDto.EmailId = responseOfCreateUserApi.Result[0].EmailId
	byteValueOfStruct, _ = json.Marshal(updateUserDto)

	log.Println("Hitting the Update User API")
	responseOfUpdateUserApi := HitUpdateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), false, responseOfUpdateUserApi.Result.SuperAdmin)
	assert.Equal(suite.T(), updateUserDto.EmailId, responseOfUpdateUserApi.Result.EmailId)
	assert.Equal(suite.T(), updateUserDto.Groups[0], responseOfUpdateUserApi.Result.Groups[0])

	log.Println("Hitting the get user by id for verifying the functionality of UpdateUserApi")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), false, responseOfGetUserById.Result.SuperAdmin)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Result.EmailId, responseOfGetUserById.Result.EmailId)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Result.Groups, responseOfGetUserById.Result.Groups)

	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
}

func (suite *UserTestSuite) TestUpdateUserWithRoleFiltersOnly() {
	createUserDto, _ := CreateUserRequestPayload(SuperAdmin, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	log.Println("Hitting the Create User API")
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Getting the payload for updating the user")
	updateUserDto, roleGroupId := CreateUserRequestPayload("RoleFilterOnly", suite.authToken)
	updateUserDto.Id = int32(responseOfCreateUserApi.Result[0].Id)
	updateUserDto.EmailId = responseOfCreateUserApi.Result[0].EmailId
	byteValueOfStruct, _ = json.Marshal(updateUserDto)

	log.Println("Hitting the Update User API")
	responseOfUpdateUserApi := HitUpdateUserApi(byteValueOfStruct, suite.authToken)
	assert.Equal(suite.T(), false, responseOfUpdateUserApi.Result.SuperAdmin)
	assert.Equal(suite.T(), updateUserDto.EmailId, responseOfUpdateUserApi.Result.EmailId)
	assert.Equal(suite.T(), updateUserDto.RoleFilters[0].Action, responseOfUpdateUserApi.Result.RoleFilters[0].Action)
	assert.Equal(suite.T(), updateUserDto.RoleFilters[0].Team, responseOfUpdateUserApi.Result.RoleFilters[0].Team)

	log.Println("Hitting the get user by id for verifying the functionality of UpdateUserApi")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), false, responseOfGetUserById.Result.SuperAdmin)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Result.EmailId, responseOfGetUserById.Result.EmailId)
	assert.Equal(suite.T(), responseOfUpdateUserApi.Result.RoleFilters, responseOfGetUserById.Result.RoleFilters)

	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
}
