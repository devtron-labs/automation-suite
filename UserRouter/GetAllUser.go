package UserRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTestSuite) TestGetAllUserApi() {
	log.Println("Getting All the Users before creation any new user")
	GetAllUserResponseDto := HitGetAllUserApi(suite.authToken)
	totalUsersBeforeCreation := len(GetAllUserResponseDto.Result)

	log.Println("Hitting the Create User API")
	createUserDto, roleGroupId := CreateUserRequestPayload(GroupsAndRoleFilter, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Getting All the Users again after creating a new user")
	GetAllUserResponseDto = HitGetAllUserApi(suite.authToken)
	assert.Equal(suite.T(), totalUsersBeforeCreation+1, len(GetAllUserResponseDto.Result))

	log.Println("Deleting the Test data Created via Automation")
	HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
}
