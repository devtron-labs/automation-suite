package RequestDTOs

type CiMaterial struct {
	Source struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"source"`
	GitMaterialId   int    `json:"gitMaterialId"`
	Id              int    `json:"id"`
	GitMaterialName string `json:"gitMaterialName"`
}
type ConditionDetails struct {
	Id                  int    `json:"id"`
	ConditionOnVariable string `json:"conditionOnVariable"`
	ConditionType       string `json:"conditionType"`
	ConditionOperator   string `json:"conditionOperator"`
	ConditionalValue    string `json:"conditionalValue"`
}

type InputVariables struct {
	Id                        int    `json:"id"`
	Name                      string `json:"name"`
	Format                    string `json:"format"`
	Description               string `json:"description"`
	IsExposed                 bool   `json:"isExposed"`
	AllowEmptyValue           bool   `json:"allowEmptyValue"`
	DefaultValue              string `json:"defaultValue"`
	VariableType              string `json:"variableType"`
	VariableStepIndexInPlugin int    `json:"variableStepIndexInPlugin"`
	RefVariableStepIndex      int    `json:"refVariableStepIndex"`
	Value                     string `json:"value"`
	RefVariableName           string `json:"refVariableName"`
	RefVariableStage          string `json:"refVariableStage"`
}
type PortMap struct {
	PortOnLocal     int `json:"portOnLocal"`
	PortOnContainer int `json:"portOnContainer"`
}
type MountPathMap struct {
	FilePathOnDisk      string `json:"filePathOnDisk"`
	FilePathOnContainer string `json:"filePathOnContainer"`
}
type Step struct {
	Id                  int                 `json:"id"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	Index               int                 `json:"index"`
	StepType            string              `json:"stepType"`
	OutputDirectoryPath []string            `json:"outputDirectoryPath"`
	InlineStepDetail    InlineStepDetail    `json:"inlineStepDetail"`
	PluginRefStepDetail PluginRefStepDetail `json:"pluginRefStepDetail"`
}

type PluginRefStepDetail struct {
	Id               int                `json:"id"`
	PluginId         int                `json:"pluginId"`
	ConditionDetails []ConditionDetails `json:"conditionDetails"`
	InputVariables   []InputVariables   `json:"inputVariables"`
}

type InlineStepDetail struct {
	ScriptType               string             `json:"scriptType"`
	Script                   string             `json:"script"`
	StoreScriptAt            string             `json:"storeScriptAt"`
	CommandArgsMap           []CommandArgsMap   `json:"commandArgsMap"`
	InputVariables           []InputVariables   `json:"inputVariables"`
	OutputVariables          []InputVariables   `json:"outputVariables"`
	ConditionDetails         []ConditionDetails `json:"conditionDetails"`
	MountCodeToContainer     bool               `json:"mountCodeToContainer,omitempty"`
	MountCodeToContainerPath string             `json:"mountCodeToContainerPath,omitempty"`
	MountDirectoryFromHost   bool               `json:"mountDirectoryFromHost"`
	ContainerImagePath       string             `json:"containerImagePath,omitempty"`
	MountPathMap             []MountPathMap     `json:"mountPathMap,omitempty"`
	PortMap                  []PortMap          `json:"portMap,omitempty"`
	IsMountCustomScript      bool               `json:"isMountCustomScript,omitempty"`
}

type CommandArgsMap struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}
type Args struct {
	Arg string `json:"args"`
}

type BuildStage struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Steps []Step `json:"steps"`
}
type CiPipeline struct {
	IsManual   bool `json:"isManual"`
	DockerArgs struct {
		Progress string `json:"--progress"`
	} `json:"dockerArgs"`
	IsExternal       bool `json:"isExternal"`
	ParentCiPipeline int  `json:"parentCiPipeline"`
	ParentAppId      int  `json:"parentAppId"`
	ExternalCiConfig struct {
		Id         int    `json:"id"`
		WebhookUrl string `json:"webhookUrl"`
		Payload    string `json:"payload"`
		AccessKey  string `json:"accessKey"`
	} `json:"externalCiConfig"`
	CiMaterial     []CiMaterial `json:"ciMaterial"`
	Name           string       `json:"name"`
	Id             int          `json:"id"`
	Active         bool         `json:"active"`
	LinkedCount    int          `json:"linkedCount"`
	ScanEnabled    bool         `json:"scanEnabled"`
	AppWorkflowId  int          `json:"appWorkflowId"`
	PreBuildStage  BuildStage   `json:"preBuildStage"`
	PostBuildStage BuildStage   `json:"postBuildStage"`
}
type GetWorkflowDetails struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Result CiPipeline `json:"result"`
}

type CreateWorkflowRequestDto struct {
	AppId         int        `json:"appId"`
	AppWorkflowId int        `json:"appWorkflowId"`
	Action        int        `json:"action"`
	CiPipeline    CiPipeline `json:"ciPipeline"`
}
