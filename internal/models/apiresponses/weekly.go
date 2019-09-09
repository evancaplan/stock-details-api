package apiresponses

type WeeklyResponse struct {
	MetaData   MetaData `json:"Meta Data"`
	TimeSeries Details `json:"Weekly Time Series"`
}
