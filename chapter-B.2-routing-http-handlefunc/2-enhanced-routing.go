package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /articles", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("list semua artikel"))
	})

	http.HandleFunc("POST /articles", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("buat artikel baru"))
	})

	http.HandleFunc("GET /articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("article id: " + id))
	})

	log.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
