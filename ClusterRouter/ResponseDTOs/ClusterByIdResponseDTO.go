package ResponseDTOs

import "automation-suite/testUtils"

type ClusterByIdResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Result Cluster            `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}
