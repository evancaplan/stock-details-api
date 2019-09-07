package apiresponses

type BestMatchResponse struct {
	BestMatches []struct {
		Symbol        string `json:"1. symbol"`
		Name          string `json:"2. name"`
		Type          string `json:"3. type"`
		Region        string `json:"4. region"`
		MarketOpen    string `json:"5. marketOpen"`
		MarketClose   string `json:"6. marketClose"`
		TimeZone      string `json:"7. timezone"`
		Currency      string `json:"8. previous close"`
		Change        string `json:"9. change"`
		ChangePercent string `json:"10. change percent"`
	} `json:"bestMatches"`
}
