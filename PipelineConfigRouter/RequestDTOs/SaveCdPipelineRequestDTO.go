package RequestDTOs

import "automation-suite/PipelineConfigRouter"

type Stage struct {
	Config      string `json:"config"`
	TriggerType string `json:"triggerType"`
	Switch      string `json:"switch"`
}

type Strategies struct {
	DeploymentTemplate string `json:"deploymentTemplate"`
	Config             struct {
		Deployment struct {
			Strategy struct {
				Rolling PipelineConfigRouter.Rolling `json:"rolling"`
			} `json:"strategy"`
		} `json:"deployment"`
	} `json:"config"`
	Default bool `json:"default"`
}

type PreStageConfigMapSecretNames struct {
	ConfigMaps []string `json:"configMaps"`
	Secrets    []string `json:"secrets"`
}

type PostStageConfigMapSecretNames struct {
	ConfigMaps []interface{} `json:"configMaps"`
	Secrets    []interface{} `json:"secrets"`
}

type Pipeline struct {
	AppWorkflowId                 int                           `json:"appWorkflowId"`
	EnvironmentId                 int                           `json:"environmentId"`
	CiPipelineId                  int                           `json:"ciPipelineId"`
	TriggerType                   string                        `json:"triggerType"`
	Name                          string                        `json:"name"`
	Strategies                    []Strategies                  `json:"strategies"`
	Namespace                     string                        `json:"namespace"`
	PreStage                      Stage                         `json:"preStage"`
	PostStage                     Stage                         `json:"postStage"`
	PreStageConfigMapSecretNames  PreStageConfigMapSecretNames  `json:"preStageConfigMapSecretNames"`
	PostStageConfigMapSecretNames PostStageConfigMapSecretNames `json:"postStageConfigMapSecretNames"`
	RunPreStageInEnv              bool                          `json:"runPreStageInEnv"`
	RunPostStageInEnv             bool                          `json:"runPostStageInEnv"`
	IsClusterCdActive             bool                          `json:"isClusterCdActive"`
	ParentPipelineId              int                           `json:"parentPipelineId"`
	ParentPipelineType            string                        `json:"parentPipelineType"`
	DeploymentTemplate            string                        `json:"deploymentTemplate"`
}

type SaveCdPipelineRequestDTO struct {
	AppId     int        `json:"appId"`
	Pipelines []Pipeline `json:"pipelines"`
}
