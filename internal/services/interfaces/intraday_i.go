package interfaces

import "net/http"

type Intraday interface {
	findOneMin(w http.ResponseWriter, ticker string, interval string)
	findFiveMin(w http.ResponseWriter, ticker string, interval string)
	findFifteenMin(w http.ResponseWriter, ticker string, interval string)
	findThirtyMin(w http.ResponseWriter, ticker string, interval string)
	findSixtyMin(w http.ResponseWriter, ticker string, interval string)
}
