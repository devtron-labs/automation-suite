package RequestDTOs

type SaveAppCiPipelineRequestDTO struct {
	AfterDockerBuild  []interface{} `json:"afterDockerBuild"`
	AppId             int           `json:"appId"`
	AppName           string        `json:"appName"`
	BeforeDockerBuild []interface{} `json:"beforeDockerBuild"`
	CiBuildConfig     CiBuildConfig `json:"ciBuildConfig"`
	DockerRegistry    string        `json:"dockerRegistry"`
	DockerRepository  string        `json:"dockerRepository"`
	Id                interface{}   `json:"id"`
}

type CiBuildConfig struct {
	BuildPackConfig   interface{}       `json:"buildPackConfig"`
	CiBuildType       string            `json:"ciBuildType"`
	DockerBuildConfig DockerBuildConfig `json:"dockerBuildConfig"`
	GitMaterialId     int               `json:"gitMaterialId"`
}

type DockerBuildConfig struct {
	Args struct {
	} `json:"args"`
	DockerfileContent      string `json:"dockerfileContent"`
	DockerfilePath         string `json:"dockerfilePath"`
	DockerfileRelativePath string `json:"dockerfileRelativePath"`
	DockerfileRepository   string `json:"dockerfileRepository"`
	TargetPlatform         string `json:"targetPlatform"`
}
