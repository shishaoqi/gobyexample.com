package main

import (
	"fmt"
	"sync"
)

type Countainer struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Countainer) inc(idx string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[idx]++
}

func main() {
	c := Countainer{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	doIncrement := func(idx string, n int) {
		for i := 0; i < n; i++ {
			c.inc(idx)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
