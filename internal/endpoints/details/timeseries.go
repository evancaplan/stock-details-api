package details

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stock-details-api/internal/enums"

	apiresponses "github.com/stock-details-api/internal/models/apiresponses"
	mapper "github.com/stock-details-api/internal/models/mappers"

	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/utils"
)

func findTimeSeriesByTicker(w http.ResponseWriter, r *http.Request) {

	log := utils.GetLogger()
	log.Info("details.findTimeSeriesByTicker() reached ...")

	ticker := (chi.URLParam(r, "ticker"))
	interval := (chi.URLParam(r, "interval"))

	if interval == enums.Daily.String() || interval == enums.DailyAdjusted.String() {
		findDailyTimeSeriesByTicker(w, ticker, interval)
	}
	if interval == enums.Weekly.String() || interval == enums.WeeklyAdjusted.String() {
		findWeeklyTimeSeriesByTicker(w, ticker, interval)
	}
	if interval == enums.Monthly.String() || interval == enums.MonthlyAdjusted.String() {
		findMonthlyTimeSeriesByTicker(w, ticker, interval)
	}
}

func findDailyTimeSeriesByTicker(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findDailyTimeSeriesByTicker() reached ...")

	println("Get parameters: ticker: ", ticker, " interval: ", interval)

	endpoint := utils.CreateEndpoint(ticker, interval)
	dailyClient := http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	println("Making request to : ", endpoint)

	req.Header.Add("Content-Type", "application/json")
	resp, err := dailyClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tsr apiresponses.DailyResponse
	println("Unmarshalling Response ...")
	err = json.Unmarshal(respBytes, &tsr)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries, true)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findWeeklyTimeSeriesByTicker(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findDailyTimeSeriesByTicker() reached ...")

	println("Get parameters: ticker: ", ticker, " interval: ", interval)

	endpoint := utils.CreateEndpoint(ticker, interval)
	dailyClient := http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	println("Making request to : ", endpoint)

	req.Header.Add("Content-Type", "application/json")
	resp, err := dailyClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tsr apiresponses.WeeklyResponse
	println("Unmarshalling Response ...")
	err = json.Unmarshal(respBytes, &tsr)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries, true)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findMonthlyTimeSeriesByTicker(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findDailyTimeSeriesByTicker() reached ...")

	println("Get parameters: ticker: ", ticker, " interval: ", interval)

	endpoint := utils.CreateEndpoint(ticker, interval)
	dailyClient := http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	println("Making request to : ", endpoint)

	req.Header.Add("Content-Type", "application/json")
	resp, err := dailyClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tsr apiresponses.MonthlyResponse
	println("Unmarshalling Response ...")
	err = json.Unmarshal(respBytes, &tsr)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries, true)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}
