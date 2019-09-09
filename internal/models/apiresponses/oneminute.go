package apiresponses

type OneMinute struct {
	IntradayMetadata IntradayMetadata `json:"Meta Data"`
	TimeSeries       Details          `json:"Time Series (1min)"`
}
