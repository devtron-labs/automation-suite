package ResponseDTOs

import "time"

type DiscoverAppApiResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Id                           int       `json:"id"`
		AppStoreApplicationVersionId int       `json:"appStoreApplicationVersionId"`
		Name                         string    `json:"name"`
		ChartRepoId                  int       `json:"chart_repo_id"`
		ChartName                    string    `json:"chart_name"`
		Icon                         string    `json:"icon"`
		Active                       bool      `json:"active"`
		ChartGitLocation             string    `json:"chart_git_location"`
		CreatedOn                    time.Time `json:"created_on"`
		UpdatedOn                    time.Time `json:"updated_on"`
		Version                      string    `json:"version"`
		Deprecated                   bool      `json:"deprecated"`
	} `json:"result"`
}
