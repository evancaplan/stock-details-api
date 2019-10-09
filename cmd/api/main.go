package main

import (
	"net/http"

	"github.com/stock-details-api/internal/handlers"
	"github.com/stock-details-api/internal/utils"
)

func main() {
	log := utils.GetLogger()
	log.Info("Initializing handlers.Routes()!!!")
	router := handlers.Routes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
