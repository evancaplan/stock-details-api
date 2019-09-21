package apiresponses

type ForexDetails struct {
	fromCode string `1. json:"From_Ccurrency Code`
	fromName string `2. json:"From_Currency Name`
	toCode   string `3. json:"To_Currency Code"`
	toName   string `4. json:"To_Currency Name`
	rate     string `5. json:"Exchange Rate"`
	refresh  string `6. json:"Last Refreshed"`
	timeZone string `7. json:"Time Zone`
	bid      string `json:"8. Bid Price`
	ask      string `json:"9. Ask Price"`
}
