package main

import (
	"net/http"
	"log"
	"github.com/kex5n/golangApp/handlers"
)

func main() {
	log.Print("Starting the service...")
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}