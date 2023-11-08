package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random int:", randomizer.Int())
	fmt.Println("random float32:", randomizer.Float32())
	fmt.Println("random uint:", randomizer.Uint32())
}
