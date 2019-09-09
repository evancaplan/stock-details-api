package utils

import (
	"github.com/stock-details-api/internal/enums"
	"github.com/stock-details-api/internal/constants"
)

func CreateEndpoint(ticker string, interval string) string {
	switch interval {
	case enums.Intraday: 
		return createIntradayTimeSeriesEndpoint(ticker, interval)
	default:
		return createTimeSeriesEndpoint(ticker, interval)
}
	
}

func createIntradayTimeSeriesEndpoint(ticker string, interval string) {
	switch interval {
		
	}
}

func createTimeSeriesEndpoint(ticker string, interval string) {
	switch interval {
default:
	return constants.BaseAlphaVantageUri + "function=" + enums.Daily.Endpoint() + "&symbol=" + ticker + "&apikey=" + constants.AlphaVantageApiKey
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
}
}
