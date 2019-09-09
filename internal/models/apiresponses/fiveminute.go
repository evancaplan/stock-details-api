package apiresponses

type FiveMinute struct {
	IntradayMetadata IntradayMetadata `json:"Meta Data"`
	TimeSeries       Details          `json:"Time Series (5min)"`
}
