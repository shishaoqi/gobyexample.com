package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	for i := 0; i < 10; i++ {
		ping(pings, fmt.Sprintf("passed message %d", i))
		pong(pings, pongs)

		fmt.Println(<-pongs)
	}
}
