package util

import (
	"io/ioutil"
	"net/http"

	"github.com/stock-details-api/internal/utils"
)

func GetResponseBytes(ticker string, interval string) ([]byte, error) {

	log := utils.GetLogger()

	println("Get parameters: ticker: ", ticker, " interval: ", interval)

	endpoint := CreateEndpoint(ticker, interval)
	dailyClient := http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	println("Making request to : ", endpoint)

	req.Header.Add("Content-Type", "application/json")
	resp, err := dailyClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return respBytes, err
}
