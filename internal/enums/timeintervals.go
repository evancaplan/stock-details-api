package enums

type TimeInterval int

const (
	Daily TimeInterval = iota
	DailyAdjusted
	Weekly
	WeeklyAdjusted
	Monthly
	MonthlyAdjusted
	Intraday
)

func (ti TimeInterval) String() string {
	return [...]string{"daily", "daily-adjusted", "weekly", "weekly-adjusted", "monthly", "monthly-adjusted", "intraday"}[ti]
}

func (ti TimeInterval) Endpoint() string {
	return [...]string{"TIME_SERIES_DAILY", "TIME_SERIES_DAILY_ADJUSTED", "TIME_SERIES_WEEKLY", "TIME_SERIES_WEEKLY_ADJUSTED", "TIME_SERIES_MONTHLY", "TIME_SERIES_MONTHLY_ADJUSTED", "TIME_SERIES_INTRADAY"}[ti]
}

func (ti TimeInterval) TimeSeriesKey() string {
	return [...]string{"Time Series (Daily)", "Time Series (Daily)", "Weekly Time Series", "Weekly Time Series", "Monthly Time Series", "Monthly Time Series", ""}[ti]
}

func GetTimeSeriesKey(interval string) string {
	println(interval)
	switch interval {
	case Daily.String():
		return Daily.TimeSeriesKey()
	case DailyAdjusted.String():
		return DailyAdjusted.TimeSeriesKey()
	case Weekly.String():
		return Weekly.TimeSeriesKey()
	case WeeklyAdjusted.String():
		return WeeklyAdjusted.TimeSeriesKey()
	case Monthly.String():
		return Monthly.TimeSeriesKey()
	case MonthlyAdjusted.String():
		return MonthlyAdjusted.TimeSeriesKey()
	}
	return "No Results Found"
}
