package ResponseDTO

import (
	Base "automation-suite/testUtils"
)

type CreateLinkResponseDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Success bool `json:"success"`
	} `json:"result"`
	Errors []Base.Errors `json:"errors"`
}
