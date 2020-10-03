package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()

	duration := t2.Sub(t1)

	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())
}
