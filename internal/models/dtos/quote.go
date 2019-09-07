package dtos

type Quote struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Open             string `json:"open"`
	High             string `json:"high"`
	Low              string `json:"loq"`
	Price            string `json:"price"`
	Volume           string `json:"volume"`
	LatestTradingDay string `json:"lastTradingDay"`
	PreviousClose    string `json:"previousClose"`
	Change           string `json:"change"`
	ChangePercent    string `json:"changePercent"`
}
