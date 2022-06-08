package ResponseDTOs

type GetWorkflowStatusResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		CdWorkflowStatus []CdWorkflowStatus `json:"cdWorkflowStatus"`
		CiWorkflowStatus []CiWorkflowStatus `json:"ciWorkflowStatus"`
	} `json:"result"`
}

type CdWorkflowStatus struct {
	CiPipelineId int    `json:"ci_pipeline_id"`
	PipelineId   int    `json:"pipeline_id"`
	DeployStatus string `json:"deploy_status"`
	PreStatus    string `json:"pre_status"`
	PostStatus   string `json:"post_status"`
}

type CiWorkflowStatus struct {
	CiPipelineId   int    `json:"ciPipelineId"`
	CiPipelineName string `json:"ciPipelineName"`
	CiStatus       string `json:"ciStatus"`
}
