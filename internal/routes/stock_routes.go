package routes

import (
	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/services/stocks"
	"github.com/stock-details-api/internal/utils"
)

func StockRoutes() *chi.Mux {
	log := utils.GetLogger()
	log.Info("In details.Routes(), registering new details routes")
	router := chi.NewRouter()
	router.Get("/{search}/quote", stocks.QuoteService{}.FindQuote)
	router.Get("/{ticker}/{interval}", stocks.ExtendedTimeService{}.FindTimeSeriesByTicker)
	router.Get("/{ticker}/intraday/{interval}", stocks.IntradayService{}.FindIntradayTimeSeriesByTicker)
	return router
}
