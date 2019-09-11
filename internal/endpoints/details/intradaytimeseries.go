package details

import (
	"encoding/json"
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
	log.Info("details.findOneMin() reached ...")

	respBytes, err := getResponseBytes(ticker, interval)
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
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findFiveMin(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findFiveMin() reached ...")

	respBytes, err := getResponseBytes(ticker, interval)
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
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findFifteenMin(w http.ResponseWriter, ticker string, interval string) {
	log := utils.GetLogger()
	log.Info("details.findFifteenMin() reached ...")

	respBytes, err := getResponseBytes(ticker, interval)
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
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findThirtyMin(w http.ResponseWriter, ticker string, interval string) {
	log := utils.GetLogger()
	log.Info("details.findThirtyMin() reached ...")

	respBytes, err := getResponseBytes(ticker, interval)
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
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func findSixtyMin(w http.ResponseWriter, ticker string, interval string) {
	log := utils.GetLogger()
	log.Info("details.findSixtyMin() reached ...")

	respBytes, err := getResponseBytes(ticker, interval)
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
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}
