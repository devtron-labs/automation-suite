package ResponseDTOs

import (
	"automation-suite/ChartRepositoryRouter/RequestDTOs"
	"automation-suite/testUtils"
)

type CreateChartRepoResponseDTO struct {
	Code   int                             `json:"code"`
	Status string                          `json:"status"`
	Result RequestDTOs.ChartRepoRequestDTO `json:"result"`
	Errors []testUtils.Errors              `json:"errors"`
}
