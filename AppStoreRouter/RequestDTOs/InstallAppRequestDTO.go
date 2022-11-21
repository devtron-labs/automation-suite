package RequestDTOs

type InstallAppRequestDTO struct {
	TeamId             int            `json:"teamId"`
	ReferenceValueId   int            `json:"referenceValueId"`
	ReferenceValueKind string         `json:"referenceValueKind"`
	EnvironmentId      int            `json:"environmentId"`
	Namespace          string         `json:"namespace"`
	AppStoreVersion    int            `json:"appStoreVersion"`
	ValuesOverride     ValuesOverride `json:"valuesOverride"`
	ValuesOverrideYaml string         `json:"valuesOverrideYaml"`
	AppName            string         `json:"appName"`
}

type Global struct {
	ImageRegistry    string        `json:"imageRegistry"`
	ImagePullSecrets []interface{} `json:"imagePullSecrets"`
	StorageClass     string        `json:"storageClass"`
}

type Auth struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	FernetKey      string `json:"fernetKey"`
	SecretKey      string `json:"secretKey"`
	ExistingSecret string `json:"existingSecret"`
}

type Dags struct {
	ExistingConfigmap string `json:"existingConfigmap"`
	Image             struct {
		Registry    string        `json:"registry"`
		Repository  string        `json:"repository"`
		Tag         string        `json:"tag"`
		PullPolicy  string        `json:"pullPolicy"`
		PullSecrets []interface{} `json:"pullSecrets"`
	} `json:"image"`
}

type Image struct {
	Registry    string        `json:"registry"`
	Repository  string        `json:"repository"`
	Tag         string        `json:"tag"`
	PullPolicy  string        `json:"pullPolicy"`
	PullSecrets []interface{} `json:"pullSecrets"`
	Debug       bool          `json:"debug"`
}

type LivenessProbe struct {
	Enabled             bool `json:"enabled"`
	InitialDelaySeconds int  `json:"initialDelaySeconds"`
	PeriodSeconds       int  `json:"periodSeconds"`
	TimeoutSeconds      int  `json:"timeoutSeconds"`
	FailureThreshold    int  `json:"failureThreshold"`
	SuccessThreshold    int  `json:"successThreshold"`
}

type ReadinessProbe struct {
	Enabled             bool `json:"enabled"`
	InitialDelaySeconds int  `json:"initialDelaySeconds"`
	PeriodSeconds       int  `json:"periodSeconds"`
	TimeoutSeconds      int  `json:"timeoutSeconds"`
	FailureThreshold    int  `json:"failureThreshold"`
	SuccessThreshold    int  `json:"successThreshold"`
}

type StartupProbe struct {
	Enabled             bool `json:"enabled"`
	InitialDelaySeconds int  `json:"initialDelaySeconds"`
	PeriodSeconds       int  `json:"periodSeconds"`
	TimeoutSeconds      int  `json:"timeoutSeconds"`
	FailureThreshold    int  `json:"failureThreshold"`
	SuccessThreshold    int  `json:"successThreshold"`
}

