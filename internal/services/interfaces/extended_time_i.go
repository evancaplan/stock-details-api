package interfaces

type ExtendedTime interface {
	findDaily(ticker string, interval string)
	findWeekly(ticker string, interval string)
	findMonthly(ticker string, interval string)
}
