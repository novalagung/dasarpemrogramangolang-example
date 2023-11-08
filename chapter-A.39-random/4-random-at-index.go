package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random int at index 3:", randomizer.Intn(3))
}
