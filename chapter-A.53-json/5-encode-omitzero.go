package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date,omitzero"`
	Score     int       `json:"score,omitzero"`
}

func main() {
	event := Event{Name: "Go Conference"}

	jsonData, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(jsonData))
	// {"name":"Go Conference"}
}
