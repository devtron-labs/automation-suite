package ResponseDTOs

import (
	"automation-suite/testUtils"
	"time"
)

type FetchMaterialsResponseDTO struct {
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
		Regex           string    `json:"regex"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
