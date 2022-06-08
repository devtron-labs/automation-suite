package ResponseDTOs

type CreateGitopsConfigResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		SuccessfulStages []string `json:"successfulStages"`
		StageErrorMap    struct {
			ErrorInConnectingWithGITHUB string `json:"error in connecting with GITHUB"`
		} `json:"stageErrorMap"`
		DeleteRepoFailed bool `json:"deleteRepoFailed"`
	} `json:"result"`
}
