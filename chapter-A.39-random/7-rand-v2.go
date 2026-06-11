package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// global random, auto-seeded
	fmt.Println("random int:", rand.Int())
	fmt.Println("random N:", rand.IntN(100))

	// reproducible random dengan seed tetap
	r := rand.New(rand.NewPCG(42, 0))
	fmt.Println("reproducible ke-1:", r.Int64())
	fmt.Println("reproducible ke-2:", r.Int64())
}
