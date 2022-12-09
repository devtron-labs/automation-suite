package ResponseDTOs

import (
	"automation-suite/testUtils"
	"time"
)

type TerminalPodEventsResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Errors []testUtils.Errors `json:"errors"`
	Result PodEventsResult    `json:"result"`
}

type PodEventsResult struct {
	Events Events `json:"events"`
}

type Events struct {
	Metadata EventMetadata `json:"metadata"`
	Items    []Items       `json:"items"`
}

type EventMetadata struct {
	ResourceVersion string `json:"resourceVersion"`
}
type FAction struct {
}
type FEventTime struct {
}
type FNote struct {
}
type FReason struct {
}
type FRegarding struct {
}
type FReportingController struct {
}
type FReportingInstance struct {
}
type FType struct {
}
type FieldsV1 struct {
	FAction              FAction              `json:"f:action"`
	FEventTime           FEventTime           `json:"f:eventTime"`
	FNote                FNote                `json:"f:note"`
	FReason              FReason              `json:"f:reason"`
	FRegarding           FRegarding           `json:"f:regarding"`
	FReportingController FReportingController `json:"f:reportingController"`
	FReportingInstance   FReportingInstance   `json:"f:reportingInstance"`
	FType                FType                `json:"f:type"`
}
type ManagedFields struct {
	Manager    string    `json:"manager"`
	Operation  string    `json:"operation"`
	APIVersion string    `json:"apiVersion"`
	Time       time.Time `json:"time"`
	FieldsType string    `json:"fieldsType"`
	FieldsV1   FieldsV1  `json:"fieldsV1"`
}
type Metadata struct {
	Name              string          `json:"name"`
	Namespace         string          `json:"namespace"`
	UID               string          `json:"uid"`
	ResourceVersion   string          `json:"resourceVersion"`
	CreationTimestamp time.Time       `json:"creationTimestamp"`
	ManagedFields     []ManagedFields `json:"managedFields"`
}
type InvolvedObject struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	UID             string `json:"uid"`
	APIVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
}
type Source struct {
}
type InvolvedObject0 struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	UID             string `json:"uid"`
	APIVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
	FieldPath       string `json:"fieldPath"`
}
type Source0 struct {
	Component string `json:"component"`
	Host      string `json:"host"`
}
type InvolvedObject1 struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	UID             string `json:"uid"`
	APIVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
	FieldPath       string `json:"fieldPath"`
}
type Source1 struct {
	Component string `json:"component"`
	Host      string `json:"host"`
}
type InvolvedObject2 struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	UID             string `json:"uid"`
	APIVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
	FieldPath       string `json:"fieldPath"`
}
type Source2 struct {
	Component string `json:"component"`
	Host      string `json:"host"`
}
type InvolvedObject3 struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	UID             string `json:"uid"`
	APIVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
	FieldPath       string `json:"fieldPath"`
}
type Source3 struct {
	Component string `json:"component"`
	Host      string `json:"host"`
}
type Items struct {
	Metadata           Metadata        `json:"metadata"`
	InvolvedObject     InvolvedObject  `json:"involvedObject,omitempty"`
	Reason             string          `json:"reason"`
	Message            string          `json:"message"`
	Source             Source          `json:"source,omitempty"`
	FirstTimestamp     interface{}     `json:"firstTimestamp"`
	LastTimestamp      interface{}     `json:"lastTimestamp"`
	Type               string          `json:"type"`
	EventTime          time.Time       `json:"eventTime"`
	Action             string          `json:"action,omitempty"`
	ReportingComponent string          `json:"reportingComponent"`
	ReportingInstance  string          `json:"reportingInstance"`
	InvolvedObject0    InvolvedObject0 `json:"involvedObject,omitempty"`
	Source0            Source0         `json:"source,omitempty"`
	Count              int             `json:"count,omitempty"`
	InvolvedObject1    InvolvedObject1 `json:"involvedObject,omitempty"`
	Source1            Source1         `json:"source,omitempty"`
	InvolvedObject2    InvolvedObject2 `json:"involvedObject,omitempty"`
	Source2            Source2         `json:"source,omitempty"`
	InvolvedObject3    InvolvedObject3 `json:"involvedObject,omitempty"`
	Source3            Source3         `json:"source,omitempty"`
}
