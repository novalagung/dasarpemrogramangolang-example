package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo dari GET /")
	})

	http.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "user id: %s\n", id)
	})

	http.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user")
	})

	fmt.Println("starting web server at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
