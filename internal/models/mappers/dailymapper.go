package mappers

import (
	"sort"
	"time"

	"github.com/stock-details-api/internal/utils"

	apiresponses "github.com/stock-details-api/internal/models/apiresponses"
	dtos "github.com/stock-details-api/internal/models/dtos"
)

func MapDailyDTOS(d apiresponses.DailyResponse) []dtos.Daily {
	var DailyDetails []dtos.Daily

	for date, details := range d.TimeSeries {
		var dailyDetail dtos.Daily
		dailyDetail.Date = date
		dailyDetail.High = details.High
		dailyDetail.Low = details.Low
		dailyDetail.Open = details.Open
		dailyDetail.Volume = details.Volume

		DailyDetails = append(DailyDetails, dailyDetail)

	}

	return sortDTOSNew(DailyDetails)
}

func sortDTOSNew(dtos []dtos.Daily) []dtos.Daily {
	println("Sorting Daily DTOS...")
	log := utils.GetLogger()

	layout := "2006-01-02"

	sort.Slice(dtos, func(i, j int) bool {
		timeOne, err := time.Parse(layout, dtos[i].Date)
		if err != nil {
			log.Fatal(err)
		}

		timeTwo, err := time.Parse(layout, dtos[j].Date)

		if err != nil {
			log.Fatal(err)
		}
		return timeOne.After(timeTwo)
	})

	return dtos
}
