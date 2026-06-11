package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second) // wait for 10 seconds
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello !!", t)
		}
	}
}
