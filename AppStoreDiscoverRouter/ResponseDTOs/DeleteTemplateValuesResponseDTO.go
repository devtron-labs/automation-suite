package ResponseDTOs

import Base "automation-suite/testUtils"

type DeleteTemplateValuesResponseDTO struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result bool          `json:"result"`
	Errors []Base.Errors `json:"errors"`
}
