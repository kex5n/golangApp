package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRoouter()
	r.HandleFunc("/home", home).Method("GET")
	return r
}