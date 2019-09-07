package dtos

type Daily struct {
	Date string `json:"date"`
	Open string `json:"open"`
	High string `json:"high"`
	Low string `json:"low"`
	Volume string `json:"volume"`
}
