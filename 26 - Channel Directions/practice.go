package main

import "fmt"

func ping(ping chan<- string, msg string) {
	ping <- msg
}

func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hai")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
