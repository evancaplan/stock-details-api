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

func findIntradayTimeSeriesByTicker(w http.ResponseWriter, r *http.Request) {

	log := utils.GetLogger()
	log.Info("details.findTimeSeriesByTicker() reached ...")

	ticker := (chi.URLParam(r, "ticker"))
	interval := (chi.URLParam(r, "interval"))

	if interval == enums.OneMin.String() {
		findOneMin(w, ticker, interval)
	}
	if interval == enums.FiveMin.String() {
		findFiveMin(w, ticker, interval)
	}
	if interval == enums.FifteenMin.String() {
		findFifteenMin(w, ticker, interval)
	}
	if interval == enums.ThirtyMin.String() {
		findThirtyMin(w, ticker, interval)
	}
	if interval == enums.SixtyMin.String() {
		findSixtyMin(w, ticker, interval)
	}
}

func findOneMin(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findOneMinTimeSeriesByTicker() reached ...")

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

	var tsr apiresponses.OneMinute
	println("Unmarshalling Response ...")
	err = json.Unmarshal(respBytes, &tsr)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries, false)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findFiveMin(w http.ResponseWriter, ticker string, interval string) {

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

	var tsr apiresponses.FiveMinute
	println("Unmarshalling Response ...")
	err = json.Unmarshal(respBytes, &tsr)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries, false)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findFifteenMin(w http.ResponseWriter, ticker string, interval string) {

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

	var tsr apiresponses.FifteenMinute
	println("Unmarshalling Response ...")
	err = json.Unmarshal(respBytes, &tsr)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries, false)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findThirtyMin(w http.ResponseWriter, ticker string, interval string) {

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

	var tsr apiresponses.ThirtyMinute
	println("Unmarshalling Response ...")
	err = json.Unmarshal(respBytes, &tsr)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries, false)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findSixtyMin(w http.ResponseWriter, ticker string, interval string) {

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

	var tsr apiresponses.SixtyMinute
	println("Unmarshalling Response ...")
	err = json.Unmarshal(respBytes, &tsr)
	if err != nil {
		log.Fatal(err)
	}

	println("Mapping DTOS ...")
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries, false)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}
