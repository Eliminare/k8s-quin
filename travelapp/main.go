package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Destination struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var destinations = []Destination{
	{ID: 1, Name: "Paris"},
	{ID: 2, Name: "New York"},
	{ID: 3, Name: "Tokyo"},
}

func destinationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(destinations)
}

func main() {
	http.HandleFunc("/destinations", destinationsHandler)
	log.Println("Travel app listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
