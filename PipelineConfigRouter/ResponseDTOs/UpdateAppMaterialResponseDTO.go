package ResponseDTOs

import (
	"automation-suite/PipelineConfigRouter/RequestDTOs"
	"automation-suite/testUtils"
)

type UpdateAppMaterialResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		AppId    int                         `json:"appId"`
		Material RequestDTOs.UpdatedMaterial `json:"material"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
