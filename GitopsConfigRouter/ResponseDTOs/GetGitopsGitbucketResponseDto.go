package ResponseDTOs

import (
	"automation-suite/GitopsConfigRouter/RequestDTOs"
	"time"
)

type UpdateGitopsConfigResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		SuccessfulStages []string `json:"successfulStages"`
		StageErrorMap    struct {
		} `json:"stageErrorMap"`
		ValidatedOn      time.Time `json:"validatedOn"`
		DeleteRepoFailed bool      `json:"deleteRepoFailed"`
	} `json:"result"`
}

type FetchAllGitopsConfigResponseDto struct {
	Code   int                                        `json:"code"`
	Status string                                     `json:"status"`
	Result []RequestDTOs.CreateGitopsConfigRequestDto `json:"result"`
}
