package apiresponses

type ThirtyMinute struct {
	IntradayMetadata IntradayMetadata `json:"Meta Data"`
	TimeSeries       Details          `json:"Time Series (30min)"`
}
