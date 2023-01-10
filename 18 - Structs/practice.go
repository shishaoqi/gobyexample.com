package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string) *Person {
	p := Person{
		name: "",
	}

	return &p
}

func main() {
	onePerson := NewPerson("shi")
	fmt.Println(onePerson)
}
