package details

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/constants"
	"github.com/stock-details-api/internal/models"
	"github.com/stock-details-api/internal/utils"
)

func findQuoteByTicker(w http.ResponseWriter, r *http.Request) {
	log := utils.GetLogger()
	log.Info("details.findQuoteByTicker() reached ...")
	println("GET parameters string : ", chi.URLParam(r, "ticker"))

	stockClient := http.Client{}
	req, err := http.NewRequest("GET", constants.BaseAlphaVantageUri+"function=GLOBAL_QUOTE&symbol="+chi.URLParam(r, "ticker")+"&apikey="+constants.AlphaVantageApiKey, nil)
	if err != nil {
		log.Fatal(err)
	}
	println("Making request to : ", constants.BaseAlphaVantageUri+"function=GLOBAL_QUOTE&symbol="+chi.URLParam(r, "ticker")+"&apikey="+constants.AlphaVantageApiKey)

	req.Header.Add("Content-Type", "application/json")
	resp, err := stockClient.Do(req)
	if err != nil {
		print(err)
		log.Fatal(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Quote models.QuoteResponse
	err = json.Unmarshal(respBytes, &Quote)
	if err != nil {
		log.Fatal(err)
	}

	Q := mapGlobalQuoteJSON(Quote.GlobalQuote)
	QuoteJSON, err := json.Marshal(Q)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(QuoteJSON)
}

func mapGlobalQuoteJSON(q models.GlobalQuote) models.Quote {
	var quote models.Quote

	quote.Change = q.Change
	quote.ChangePercent = q.ChangePercent
	quote.High = q.High
	quote.LatestTradingDay = q.LatestTradingDay
	quote.Low = q.Low
	quote.Open = q.Open
	quote.PreviousClose = q.PreviousClose
	quote.Price = q.Price
	quote.Symbol = q.Symbol
	quote.Volume = q.Volume

	return quote
}
