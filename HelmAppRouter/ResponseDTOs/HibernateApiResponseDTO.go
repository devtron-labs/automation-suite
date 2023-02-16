package ResponseDTOs

import "automation-suite/testUtils"

type HibernateApiResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		ErrorMessage string       `json:"errorMessage"`
		Success      bool         `json:"success"`
		TargetObject TargetObject `json:"targetObject"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}

type TargetObject struct {
	Group     string `json:"group"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Version   string `json:"version"`
}
