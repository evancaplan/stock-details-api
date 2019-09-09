package apiresponses

type DailyResponse struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		OutputSize    string `json:"4. Output Size"`
		Timezone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries Details `json:"Time Series (Daily)"`
}
