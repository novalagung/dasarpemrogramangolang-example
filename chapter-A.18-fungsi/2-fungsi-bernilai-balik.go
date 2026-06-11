package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
}

func randomWithRange(minNum, maxNum int) int {
	var value = randomizer.Int()%(maxNum-minNum+1) + minNum
	return value
}
