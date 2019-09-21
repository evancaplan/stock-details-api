package apiresponses

type FifteenMinute struct {
	IntradayMetadata IntradayMetadata `json:"Meta Data"`
	TimeSeries       Details          `json:"Time Series (15min)"`
}
