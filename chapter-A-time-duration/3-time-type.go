package main

import (
	"fmt"
	"time"
)

func main() {
	n := 5
	duration := time.Duration(n) * time.Second
	fmt.Println(duration.Seconds())
}
