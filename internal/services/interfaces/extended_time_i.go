package interfaces

import "net/http"

type ExtendedTime interface {
	findDaily(w http.ResponseWriter, ticker string, interval string)
	findWeekly(w http.ResponseWriter, ticker string, interval string)
	findMonthly(w http.ResponseWriter, ticker string, interval string)
}
