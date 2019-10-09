package interfaces

type Intraday interface {
	findOneMin(ticker string, interval string)
	findFiveMin(ticker string, interval string)
	findFifteenMin(ticker string, interval string)
	findThirtyMin(ticker string, interval string)
	findSixtyMin(ticker string, interval string)
}