type Web struct {
	Image               Image         `json:"image"`
	BaseUrl             string        `json:"baseUrl"`
	ExistingConfigmap   string        `json:"existingConfigmap"`
	Command             []interface{} `json:"command"`
	Args                []interface{} `json:"args"`
	ExtraEnvVars        []interface{} `json:"extraEnvVars"`
	ExtraEnvVarsCM      string        `json:"extraEnvVarsCM"`
	ExtraEnvVarsSecret  string        `json:"extraEnvVarsSecret"`
	ExtraEnvVarsSecrets []interface{} `json:"extraEnvVarsSecrets"`
	ContainerPorts      struct {
		Http int `json:"http"`
	} `json:"containerPorts"`
	ReplicaCount        int            `json:"replicaCount"`
	LivenessProbe       LivenessProbe  `json:"livenessProbe"`
	ReadinessProbe      ReadinessProbe `json:"readinessProbe"`
	StartupProbe        StartupProbe   `json:"startupProbe"`
	CustomLivenessProbe struct {
	} `json:"customLivenessProbe"`
	CustomReadinessProbe struct {
	} `json:"customReadinessProbe"`
	CustomStartupProbe struct {
	} `json:"customStartupProbe"`
	Resources struct {
		Limits struct {
		} `json:"limits"`
		Requests struct {
		} `json:"requests"`
	} `json:"resources"`
	PodSecurityContext struct {
		Enabled bool `json:"enabled"`
		FsGroup int  `json:"fsGroup"`
	} `json:"podSecurityContext"`
	ContainerSecurityContext struct {
		Enabled      bool `json:"enabled"`
		RunAsUser    int  `json:"runAsUser"`
		RunAsNonRoot bool `json:"runAsNonRoot"`
	} `json:"containerSecurityContext"`
	LifecycleHooks struct {
	} `json:"lifecycleHooks"`
	HostAliases []interface{} `json:"hostAliases"`
	PodLabels   struct {
	} `json:"podLabels"`
	PodAnnotations struct {
	} `json:"podAnnotations"`
	Affinity struct {
	} `json:"affinity"`
	NodeAffinityPreset struct {
		Key    string        `json:"key"`
		Type   string        `json:"type"`
		Values []interface{} `json:"values"`
	} `json:"nodeAffinityPreset"`
	NodeSelector struct {
	} `json:"nodeSelector"`
	PodAffinityPreset             string        `json:"podAffinityPreset"`
	PodAntiAffinityPreset         string        `json:"podAntiAffinityPreset"`
	Tolerations                   []interface{} `json:"tolerations"`
	TopologySpreadConstraints     []interface{} `json:"topologySpreadConstraints"`
	PriorityClassName             string        `json:"priorityClassName"`
	SchedulerName                 string        `json:"schedulerName"`
	TerminationGracePeriodSeconds string        `json:"terminationGracePeriodSeconds"`
	UpdateStrategy                struct {
		Type          string `json:"type"`
		RollingUpdate struct {
		} `json:"rollingUpdate"`
	} `json:"updateStrategy"`
	Sidecars          []interface{} `json:"sidecars"`
	InitContainers    []interface{} `json:"initContainers"`
	ExtraVolumeMounts []interface{} `json:"extraVolumeMounts"`
	ExtraVolumes      []interface{} `json:"extraVolumes"`
	Pdb               struct {
		Create         bool   `json:"create"`
		MinAvailable   int    `json:"minAvailable"`
		MaxUnavailable string `json:"maxUnavailable"`
	} `json:"pdb"`
}

