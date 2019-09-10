package enums

type IntradayTimeInterval int

const (
	OneMin IntradayTimeInterval = iota
	FiveMin
	FifteenMin
	ThirtyMin
	SixtyMin
)

func (iti IntradayTimeInterval) String() string {
	return [...]string{"onemin", "fivemin", "fifteenmin", "thirtymin", "sixtymin"}[iti]
}

func (iti IntradayTimeInterval) Endpoint() string {
	return [...]string{"1min", "5min", "15min", "30min", "60min"}[iti]
}

func (iti IntradayTimeInterval) TimeSeriesKey() string {
	return [...]string{"Time Series (1min)", "Time Series (5min)", "Time Series (15min)", "Time Series (30min)", "Time Series (60min)"}[iti]
}

func GetIntradayTimeSeriesKey(interval string) string {
	switch interval {
	case OneMin.String():
		return OneMin.TimeSeriesKey()
	case FiveMin.String():
		return FiveMin.TimeSeriesKey()
	case FifteenMin.String():
		return FifteenMin.TimeSeriesKey()
	case ThirtyMin.String():
		return ThirtyMin.TimeSeriesKey()
	case SixtyMin.String():
		return SixtyMin.TimeSeriesKey()
	}
	return "No Results Found"
}
