package ResponseDTOs

type GlobalChecklistResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		AppChecklist   AppChecklist   `json:"appChecklist"`
		ChartChecklist ChartChecklist `json:"chartChecklist"`
		IsAppCreated   bool           `json:"isAppCreated"`
	} `json:"result"`
}

type ChartChecklist struct {
	Project     int `json:"project"`
	Environment int `json:"environment"`
}

type AppChecklist struct {
	Project     int `json:"project"`
	Git         int `json:"git"`
	Environment int `json:"environment"`
	Docker      int `json:"docker"`
	HostUrl     int `json:"hostUrl"`
}
