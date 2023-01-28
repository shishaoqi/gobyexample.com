package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *container) inc(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[s]++
}

func main() {
	con := container{
		counter: map[string]int{"a": 0, "b": 0},
	}

	doIncrement := func(idx string, n int, wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			con.inc(idx)
		}
		wg.Done()
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go doIncrement("a", 10000, &wg)
	go doIncrement("a", 10000, &wg)
	go doIncrement("b", 10000, &wg)

	wg.Wait()
	fmt.Println(con)
}
