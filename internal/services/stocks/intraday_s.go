package stocks

import (
	"encoding/json"
	enums2 "github.com/stock-details-api/internal/services/stocks/enums"
	"github.com/stock-details-api/internal/services/stocks/util"
	"net/http"

	"github.com/stock-details-api/internal/services/stocks/models/apiresponses"
	mapper "github.com/stock-details-api/internal/services/stocks/models/mappers"

	"github.com/go-chi/chi"
	"github.com/stock-details-api/internal/utils"
)

type Intraday interface {
	FindIntradayTimeSeriesByTicker(w http.ResponseWriter, r *http.Request)
}

type IntradayService struct {}

func(is IntradayService) FindIntradayTimeSeriesByTicker(w http.ResponseWriter, r *http.Request) {

	log := utils.GetLogger()
	log.Info("details.findTimeSeriesByTicker() reached ...")

	ticker := (chi.URLParam(r, "ticker"))
	interval := (chi.URLParam(r, "interval"))

	if interval == enums2.OneMin.String() {
		is.findOneMin(w, ticker, interval)
	}
	if interval == enums2.FiveMin.String() {
		is.findFiveMin(w, ticker, interval)
	}
	if interval == enums2.FifteenMin.String() {
		is.findFifteenMin(w, ticker, interval)
	}
	if interval == enums2.ThirtyMin.String() {
		is.findThirtyMin(w, ticker, interval)
	}
	if interval == enums2.SixtyMin.String() {
		is.findSixtyMin(w, ticker, interval)
	}
}

func(is IntradayService) findOneMin(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findOneMin() reached ...")

	respBytes, err := util.GetResponseBytes(ticker, interval)
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

func(is IntradayService) findFiveMin(w http.ResponseWriter, ticker string, interval string) {

	log := utils.GetLogger()
	log.Info("details.findFiveMin() reached ...")

	respBytes, err := util.GetResponseBytes(ticker, interval)
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

func(is IntradayService) findFifteenMin(w http.ResponseWriter, ticker string, interval string) {
	log := utils.GetLogger()
	log.Info("details.findFifteenMin() reached ...")

	respBytes, err := util.GetResponseBytes(ticker, interval)
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

func(is IntradayService) findThirtyMin(w http.ResponseWriter, ticker string, interval string) {
	log := utils.GetLogger()
	log.Info("details.findThirtyMin() reached ...")

	respBytes, err := util.GetResponseBytes(ticker, interval)
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

func(is IntradayService) findSixtyMin(w http.ResponseWriter, ticker string, interval string) {
	log := utils.GetLogger()
	log.Info("details.findSixtyMin() reached ...")

	respBytes, err := util.GetResponseBytes(ticker, interval)
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
