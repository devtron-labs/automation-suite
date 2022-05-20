package UserRouter

import (
	"automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTestSuite) TestGetUserByIdWithValidArg() {
	createUserDto, _ := CreateUserRequestPayload(SuperAdmin, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the Get User By Id")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), responseOfCreateUserApi.Result[0].SuperAdmin, responseOfGetUserById.Result.SuperAdmin)
	assert.Equal(suite.T(), responseOfCreateUserApi.Result[0].EmailId, responseOfGetUserById.Result.EmailId)
	assert.Empty(suite.T(), responseOfGetUserById.Result.Groups)
	assert.Empty(suite.T(), responseOfGetUserById.Result.RoleFilters)

	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
}

func (suite *UserTestSuite) TestGetUserByIdWithInvalidArg() {
	randomId := testUtils.GetRandomNumberOf9Digit()
	log.Println("Hitting the Get User By Id with invalid ID")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(randomId), suite.authToken)
	assert.Equal(suite.T(), 404, responseOfGetUserById.Code)
	assert.Equal(suite.T(), "Failed to get by id", responseOfGetUserById.Errors[0].UserMessage)
}