type ValuesOverride struct {
	Global           Global        `json:"global"`
	KubeVersion      string        `json:"kubeVersion"`
	NameOverride     string        `json:"nameOverride"`
	FullnameOverride string        `json:"fullnameOverride"`
	ClusterDomain    string        `json:"clusterDomain"`
	ExtraDeploy      []interface{} `json:"extraDeploy"`
	CommonLabels     struct {
	} `json:"commonLabels"`
	CommonAnnotations struct {
	} `json:"commonAnnotations"`
	DiagnosticMode struct {
		Enabled bool     `json:"enabled"`
		Command []string `json:"command"`
		Args    []string `json:"args"`
	} `json:"diagnosticMode"`
	Auth                Auth          `json:"auth"`
	Executor            string        `json:"executor"`
	LoadExamples        bool          `json:"loadExamples"`
	Configuration       string        `json:"configuration"`
	ExistingConfigmap   string        `json:"existingConfigmap"`
	Dags                Dags          `json:"dags"`
	ExtraEnvVars        []interface{} `json:"extraEnvVars"`
	ExtraEnvVarsCM      string        `json:"extraEnvVarsCM"`
	ExtraEnvVarsSecret  string        `json:"extraEnvVarsSecret"`
	ExtraEnvVarsSecrets []interface{} `json:"extraEnvVarsSecrets"`
	Sidecars            []interface{} `json:"sidecars"`
	InitContainers      []interface{} `json:"initContainers"`
	ExtraVolumeMounts   []interface{} `json:"extraVolumeMounts"`
	ExtraVolumes        []interface{} `json:"extraVolumes"`
	Web                 Web           `json:"web"`
	Scheduler           struct {
		Image               Image         `json:"image"`
		ReplicaCount        int           `json:"replicaCount"`
		Command             []interface{} `json:"command"`
		Args                []interface{} `json:"args"`
		ExtraEnvVars        []interface{} `json:"extraEnvVars"`
		ExtraEnvVarsCM      string        `json:"extraEnvVarsCM"`
		ExtraEnvVarsSecret  string        `json:"extraEnvVarsSecret"`
		ExtraEnvVarsSecrets []interface{} `json:"extraEnvVarsSecrets"`
		CustomLivenessProbe struct {
		} `json:"customLivenessProbe"`
		CustomReadinessProbe struct {
		} `json:"customReadinessProbe"`
		CustomStartupProbe struct {
		} `json:"customStartupProbe"`
		Resources struct {
			Limits struct {
			} `json:"limits"`
			Requests struct {
			} `json:"requests"`
		} `json:"resources"`
		PodSecurityContext struct {
			Enabled bool `json:"enabled"`
			FsGroup int  `json:"fsGroup"`
		} `json:"podSecurityContext"`
		ContainerSecurityContext struct {
			Enabled      bool `json:"enabled"`
			RunAsUser    int  `json:"runAsUser"`
			RunAsNonRoot bool `json:"runAsNonRoot"`
		} `json:"containerSecurityContext"`
		LifecycleHooks struct {
		} `json:"lifecycleHooks"`
		HostAliases []interface{} `json:"hostAliases"`
		PodLabels   struct {
		} `json:"podLabels"`
		PodAnnotations struct {
		} `json:"podAnnotations"`
		Affinity struct {
		} `json:"affinity"`
		NodeAffinityPreset struct {
			Key    string        `json:"key"`
			Type   string        `json:"type"`
			Values []interface{} `json:"values"`
		} `json:"nodeAffinityPreset"`
		NodeSelector struct {
		} `json:"nodeSelector"`
		PodAffinityPreset             string        `json:"podAffinityPreset"`
		PodAntiAffinityPreset         string        `json:"podAntiAffinityPreset"`
		Tolerations                   []interface{} `json:"tolerations"`
		TopologySpreadConstraints     []interface{} `json:"topologySpreadConstraints"`
		PriorityClassName             string        `json:"priorityClassName"`
		SchedulerName                 string        `json:"schedulerName"`
		TerminationGracePeriodSeconds string        `json:"terminationGracePeriodSeconds"`
		UpdateStrategy                struct {
			Type          string `json:"type"`
			RollingUpdate struct {
			} `json:"rollingUpdate"`
		} `json:"updateStrategy"`
		Sidecars          []interface{} `json:"sidecars"`
		InitContainers    []interface{} `json:"initContainers"`
		ExtraVolumeMounts []interface{} `json:"extraVolumeMounts"`
		ExtraVolumes      []interface{} `json:"extraVolumes"`
		Pdb               struct {
			Create         bool   `json:"create"`
			MinAvailable   int    `json:"minAvailable"`
			MaxUnavailable string `json:"maxUnavailable"`
		} `json:"pdb"`
	} `json:"scheduler"`
	Worker struct {
		Image struct {
			Registry    string        `json:"registry"`
			Repository  string        `json:"repository"`
			Tag         string        `json:"tag"`
			PullPolicy  string        `json:"pullPolicy"`
			PullSecrets []interface{} `json:"pullSecrets"`
			Debug       bool          `json:"debug"`
		} `json:"image"`
		Command             []interface{} `json:"command"`
		Args                []interface{} `json:"args"`
		ExtraEnvVars        []interface{} `json:"extraEnvVars"`
		ExtraEnvVarsCM      string        `json:"extraEnvVarsCM"`
		ExtraEnvVarsSecret  string        `json:"extraEnvVarsSecret"`
		ExtraEnvVarsSecrets []interface{} `json:"extraEnvVarsSecrets"`
		ContainerPorts      struct {
			Http int `json:"http"`
		} `json:"containerPorts"`
		ReplicaCount  int `json:"replicaCount"`
		LivenessProbe struct {
			Enabled             bool `json:"enabled"`
			InitialDelaySeconds int  `json:"initialDelaySeconds"`
			PeriodSeconds       int  `json:"periodSeconds"`
			TimeoutSeconds      int  `json:"timeoutSeconds"`
			FailureThreshold    int  `json:"failureThreshold"`
			SuccessThreshold    int  `json:"successThreshold"`
		} `json:"livenessProbe"`
		ReadinessProbe struct {
			Enabled             bool `json:"enabled"`
			InitialDelaySeconds int  `json:"initialDelaySeconds"`
			PeriodSeconds       int  `json:"periodSeconds"`
			TimeoutSeconds      int  `json:"timeoutSeconds"`
			FailureThreshold    int  `json:"failureThreshold"`
			SuccessThreshold    int  `json:"successThreshold"`
		} `json:"readinessProbe"`
		StartupProbe struct {
			Enabled             bool `json:"enabled"`
			InitialDelaySeconds int  `json:"initialDelaySeconds"`
			PeriodSeconds       int  `json:"periodSeconds"`
			TimeoutSeconds      int  `json:"timeoutSeconds"`
			FailureThreshold    int  `json:"failureThreshold"`
			SuccessThreshold    int  `json:"successThreshold"`
		} `json:"startupProbe"`
		CustomLivenessProbe struct {
		} `json:"customLivenessProbe"`
		CustomReadinessProbe struct {
		} `json:"customReadinessProbe"`
		CustomStartupProbe struct {
		} `json:"customStartupProbe"`
		Resources struct {
			Limits struct {
			} `json:"limits"`
			Requests struct {
			} `json:"requests"`
		} `json:"resources"`
		PodSecurityContext struct {
			Enabled bool `json:"enabled"`
			FsGroup int  `json:"fsGroup"`
		} `json:"podSecurityContext"`
		ContainerSecurityContext struct {
			Enabled      bool `json:"enabled"`
			RunAsUser    int  `json:"runAsUser"`
			RunAsNonRoot bool `json:"runAsNonRoot"`
		} `json:"containerSecurityContext"`
		LifecycleHooks struct {
		} `json:"lifecycleHooks"`
		HostAliases []interface{} `json:"hostAliases"`
		PodLabels   struct {
		} `json:"podLabels"`
		PodAnnotations struct {
		} `json:"podAnnotations"`
		Affinity struct {
		} `json:"affinity"`
		NodeAffinityPreset struct {
			Key    string        `json:"key"`
			Type   string        `json:"type"`
			Values []interface{} `json:"values"`
		} `json:"nodeAffinityPreset"`
		NodeSelector struct {
		} `json:"nodeSelector"`
		PodAffinityPreset             string        `json:"podAffinityPreset"`
		PodAntiAffinityPreset         string        `json:"podAntiAffinityPreset"`
		Tolerations                   []interface{} `json:"tolerations"`
		TopologySpreadConstraints     []interface{} `json:"topologySpreadConstraints"`
		PriorityClassName             string        `json:"priorityClassName"`
		SchedulerName                 string        `json:"schedulerName"`
		TerminationGracePeriodSeconds string        `json:"terminationGracePeriodSeconds"`
		UpdateStrategy                struct {
			Type          string `json:"type"`
			RollingUpdate struct {
			} `json:"rollingUpdate"`
		} `json:"updateStrategy"`
		Sidecars                  []interface{} `json:"sidecars"`
		InitContainers            []interface{} `json:"initContainers"`
		ExtraVolumeMounts         []interface{} `json:"extraVolumeMounts"`
		ExtraVolumes              []interface{} `json:"extraVolumes"`
		ExtraVolumeClaimTemplates []interface{} `json:"extraVolumeClaimTemplates"`
		PodTemplate               struct {
		} `json:"podTemplate"`
		Pdb struct {
			Create         bool   `json:"create"`
			MinAvailable   int    `json:"minAvailable"`
			MaxUnavailable string `json:"maxUnavailable"`
		} `json:"pdb"`
		Autoscaling struct {
			Enabled      bool `json:"enabled"`
			MinReplicas  int  `json:"minReplicas"`
			MaxReplicas  int  `json:"maxReplicas"`
			TargetCPU    int  `json:"targetCPU"`
			TargetMemory int  `json:"targetMemory"`
		} `json:"autoscaling"`
	} `json:"worker"`
	Git struct {
		Image struct {
			Registry    string        `json:"registry"`
			Repository  string        `json:"repository"`
			Tag         string        `json:"tag"`
			PullPolicy  string        `json:"pullPolicy"`
			PullSecrets []interface{} `json:"pullSecrets"`
		} `json:"image"`
		Dags struct {
			Enabled      bool `json:"enabled"`
			Repositories []struct {
				Repository string `json:"repository"`
				Branch     string `json:"branch"`
				Name       string `json:"name"`
				Path       string `json:"path"`
			} `json:"repositories"`
		} `json:"dags"`
		Plugins struct {
			Enabled      bool `json:"enabled"`
			Repositories []struct {
				Repository string `json:"repository"`
				Branch     string `json:"branch"`
				Name       string `json:"name"`
				Path       string `json:"path"`
			} `json:"repositories"`
		} `json:"plugins"`
		Clone struct {
			Command            []interface{} `json:"command"`
			Args               []interface{} `json:"args"`
			ExtraVolumeMounts  []interface{} `json:"extraVolumeMounts"`
			ExtraEnvVars       []interface{} `json:"extraEnvVars"`
			ExtraEnvVarsCM     string        `json:"extraEnvVarsCM"`
			ExtraEnvVarsSecret string        `json:"extraEnvVarsSecret"`
			Resources          struct {
			} `json:"resources"`
		} `json:"clone"`
		Sync struct {
			Interval           int           `json:"interval"`
			Command            []interface{} `json:"command"`
			Args               []interface{} `json:"args"`
			ExtraVolumeMounts  []interface{} `json:"extraVolumeMounts"`
			ExtraEnvVars       []interface{} `json:"extraEnvVars"`
			ExtraEnvVarsCM     string        `json:"extraEnvVarsCM"`
			ExtraEnvVarsSecret string        `json:"extraEnvVarsSecret"`
			Resources          struct {
			} `json:"resources"`
		} `json:"sync"`
	} `json:"git"`
	Ldap struct {
		Enabled              bool   `json:"enabled"`
		Uri                  string `json:"uri"`
		Basedn               string `json:"basedn"`
		SearchAttribute      string `json:"searchAttribute"`
		Binddn               string `json:"binddn"`
		Bindpw               string `json:"bindpw"`
		UserRegistration     string `json:"userRegistration"`
		UserRegistrationRole string `json:"userRegistrationRole"`
		RolesMapping         string `json:"rolesMapping"`
		RolesSyncAtLogin     string `json:"rolesSyncAtLogin"`
		Tls                  struct {
			Enabled               bool   `json:"enabled"`
			AllowSelfSigned       bool   `json:"allowSelfSigned"`
			CertificatesSecret    string `json:"certificatesSecret"`
			CertificatesMountPath string `json:"certificatesMountPath"`
			CAFilename            string `json:"CAFilename"`
		} `json:"tls"`
	} `json:"ldap"`
	Service struct {
		Type  string `json:"type"`
		Ports struct {
			Http int `json:"http"`
		} `json:"ports"`
		NodePorts struct {
			Http string `json:"http"`
		} `json:"nodePorts"`
		SessionAffinity       string `json:"sessionAffinity"`
		SessionAffinityConfig struct {
		} `json:"sessionAffinityConfig"`
		ClusterIP                string        `json:"clusterIP"`
		LoadBalancerIP           string        `json:"loadBalancerIP"`
		LoadBalancerSourceRanges []interface{} `json:"loadBalancerSourceRanges"`
		ExternalTrafficPolicy    string        `json:"externalTrafficPolicy"`
		Annotations              struct {
		} `json:"annotations"`
		ExtraPorts []interface{} `json:"extraPorts"`
	} `json:"service"`
	Ingress struct {
		Enabled          bool   `json:"enabled"`
		IngressClassName string `json:"ingressClassName"`
		PathType         string `json:"pathType"`
		ApiVersion       string `json:"apiVersion"`
		Hostname         string `json:"hostname"`
		Path             string `json:"path"`
		Annotations      struct {
		} `json:"annotations"`
		Tls        bool          `json:"tls"`
		SelfSigned bool          `json:"selfSigned"`
		ExtraHosts []interface{} `json:"extraHosts"`
		ExtraPaths []interface{} `json:"extraPaths"`
		ExtraTls   []interface{} `json:"extraTls"`
		Secrets    []interface{} `json:"secrets"`
		ExtraRules []interface{} `json:"extraRules"`
	} `json:"ingress"`
	ServiceAccount struct {
		Create                       bool   `json:"create"`
		Name                         string `json:"name"`
		AutomountServiceAccountToken bool   `json:"automountServiceAccountToken"`
		Annotations                  struct {
		} `json:"annotations"`
	} `json:"serviceAccount"`
	Rbac struct {
		Create bool          `json:"create"`
		Rules  []interface{} `json:"rules"`
	} `json:"rbac"`
	Metrics struct {
		Enabled bool `json:"enabled"`
		Image   struct {
			Registry    string        `json:"registry"`
			Repository  string        `json:"repository"`
			Tag         string        `json:"tag"`
			PullPolicy  string        `json:"pullPolicy"`
			PullSecrets []interface{} `json:"pullSecrets"`
		} `json:"image"`
		ExtraEnvVars       []interface{} `json:"extraEnvVars"`
		ExtraEnvVarsCM     string        `json:"extraEnvVarsCM"`
		ExtraEnvVarsSecret string        `json:"extraEnvVarsSecret"`
		ContainerPorts     struct {
			Http int `json:"http"`
		} `json:"containerPorts"`
		Resources struct {
			Limits struct {
			} `json:"limits"`
			Requests struct {
			} `json:"requests"`
		} `json:"resources"`
		PodSecurityContext struct {
			Enabled bool `json:"enabled"`
			FsGroup int  `json:"fsGroup"`
		} `json:"podSecurityContext"`
		ContainerSecurityContext struct {
			Enabled      bool `json:"enabled"`
			RunAsUser    int  `json:"runAsUser"`
			RunAsNonRoot bool `json:"runAsNonRoot"`
		} `json:"containerSecurityContext"`
		LifecycleHooks struct {
		} `json:"lifecycleHooks"`
		HostAliases []interface{} `json:"hostAliases"`
		PodLabels   struct {
		} `json:"podLabels"`
		PodAnnotations struct {
		} `json:"podAnnotations"`
		PodAffinityPreset     string `json:"podAffinityPreset"`
		PodAntiAffinityPreset string `json:"podAntiAffinityPreset"`
		NodeAffinityPreset    struct {
			Type   string        `json:"type"`
			Key    string        `json:"key"`
			Values []interface{} `json:"values"`
		} `json:"nodeAffinityPreset"`
		Affinity struct {
		} `json:"affinity"`
		NodeSelector struct {
		} `json:"nodeSelector"`
		Tolerations   []interface{} `json:"tolerations"`
		SchedulerName string        `json:"schedulerName"`
		Service       struct {
			Ports struct {
				Http int `json:"http"`
			} `json:"ports"`
			ClusterIP       string `json:"clusterIP"`
			SessionAffinity string `json:"sessionAffinity"`
			Annotations     struct {
				PrometheusIoScrape string `json:"prometheus.io/scrape"`
				PrometheusIoPort   string `json:"prometheus.io/port"`
			} `json:"annotations"`
		} `json:"service"`
		ServiceMonitor struct {
			Enabled       bool   `json:"enabled"`
			Namespace     string `json:"namespace"`
			Interval      string `json:"interval"`
			ScrapeTimeout string `json:"scrapeTimeout"`
			Labels        struct {
			} `json:"labels"`
			Selector struct {
			} `json:"selector"`
			Relabelings       []interface{} `json:"relabelings"`
			MetricRelabelings []interface{} `json:"metricRelabelings"`
			HonorLabels       bool          `json:"honorLabels"`
			JobLabel          string        `json:"jobLabel"`
		} `json:"serviceMonitor"`
	} `json:"metrics"`
	Postgresql struct {
		Enabled bool `json:"enabled"`
		Auth    struct {
			EnablePostgresUser bool   `json:"enablePostgresUser"`
			Username           string `json:"username"`
			Password           string `json:"password"`
			Database           string `json:"database"`
			ExistingSecret     string `json:"existingSecret"`
		} `json:"auth"`
		Architecture string `json:"architecture"`
	} `json:"postgresql"`
	ExternalDatabase struct {
		Host                      string `json:"host"`
		Port                      int    `json:"port"`
		User                      string `json:"user"`
		Database                  string `json:"database"`
		Password                  string `json:"password"`
		ExistingSecret            string `json:"existingSecret"`
		ExistingSecretPasswordKey string `json:"existingSecretPasswordKey"`
	} `json:"externalDatabase"`
	Redis struct {
		Enabled bool `json:"enabled"`
		Auth    struct {
			Enabled        bool   `json:"enabled"`
			Password       string `json:"password"`
			ExistingSecret string `json:"existingSecret"`
		} `json:"auth"`
		Architecture string `json:"architecture"`
	} `json:"redis"`
	ExternalRedis struct {
		Host                      string `json:"host"`
		Port                      int    `json:"port"`
		Username                  string `json:"username"`
		Password                  string `json:"password"`
		ExistingSecret            string `json:"existingSecret"`
		ExistingSecretPasswordKey string `json:"existingSecretPasswordKey"`
	} `json:"externalRedis"`
}
