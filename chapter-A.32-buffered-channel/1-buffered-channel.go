package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}

	time.Sleep(1 * time.Second)
}
