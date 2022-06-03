package ResponseDTOs

type TriggerChartSyncManualRespDTo struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
		Status string `json:"status"`
	} `json:"result"`
}
