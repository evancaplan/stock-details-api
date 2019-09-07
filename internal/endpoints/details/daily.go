package details

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	apiresponses "github.com/stock-details-api/internal/models/apiresponses"
	mapper "github.com/stock-details-api/internal/models/mappers"

	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/constants"
	"github.com/stock-details-api/internal/utils"
)

func findDailyDetails(w http.ResponseWriter, r *http.Request) {

	log := utils.GetLogger()
	log.Info("details.findDailyDetails() reached ...")

	ticker := (chi.URLParam(r, "ticker"))

	println("Get parameters: ticker: ", ticker)

	dailyClient := http.Client{}
	req, err := http.NewRequest("GET", constants.BaseAlphaVantageUri+"function=TIME_SERIES_DAILY&symbol="+ticker+"&apikey="+constants.AlphaVantageApiKey, nil)
	if err != nil {
		log.Fatal(err)
	}

	println("Making request to : ", constants.BaseAlphaVantageUri+"function=TIME_SERIES_DAILY&symbol="+ticker+"&apikey="+constants.AlphaVantageApiKey)

	req.Header.Add("Content-Type", "application/json")
	resp, err := dailyClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	println("Unmarshalling Response ...")
	var dR apiresponses.DailyResponse
	err = json.Unmarshal(respBytes, &dR)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapDailyDTOS(dR)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)

}
