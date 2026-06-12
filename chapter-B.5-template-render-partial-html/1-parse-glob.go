package main

import (
	"html/template"
	"log"
	"net/http"
)

type M map[string]any

func main() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		err := tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("server started at localhost:9000")
	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
