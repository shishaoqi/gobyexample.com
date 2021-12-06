package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person{
	p := person{
		name: name,
	}
	p.age = 30
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name:"Alice", age:30})
	fmt.Println(person{name:"Fred"})
	fmt.Println(&person{name:"Andy", age: 40})

	fmt.Println(newPerson("jon"))
	ps := person{name: "Sean", age: 50}
	fmt.Println(ps.name)
	psPtr := &ps
	fmt.Println(psPtr.age)

	psPtr.age = 51
	fmt.Println(psPtr.age)

}
