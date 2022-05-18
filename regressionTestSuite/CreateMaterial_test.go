package regressionTestSuite

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateAppMaterialWithValidPayload(t *testing.T) {
	authToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImYyMjFhMjIwZDhlZDNmYTZlMjBiNDQxYzM4MmYzYmRiOTIyNWQyMzAifQ.eyJpc3MiOiJodHRwczovL3N0YWdpbmcuZGV2dHJvbi5pbmZvL29yY2hlc3RyYXRvci9hcGkvZGV4Iiwic3ViIjoiQ2hVeE1ESTJPVEk1TmpBd056RXhNekU0TVRFMU5qY1NCbWR2YjJkc1pRIiwiYXVkIjoiYXJnby1jZCIsImV4cCI6MTY1MjQyNjE1NiwiaWF0IjoxNjUyMzM5NzU2LCJhdF9oYXNoIjoiNWFQYmZzcFRKZGRDS0ZVaEozcDA5dyIsImVtYWlsIjoibmlrZXNoQGRldnRyb24uYWkiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6Ik5pa2VzaCBSYXRob2QifQ.F9QXzcy-nwge_csGaX_ZRXmXRTLa92a6zP4ODkEJh-qAHrMHCgW7SjnZ00PUVqHtLszEZ5-rGPX028QXZ_Ky3WrtbDmKa2KGMnm9FcDTChml9ODtaw13nA53-NoRjQSI2bHG0emYFBUM9Gnjg4_t54XXGKvdKK_WB7keRYeOyQqvPpg75Tr1uAbRRASgDFgSotsgNYQizyv95m9hxzO_dFbQx-gdwxXnalIOU0lrbSHBaX4GGldC_cRLdEtC_LxFdApfgw7S8iyrmgffdB_jGezmEkdecQZfNii_NHoARryHieo8CQJsfFnXDmRRT_DYKVpEOaqb94RIZZ9SfHZC7w"
	appId := 201
	gitopsConfig, _ := GetGitopsConfig()

	createAppMaterialRequestDto := GetAppMaterialRequestDto(appId, gitopsConfig.Url, 1, false)
	byteValueOfStruct2, _ := json.Marshal(createAppMaterialRequestDto)
	log.Println("Hitting The post team API")
	createAppMaterialResponseDto := HitCreateAppMaterialApi(byteValueOfStruct2, appId, gitopsConfig.Url, 1, false, authToken)

	log.Println("Validating the Response of the Create Gitops Config API...")
	assert.Equal(t, appId, createAppMaterialResponseDto.Result.AppId)

	log.Println("getting payload for Delete Team API")
	byteValueOfDeleteApp := GetPayLoadForDeleteAppMaterialAPI(createAppMaterialResponseDto.Result.AppId, createAppMaterialResponseDto.Result.Material[0])
	log.Println("Hitting the Delete team API for Removing the data created via automation")
	HitDeleteAppMaterialApi(byteValueOfDeleteApp, authToken)
}
