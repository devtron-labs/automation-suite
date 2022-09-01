package ResponseDTOs

type AppListByTeamIdsResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		ProjectId   int    `json:"projectId"`
		ProjectName string `json:"projectName"`
		AppList     []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"appList"`
	} `json:"result"`
}
