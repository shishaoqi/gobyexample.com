package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "tow"
	close(queue) // 注释掉后，报 atal error: all goroutines are asleep - deadlock!

	for elem := range queue {
		fmt.Println(elem)
	}
}
