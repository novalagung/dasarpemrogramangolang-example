package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random ke-1:", randomizer.Int())
	fmt.Println("random ke-2:", randomizer.Int())
	fmt.Println("random ke-3:", randomizer.Int())
}
