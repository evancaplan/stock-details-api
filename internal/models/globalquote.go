package models

import (
	"time"
)

type GlobalQuote struct {
	Symbol 		string 	`json:"01. symbol"`
	Open		int		`json:"02. open"`
	High		int		`json:"03. high"`
	Low			int		`json:"04. low"`
	Price		float32	`json:"05. price"`
	Volume		int		`json:"06. volume"`
	LatestTradingDay		time.Time	`json:"07. latest trading day"`
	PreviousClose float32 `json:"08. previous close"`
	Change		float64 `json:"09. change"`
	ChangePercent string `json:"10. change percent"`
}