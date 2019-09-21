package stocks

import (
	"encoding/json"
	"github.com/stock-details-api/internal/constants"
	"github.com/stock-details-api/internal/services/stocks/models/apiresponses"
	"github.com/stock-details-api/internal/utils"
	"io/ioutil"
	"net/http"
)

type Search interface {
	FindClosestMatch(search string) apiresponses.BestMatchResponse
}

type SearchService struct{}

func(ss SearchService) FindClosestMatch(search string) apiresponses.BestMatchResponse {
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