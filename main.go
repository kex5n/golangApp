package main

import (
	"net/http"
	"log"
	"github.com/kex5n/golangApp/handlers"
	"os"
)

func main() {
	log.Print("Starting the service...¥ncommit: %s, build time: %s, release: %s",
	version.Commit, version.BuildTime, version.Release,
)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}