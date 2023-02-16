package RequestDTOs

type TriggerCiPipelineRequestDTO struct {
	PipelineId          int                   `json:"pipelineId"`
	CiPipelineMaterials []CiPipelineMaterials `json:"ciPipelineMaterials"`
	InvalidateCache     bool                  `json:"invalidateCache"`
}

type GitCommit struct {
	Commit string `json:"Commit"`
}

type CiPipelineMaterials struct {
	Id        int       `json:"Id"`
	GitCommit GitCommit `json:"GitCommit"`
}
