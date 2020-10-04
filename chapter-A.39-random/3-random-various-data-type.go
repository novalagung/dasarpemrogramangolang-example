package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random int:", rand.Int())
	fmt.Println("random float32:", rand.Float32())
	fmt.Println("random uint:", rand.Uint32())
}
