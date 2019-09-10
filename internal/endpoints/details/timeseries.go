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

	respBytes, err := getResponseBytes(ticker, interval)
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
	log.Info("details.findWeeklyTimeSeriesByTicker() reached ...")

	respBytes, err := getResponseBytes(ticker, interval)
	if err != nil {
		log.Fatal(err)
	}

	println("Unmarshalling Response ...")
	var tsr apiresponses.WeeklyResponse
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
	log.Info("details.findWeeklyTimeSeriesByTicker() reached ...")

	respBytes, err := getResponseBytes(ticker, interval)
	if err != nil {
		log.Fatal(err)
	}

	println("Unmarshalling Response ...")
	var tsr apiresponses.MonthlyResponse
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
