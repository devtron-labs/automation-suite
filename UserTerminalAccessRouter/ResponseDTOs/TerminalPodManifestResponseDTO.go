package ResponseDTOs

import (
	"automation-suite/testUtils"
	"time"
)

type TerminalPodManifestResponseDTO struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Errors []testUtils.Errors `json:"errors"`
	Result struct {
		Manifest struct {
			ApiVersion string `json:"apiVersion"`
			Kind       string `json:"kind"`
			Metadata   struct {
				CreationTimestamp time.Time `json:"creationTimestamp"`
				ManagedFields     []struct {
					ApiVersion string `json:"apiVersion"`
					FieldsType string `json:"fieldsType"`
					FieldsV1   struct {
						FSpec struct {
							FContainers struct {
								KNameDevtronDebugTerminal struct {
									Field1 struct {
									} `json:"."`
									FArgs struct {
									} `json:"f:args"`
									FCommand struct {
									} `json:"f:command"`
									FImage struct {
									} `json:"f:image"`
									FImagePullPolicy struct {
									} `json:"f:imagePullPolicy"`
									FName struct {
									} `json:"f:name"`
									FResources struct {
									} `json:"f:resources"`
									FTerminationMessagePath struct {
									} `json:"f:terminationMessagePath"`
									FTerminationMessagePolicy struct {
									} `json:"f:terminationMessagePolicy"`
								} `json:"k:{"name":"devtron-debug-terminal"}"`
							} `json:"f:containers"`
							FDnsPolicy struct {
							} `json:"f:dnsPolicy"`
							FEnableServiceLinks struct {
							} `json:"f:enableServiceLinks"`
							FNodeSelector struct {
							} `json:"f:nodeSelector"`
							FRestartPolicy struct {
							} `json:"f:restartPolicy"`
							FSchedulerName struct {
							} `json:"f:schedulerName"`
							FSecurityContext struct {
							} `json:"f:securityContext"`
							FServiceAccount struct {
							} `json:"f:serviceAccount"`
							FServiceAccountName struct {
							} `json:"f:serviceAccountName"`
							FTerminationGracePeriodSeconds struct {
							} `json:"f:terminationGracePeriodSeconds"`
						} `json:"f:spec,omitempty"`
						FStatus struct {
							FConditions struct {
								KTypeContainersReady struct {
									Field1 struct {
									} `json:"."`
									FLastProbeTime struct {
									} `json:"f:lastProbeTime"`
									FLastTransitionTime struct {
									} `json:"f:lastTransitionTime"`
									FStatus struct {
									} `json:"f:status"`
									FType struct {
									} `json:"f:type"`
								} `json:"k:{"type":"ContainersReady"}"`
								KTypeInitialized struct {
									Field1 struct {
									} `json:"."`
									FLastProbeTime struct {
									} `json:"f:lastProbeTime"`
									FLastTransitionTime struct {
									} `json:"f:lastTransitionTime"`
									FStatus struct {
									} `json:"f:status"`
									FType struct {
									} `json:"f:type"`
								} `json:"k:{"type":"Initialized"}"`
								KTypeReady struct {
									Field1 struct {
									} `json:"."`
									FLastProbeTime struct {
									} `json:"f:lastProbeTime"`
									FLastTransitionTime struct {
									} `json:"f:lastTransitionTime"`
									FStatus struct {
									} `json:"f:status"`
									FType struct {
									} `json:"f:type"`
								} `json:"k:{"type":"Ready"}"`
							} `json:"f:conditions"`
							FContainerStatuses struct {
							} `json:"f:containerStatuses"`
							FHostIP struct {
							} `json:"f:hostIP"`
							FPhase struct {
							} `json:"f:phase"`
							FPodIP struct {
							} `json:"f:podIP"`
							FPodIPs struct {
								Field1 struct {
								} `json:"."`
								KIp10244078 struct {
									Field1 struct {
									} `json:"."`
									FIp struct {
									} `json:"f:ip"`
								} `json:"k:{"ip":"10.244.0.78"}"`
							} `json:"f:podIPs"`
							FStartTime struct {
							} `json:"f:startTime"`
						} `json:"f:status,omitempty"`
					} `json:"fieldsV1"`
					Manager     string    `json:"manager"`
					Operation   string    `json:"operation"`
					Time        time.Time `json:"time"`
					Subresource string    `json:"subresource,omitempty"`
				} `json:"managedFields"`
				Name            string `json:"name"`
				Namespace       string `json:"namespace"`
				ResourceVersion string `json:"resourceVersion"`
				Uid             string `json:"uid"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Args            []string `json:"args"`
					Command         []string `json:"command"`
					Image           string   `json:"image"`
					ImagePullPolicy string   `json:"imagePullPolicy"`
					Name            string   `json:"name"`
					Resources       struct {
					} `json:"resources"`
					TerminationMessagePath   string `json:"terminationMessagePath"`
					TerminationMessagePolicy string `json:"terminationMessagePolicy"`
					VolumeMounts             []struct {
						MountPath string `json:"mountPath"`
						Name      string `json:"name"`
						ReadOnly  bool   `json:"readOnly"`
					} `json:"volumeMounts"`
				} `json:"containers"`
				DnsPolicy          string `json:"dnsPolicy"`
				EnableServiceLinks bool   `json:"enableServiceLinks"`
				NodeName           string `json:"nodeName"`
				NodeSelector       struct {
					KubernetesIoHostname string `json:"kubernetes.io/hostname"`
				} `json:"nodeSelector"`
				PreemptionPolicy string `json:"preemptionPolicy"`
				Priority         int    `json:"priority"`
				RestartPolicy    string `json:"restartPolicy"`
				SchedulerName    string `json:"schedulerName"`
				SecurityContext  struct {
				} `json:"securityContext"`
				ServiceAccount                string `json:"serviceAccount"`
				ServiceAccountName            string `json:"serviceAccountName"`
				TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
				Tolerations                   []struct {
					Effect            string `json:"effect"`
					Key               string `json:"key"`
					Operator          string `json:"operator"`
					TolerationSeconds int    `json:"tolerationSeconds"`
				} `json:"tolerations"`
				Volumes []struct {
					Name      string `json:"name"`
					Projected struct {
						DefaultMode int `json:"defaultMode"`
						Sources     []struct {
							ServiceAccountToken struct {
								ExpirationSeconds int    `json:"expirationSeconds"`
								Path              string `json:"path"`
							} `json:"serviceAccountToken,omitempty"`
							ConfigMap struct {
								Items []struct {
									Key  string `json:"key"`
									Path string `json:"path"`
								} `json:"items"`
								Name string `json:"name"`
							} `json:"configMap,omitempty"`
							DownwardAPI struct {
								Items []struct {
									FieldRef struct {
										ApiVersion string `json:"apiVersion"`
										FieldPath  string `json:"fieldPath"`
									} `json:"fieldRef"`
									Path string `json:"path"`
								} `json:"items"`
							} `json:"downwardAPI,omitempty"`
						} `json:"sources"`
					} `json:"projected"`
				} `json:"volumes"`
			} `json:"spec"`
			Status struct {
				Conditions []struct {
					LastProbeTime      interface{} `json:"lastProbeTime"`
					LastTransitionTime time.Time   `json:"lastTransitionTime"`
					Status             string      `json:"status"`
					Type               string      `json:"type"`
				} `json:"conditions"`
				ContainerStatuses []struct {
					ContainerID string `json:"containerID"`
					Image       string `json:"image"`
					ImageID     string `json:"imageID"`
					LastState   struct {
					} `json:"lastState"`
					Name         string `json:"name"`
					Ready        bool   `json:"ready"`
					RestartCount int    `json:"restartCount"`
					Started      bool   `json:"started"`
					State        struct {
						Running struct {
							StartedAt time.Time `json:"startedAt"`
						} `json:"running"`
					} `json:"state"`
				} `json:"containerStatuses"`
				HostIP string `json:"hostIP"`
				Phase  string `json:"phase"`
				PodIP  string `json:"podIP"`
				PodIPs []struct {
					Ip string `json:"ip"`
				} `json:"podIPs"`
				QosClass  string    `json:"qosClass"`
				StartTime time.Time `json:"startTime"`
			} `json:"status"`
		} `json:"manifest"`
	} `json:"result"`
}
