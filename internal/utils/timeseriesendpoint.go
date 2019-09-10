package utils

import (
	"github.com/stock-details-api/internal/constants"
	"github.com/stock-details-api/internal/enums"
)

func CreateEndpoint(ticker string, interval string) string {
	println(interval + " Suck my clit and balls")
	if isIntradayTimeSeries(interval) {
		return createIntradayTimeSeriesEndpoint(ticker, interval)
	}
	println("started from the top now we here")
	return createTimeSeriesEndpoint(ticker, interval)

}

func isIntradayTimeSeries(interval string) bool {
	return interval == enums.OneMin.String() ||
		interval == enums.FiveMin.String() ||
		interval == enums.FifteenMin.String() ||
		interval == enums.ThirtyMin.String() ||
		interval == enums.SixtyMin.String()
}

func createIntradayTimeSeriesEndpoint(ticker string, interval string) string {
	switch interval {
	case enums.FiveMin.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums.FiveMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	case enums.FifteenMin.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums.FifteenMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	case enums.ThirtyMin.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums.ThirtyMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	case enums.SixtyMin.String():
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums.SixtyMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
	default:
		return constants.BaseAlphaVantageUri + "function=" + enums.Intraday.Endpoint() + "&symbol=" + ticker + "&interval=" + enums.OneMin.Endpoint() + "&apikey=" + constants.AlphaVantageApiKey
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
