package stocks_test

import (
	"github.com/stock-details-api/internal/services/mocks"
	"io"
	"net/http"
	"testing"
)

func TestExtendedTimeService_FindTimeSeriesByTicker(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	extendedTimeService := new(mocks.ExtendedTimeStock)

	extendedTimeService.On("FindTimeSeriesByTicker")
}