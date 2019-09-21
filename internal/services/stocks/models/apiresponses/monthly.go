package apiresponses

type MonthlyResponse struct {
	MetaData   MetaData `json:"Meta Data"`
	TimeSeries Details  `json:"Monthly Time Series"`
}