package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		i := i // 注释掉，有神奇效果。为什么呢？
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
