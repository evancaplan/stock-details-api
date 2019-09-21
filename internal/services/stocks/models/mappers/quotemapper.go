package mappers

import "github.com/stock-details-api/internal/services/stocks/models/dtos"
import "github.com/stock-details-api/internal/services/stocks/models/apiresponses"

func MapQuoteDTO(q apiresponses.QuoteResponse, name string) dtos.Quote {
	var quote dtos.Quote

	quote.Change = q.GlobalQuote.Change
	quote.ChangePercent = q.GlobalQuote.ChangePercent
	quote.High = q.GlobalQuote.High
	quote.LatestTradingDay = q.GlobalQuote.LatestTradingDay
	quote.Low = q.GlobalQuote.Low
	quote.Open = q.GlobalQuote.Open
	quote.PreviousClose = q.GlobalQuote.PreviousClose
	quote.Price = q.GlobalQuote.Price
	quote.Symbol = q.GlobalQuote.Symbol
	quote.Volume = q.GlobalQuote.Volume
	quote.Name = name

	return quote
}
