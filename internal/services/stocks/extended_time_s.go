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

type ExtendedTimeStock interface {
	FindTimeSeriesByTicker(w http.ResponseWriter, r *http.Request)
}

type ExtendedTimeService struct{}

func(ets ExtendedTimeService) FindTimeSeriesByTicker(w http.ResponseWriter, r *http.Request) {

	log := utils.GetLogger()
	log.Info("details.findTimeSeriesByTicker() reached ...")

	ticker := chi.URLParam(r, "ticker")
	interval := chi.URLParam(r, "interval")

	if interval == enums.Daily.String() || interval == enums.DailyAdjusted.String() {
		_, err := w.Write(ets.findDaily(ticker, interval))
		if err != nil {
			log.Fatal(err)
		}
	}
	if interval == enums.Weekly.String() || interval == enums.WeeklyAdjusted.String() {
		_, err := w.Write(ets.findWeekly(ticker, interval))
		if err != nil {
			log.Fatal(err)
		}
	}
	if interval == enums.Monthly.String() || interval == enums.MonthlyAdjusted.String() {
		_, err := w.Write(ets.findMonthly(ticker, interval))
		if err != nil {
			log.Fatal(err)
		}
	}
}


func(ets ExtendedTimeService) findDaily(ticker string, interval string) []byte {

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
	var Details = mapper.TimeSeriesMapper{}.Map(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	return DetailsJSON
}

func(ets ExtendedTimeService) findWeekly(ticker string, interval string) []byte {

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
	var Details = mapper.TimeSeriesMapper{}.Map(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	return DetailsJSON
}

func(ets ExtendedTimeService) findMonthly(ticker string, interval string) []byte {
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
	var Details = mapper.TimeSeriesMapper{}.Map(tsr.TimeSeries)
	DetailsJSON, err := json.Marshal(Details)
	if err != nil {
		log.Fatal(err)
	}

	return DetailsJSON
}
