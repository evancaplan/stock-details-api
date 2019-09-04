package details

import (
	"github.com/stock-details-api/internal/models"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/stock-details-api/internal/constants"
	"time"
	"net/http"
	"encoding/json"
	"github.com/stock-details-api/internal/utils"
)

func FindQuoteByTicker(w http.ResponseWriter, r *http.Request) {
	log:= utils.GetLogger()
	log.Info("details.findQuoteByTicker() reached ...")
	vars := mux.Vars(r)

	stockClient:= http.Client{Timeout: time.Second * 2,}

	req, err :=http.NewRequest(http.MethodGet, constants.BaseAlphaVantageUri + "/unction=GLOBAL_QUOTE&symbol=" + vars["ticker"] + "/&apikey=" + constants.AlphaVantageApiKey, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stockClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	quoteResponse := models.GlobalQuote{}
	jsonErr := json.Unmarshal(body, &quoteResponse)
	if jsonErr != nil {
		log.Fatal(err)
	}

	quoteResponseJSON, err := json.Marshal(&quoteResponse)
	if err != nil {
		log.Fatal(err)
	 }

	print(string(quoteResponseJSON))
}