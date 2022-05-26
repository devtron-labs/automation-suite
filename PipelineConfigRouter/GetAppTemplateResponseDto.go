package PipelineConfigRouter

import Base "automation-suite/testUtils"

/////////////////=== GetAppTemplateResponseDto ====//////////////

type ContainerPort struct {
	EnvoyPort        int    `json:"envoyPort"`
	IdleTimeout      string `json:"idleTimeout"`
	Name             string `json:"name"`
	Port             int    `json:"port"`
	ServicePort      int    `json:"servicePort"`
	SupportStreaming bool   `json:"supportStreaming"`
	UseHTTP2         bool   `json:"useHTTP2"`
}

type ReadinessProbe struct {
	Path                string        `json:"Path"`
	Command             []interface{} `json:"command"`
	FailureThreshold    int           `json:"failureThreshold"`
	HttpHeaders         []interface{} `json:"httpHeaders"`
	InitialDelaySeconds int           `json:"initialDelaySeconds"`
	PeriodSeconds       int           `json:"periodSeconds"`
	Port                int           `json:"port"`
	Scheme              string        `json:"scheme"`
	SuccessThreshold    int           `json:"successThreshold"`
	Tcp                 bool          `json:"tcp"`
	TimeoutSeconds      int           `json:"timeoutSeconds"`
}

type Autoscaling struct {
	MaxReplicas                       int `json:"MaxReplicas"`
	MinReplicas                       int `json:"MinReplicas"`
	TargetCPUUtilizationPercentage    int `json:"TargetCPUUtilizationPercentage"`
	TargetMemoryUtilizationPercentage int `json:"TargetMemoryUtilizationPercentage"`
	Annotations                       struct {
	} `json:"annotations"`
	Behavior struct {
	} `json:"behavior"`
	Enabled      bool          `json:"enabled"`
	ExtraMetrics []interface{} `json:"extraMetrics"`
	Labels       struct {
	} `json:"labels"`
}

