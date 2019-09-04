package details

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stock-details-api/internal/models"

	"github.com/go-chi/chi"

	"github.com/stock-details-api/internal/constants"
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

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var Quote models.QuoteResponse
	err = json.Unmarshal(respBytes, &Quote)
	if err != nil {
		log.Fatal(err)
	}

	QuoteJSON, err := json.Marshal(Quote)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(QuoteJSON)
}