package ResponseDTOs

import "automation-suite/testUtils"

type GetApplicationDetailResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		AppDetail AppDetail `json:"appDetail"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}

type AppDetail struct {
	ApplicationStatus  string             `json:"applicationStatus"`
	ReleaseStatus      ReleaseStatus      `json:"releaseStatus"`
	ChartMetadata      ChartMetadata      `json:"chartMetadata"`
	EnvironmentDetails EnvironmentDetails `json:"environmentDetails"`
}

type ReleaseStatus struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}
