package forex

import (
	"github.com/stock-details-api/internal/utils"
	"net/http"
)

type ExchangeRate interface {
	getExchangeRate(w http.ResponseWriter, r *http.Request)
}

type ExchangeRateService struct {}

func(ers ExchangeRateService) getExchangeRate(w http.ResponseWriter, r *http.Request) {
	log := utils.GetLogger()
	log.Info("details.findQuoteByTicker() reached ...")

}
