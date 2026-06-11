package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	val int
}

func (c *counter) Add() {
	c.val++
}

func (c *counter) Value() int {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
