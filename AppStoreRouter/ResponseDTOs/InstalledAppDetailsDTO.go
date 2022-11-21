package ResponseDTOs

import (
	Base "automation-suite/testUtils"
	"time"
)

type InstalledAppDetailsResponseDTO struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Result AppDetails    `json:"result"`
	Error  []Base.Errors `json:"errors"`
}

type AppDetails struct {
	InstalledAppId                int          `json:"installedAppId"`
	AppId                         int          `json:"appId"`
	AppStoreInstalledAppVersionId int          `json:"appStoreInstalledAppVersionId"`
	AppStoreChartName             string       `json:"appStoreChartName"`
	AppStoreChartId               int          `json:"appStoreChartId"`
	AppStoreAppName               string       `json:"appStoreAppName"`
	AppStoreAppVersion            string       `json:"appStoreAppVersion"`
	AppName                       string       `json:"appName"`
	EnvironmentId                 int          `json:"environmentId"`
	EnvironmentName               string       `json:"environmentName"`
	Namespace                     string       `json:"namespace"`
	LastDeployedTime              time.Time    `json:"lastDeployedTime"`
	LastDeployedBy                string       `json:"lastDeployedBy"`
	Deprecated                    bool         `json:"deprecated"`
	K8SVersion                    string       `json:"k8sVersion"`
	CiArtifactId                  int          `json:"ciArtifactId"`
	ClusterId                     int          `json:"clusterId"`
	DeploymentAppType             string       `json:"deploymentAppType"`
	InstanceDetail                interface{}  `json:"instanceDetail"`
	ResourceTree                  ResourceTree `json:"resourceTree"`
}

type ResourceTree struct {
	Conditions               interface{}   `json:"conditions"`
	Hosts                    []Hosts       `json:"hosts"`
	NewGenerationReplicaSets []string      `json:"newGenerationReplicaSets"`
	Nodes                    []Nodes       `json:"nodes"`
	PodMetadata              []PodMetadata `json:"podMetadata"`
	Status                   string        `json:"status"`
}

type Hosts struct {
	Name          string          `json:"name"`
	ResourcesInfo []ResourcesInfo `json:"resourcesInfo"`
	SystemInfo    SystemInfo      `json:"systemInfo"`
}

type Nodes struct {
	CreatedAt       time.Time      `json:"createdAt"`
	Kind            string         `json:"kind"`
	Name            string         `json:"name"`
	Namespace       string         `json:"namespace"`
	ResourceVersion string         `json:"resourceVersion"`
	Uid             string         `json:"uid"`
	Version         string         `json:"version"`
	ParentRefs      []ParentRefs   `json:"parentRefs,omitempty"`
	Health          Health         `json:"health,omitempty"`
	Images          []string       `json:"images,omitempty"`
	Info            []Info         `json:"info,omitempty"`
	NetworkingInfo  NetworkingInfo `json:"networkingInfo,omitempty"`
	Group           string         `json:"group,omitempty"`
}

type PodMetadata struct {
	Containers     []string      `json:"containers"`
	InitContainers []interface{} `json:"initContainers"`
	IsNew          bool          `json:"isNew"`
	Name           string        `json:"name"`
	Uid            string        `json:"uid"`
}

type Labels struct {
	AppKubernetesIoComponent       string `json:"app.kubernetes.io/component"`
	AppKubernetesIoInstance        string `json:"app.kubernetes.io/instance"`
	AppKubernetesIoManagedBy       string `json:"app.kubernetes.io/managed-by"`
	AppKubernetesIoName            string `json:"app.kubernetes.io/name"`
	ControllerRevisionHash         string `json:"controller-revision-hash,omitempty"`
	HelmShChart                    string `json:"helm.sh/chart"`
	StatefulsetKubernetesIoPodName string `json:"statefulset.kubernetes.io/pod-name,omitempty"`
	PodTemplateHash                string `json:"pod-template-hash,omitempty"`
}

type TargetLabels struct {
	AppKubernetesIoComponent string `json:"app.kubernetes.io/component,omitempty"`
	AppKubernetesIoInstance  string `json:"app.kubernetes.io/instance"`
	AppKubernetesIoName      string `json:"app.kubernetes.io/name"`
}

type SystemInfo struct {
	Architecture            string `json:"architecture"`
	BootID                  string `json:"bootID"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KernelVersion           string `json:"kernelVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	MachineID               string `json:"machineID"`
	OperatingSystem         string `json:"operatingSystem"`
	OsImage                 string `json:"osImage"`
	SystemUUID              string `json:"systemUUID"`
}

type ParentRefs struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Uid       string `json:"uid"`
	Group     string `json:"group,omitempty"`
}

type Health struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type ResourcesInfo struct {
	Capacity             int64  `json:"capacity"`
	RequestedByApp       int64  `json:"requestedByApp"`
	RequestedByNeighbors int64  `json:"requestedByNeighbors"`
	ResourceName         string `json:"resourceName"`
}

type Info struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NetworkingInfo struct {
	Labels       Labels       `json:"labels,omitempty"`
	TargetLabels TargetLabels `json:"targetLabels,omitempty"`
}
