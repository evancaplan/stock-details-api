package mappers

import (
	"sort"
	"time"

	"github.com/stock-details-api/internal/utils"

	apiresponses "github.com/stock-details-api/internal/models/apiresponses"
	dtos "github.com/stock-details-api/internal/models/dtos"
)

func MapTimeSeriesDTOS(d apiresponses.Details) []dtos.ChartData {
	var MappedDetails []dtos.ChartData

	for date, details := range d {
		var data dtos.ChartData
		data.Date = date
		data.High = details.High
		data.Low = details.Low
		data.Open = details.Open

		MappedDetails = append(MappedDetails, data)

	}

	return sortDTOSNew(MappedDetails)
}

func sortDTOSNew(dtos []dtos.ChartData) []dtos.ChartData {
	println("Sorting ChartData DTOS...")
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
