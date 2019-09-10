package details

import (
	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/utils"
)

func Routes() *chi.Mux {
	log := utils.GetLogger()
	log.Info("In details.Routes(), registering new details routes")
	router := chi.NewRouter()
	router.Get("/{search}/quote", findQuoteBySearch)
	router.Get("/{ticker}/{interval}", findTimeSeriesByTicker)
	router.Get("/{ticker}/intraday/{interval}", findIntradayTimeSeriesByTicker)
	return router
}
