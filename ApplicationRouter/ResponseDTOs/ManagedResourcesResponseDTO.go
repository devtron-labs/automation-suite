package ResponseDTOs

import "automation-suite/testUtils"

type ManagedResourcesResponseDTO struct {
	Result struct {
		Items []Items `json:"items"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}

type Items struct {
	Kind                string `json:"kind"`
	Namespace           string `json:"namespace"`
	Name                string `json:"name"`
	TargetState         string `json:"targetState"`
	LiveState           string `json:"liveState"`
	NormalizedLiveState string `json:"normalizedLiveState"`
	PredictedLiveState  string `json:"predictedLiveState"`
	ResourceVersion     string `json:"resourceVersion"`
	Group               string `json:"group,omitempty"`
}
