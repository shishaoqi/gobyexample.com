package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(i int, s int) {
	fmt.Printf("worker %d starting\n", i)
	d := time.Duration(s) * time.Second
	time.Sleep(d)
	fmt.Printf("worker %d use %d second to finished\n", i, s)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		i := i // 如果没这句，要把下面的 n 改写 i
		go func(n int) {
			rand.Seed(time.Now().UnixNano())
			s := rand.Intn(9)
			defer wg.Done()
			worker(i, s)
		}(i)
	}

	wg.Wait()
}
