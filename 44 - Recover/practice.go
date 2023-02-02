package main

import (
	"fmt"
)

func occurPanic() {
	panic("error problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover! Error: ", r)
		}
	}()

	occurPanic()

	fmt.Println("After panic")
}
