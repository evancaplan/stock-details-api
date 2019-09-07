package details

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	apiresponses "github.com/stock-details-api/internal/models/apiresponses"
	modelMapper "github.com/stock-details-api/internal/models/mappers"

	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/constants"
	"github.com/stock-details-api/internal/utils"
)

func findQuoteBySearch(w http.ResponseWriter, r *http.Request) {

	log := utils.GetLogger()
	log.Info("details.findQuoteByTicker() reached ...")

	bestMatches := findClosestMatchToSearch(chi.URLParam(r, "search"))

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

func findClosestMatchToSearch(search string) apiresponses.BestMatchResponse {
	log := utils.GetLogger()
	println("Finding closest match to: ", search)

	searchClient := http.Client{}
	req, err := http.NewRequest("GET", constants.BaseAlphaVantageUri+"function=SYMBOL_SEARCH&keywords="+search+"&apikey="+constants.AlphaVantageApiKey, nil)
	if err != nil {
		log.Fatal(err)
	}
	println("Making request to : ", constants.BaseAlphaVantageUri+"function=SYMBOL_SEARCH&keywords="+search+"&apikey="+constants.AlphaVantageApiKey)

	req.Header.Add("Content-Type", "application/json")
	resp, err := searchClient.Do(req)
	if err != nil {
		print(err)
		log.Fatal(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var bestMatches apiresponses.BestMatchResponse
	err = json.Unmarshal(respBytes, &bestMatches)
	if err != nil {
		log.Fatal(err)
	}

	return bestMatches

}
