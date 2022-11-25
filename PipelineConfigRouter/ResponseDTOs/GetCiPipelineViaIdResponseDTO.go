package ResponseDTOs

import (
	Base "automation-suite/testUtils"
	"time"
)

type GetCiPipelineViaIdResponseDTO struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Errors []Base.Errors `json:"errors"`
	Result struct {
		Id               int           `json:"id"`
		AppId            int           `json:"appId"`
		DockerRegistry   string        `json:"dockerRegistry"`
		DockerRepository string        `json:"dockerRepository"`
		CiBuildConfig    CiBuildConfig `json:"ciBuildConfig"`
		AppName          string        `json:"appName"`
		Materials        []Materials   `json:"materials"`
		ScanEnabled      bool          `json:"scanEnabled"`
		CreatedOn        time.Time     `json:"CreatedOn"`
		CreatedBy        int           `json:"CreatedBy"`
		UpdatedOn        time.Time     `json:"UpdatedOn"`
		UpdatedBy        int           `json:"UpdatedBy"`
	} `json:"result"`
}

type CiBuildConfig struct {
	Id                int               `json:"id"`
	GitMaterialId     int               `json:"gitMaterialId"`
	CiBuildType       string            `json:"ciBuildType"`
	DockerBuildConfig DockerBuildConfig `json:"dockerBuildConfig"`
	BuildPackConfig   interface{}       `json:"buildPackConfig"`
}

type DockerBuildConfig struct {
	DockerfileRelativePath string `json:"dockerfileRelativePath"`
	DockerfileContent      string `json:"dockerfileContent"`
}

type Materials struct {
	GitMaterialId int    `json:"gitMaterialId"`
	MaterialName  string `json:"materialName"`
}
