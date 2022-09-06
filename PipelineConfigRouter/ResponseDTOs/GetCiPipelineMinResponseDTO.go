package ResponseDTOs

import "automation-suite/testUtils"

type GetCiPipelineMinResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Result []CiPipelineMin    `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}

type CiPipelineMin struct {
	Name             string `json:"name"`
	Id               int    `json:"id"`
	ParentCiPipeline int    `json:"parentCiPipeline"`
	ParentAppId      int    `json:"parentAppId"`
	PipelineType     string `json:"pipelineType"`
	ScanEnabled      bool   `json:"scanEnabled"`
}
