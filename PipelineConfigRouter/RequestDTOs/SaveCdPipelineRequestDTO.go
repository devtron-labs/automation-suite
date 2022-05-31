package RequestDTOs

type Stage struct {
	Config      string `json:"config"`
	TriggerType string `json:"triggerType"`
	Switch      string `json:"switch"`
}
type Rolling struct {
	MaxSurge       string `json:"maxSurge"`
	MaxUnavailable int    `json:"maxUnavailable"`
}
type Strategies struct {
	DeploymentTemplate string `json:"deploymentTemplate"`
	Config             struct {
		Deployment struct {
			Strategy struct {
				Rolling Rolling `json:"rolling"`
			} `json:"strategy"`
		} `json:"deployment"`
	} `json:"config"`
	Default bool `json:"default"`
}

type StageConfigMapSecretNames struct {
	ConfigMaps []string `json:"configMaps"`
	Secrets    []string `json:"secrets"`
}

type Pipeline struct {
	AppWorkflowId                 int                       `json:"appWorkflowId"`
	EnvironmentId                 int                       `json:"environmentId"`
	CiPipelineId                  int                       `json:"ciPipelineId"`
	TriggerType                   string                    `json:"triggerType"`
	Name                          string                    `json:"name"`
	Strategies                    []Strategies              `json:"strategies"`
	Namespace                     string                    `json:"namespace"`
	PreStage                      Stage                     `json:"preStage"`
	PostStage                     Stage                     `json:"postStage"`
	PreStageConfigMapSecretNames  StageConfigMapSecretNames `json:"preStageConfigMapSecretNames"`
	PostStageConfigMapSecretNames StageConfigMapSecretNames `json:"postStageConfigMapSecretNames"`
	RunPreStageInEnv              bool                      `json:"runPreStageInEnv"`
	RunPostStageInEnv             bool                      `json:"runPostStageInEnv"`
	IsClusterCdActive             bool                      `json:"isClusterCdActive"`
	ParentPipelineId              int                       `json:"parentPipelineId"`
	ParentPipelineType            string                    `json:"parentPipelineType"`
	DeploymentTemplate            string                    `json:"deploymentTemplate"`
}

type SaveCdPipelineRequestDTO struct {
	AppId     int        `json:"appId"`
	Pipelines []Pipeline `json:"pipelines"`
}
