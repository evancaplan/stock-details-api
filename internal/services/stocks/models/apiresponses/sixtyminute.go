package apiresponses

type SixtyMinute struct {
	IntradayMetadata IntradayMetadata `json:"Meta Data"`
	TimeSeries       Details          `json:"Time Series (60min)"`
}
