package stocks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stock-details-api/internal/services/stocks/models/apiresponses"
	modelMapper "github.com/stock-details-api/internal/services/stocks/models/mappers"

	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/constants"
	"github.com/stock-details-api/internal/utils"
)

type Quote interface {
	FindQuote(w http.ResponseWriter, r *http.Request)
}

type QuoteService struct {}

func(qs QuoteService) FindQuote(w http.ResponseWriter, r *http.Request) {

	log := utils.GetLogger()
	log.Info("details.findQuoteByTicker() reached ...")

	bestMatches := SearchService{}.FindClosestMatch(chi.URLParam(r, "search"))

	ticker := bestMatches.BestMatches[0].Symbol

	println("GET parameters string : ", ticker)

	stockClient := http.Client{}
	req, err := http.NewRequest("GET", constants.BaseAlphaVantageUri+"function=GLOBAL_QUOTE&symbol="+ticker+"&apikey="+constants.AlphaVantageApiKey, nil)
	if err != nil {
		log.Fatal(err)
	}
	println("Making request to : ", constants.BaseAlphaVantageUri+"function=GLOBAL_QUOTE&symbol="+ticker+"&apikey="+constants.AlphaVantageApiKey)

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

	var Quote apiresponses.QuoteResponse
	err = json.Unmarshal(respBytes, &Quote)
	if err != nil {
		log.Fatal(err)
	}

	name := bestMatches.BestMatches[0].Name

	Q := modelMapper.MapQuoteDTO(Quote, name)
	QuoteJSON, err := json.Marshal(Q)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(QuoteJSON)
}


