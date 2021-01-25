package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kex5n/golangApp/version"
)

func home(w http.ResponseWriter, _ *http.Request) {
	info := struct {
		Buildtime string `json:"buildTime"`
		Commit string `json:"commit"`
		Release string `json:"release"`
	}{
		version.BuildTime, version.Commit, version.Release,
	}

	body, err := json.Marshal(info)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavalable), http.StatusServiceUnavalable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}