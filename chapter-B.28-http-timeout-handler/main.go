package main

import (
	"log"
	"net/http"
	"time"
)

func handleFast(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("response cepat"))
}

func handleSlow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Write([]byte("response lambat selesai"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /fast", http.TimeoutHandler(
		http.HandlerFunc(handleFast),
		3*time.Second,
		"request timeout",
	))

	mux.Handle("GET /slow", http.TimeoutHandler(
		http.HandlerFunc(handleSlow),
		3*time.Second,
		"request timeout: proses terlalu lama",
	))

	log.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
