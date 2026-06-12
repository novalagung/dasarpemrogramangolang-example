package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/old-page", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-page", http.StatusMovedPermanently)
	})

	mux.HandleFunc("/new-page", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the new page!"))
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	})

	mux.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			http.Redirect(w, r, "/thank-you", http.StatusSeeOther)
			return
		}
		w.Write([]byte(`<form method="POST"><button type="submit">Submit</button></form>`))
	})

	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dashboard"))
	})

	mux.HandleFunc("/thank-you", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Thank you!"))
	})

	log.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
