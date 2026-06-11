package main

import "fmt"

func main() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
	}

	fmt.Println("sebelum clear:", len(chicken), chicken)

	clear(chicken)

	fmt.Println("sesudah clear:", len(chicken), chicken)
}
