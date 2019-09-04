package models

import (
	"time"
)

type Quote struct {
	Symbol 		string 	`json:"symbol"`
	Open		int		`json:"open"`
	High		int		`json:"high"`
	Low			int		`json:"loq"`
	Price		float32	`json:"price"`
	Volume		int		`json:"volume"`
	LatestTradingDay		time.Time	`json:"lastTradingDay"`
	PreviousClose float32 `json:"previousClose"`
	Change		float64 `json:"change"`
	ChangePercent string `json:"changePercent"`
}
