package domain

// ChartData holds the chart data
type ChartData struct {
	Category string `db:"category" json:"category"`
	Count    int64  `db:"count" json:"count"`
}
