package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Richard Grayson", 24},
		{"Jason Todd", 23},
		{"Tim Drake", 22},
		{"Damian Wayne", 21},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", ActionIndex)

	log.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
