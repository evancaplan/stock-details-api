package stocks

import (
	"encoding/json"
	"github.com/stock-details-api/internal/services/stocks/util"
	"net/http"

	"github.com/stock-details-api/internal/enums"

	"github.com/stock-details-api/internal/services/stocks/models/apiresponses"
	mapper "github.com/stock-details-api/internal/services/stocks/models/mappers"

	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/utils"
)

type ExtendedTime interface {
	FindTimeSeriesByTicker(w http.ResponseWriter, r *http.Request)
}

type ExtendedTimeService struct{}

func(ets ExtendedTimeService) FindTimeSeriesByTicker(w http.ResponseWriter, r *http.Request) {

	log := utils.GetLogger()
	log.Info("details.findTimeSeriesByTicker() reached ...")

	ticker := (chi.URLParam(r, "ticker"))
	interval := (chi.URLParam(r, "interval"))

	if interval == enums.Daily.String() || interval == enums.DailyAdjusted.String() {
		ets.findDaily(w, ticker, interval)
	}
	if interval == enums.Weekly.String() || interval == enums.WeeklyAdjusted.String() {
		ets.findWeekly(w, ticker, interval)
	}
	if interval == enums.Monthly.String() || interval == enums.MonthlyAdjusted.String() {
		ets.findMonthly(w, ticker, interval)
	}
}


func(ets ExtendedTimeService) findDaily(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findDailyTimeSeriesByTicker() reached ...")

	respBytes, err := util.GetResponseBytes(ticker, interval)
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
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func(ets ExtendedTimeService) findWeekly(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findWeeklyTimeSeriesByTicker() reached ...")

	respBytes, err := util.GetResponseBytes(ticker, interval)
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
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}

func(ets ExtendedTimeService) findMonthly(w http.ResponseWriter, ticker string, interval string) {
	log := utils.GetLogger()
	log.Info("details.findWeeklyTimeSeriesByTicker() reached ...")

	respBytes, err := util.GetResponseBytes(ticker, interval)
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
	var Details = mapper.MapTimeSeriesDTOS(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(DetailsJSON)
}
