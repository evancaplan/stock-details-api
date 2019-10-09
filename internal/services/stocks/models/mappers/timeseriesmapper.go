package mappers

import (
	"github.com/stock-details-api/internal/services/stocks/models/apiresponses"
	"github.com/stock-details-api/internal/services/stocks/models/dtos"
)

type TimeSeriesMapper struct{}

func (t TimeSeriesMapper) Map(d apiresponses.Details) []dtos.ChartData {
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
