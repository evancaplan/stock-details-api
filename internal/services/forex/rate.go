package forex

import (
	"github.com/stock-details-api/internal/utils"
	"net/http"
)

func getExchangeRate(w http.ResponseWriter, r *http.Request) {
	log := utils.GetLogger()
	log.Info("details.findQuoteByTicker() reached ...")

	
}
