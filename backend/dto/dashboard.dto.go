package dto

type DonutChartData struct {
	FilterBy string           `json:"filter_by"`
	Data     map[string]int64 `json:"data"`
}

type BarChartData struct {
	Data []ReviewerStats `json:"data"`
}

type ReviewerStats struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}
