package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func index(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("Welcome to the API!!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api", index).Methods("GET")

	http.ListenAndServe(":5000", router)
}
