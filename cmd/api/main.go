package main

import (
	"net/http"

	"github.com/stock-details-api/internal/routes"
	"github.com/stock-details-api/internal/utils"
)

func main() {
	log := utils.GetLogger()
	log.Info("Initializing routes.Routes()!!!")
	router := routes.Routes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
