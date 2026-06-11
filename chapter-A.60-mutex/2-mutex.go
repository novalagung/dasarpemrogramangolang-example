package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	mu  sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
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
				meter.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
