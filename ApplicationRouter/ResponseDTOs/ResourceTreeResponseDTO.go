package ResponseDTOs

import (
	"automation-suite/testUtils"
	"time"
)

type ResourceTreeResponseDTO struct {
	Result struct {
		Nodes                    []Nodes       `json:"nodes"`
		Hosts                    []Hosts       `json:"hosts"`
		NewGenerationReplicaSets []string      `json:"newGenerationReplicaSets"`
		Status                   string        `json:"status"`
		PodMetadata              []PodMetadata `json:"podMetadata"`
		Conditions               interface{}   `json:"conditions"`
	} `json:"result"`
	Errors []testUtils.Errors `json:"errors"`
}

type PodMetadata struct {
	Name           string      `json:"name"`
	Uid            string      `json:"uid"`
	Containers     []string    `json:"containers"`
	InitContainers interface{} `json:"initContainers"`
	IsNew          bool        `json:"isNew"`
}

type Hosts struct {
	Name          string          `json:"name"`
	ResourcesInfo []ResourcesInfo `json:"resourcesInfo"`
	SystemInfo    SystemInfo      `json:"systemInfo"`
}

type ResourcesInfo struct {
	ResourceName         string `json:"resourceName"`
	RequestedByApp       int64  `json:"requestedByApp"`
	RequestedByNeighbors int64  `json:"requestedByNeighbors"`
	Capacity             int64  `json:"capacity"`
}

type SystemInfo struct {
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OsImage                 string `json:"osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
}

type Nodes struct {
	Version         string       `json:"version"`
	Kind            string       `json:"kind"`
	Namespace       string       `json:"namespace"`
	Name            string       `json:"name"`
	Uid             string       `json:"uid"`
	ParentRefs      []ParentRefs `json:"parentRefs,omitempty"`
	ResourceVersion string       `json:"resourceVersion"`
	CreatedAt       time.Time    `json:"createdAt"`
	Info            []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"info,omitempty"`
	NetworkingInfo struct {
		Labels struct {
			App                     string `json:"app"`
			AppId                   string `json:"appId"`
			EnvId                   string `json:"envId"`
			Release                 string `json:"release"`
			RolloutsPodTemplateHash string `json:"rollouts-pod-template-hash"`
		} `json:"labels,omitempty"`
		TargetLabels struct {
			App string `json:"app"`
		} `json:"targetLabels,omitempty"`
	} `json:"networkingInfo,omitempty"`
	Images []string `json:"images,omitempty"`
	Health struct {
		Status  string `json:"status"`
		Message string `json:"message,omitempty"`
	} `json:"health,omitempty"`
	Group string `json:"group,omitempty"`
}

type ParentRefs struct {
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Uid       string `json:"uid"`
	Group     string `json:"group,omitempty"`
}
