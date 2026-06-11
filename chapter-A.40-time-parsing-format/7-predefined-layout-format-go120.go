package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Format(time.DateTime))
	// contoh: 2026-06-11 14:30:00

	fmt.Println(now.Format(time.DateOnly))
	// contoh: 2026-06-11

	fmt.Println(now.Format(time.TimeOnly))
	// contoh: 14:30:00
}
