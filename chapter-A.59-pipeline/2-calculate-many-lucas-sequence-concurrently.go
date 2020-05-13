package main

import (
	"log"
	"math/big"
	"sync"
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

func generateRandomSeeds() <-chan lucasSeed {
	chanOut := make(chan lucasSeed)

	go func() {
		for i := 0; i < totalSeeds; i++ {
			chanOut <- lucasSeed{
				L0: gubrak.RandomInt(randomSeedMin, randomSeedMax),
				L1: gubrak.RandomInt(randomSeedMin, randomSeedMax),
			}
		}
		close(chanOut)
	}()

	return chanOut
}

type lucasSequence struct {
	Seed   lucasSeed
	Number int
}

func findLucasNumberAtSequenceX(chanIn <-chan lucasSeed) <-chan lucasSequence {
	chanOut := make(chan lucasSequence)

	go func() {
		for seed := range chanIn {
			current, addition := seed.L0, seed.L1
			for i := 0; i < sequenceIndexToFind; i++ {
				current, addition = addition, current+addition
			}

			chanOut <- lucasSequence{
				Seed:   seed,
				Number: current,
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func mergeChanLucasNumber(chanInMany ...<-chan lucasSequence) <-chan lucasSequence {
	wg := new(sync.WaitGroup)
	chanOut := make(chan lucasSequence)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan lucasSequence) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

type lucasSequencePrimeResult struct {
	Sequence lucasSequence
	IsPrime  bool
}

func isPrime(chanIn <-chan lucasSequence) <-chan lucasSequencePrimeResult {
	chanOut := make(chan lucasSequencePrimeResult)

	go func() {
		for sequence := range chanIn {
			isPrime := big.NewInt(int64(sequence.Number)).ProbablyPrime(0)

			chanOut <- lucasSequencePrimeResult{
				Sequence: sequence,
				IsPrime:  isPrime,
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func mergeChanIsPrime(chanInMany ...<-chan lucasSequencePrimeResult) <-chan lucasSequencePrimeResult {
	wg := new(sync.WaitGroup)
	chanOut := make(chan lucasSequencePrimeResult)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan lucasSequencePrimeResult) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func main() {
	log.Println("start")
	start := time.Now()

	// pipeline 1: generate random seeds
	chanSeed := generateRandomSeeds()

	// pipeline 2: find lucas sequences using three worker, then merge it
	chanNumber1 := findLucasNumberAtSequenceX(chanSeed)
	chanNumber2 := findLucasNumberAtSequenceX(chanSeed)
	chanNumber3 := findLucasNumberAtSequenceX(chanSeed)
	chanNumber := mergeChanLucasNumber(chanNumber1, chanNumber2, chanNumber3)

	// pipeline 3: detect prime from sequences using three worker, then merge it
	chanResult1 := isPrime(chanNumber)
	chanResult2 := isPrime(chanNumber)
	chanResult3 := isPrime(chanNumber)
	chanResult := mergeChanIsPrime(chanResult1, chanResult2, chanResult3)

	// print results
	results := make([]lucasSequencePrimeResult, 0)
	for eachResult := range chanResult {
		results = append(results, eachResult)
	}

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
