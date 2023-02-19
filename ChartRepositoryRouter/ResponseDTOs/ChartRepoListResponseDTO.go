package ResponseDTOs

import "automation-suite/ChartRepositoryRouter/RequestDTOs"

type GetChartRepoListResponseDTO struct {
	Code   int                               `json:"code"`
	Status string                            `json:"status"`
	Result []RequestDTOs.ChartRepoRequestDTO `json:"result"`
}
