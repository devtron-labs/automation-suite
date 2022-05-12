package AppLabelsRouter

import (
	"log"
	"testing"
)

func TestGetAppLabelsWithValidAppId(t *testing.T) {
	authToken := ""
	attributesApiResp := HitGetAppMetaInfoByIdApi("83", authToken)
	log.Println(attributesApiResp)
}
