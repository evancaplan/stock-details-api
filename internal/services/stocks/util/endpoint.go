package util

import (
	"github.com/stock-details-api/internal/constants"
	"github.com/stock-details-api/internal/enums"
	enums2 "github.com/stock-details-api/internal/services/stocks/enums"
)

func CreateEndpoint(ticker string, interval string) string {
	if isIntradayTimeSeries(interval) {
		return createIntradayTimeSeriesEndpoint(ticker, interval)
	}
	println("started from the top now we here")
	return createTimeSeriesEndpoint(ticker, interval)

}

func isIntradayTimeSeries(interval string) bool {
	return interval == enums2.OneMin.String() ||
		interval == enums2.FiveMin.String() ||
		interval == enums2.FifteenMin.String() ||
		interval == enums2.ThirtyMin.String() ||
		interval == enums2.SixtyMin.String()
}

func createIntradayTimeSeriesEndpoint(ticker string, interval string) string {
	switch interval {
	case enums2.FiveMin.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums2.FiveMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	case enums2.FifteenMin.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums2.FifteenMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	case enums2.ThirtyMin.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums2.ThirtyMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	case enums2.SixtyMin.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums2.SixtyMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	default:
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums2.OneMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	}
}

func createTimeSeriesEndpoint(ticker string, interval string) string {
	switch interval {
	case enums.DailyAdjusted.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.DailyAdjusted.Endpoint() + "&symbol=" + ticker + "&apikey=" + constants.AlphaVantageApiKey
	case enums.Weekly.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Weekly.Endpoint() + "&symbol=" + ticker + "&apikey=" + constants.AlphaVantageApiKey
	case enums.WeeklyAdjusted.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.WeeklyAdjusted.Endpoint() + "&symbol=" + ticker + "&apikey=" + constants.AlphaVantageApiKey
	case enums.Monthly.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Monthly.Endpoint() + "&symbol=" + ticker + "&apikey=" + constants.AlphaVantageApiKey
	case enums.MonthlyAdjusted.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.MonthlyAdjusted.Endpoint() + "&symbol=" + ticker + "&apikey=" + constants.AlphaVantageApiKey
	default:
		return constants.BaseAlphaVantageUri + "function=" + enums.Daily.Endpoint() + "&symbol=" + ticker + "&apikey=" + constants.AlphaVantageApiKey
	}
}
