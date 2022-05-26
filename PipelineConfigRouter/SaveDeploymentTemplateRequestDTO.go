package PipelineConfigRouter

type Spec struct {
	Affinity struct {
		Key    interface{} `json:"Key"`
		Values string      `json:"Values"`
		Key1   string      `json:"key"`
	} `json:"Affinity"`
}

type Ingress struct {
	Annotations struct {
	} `json:"annotations"`
	ClassName string `json:"className"`
	Enabled   bool   `json:"enabled"`
	Hosts     []struct {
		Host     string   `json:"host"`
		PathType string   `json:"pathType"`
		Paths    []string `json:"paths"`
	} `json:"hosts"`
	Labels struct {
	} `json:"labels"`
	Tls []interface{} `json:"tls"`
}
type SaveTemplateRequestDTO struct {
	AppId          int `json:"appId"`
	ChartRefId     int `json:"chartRefId"`
	ValuesOverride struct {
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
	} `json:"valuesOverride"`
	DefaultAppOverride struct {
		ContainerPort []struct {
			EnvoyPort        int    `json:"envoyPort"`
			IdleTimeout      string `json:"idleTimeout"`
			Name             string `json:"name"`
			Port             int    `json:"port"`
			ServicePort      int    `json:"servicePort"`
			SupportStreaming bool   `json:"supportStreaming"`
			UseHTTP2         bool   `json:"useHTTP2"`
		} `json:"ContainerPort"`
		EnvVariables    []interface{}  `json:"EnvVariables"`
		GracePeriod     int            `json:"GracePeriod"`
		LivenessProbe   LivenessProbe  `json:"LivenessProbe"`
		MaxSurge        int            `json:"MaxSurge"`
		MaxUnavailable  int            `json:"MaxUnavailable"`
		MinReadySeconds int            `json:"MinReadySeconds"`
		ReadinessProbe  ReadinessProbe `json:"ReadinessProbe"`
		Spec            Spec           `json:"Spec"`
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
}
