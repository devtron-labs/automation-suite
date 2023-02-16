package ResponseDTOs

import (
	"automation-suite/PipelineConfigRouter/RequestDTOs"
	"automation-suite/testUtils"
)

type SaveCdPipelineResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Pipelines []RequestDTOs.Pipeline `json:"pipelines"`
		AppId     int                    `json:"appId"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
