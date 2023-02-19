package ResponseDTOs

import (
	Base "automation-suite/testUtils"
	"time"
)

type TimelineStatus string
type GetAppDeploymentStatusTimelineDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		DeploymentStartedOn  time.Time `json:"deploymentStartedOn"`
		DeploymentFinishedOn time.Time `json:"deploymentFinishedOn"`
		TriggeredBy          string    `json:"triggeredBy"`
		Timelines            []*struct {
			Id                           int            `json:"id"`
			InstalledAppVersionHistoryId int            `json:"InstalledAppVersionHistoryId,omitempty"`
			CdWorkflowRunnerId           int            `json:"cdWorkflowRunnerId"`
			Status                       TimelineStatus `json:"status"`
			StatusDetail                 string         `json:"statusDetail"`
			StatusTime                   time.Time      `json:"statusTime"`
			ResourceDetails              []*struct {
				Id                           int    `json:"id"`
				InstalledAppVersionHistoryId int    `json:"installedAppVersionHistoryId,omitempty"`
				CdWorkflowRunnerId           int    `json:"cdWorkflowRunnerId,omitempty"`
				ResourceName                 string `json:"resourceName"`
				ResourceKind                 string `json:"resourceKind"`
				ResourceGroup                string `json:"resourceGroup"`
				ResourceStatus               string `json:"resourceStatus"`
				ResourcePhase                string `json:"resourcePhase"`
				StatusMessage                string `json:"statusMessage"`
				TimelineStage                string `json:"timelineStage,omitempty"`
			} `json:"resourceDetails,omitempty"`
		} `json:"timelines"`
		StatusLastFetchedAt time.Time `json:"statusLastFetchedAt"`
		StatusFetchCount    int       `json:"statusFetchCount"`
	} `json:"result"`
	Error []Base.Errors `json:"errors"`
}
