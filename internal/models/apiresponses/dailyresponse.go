package apiresponses

type DailyResponse struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		OutputSize    string `json:"4. Output Size"`
		Timezone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries map[string]struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	} `json:"Time Series (Daily)"`
}

// type Details struct {
// 	Open   string `json:"1. open"`
// 	High   string `json:"2. high"`
// 	Low    string `json:"3. low"`
// 	Close  string `json:"4. close"`
// 	Volume string `json:"5. volume"`
// }
