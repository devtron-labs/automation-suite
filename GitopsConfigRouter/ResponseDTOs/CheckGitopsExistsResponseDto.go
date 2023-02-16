package ResponseDTOs

type CheckGitopsExistsResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Exists bool `json:"exists"`
	} `json:"result"`
}
