package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomizer := rand.New(rand.NewSource(10))
	fmt.Println("random ke-1:", randomizer.Int()) // 5221277731205826435
	fmt.Println("random ke-2:", randomizer.Int()) // 3852159813000522384
	fmt.Println("random ke-3:", randomizer.Int()) // 8532807521486154107
}
