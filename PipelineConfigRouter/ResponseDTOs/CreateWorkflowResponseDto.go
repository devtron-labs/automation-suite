package ResponseDTOs

import "automation-suite/PipelineConfigRouter/RequestDTOs"

type Material struct {
	GitMaterialId int    `json:"gitMaterialId"`
	MaterialName  string `json:"materialName"`
}
type CreateWorkflowResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Id                int    `json:"id"`
		AppId             int    `json:"appId"`
		DockerRegistry    string `json:"dockerRegistry"`
		DockerRepository  string `json:"dockerRepository"`
		DockerBuildConfig struct {
			GitMaterialId          int    `json:"gitMaterialId"`
			DockerfileRelativePath string `json:"dockerfileRelativePath"`
		} `json:"dockerBuildConfig"`
		CiPipelines   []RequestDTOs.CiPipeline `json:"ciPipelines"`
		AppName       string                   `json:"appName"`
		Materials     []Material               `json:"materials"`
		AppWorkflowId int                      `json:"appWorkflowId"`
		ScanEnabled   bool                     `json:"scanEnabled"`
	} `json:"result"`
}