type Envoyproxy struct {
	ConfigMapName string `json:"configMapName"`
	Image         string `json:"image"`
	Resources     struct {
		Limits struct {
			Cpu    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"limits"`
		Requests struct {
			Cpu    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"requests"`
	} `json:"resources"`
}

type KedaAutoscaling struct {
	Advanced struct {
	} `json:"advanced"`
	AuthenticationRef struct {
	} `json:"authenticationRef"`
	Enabled                bool   `json:"enabled"`
	EnvSourceContainerName string `json:"envSourceContainerName"`
	MaxReplicaCount        int    `json:"maxReplicaCount"`
	MinReplicaCount        int    `json:"minReplicaCount"`
	TriggerAuthentication  struct {
		Enabled bool   `json:"enabled"`
		Name    string `json:"name"`
		Spec    struct {
		} `json:"spec"`
	} `json:"triggerAuthentication"`
	Triggers []interface{} `json:"triggers"`
}

type Service struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Properties  struct {
		Annotations struct {
			Type        string `json:"type"`
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"annotations"`
		Type struct {
			Type        string   `json:"type"`
			Description string   `json:"description"`
			Title       string   `json:"title"`
			Enum        []string `json:"enum"`
		} `json:"type"`
	} `json:"properties"`
}

type GetAppTemplateResponseDto struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Errors []Base.Errors `json:"errors"`
	Result struct {
		GlobalConfig struct {
			DefaultAppOverride struct {
				ContainerPort   []ContainerPort `json:"ContainerPort"`
				EnvVariables    []interface{}   `json:"EnvVariables"`
				GracePeriod     int             `json:"GracePeriod"`
				LivenessProbe   LivenessProbe   `json:"LivenessProbe"`
				MaxSurge        int             `json:"MaxSurge"`
				MaxUnavailable  int             `json:"MaxUnavailable"`
				MinReadySeconds int             `json:"MinReadySeconds"`
				ReadinessProbe  ReadinessProbe  `json:"ReadinessProbe"`
				Spec            Spec            `json:"Spec"`
				Args            struct {
					Enabled bool     `json:"enabled"`
					Value   []string `json:"value"`
				} `json:"args"`
				Autoscaling Autoscaling `json:"autoscaling"`
				Command     struct {
					Enabled bool          `json:"enabled"`
					Value   []interface{} `json:"value"`
				} `json:"command"`
				ContainerSecurityContext struct {
				} `json:"containerSecurityContext"`
				Containers        []interface{} `json:"containers"`
				DbMigrationConfig struct {
					Enabled bool `json:"enabled"`
				} `json:"dbMigrationConfig"`
				Envoyproxy Envoyproxy `json:"envoyproxy"`
				Image      struct {
					PullPolicy string `json:"pullPolicy"`
				} `json:"image"`
				ImagePullSecrets []interface{} `json:"imagePullSecrets"`
				Ingress          Ingress       `json:"ingress"`
				IngressInternal  struct {
					Annotations struct {
					} `json:"annotations"`
					ClassName string `json:"className"`
					Enabled   bool   `json:"enabled"`
					Hosts     []struct {
						Host     string   `json:"host"`
						PathType string   `json:"pathType"`
						Paths    []string `json:"paths"`
					} `json:"hosts"`
					Tls []interface{} `json:"tls"`
				} `json:"ingressInternal"`
				InitContainers                    []interface{}   `json:"initContainers"`
				KedaAutoscaling                   KedaAutoscaling `json:"kedaAutoscaling"`
				PauseForSecondsBeforeSwitchActive int             `json:"pauseForSecondsBeforeSwitchActive"`
				PodAnnotations                    struct {
				} `json:"podAnnotations"`
				PodLabels struct {
				} `json:"podLabels"`
				PodSecurityContext struct {
				} `json:"podSecurityContext"`
				Prometheus struct {
					Release string `json:"release"`
				} `json:"prometheus"`
				RawYaml      []interface{} `json:"rawYaml"`
				ReplicaCount int           `json:"replicaCount"`
				Resources    struct {
					Limits struct {
						Cpu    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"limits"`
					Requests struct {
						Cpu    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"requests"`
				} `json:"resources"`
				Secret struct {
					Data struct {
					} `json:"data"`
					Enabled bool `json:"enabled"`
				} `json:"secret"`
				Server struct {
					Deployment struct {
						Image    string `json:"image"`
						ImageTag string `json:"image_tag"`
					} `json:"deployment"`
				} `json:"server"`
				Service struct {
					Annotations struct {
					} `json:"annotations"`
					LoadBalancerSourceRanges []interface{} `json:"loadBalancerSourceRanges"`
					Type                     string        `json:"type"`
				} `json:"service"`
				ServiceAccount struct {
					Annotations struct {
					} `json:"annotations"`
					Create bool   `json:"create"`
					Name   string `json:"name"`
				} `json:"serviceAccount"`
				Servicemonitor struct {
					AdditionalLabels struct {
					} `json:"additionalLabels"`
				} `json:"servicemonitor"`
				Tolerations                     []interface{} `json:"tolerations"`
				TopologySpreadConstraints       []interface{} `json:"topologySpreadConstraints"`
				VolumeMounts                    []interface{} `json:"volumeMounts"`
				Volumes                         []interface{} `json:"volumes"`
				WaitForSecondsBeforeScalingDown int           `json:"waitForSecondsBeforeScalingDown"`
			} `json:"defaultAppOverride"`
			Readme string `json:"readme"`
			Schema struct {
				Schema     string `json:"$schema"`
				Type       string `json:"type"`
				Properties struct {
					ContainerPort struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Items       struct {
							Type       string `json:"type"`
							Properties struct {
								EnvoyPort struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"envoyPort"`
								IdleTimeout struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"idleTimeout"`
								Name struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"name"`
								Port struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"port"`
								ServicePort struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"servicePort"`
								SupportStreaming struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"supportStreaming"`
								UseHTTP2 struct {
									Type        string `json:"type"`
									Description string `json:"description"`
									Title       string `json:"title"`
								} `json:"useHTTP2"`
							} `json:"properties"`
						} `json:"items"`
					} `json:"ContainerPort"`
					EnvVariables struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"EnvVariables"`
					GracePeriod struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"GracePeriod"`
					LivenessProbe struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Path struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"Path"`
							Command struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"command"`
							FailureThreshold struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"failureThreshold"`
							HttpHeaders struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"httpHeaders"`
							InitialDelaySeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"initialDelaySeconds"`
							PeriodSeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"periodSeconds"`
							Port struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"port"`
							Scheme struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"scheme"`
							SuccessThreshold struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"successThreshold"`
							Tcp struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"tcp"`
							TimeoutSeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"timeoutSeconds"`
						} `json:"properties"`
					} `json:"LivenessProbe"`
					MaxSurge struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"MaxSurge"`
					MaxUnavailable struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"MaxUnavailable"`
					MinReadySeconds struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"MinReadySeconds"`
					ReadinessProbe struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Path struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"Path"`
							Command struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"command"`
							FailureThreshold struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"failureThreshold"`
							HttpHeader struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"httpHeader"`
							InitialDelaySeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"initialDelaySeconds"`
							PeriodSeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"periodSeconds"`
							Port struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"port"`
							Scheme struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"scheme"`
							SuccessThreshold struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"successThreshold"`
							Tcp struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"tcp"`
							TimeoutSeconds struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"timeoutSeconds"`
						} `json:"properties"`
					} `json:"ReadinessProbe"`
					Spec struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Affinity struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Key struct {
										AnyOf []struct {
											Type        string `json:"type"`
											Description string `json:"description,omitempty"`
											Title       string `json:"title,omitempty"`
										} `json:"anyOf"`
									} `json:"Key"`
									Values struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"Values"`
									Key1 struct {
										Type string `json:"type"`
									} `json:"key"`
								} `json:"properties"`
							} `json:"Affinity"`
						} `json:"properties"`
					} `json:"Spec"`
					Args struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
							Value struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Items       []struct {
									Type string `json:"type"`
								} `json:"items"`
							} `json:"value"`
						} `json:"properties"`
					} `json:"args"`
					Autoscaling struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							MaxReplicas struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"MaxReplicas"`
							MinReplicas struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"MinReplicas"`
							TargetCPUUtilizationPercentage struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"TargetCPUUtilizationPercentage"`
							TargetMemoryUtilizationPercentage struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"TargetMemoryUtilizationPercentage"`
							Behavior struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"behavior"`
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
							ExtraMetrics struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"extraMetrics"`
						} `json:"properties"`
					} `json:"autoscaling"`
					Command struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
							} `json:"enabled"`
							Value struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"value"`
						} `json:"properties"`
					} `json:"command"`
					ContainerSecurityContext struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"containerSecurityContext"`
					Containers struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
					} `json:"containers"`
					DbMigrationConfig struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
						} `json:"properties"`
					} `json:"dbMigrationConfig"`
					Envoyproxy struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							ConfigMapName struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"configMapName"`
							Image struct {
								Type        string `json:"type"`
								Description string `json:"description"`
							} `json:"image"`
							Resources struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Limits struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
										Properties  struct {
											Cpu struct {
												Type        string `json:"type"`
												Format      string `json:"format"`
												Description string `json:"description"`
												Title       string `json:"title"`
											} `json:"cpu"`
											Memory struct {
												Type        string `json:"type"`
												Format      string `json:"format"`
												Description string `json:"description"`
												Title       string `json:"title"`
											} `json:"memory"`
										} `json:"properties"`
									} `json:"limits"`
									Requests struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
										Properties  struct {
											Cpu struct {
												Type        string `json:"type"`
												Format      string `json:"format"`
												Description string `json:"description"`
												Title       string `json:"title"`
											} `json:"cpu"`
											Memory struct {
												Type        string `json:"type"`
												Format      string `json:"format"`
												Description string `json:"description"`
												Title       string `json:"title"`
											} `json:"memory"`
										} `json:"properties"`
									} `json:"requests"`
								} `json:"properties"`
							} `json:"resources"`
						} `json:"properties"`
					} `json:"envoyproxy"`
					Image struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							PullPolicy struct {
								Type        string   `json:"type"`
								Description string   `json:"description"`
								Title       string   `json:"title"`
								Enum        []string `json:"enum"`
							} `json:"pullPolicy"`
						} `json:"properties"`
					} `json:"image"`
					ImagePullSecrets struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"imagePullSecrets"`
					Ingress struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Annotations struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"annotations"`
							ClassName struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Default     string `json:"default"`
							} `json:"className"`
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
							Hosts struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Items       []struct {
									Type       string `json:"type"`
									Properties struct {
										Host struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
										} `json:"host"`
										PathType struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
										} `json:"pathType"`
										Paths struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
											Items       []struct {
												Type string `json:"type"`
											} `json:"items"`
										} `json:"paths"`
									} `json:"properties"`
								} `json:"items"`
							} `json:"hosts"`
							Tls struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"tls"`
						} `json:"properties"`
					} `json:"ingress"`
					IngressInternal struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Properties  struct {
							Annotations struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"annotations"`
							ClassName struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Default     string `json:"default"`
							} `json:"className"`
							Enabled struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"enabled"`
							Hosts struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Items       []struct {
									Type       string `json:"type"`
									Properties struct {
										Host struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
										} `json:"host"`
										PathType struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
										} `json:"pathType"`
										Paths struct {
											Type        string `json:"type"`
											Description string `json:"description"`
											Title       string `json:"title"`
											Items       []struct {
												Type string `json:"type"`
											} `json:"items"`
										} `json:"paths"`
									} `json:"properties"`
								} `json:"items"`
							} `json:"hosts"`
							Tls struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"tls"`
						} `json:"properties"`
					} `json:"ingressInternal"`
					InitContainers struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"initContainers"`
					KedaAutoscaling struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Advanced struct {
								Type string `json:"type"`
							} `json:"advanced"`
							AuthenticationRef struct {
								Type string `json:"type"`
							} `json:"authenticationRef"`
							Enabled struct {
								Type string `json:"type"`
							} `json:"enabled"`
							EnvSourceContainerName struct {
								Type string `json:"type"`
							} `json:"envSourceContainerName"`
							MaxReplicaCount struct {
								Type string `json:"type"`
							} `json:"maxReplicaCount"`
							MinReplicaCount struct {
								Type string `json:"type"`
							} `json:"minReplicaCount"`
							TriggerAuthentication struct {
								Type       string `json:"type"`
								Properties struct {
									Enabled struct {
										Type string `json:"type"`
									} `json:"enabled"`
									Name struct {
										Type string `json:"type"`
									} `json:"name"`
									Spec struct {
										Type string `json:"type"`
									} `json:"spec"`
								} `json:"properties"`
							} `json:"triggerAuthentication"`
							Triggers struct {
								Type  string `json:"type"`
								Items struct {
								} `json:"items"`
							} `json:"triggers"`
						} `json:"properties"`
					} `json:"kedaAutoscaling"`
					PauseForSecondsBeforeSwitchActive struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"pauseForSecondsBeforeSwitchActive"`
					PodAnnotations struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"podAnnotations"`
					PodLabels struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"podLabels"`
					PodSecurityContext struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"podSecurityContext"`
					Prometheus struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Release struct {
								Type        string `json:"type"`
								Description string `json:"description"`
							} `json:"release"`
						} `json:"properties"`
					} `json:"prometheus"`
					RawYaml struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"rawYaml"`
					ReplicaCount struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"replicaCount"`
					Resources struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Limits struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Cpu struct {
										Type        string `json:"type"`
										Format      string `json:"format"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"cpu"`
									Memory struct {
										Type        string `json:"type"`
										Format      string `json:"format"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"memory"`
								} `json:"properties"`
							} `json:"limits"`
							Requests struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Cpu struct {
										Type        string `json:"type"`
										Format      string `json:"format"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"cpu"`
									Memory struct {
										Type        string `json:"type"`
										Format      string `json:"format"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"memory"`
								} `json:"properties"`
							} `json:"requests"`
						} `json:"properties"`
					} `json:"resources"`
					Secret struct {
						Type       string `json:"type"`
						Properties struct {
							Data struct {
								Type string `json:"type"`
							} `json:"data"`
							Enabled struct {
								Type string `json:"type"`
							} `json:"enabled"`
						} `json:"properties"`
					} `json:"secret"`
					Server struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Deployment struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
								Properties  struct {
									Image struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"image"`
									ImageTag struct {
										Type        string `json:"type"`
										Description string `json:"description"`
										Title       string `json:"title"`
									} `json:"image_tag"`
								} `json:"properties"`
							} `json:"deployment"`
						} `json:"properties"`
					} `json:"server"`
					Service        Service `json:"service"`
					ServiceAccount struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							Annotations struct {
								Type        string `json:"type"`
								Title       string `json:"title"`
								Description string `json:"description"`
							} `json:"annotations"`
							Name struct {
								Type        string `json:"type"`
								Description string `json:"description"`
								Title       string `json:"title"`
							} `json:"name"`
							Create struct {
								Type string `json:"type"`
							} `json:"create"`
						} `json:"properties"`
					} `json:"serviceAccount"`
					Servicemonitor struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
						Properties  struct {
							AdditionalLabels struct {
								Type string `json:"type"`
							} `json:"additionalLabels"`
						} `json:"properties"`
					} `json:"servicemonitor"`
					Tolerations struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"tolerations"`
					TopologySpreadConstraints struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"topologySpreadConstraints"`
					VolumeMounts struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"volumeMounts"`
					Volumes struct {
						Type  string `json:"type"`
						Items struct {
						} `json:"items"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"volumes"`
					WaitForSecondsBeforeScalingDown struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Title       string `json:"title"`
					} `json:"waitForSecondsBeforeScalingDown"`
				} `json:"properties"`
			} `json:"schema"`
		} `json:"globalConfig"`
	} `json:"result"`
}
