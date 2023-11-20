package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Pizza struct {
	Slice int `form:"slice" json:"slice"`
	Price int `form:"price" json:"price"`
}

func main() {
	http.HandleFunc("/sendpizza", func(w http.ResponseWriter, r *http.Request) {
		pizza := new(Pizza)
		dec := json.NewDecoder(r.Body)

		err := dec.Decode(&pizza)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("[ERROR]:" + err.Error())
		}
		pizzaBytes, err := json.Marshal(pizza)
		err = PushPizzaQueue("Pizza", pizzaBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("[ERROR]:" + err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(pizza)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
