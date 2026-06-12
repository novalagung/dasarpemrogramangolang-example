package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		// do the process here
		// simulate a long-time request by putting 10 seconds sleep
		time.Sleep(10 * time.Second)

		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if errors.Is(err, context.Canceled) {
				log.Println("request canceled")
			} else {
				log.Println("unknown error occurred.", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
