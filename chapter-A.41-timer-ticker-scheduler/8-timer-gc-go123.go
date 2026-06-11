package main

import (
	"fmt"
	"time"
)

func main() {
	// Sejak Go 1.23, time.After() tidak lagi menyebabkan memory leak dalam loop.
	// Untuk Go < 1.23, gunakan time.NewTimer() dengan .Stop() secara eksplisit.
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	fmt.Println("waiting...")
	<-timer.C
	fmt.Println("done")
}
