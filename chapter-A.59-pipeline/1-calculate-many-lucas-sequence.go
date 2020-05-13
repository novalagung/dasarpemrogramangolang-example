package main

import (
	"log"
	"math/big"
	"time"

	"github.com/novalagung/gubrak"
)

const randomSeedMin = 0
const randomSeedMax = 10
const totalSeeds = 10000000
const sequenceIndexToFind = 20

type lucasSeed struct {
	L0 int
	L1 int
}

func generateRandomSeeds() []lucasSeed {
	seeds := make([]lucasSeed, 0)
	for i := 0; i < totalSeeds; i++ {
		seeds = append(seeds, lucasSeed{
			L0: gubrak.RandomInt(randomSeedMin, randomSeedMax),
			L1: gubrak.RandomInt(randomSeedMin, randomSeedMax),
		})
	}
	return seeds
}

type lucasSequence struct {
	Seed   lucasSeed
	Number int
}

func findLucasNumberAtSequenceX(seed lucasSeed) lucasSequence {
	current, addition := seed.L0, seed.L1
	for i := 0; i < sequenceIndexToFind; i++ {
		current, addition = addition, current+addition
	}
	return lucasSequence{
		Seed:   seed,
		Number: current,
	}
}

type lucasSequencePrimeResult struct {
	Sequence lucasSequence
	IsPrime  bool
}

func isPrime(sequence lucasSequence) lucasSequencePrimeResult {
	isPrime := big.NewInt(int64(sequence.Number)).ProbablyPrime(0)
	return lucasSequencePrimeResult{
		Sequence: sequence,
		IsPrime:  isPrime,
	}
}

func main() {
	log.Println("start")
	start := time.Now()

	seeds := generateRandomSeeds()
	results := make([]lucasSequencePrimeResult, 0)
	for _, seed := range seeds {
		sequence := findLucasNumberAtSequenceX(seed)
		sequenceResult := isPrime(sequence)
		results = append(results, sequenceResult)
	}

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
