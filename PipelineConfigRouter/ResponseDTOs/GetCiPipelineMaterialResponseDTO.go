package ResponseDTOs

import (
	"automation-suite/testUtils"
	"time"
)

type GetCiPipelineMaterialResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id              int       `json:"id"`
		GitMaterialId   int       `json:"gitMaterialId"`
		GitMaterialUrl  string    `json:"gitMaterialUrl"`
		GitMaterialName string    `json:"gitMaterialName"`
		Type            string    `json:"type"`
		Value           string    `json:"value"`
		Active          bool      `json:"active"`
		History         []History `json:"history"`
		LastFetchTime   time.Time `json:"lastFetchTime"`
		IsRepoError     bool      `json:"isRepoError"`
		RepoErrorMsg    string    `json:"repoErrorMsg"`
		IsBranchError   bool      `json:"isBranchError"`
		BranchErrorMsg  string    `json:"branchErrorMsg"`
		Url             string    `json:"url"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}

type History struct {
	Commit      string      `json:"Commit"`
	Author      string      `json:"Author"`
	Date        time.Time   `json:"Date"`
	Message     string      `json:"Message"`
	Changes     []string    `json:"Changes"`
	WebhookData interface{} `json:"WebhookData"`
}
