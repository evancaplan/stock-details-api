package mappers

import (
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
	return MappedDetails
}
